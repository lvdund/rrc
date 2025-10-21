#!/usr/bin/env python3
"""
ASN.1 to Go Code Generator for LTE RRC Interface
Parses ASN.1 definitions and generates Go structs (one file per type)
Uses common types from utils/CommonType.go
"""

import re
import os
import sys
from dataclasses import dataclass, field
from typing import List, Dict, Optional, Set, Tuple, Any
from pathlib import Path


# ============================================================================
# Data Classes for ASN.1 Constructs
# ============================================================================

@dataclass
class ASN1Field:
    """Represents a field in a SEQUENCE or choice alternative"""
    name: str
    type_name: str
    optional: bool = False
    size_constraint: Optional[str] = None  # e.g., "1..10"
    range_constraint: Optional[str] = None  # e.g., "0..255"
    containing: Optional[str] = None  # CONTAINING clause
    default_value: Optional[str] = None
    comment: Optional[str] = None
    is_extension: bool = False  # Field is in extension marker
    extensible: bool = False  # Type itself is extensible


@dataclass
class ASN1Sequence:
    """SEQUENCE type"""
    name: str
    fields: List[ASN1Field]
    extensible: bool = False
    version_extensions: Dict[str, List[ASN1Field]] = field(default_factory=dict)


@dataclass
class ASN1Choice:
    """CHOICE type"""
    name: str
    alternatives: List[ASN1Field]
    extensible: bool = False


@dataclass
class ASN1SequenceOf:
    """SEQUENCE OF type (arrays)"""
    name: str
    element_type: str
    size_constraint: Optional[str] = None  # e.g., "1..maxDRB"


@dataclass
class ASN1Enumerated:
    """ENUMERATED type"""
    name: str
    values: List[str]
    extensible: bool = False


@dataclass
class ASN1Integer:
    """INTEGER type with constraints"""
    name: str
    range_constraint: Optional[str] = None


@dataclass
class ASN1BitString:
    """BIT STRING type"""
    name: str
    size_constraint: Optional[str] = None


@dataclass
class ASN1OctetString:
    """OCTET STRING type"""
    name: str
    size_constraint: Optional[str] = None
    containing: Optional[str] = None


@dataclass
class ASN1TypeAlias:
    """Simple type alias (TypeName ::= OtherType)"""
    name: str
    target_type: str


@dataclass
class ASN1Constant:
    """Constant definition"""
    name: str
    value: int


# ============================================================================
# ASN.1 Parser
# ============================================================================

class ASN1Parser:
    """Parser for ASN.1 definitions"""
    
    def __init__(self, content: str):
        self.content = content
        self.types: Dict[str, Any] = {}
        self.constants: Dict[str, ASN1Constant] = {}
        self.current_module = None
        
    def parse(self):
        """Main parsing entry point"""
        # Remove comments (-- comment text)
        content = re.sub(r'--[^\n]*', '', self.content)
        
        # Extract constants first (maxXXX definitions)
        self._extract_constants(content)
        
        # Extract type definitions
        self._extract_types(content)
        
        return self.types, self.constants
    
    def _extract_constants(self, content: str):
        """Extract constant definitions like maxDRB INTEGER ::= 11"""
        pattern = r'([a-z][a-zA-Z0-9_-]*)\s+INTEGER\s*::=\s*(\d+)'
        for match in re.finditer(pattern, content):
            name = match.group(1)
            value = int(match.group(2))
            self.constants[name] = ASN1Constant(name=name, value=value)
            # Also store with underscores for lookup
            name_underscore = name.replace('-', '_')
            self.constants[name_underscore] = ASN1Constant(name=name, value=value)
    
    def _extract_types(self, content: str):
        """Extract all type definitions"""
        # Pattern to match type definitions: TypeName ::= ...
        # We need to handle multi-line definitions
        lines = content.split('\n')
        i = 0
        
        while i < len(lines):
            line = lines[i].strip()
            
            # Check if this is a type definition start
            type_match = re.match(r'^([A-Z][a-zA-Z0-9_-]*)\s*::=\s*(.*)$', line)
            if type_match:
                type_name = type_match.group(1)
                rest = type_match.group(2).strip()
                
                # Collect the full definition (may span multiple lines)
                definition_lines = [rest]
                i += 1
                
                # Continue collecting until we find the end of definition
                brace_count = rest.count('{') - rest.count('}')
                paren_count = rest.count('(') - rest.count(')')
                
                while i < len(lines) and (brace_count > 0 or paren_count > 0 or 
                                         (definition_lines and not self._is_definition_complete(definition_lines))):
                    line = lines[i].strip()
                    if line:
                        definition_lines.append(line)
                        brace_count += line.count('{') - line.count('}')
                        paren_count += line.count('(') - line.count(')')
                    i += 1
                
                full_definition = ' '.join(definition_lines)
                
                # Parse based on type
                parsed_type = self._parse_type_definition(type_name, full_definition)
                if parsed_type:
                    self.types[type_name] = parsed_type
            else:
                i += 1
    
    def _is_definition_complete(self, lines: List[str]) -> bool:
        """Check if a type definition is complete"""
        text = ' '.join(lines)
        # Simple heuristic: if it ends with } or a known keyword pattern
        if re.search(r'\}\s*$', text):
            return True
        # Check for simple assignments that end immediately
        if re.match(r'^(SEQUENCE|CHOICE|INTEGER|ENUMERATED|BIT\s+STRING|OCTET\s+STRING|BOOLEAN|NULL)\s*$', text):
            return True
        return False
    
    def _parse_type_definition(self, name: str, definition: str) -> Optional[Any]:
        """Parse a type definition and return appropriate dataclass"""
        definition = definition.strip()
        
        # SEQUENCE type
        if definition.startswith('SEQUENCE {'):
            return self._parse_sequence(name, definition)
        
        # SEQUENCE OF type
        if definition.startswith('SEQUENCE (SIZE') or definition.startswith('SEQUENCE OF'):
            return self._parse_sequence_of(name, definition)
        
        # CHOICE type
        if definition.startswith('CHOICE {'):
            return self._parse_choice(name, definition)
        
        # ENUMERATED type
        if definition.startswith('ENUMERATED {') or definition.startswith('ENUMERATED{'):
            return self._parse_enumerated(name, definition)
        
        # INTEGER type
        if definition.startswith('INTEGER'):
            return self._parse_integer(name, definition)
        
        # BIT STRING type
        if definition.startswith('BIT STRING'):
            return self._parse_bit_string(name, definition)
        
        # OCTET STRING type
        if definition.startswith('OCTET STRING'):
            return self._parse_octet_string(name, definition)
        
        # BOOLEAN type
        if definition.startswith('BOOLEAN'):
            return ASN1TypeAlias(name=name, target_type='BOOLEAN')
        
        # NULL type
        if definition.startswith('NULL'):
            return ASN1TypeAlias(name=name, target_type='NULL')
        
        # Type alias (TypeName ::= OtherType)
        simple_type = re.match(r'^([A-Z][a-zA-Z0-9_-]*)(\s*\(.*\))?\s*$', definition)
        if simple_type:
            return ASN1TypeAlias(name=name, target_type=simple_type.group(1))
        
        return None
    
    def _parse_sequence(self, name: str, definition: str) -> ASN1Sequence:
        """Parse SEQUENCE definition"""
        # Extract content between { }
        match = re.search(r'SEQUENCE\s*\{(.*)\}', definition, re.DOTALL)
        if not match:
            return ASN1Sequence(name=name, fields=[])
        
        content = match.group(1)
        fields = []
        extensible = '...' in content
        
        # Split by comma, but respect nested braces
        field_defs = self._split_by_comma(content)
        
        for field_def in field_defs:
            field_def = field_def.strip()
            if not field_def or field_def == '...' or field_def.startswith('--'):
                continue
            
            # Handle extension groups [[...]]
            if field_def.startswith('[['):
                continue
            
            parsed_field = self._parse_field(field_def)
            if parsed_field:
                fields.append(parsed_field)
        
        return ASN1Sequence(name=name, fields=fields, extensible=extensible)
    
    def _parse_choice(self, name: str, definition: str) -> ASN1Choice:
        """Parse CHOICE definition"""
        match = re.search(r'CHOICE\s*\{(.*)\}', definition, re.DOTALL)
        if not match:
            return ASN1Choice(name=name, alternatives=[])
        
        content = match.group(1)
        alternatives = []
        extensible = '...' in content
        
        field_defs = self._split_by_comma(content)
        
        for field_def in field_defs:
            field_def = field_def.strip()
            if not field_def or field_def == '...' or field_def.startswith('--'):
                continue
            
            parsed_field = self._parse_field(field_def)
            if parsed_field:
                alternatives.append(parsed_field)
        
        return ASN1Choice(name=name, alternatives=alternatives, extensible=extensible)
    
    def _parse_sequence_of(self, name: str, definition: str) -> ASN1SequenceOf:
        """Parse SEQUENCE OF definition"""
        # SEQUENCE (SIZE (1..10)) OF ElementType
        match = re.match(r'SEQUENCE\s*\(SIZE\s*\(([^)]+)\)\)\s*OF\s+([A-Z][a-zA-Z0-9_-]+)', definition)
        if match:
            size = match.group(1).strip()
            elem_type = match.group(2).strip()
            return ASN1SequenceOf(name=name, element_type=elem_type, size_constraint=size)
        
        # SEQUENCE OF ElementType (without size)
        match = re.match(r'SEQUENCE\s*OF\s+([A-Z][a-zA-Z0-9_-]+)', definition)
        if match:
            elem_type = match.group(1).strip()
            return ASN1SequenceOf(name=name, element_type=elem_type)
        
        return ASN1SequenceOf(name=name, element_type='Unknown')
    
    def _parse_enumerated(self, name: str, definition: str) -> ASN1Enumerated:
        """Parse ENUMERATED definition"""
        match = re.search(r'ENUMERATED\s*\{([^}]+)\}', definition)
        if not match:
            return ASN1Enumerated(name=name, values=[])
        
        content = match.group(1)
        extensible = '...' in content
        
        # Extract enum values
        values = []
        parts = re.split(r',\s*', content)
        for part in parts:
            part = part.strip()
            if part and part != '...' and not part.startswith('--'):
                # Handle value assignments like "value1(0)" or just "value1"
                value_name = re.match(r'([a-zA-Z0-9_-]+)', part)
                if value_name:
                    values.append(value_name.group(1))
        
        return ASN1Enumerated(name=name, values=values, extensible=extensible)
    
    def _parse_integer(self, name: str, definition: str) -> ASN1Integer:
        """Parse INTEGER definition"""
        # Extract range constraint if present
        range_match = re.search(r'INTEGER\s*\(([^)]+)\)', definition)
        range_constraint = range_match.group(1).strip() if range_match else None
        return ASN1Integer(name=name, range_constraint=range_constraint)
    
    def _parse_bit_string(self, name: str, definition: str) -> ASN1BitString:
        """Parse BIT STRING definition"""
        size_match = re.search(r'BIT\s+STRING\s*\(SIZE\s*\(([^)]+)\)\)', definition)
        size_constraint = size_match.group(1).strip() if size_match else None
        return ASN1BitString(name=name, size_constraint=size_constraint)
    
    def _parse_octet_string(self, name: str, definition: str) -> ASN1OctetString:
        """Parse OCTET STRING definition"""
        size_match = re.search(r'OCTET\s+STRING\s*\(SIZE\s*\(([^)]+)\)\)', definition)
        size_constraint = size_match.group(1).strip() if size_match else None
        
        containing_match = re.search(r'CONTAINING\s+([A-Z][a-zA-Z0-9_-]+)', definition)
        containing = containing_match.group(1) if containing_match else None
        
        return ASN1OctetString(name=name, size_constraint=size_constraint, containing=containing)
    
    def _parse_field(self, field_def: str) -> Optional[ASN1Field]:
        """Parse a single field definition"""
        field_def = field_def.strip()
        if not field_def:
            return None
        
        # Check for OPTIONAL
        optional = 'OPTIONAL' in field_def
        field_def = re.sub(r'\s*OPTIONAL\s*', ' ', field_def)
        
        # Extract comment (-- Cond XXX or -- Need XX)
        comment_match = re.search(r'--\s*(.+)$', field_def)
        comment = comment_match.group(1).strip() if comment_match else None
        field_def = re.sub(r'--.*$', '', field_def).strip()
        
        # Extract CONTAINING clause
        containing_match = re.search(r'CONTAINING\s+([A-Z][a-zA-Z0-9_-]+)', field_def)
        containing = containing_match.group(1) if containing_match else None
        
        # Remove CONTAINING clause
        field_def = re.sub(r'\(CONTAINING[^)]+\)', '', field_def)
        
        # Handle inline SEQUENCE/CHOICE/ENUMERATED
        if 'SEQUENCE {' in field_def or 'CHOICE {' in field_def or 'ENUMERATED {' in field_def:
            # Extract field name before the inline type
            name_match = re.match(r'([a-zA-Z0-9_-]+)\s+', field_def)
            if name_match:
                field_name = name_match.group(1)
                # Determine type
                if 'SEQUENCE {' in field_def:
                    type_name = 'SEQUENCE'
                elif 'CHOICE {' in field_def:
                    type_name = 'CHOICE'
                elif 'ENUMERATED {' in field_def:
                    # Parse inline enumerated
                    enum_match = re.search(r'ENUMERATED\s*\{([^}]+)\}', field_def)
                    type_name = 'ENUMERATED'
                    extensible = '...' in field_def
                    return ASN1Field(name=field_name, type_name=type_name, optional=optional, 
                                   comment=comment, extensible=extensible)
                return ASN1Field(name=field_name, type_name=type_name, optional=optional, comment=comment)
        
        # Simple field: name type [constraints]
        # Need to handle multi-word types like "BIT STRING", "OCTET STRING"
        match = re.match(r'([a-zA-Z0-9_-]+)\s+(BIT\s+STRING|OCTET\s+STRING|[A-Z][a-zA-Z0-9_-]*|INTEGER|BOOLEAN|NULL|ENUMERATED)', field_def)
        if match:
            field_name = match.group(1)
            type_name = match.group(2)
            
            # Extract constraints
            size_constraint = None
            range_constraint = None
            
            # SIZE constraint
            size_match = re.search(r'SIZE\s*\(([^)]+)\)', field_def)
            if size_match:
                size_constraint = size_match.group(1).strip()
            
            # Range constraint for INTEGER
            if 'INTEGER' in field_def:
                range_match = re.search(r'INTEGER\s*\(([^)]+)\)', field_def)
                if range_match:
                    range_constraint = range_match.group(1).strip()
            
            # Check if type has inline ENUMERATED values
            extensible = False
            if 'ENUMERATED' in field_def and '{' in field_def:
                extensible = '...' in field_def
            
            return ASN1Field(
                name=field_name,
                type_name=type_name,
                optional=optional,
                size_constraint=size_constraint,
                range_constraint=range_constraint,
                containing=containing,
                comment=comment,
                extensible=extensible
            )
        
        return None
    
    def _split_by_comma(self, text: str) -> List[str]:
        """Split text by comma, respecting nested braces and parentheses"""
        parts = []
        current = []
        brace_depth = 0
        paren_depth = 0
        bracket_depth = 0
        
        i = 0
        while i < len(text):
            char = text[i]
            
            if char == '{':
                brace_depth += 1
            elif char == '}':
                brace_depth -= 1
            elif char == '(':
                paren_depth += 1
            elif char == ')':
                paren_depth -= 1
            elif char == '[':
                bracket_depth += 1
            elif char == ']':
                bracket_depth -= 1
            elif char == ',' and brace_depth == 0 and paren_depth == 0 and bracket_depth == 0:
                parts.append(''.join(current))
                current = []
                i += 1
                continue
            
            current.append(char)
            i += 1
        
        if current:
            parts.append(''.join(current))
        
        return parts


# ============================================================================
# Go Code Generator
# ============================================================================

class GoGenerator:
    """Generate Go code from ASN.1 types using CommonType.go primitives"""
    
    def __init__(self, types: Dict[str, Any], constants: Dict[str, ASN1Constant], output_dir: str = 'ies'):
        self.types = types
        self.constants = constants
        self.output_dir = Path(output_dir)
        self.package_name = 'ies'
        
    def generate(self):
        """Generate all Go files"""
        # Create output directory
        self.output_dir.mkdir(exist_ok=True)
        
        print(f"Generating Go code to {self.output_dir}/")
        print(f"Found {len(self.types)} types and {len(self.constants)} constants")
        
        # Generate one file per type
        count = 0
        for type_name, type_def in self.types.items():
            self._generate_type_file(type_name, type_def)
            count += 1
            if count % 100 == 0:
                print(f"  Generated {count}/{len(self.types)} files...")
        
        print(f"✓ Generated {len(self.types)} type files")
    
    def _generate_type_file(self, type_name: str, type_def: Any):
        """Generate a single .go file for a type"""
        go_name = self._to_go_name(type_name)
        filename = f"{go_name}.go"
        filepath = self.output_dir / filename
        
        with open(filepath, 'w') as f:
            f.write(f"package {self.package_name}\n\n")
            f.write('import "rrc/utils"\n\n')
            
            # Generate based on type
            if isinstance(type_def, ASN1Sequence):
                self._generate_sequence(f, type_def)
            elif isinstance(type_def, ASN1Choice):
                self._generate_choice(f, type_def)
            elif isinstance(type_def, ASN1SequenceOf):
                self._generate_sequence_of(f, type_def)
            elif isinstance(type_def, ASN1Enumerated):
                self._generate_enumerated(f, type_def)
            elif isinstance(type_def, ASN1Integer):
                self._generate_integer(f, type_def)
            elif isinstance(type_def, ASN1BitString):
                self._generate_bit_string(f, type_def)
            elif isinstance(type_def, ASN1OctetString):
                self._generate_octet_string(f, type_def)
            elif isinstance(type_def, ASN1TypeAlias):
                self._generate_type_alias(f, type_def)
    
    def _generate_sequence(self, f, seq: ASN1Sequence):
        """Generate Go struct for SEQUENCE using CommonType wrappers"""
        go_name = self._to_go_name(seq.name)
        
        f.write(f"// {seq.name} ::= SEQUENCE\n")
        if seq.extensible:
            f.write("// Extensible\n")
        f.write(f"type {go_name} struct {{\n")
        
        for field in seq.fields:
            self._generate_field(f, field)
        
        f.write("}\n")
    
    def _generate_choice(self, f, choice: ASN1Choice):
        """Generate Go interface for CHOICE"""
        go_name = self._to_go_name(choice.name)
        
        f.write(f"// {choice.name} ::= CHOICE\n")
        if choice.extensible:
            f.write("// Extensible\n")
        
        # Generate interface
        f.write(f"type {go_name} interface {{\n")
        f.write(f"\tis{go_name}()\n")
        f.write("}\n\n")
        
        # Generate concrete types for each alternative
        for alt in choice.alternatives:
            alt_go_name = f"{go_name}{self._to_go_name(alt.name)}"
            f.write(f"type {alt_go_name} struct {{\n")
            f.write(f"\tValue {self._field_type_to_go(alt)}\n")
            f.write("}\n\n")
            f.write(f"func (*{alt_go_name}) is{go_name}() {{}}\n\n")
    
    def _generate_sequence_of(self, f, seq_of: ASN1SequenceOf):
        """Generate Go type for SEQUENCE OF using Sequence[T] wrapper"""
        go_name = self._to_go_name(seq_of.name)
        elem_go_type = self._to_go_name(seq_of.element_type)
        
        f.write(f"// {seq_of.name} ::= SEQUENCE OF {seq_of.element_type}\n")
        if seq_of.size_constraint:
            f.write(f"// SIZE ({seq_of.size_constraint})\n")
        
        # Use Sequence[T] from CommonType.go
        f.write(f"type {go_name} struct {{\n")
        f.write(f"\tValue utils.Sequence[{elem_go_type}]\n")
        f.write("}\n")
    
    def _generate_enumerated(self, f, enum: ASN1Enumerated):
        """Generate Go type for ENUMERATED using ENUMERATED wrapper"""
        go_name = self._to_go_name(enum.name)
        
        f.write(f"// {enum.name} ::= ENUMERATED\n")
        if enum.extensible:
            f.write("// Extensible\n")
        
        # Generate type
        f.write(f"type {go_name} struct {{\n")
        f.write("\tValue utils.ENUMERATED\n")
        f.write("}\n\n")
        
        # Generate constants for enum values
        f.write("const (\n")
        for i, value in enumerate(enum.values):
            const_name = f"{go_name}{self._to_go_name(value, capitalize_first=True)}"
            f.write(f"\t{const_name} = {i}\n")
        f.write(")\n")
    
    def _generate_integer(self, f, integer: ASN1Integer):
        """Generate Go type for INTEGER using INTEGER wrapper"""
        go_name = self._to_go_name(integer.name)
        
        f.write(f"// {integer.name} ::= INTEGER")
        if integer.range_constraint:
            f.write(f" ({integer.range_constraint})")
        f.write("\n")
        
        f.write(f"type {go_name} struct {{\n")
        f.write("\tValue utils.INTEGER\n")
        f.write("}\n")
    
    def _generate_bit_string(self, f, bit_str: ASN1BitString):
        """Generate Go type for BIT STRING using BITSTRING wrapper"""
        go_name = self._to_go_name(bit_str.name)
        
        f.write(f"// {bit_str.name} ::= BIT STRING")
        if bit_str.size_constraint:
            f.write(f" (SIZE ({bit_str.size_constraint}))")
        f.write("\n")
        
        f.write(f"type {go_name} struct {{\n")
        f.write("\tValue utils.BITSTRING\n")
        f.write("}\n")
    
    def _generate_octet_string(self, f, octet_str: ASN1OctetString):
        """Generate Go type for OCTET STRING using OCTETSTRING wrapper"""
        go_name = self._to_go_name(octet_str.name)
        
        f.write(f"// {octet_str.name} ::= OCTET STRING")
        if octet_str.size_constraint:
            f.write(f" (SIZE ({octet_str.size_constraint}))")
        if octet_str.containing:
            f.write(f" (CONTAINING {octet_str.containing})")
        f.write("\n")
        
        f.write(f"type {go_name} struct {{\n")
        f.write("\tValue utils.OCTETSTRING\n")
        f.write("}\n")
    
    def _generate_type_alias(self, f, alias: ASN1TypeAlias):
        """Generate Go type alias"""
        go_name = self._to_go_name(alias.name)
        target_go = self._to_go_name(alias.target_type)
        
        f.write(f"// {alias.name} ::= {alias.target_type}\n")
        
        # Special handling for built-in types
        if alias.target_type == 'BOOLEAN':
            f.write(f"type {go_name} bool\n")
        elif alias.target_type == 'NULL':
            f.write(f"type {go_name} struct{{}}\n")
        else:
            f.write(f"type {go_name} {target_go}\n")
    
    def _generate_field(self, f, field: ASN1Field):
        """Generate a struct field"""
        go_field_name = self._to_go_name(field.name)
        
        # Handle inline types with constraints
        if field.type_name in ['INTEGER', 'BIT STRING', 'OCTET STRING', 'ENUMERATED']:
            # For inline primitive types in fields, we can either:
            # 1. Just use the utils type directly
            # 2. Generate a wrapper for each field (more complex)
            # Let's use the simple approach - use utils types directly
            go_type = self._field_type_to_go(field)
        else:
            go_type = self._field_type_to_go(field)
        
        # Make pointer if optional
        if field.optional:
            go_type = f"*{go_type}"
        
        # Add comment if present
        comment_str = ""
        if field.comment:
            comment_str = f" // {field.comment}"
        
        f.write(f"\t{go_field_name} {go_type}{comment_str}\n")
    
    def _field_type_to_go(self, field: ASN1Field) -> str:
        """Convert field type to Go type, handling constraints"""
        type_name = field.type_name
        
        # Handle built-in types with CommonType wrappers
        if type_name == 'INTEGER':
            return 'utils.INTEGER'
        elif type_name == 'BOOLEAN':
            return 'bool'
        elif type_name == 'ENUMERATED':
            return 'utils.ENUMERATED'
        elif type_name == 'BIT STRING':
            return 'utils.BITSTRING'
        elif type_name == 'OCTET STRING':
            return 'utils.OCTETSTRING'
        elif type_name == 'NULL':
            return 'struct{}'
        elif type_name == 'SEQUENCE' or type_name == 'CHOICE':
            # Inline type - would need special handling
            return 'interface{}'
        else:
            # Custom type reference
            return self._to_go_name(type_name)
    
    def _to_go_name(self, asn1_name: str, capitalize_first: bool = True) -> str:
        """Convert ASN.1 name to Go name (CamelCase)"""
        # Replace hyphens with underscores, then convert to CamelCase
        parts = asn1_name.replace('-', '_').split('_')
        
        if capitalize_first:
            return ''.join(word.capitalize() for word in parts)
        else:
            # For enum values, keep original casing more
            return ''.join(word.capitalize() if i > 0 else word for i, word in enumerate(parts))
    
    def _resolve_constraint_bounds(self, constraint: Optional[str]) -> Tuple[int, int]:
        """Resolve constraint bounds, handling constant references"""
        if not constraint:
            return (0, 0)  # No constraint
        
        # Clean up whitespace
        constraint = constraint.strip()
        
        # Handle range: "min..max" or "min .. max" (with spaces)
        match = re.match(r'(-?\d+)\s*\.\.\s*(-?\d+|[a-z][a-zA-Z0-9_-]*)', constraint)
        if match:
            lb = int(match.group(1))
            ub_str = match.group(2).strip()
            
            # Check if upper bound is a number or constant reference
            if ub_str.isdigit() or (ub_str.startswith('-') and ub_str[1:].isdigit()):
                ub = int(ub_str)
            else:
                # It's a constant reference - resolve it
                ub = self._resolve_constant(ub_str)
            
            return (lb, ub)
        
        # Handle single value or constant
        if constraint.isdigit():
            val = int(constraint)
            return (val, val)
        else:
            # Constant reference
            val = self._resolve_constant(constraint)
            return (val, val)
    
    def _resolve_constant(self, const_name: str) -> int:
        """Resolve a constant name to its integer value"""
        # Try with original name
        if const_name in self.constants:
            return self.constants[const_name].value
        
        # Try with underscores
        const_name_underscore = const_name.replace('-', '_')
        if const_name_underscore in self.constants:
            return self.constants[const_name_underscore].value
        
        # Default fallback
        print(f"Warning: Could not resolve constant '{const_name}', using 0")
        return 0


# ============================================================================
# Main Entry Point
# ============================================================================

def main():
    """Main function"""
    import subprocess
    
    asn1_file = 'lte-rrc-16.13.0.asn1'
    output_dir = 'ies'
    
    if not os.path.exists(asn1_file):
        print(f"Error: {asn1_file} not found")
        sys.exit(1)
    
    print(f"Reading {asn1_file}...")
    with open(asn1_file, 'r', encoding='utf-8') as f:
        content = f.read()
    
    print("Parsing ASN.1 definitions...")
    parser = ASN1Parser(content)
    types, constants = parser.parse()
    
    print(f"Parsed {len(types)} types and {len(constants)} constants")
    
    print(f"\nGenerating Go code to {output_dir}/...")
    generator = GoGenerator(types, constants, output_dir)
    generator.generate()
    
    print("\n✓ Generation complete!")
    print(f"  Output directory: {output_dir}/")
    print(f"  Files generated: {len(types)}")
    
    # Format all Go files
    print(f"\nFormatting Go files in {output_dir}/...")
    try:
        result = subprocess.run(
            ['gofmt', '-w', output_dir],
            capture_output=True,
            text=True,
            check=True
        )
        print("✓ Go files formatted successfully")
    except subprocess.CalledProcessError as e:
        print(f"Warning: gofmt failed: {e}")
    except FileNotFoundError:
        print("Warning: gofmt not found, skipping formatting")


if __name__ == '__main__':
    main()
