#!/usr/bin/env python3
"""
ASN.1 to Go Code Generator for LTE RRC Interface
Generates Go code with CHOICE discriminators, constraint tags, and proper type handling
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
    size_constraint: Optional[str] = None
    range_constraint: Optional[str] = None
    containing: Optional[str] = None
    comment: Optional[str] = None
    is_extension: bool = False
    extensible: bool = False


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
    size_constraint: Optional[str] = None


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
# ASN.1 Parser - (keeping existing implementation)
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
        content = re.sub(r'--[^\n]*', '', self.content)
        self._extract_constants(content)
        self._extract_types(content)
        return self.types, self.constants
    
    def _extract_constants(self, content: str):
        """Extract constant definitions like maxDRB INTEGER ::= 11"""
        pattern = r'([a-z][a-zA-Z0-9_-]*)\s+INTEGER\s*::=\s*(\d+)'
        for match in re.finditer(pattern, content):
            name = match.group(1)
            value = int(match.group(2))
            self.constants[name] = ASN1Constant(name=name, value=value)
            name_underscore = name.replace('-', '_')
            self.constants[name_underscore] = ASN1Constant(name=name, value=value)
    
    def _extract_types(self, content: str):
        """Extract all type definitions"""
        lines = content.split('\n')
        i = 0
        
        while i < len(lines):
            line = lines[i].strip()
            type_match = re.match(r'^([A-Z][a-zA-Z0-9_-]*)\s*::=\s*(.*)$', line)
            if type_match:
                type_name = type_match.group(1)
                rest = type_match.group(2).strip()
                definition_lines = [rest]
                i += 1
                
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
                parsed_type = self._parse_type_definition(type_name, full_definition)
                if parsed_type:
                    self.types[type_name] = parsed_type
            else:
                i += 1
    
    def _is_definition_complete(self, lines: List[str]) -> bool:
        """Check if a type definition is complete"""
        text = ' '.join(lines)
        if re.search(r'\}\s*$', text):
            return True
        # Note: ENUMERATED is excluded here because it often has { on next line
        if re.match(r'^(SEQUENCE|CHOICE|INTEGER|BIT\s+STRING|BITSTRING|OCTET\s+STRING|OCTETSTRING|BOOLEAN|NULL)\s*$', text):
            return True
        # SEQUENCE OF is complete on one line (with or without constraints on element type)
        # Pattern: SEQUENCE (SIZE (...))? OF TypeName (...)?
        if re.match(r'^SEQUENCE\s*(\(SIZE\s*\([^)]+\)\))?\s*OF\s+', text):
            # Check if all parens are balanced
            if text.count('(') == text.count(')'):
                return True
        # BIT STRING / BITSTRING or OCTET STRING / OCTETSTRING with constraints is complete
        if re.match(r'^(BIT\s+STRING|BITSTRING|OCTET\s+STRING|OCTETSTRING)\s*\(', text):
            # Check if all parens are balanced
            if text.count('(') == text.count(')'):
                return True
        # INTEGER with constraints is complete
        if re.match(r'^INTEGER\s*\(', text):
            if text.count('(') == text.count(')'):
                return True
        # Type alias or simple type reference is complete
        # But not standalone ENUMERATED (which needs {})
        if re.match(r'^[A-Z][a-zA-Z0-9_-]*(\s*\([^)]+\))?\s*$', text) and text.strip() != 'ENUMERATED':
            return True
        return False
    
    def _parse_type_definition(self, name: str, definition: str) -> Optional[Any]:
        """Parse a type definition and return appropriate dataclass"""
        definition = definition.strip()
        
        # Use regex to handle variable whitespace
        if re.match(r'^SEQUENCE\s*\{', definition):
            return self._parse_sequence(name, definition)
        elif re.match(r'^SEQUENCE\s*\(SIZE', definition) or re.match(r'^SEQUENCE\s+OF', definition):
            return self._parse_sequence_of(name, definition)
        elif re.match(r'^CHOICE\s*\{', definition):
            return self._parse_choice(name, definition)
        elif 'ENUMERATED' in definition and '{' in definition:
            # ENUMERATED can have { on same or different line
            return self._parse_enumerated(name, definition)
        elif definition.startswith('INTEGER'):
            return self._parse_integer(name, definition)
        elif definition.startswith('BIT STRING') or definition.startswith('BITSTRING'):
            return self._parse_bit_string(name, definition)
        elif definition.startswith('OCTET STRING') or definition.startswith('OCTETSTRING'):
            return self._parse_octet_string(name, definition)
        elif definition.startswith('BOOLEAN'):
            return ASN1TypeAlias(name=name, target_type='BOOLEAN')
        elif definition.startswith('NULL'):
            return ASN1TypeAlias(name=name, target_type='NULL')
        elif definition.startswith('ENUMERATED'):
            # Standalone ENUMERATED without { - should not happen if parsed correctly
            # But if it does, treat as incomplete and log warning
            print(f"Warning: ENUMERATED type {name} has no values defined")
            return None
        else:
            simple_type = re.match(r'^([A-Z][a-zA-Z0-9_-]*)(\s*\(.*\))?\s*$', definition)
            if simple_type:
                return ASN1TypeAlias(name=name, target_type=simple_type.group(1))
        
        return None
    
    def _parse_sequence(self, name: str, definition: str) -> ASN1Sequence:
        """Parse SEQUENCE definition"""
        match = re.search(r'SEQUENCE\s*\{(.*)\}', definition, re.DOTALL)
        if not match:
            return ASN1Sequence(name=name, fields=[])
        
        content = match.group(1)
        fields = []
        extensible = '...' in content
        
        # Extract inline types first
        self._extract_inline_types_from_sequence(name, definition)
        
        field_defs = self._split_by_comma(content)
        
        for field_def in field_defs:
            field_def = field_def.strip()
            if not field_def or field_def == '...' or field_def.startswith('--'):
                continue
            
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
        
        # Extract inline types first
        self._extract_inline_types_from_choice(name, definition)
        
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
        # Check for SEQUENCE OF CHOICE (inline CHOICE)
        choice_match = re.match(r'SEQUENCE\s*\(SIZE\s*\(([^)]+)\)\)\s*OF\s+CHOICE\s*\{', definition)
        if choice_match:
            size = choice_match.group(1).strip()
            # Extract the CHOICE definition
            choice_content_match = re.search(r'CHOICE\s*\{(.*)\}', definition, re.DOTALL)
            if choice_content_match:
                choice_content = choice_content_match.group(1)
                # Create an inline CHOICE type
                inline_choice_name = f"{name}-Item"
                inline_choice = self._parse_choice_from_content(inline_choice_name, choice_content)
                if inline_choice:
                    self.types[inline_choice_name] = inline_choice
                    # Also extract nested inline types from this CHOICE
                    self._extract_inline_types_from_choice(inline_choice_name, f"CHOICE {{{choice_content}}}")
                return ASN1SequenceOf(name=name, element_type=inline_choice_name, size_constraint=size)
        
        match = re.match(r'SEQUENCE\s*\(SIZE\s*\(([^)]+)\)\)\s*OF\s+([A-Z][a-zA-Z0-9_-]+)', definition)
        if match:
            size = match.group(1).strip()
            elem_type = match.group(2).strip()
            return ASN1SequenceOf(name=name, element_type=elem_type, size_constraint=size)
        
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
        
        values = []
        parts = re.split(r',\s*', content)
        for part in parts:
            part = part.strip()
            if part and part != '...' and not part.startswith('--'):
                value_name = re.match(r'([a-zA-Z0-9_-]+)', part)
                if value_name:
                    values.append(value_name.group(1))
        
        return ASN1Enumerated(name=name, values=values, extensible=extensible)
    
    def _parse_integer(self, name: str, definition: str) -> ASN1Integer:
        """Parse INTEGER definition"""
        range_match = re.search(r'INTEGER\s*\(([^)]+)\)', definition)
        range_constraint = range_match.group(1).strip() if range_match else None
        return ASN1Integer(name=name, range_constraint=range_constraint)
    
    def _parse_bit_string(self, name: str, definition: str) -> ASN1BitString:
        """Parse BIT STRING or BITSTRING definition"""
        # Handle both "BIT STRING" and "BITSTRING"
        size_match = re.search(r'(BIT\s+STRING|BITSTRING)\s*\(SIZE\s*\(([^)]+)\)\)', definition)
        size_constraint = size_match.group(2).strip() if size_match else None
        return ASN1BitString(name=name, size_constraint=size_constraint)
    
    def _parse_octet_string(self, name: str, definition: str) -> ASN1OctetString:
        """Parse OCTET STRING or OCTETSTRING definition"""
        # Handle both "OCTET STRING" and "OCTETSTRING"
        size_match = re.search(r'(OCTET\s+STRING|OCTETSTRING)\s*\(SIZE\s*\(([^)]+)\)\)', definition)
        size_constraint = size_match.group(2).strip() if size_match else None
        
        containing_match = re.search(r'CONTAINING\s+([A-Z][a-zA-Z0-9_-]+)', definition)
        containing = containing_match.group(1) if containing_match else None
        
        return ASN1OctetString(name=name, size_constraint=size_constraint, containing=containing)
    
    def _parse_field(self, field_def: str) -> Optional[ASN1Field]:
        """Parse a single field definition"""
        field_def = field_def.strip()
        if not field_def:
            return None
        
        optional = 'OPTIONAL' in field_def
        field_def = re.sub(r'\s*OPTIONAL\s*', ' ', field_def)
        
        comment_match = re.search(r'--\s*(.+)$', field_def)
        comment = comment_match.group(1).strip() if comment_match else None
        field_def = re.sub(r'--.*$', '', field_def).strip()
        
        containing_match = re.search(r'CONTAINING\s+([A-Z][a-zA-Z0-9_-]+)', field_def)
        containing = containing_match.group(1) if containing_match else None
        
        field_def = re.sub(r'\(CONTAINING[^)]+\)', '', field_def)
        
        # Handle inline SEQUENCE OF (array/list)
        # First check for SEQUENCE OF CHOICE (special case)
        seq_of_choice_match = re.match(r'([a-zA-Z0-9_-]+)\s+SEQUENCE\s*\(SIZE\s*\(([^)]+)\)\)\s*OF\s+CHOICE\s*\{', field_def, re.DOTALL)
        if seq_of_choice_match:
            field_name = seq_of_choice_match.group(1)
            size_constraint = seq_of_choice_match.group(2).strip()
            # The element type will be ParentName-FieldName-Item (created in _extract_inline_types_from_sequence)
            # For now, use a placeholder that will be resolved later
            return ASN1Field(
                name=field_name,
                type_name='SEQUENCE_OF_INLINE_CHOICE',
                optional=optional,
                size_constraint=size_constraint,
                comment=comment
            )
        
        # Then check for regular SEQUENCE OF
        seq_of_match = re.match(r'([a-zA-Z0-9_-]+)\s+SEQUENCE\s*\(SIZE\s*\(([^)]+)\)\)\s*OF\s+([A-Z][a-zA-Z0-9_-]+)', field_def)
        if seq_of_match:
            field_name = seq_of_match.group(1)
            size_constraint = seq_of_match.group(2).strip()
            element_type = seq_of_match.group(3)
            # Return as a SEQUENCE OF field with the element type
            return ASN1Field(
                name=field_name,
                type_name=f'SEQUENCE_OF_{element_type}',
                optional=optional,
                size_constraint=size_constraint,
                comment=comment
            )
        
        # Handle inline SEQUENCE/CHOICE/ENUMERATED definitions
        if 'SEQUENCE {' in field_def or 'CHOICE {' in field_def or 'ENUMERATED {' in field_def:
            name_match = re.match(r'([a-zA-Z0-9_-]+)\s+', field_def)
            if name_match:
                field_name = name_match.group(1)
                if 'SEQUENCE {' in field_def:
                    type_name = 'SEQUENCE'
                elif 'CHOICE {' in field_def:
                    type_name = 'CHOICE'
                elif 'ENUMERATED {' in field_def:
                    type_name = 'ENUMERATED'
                    extensible = '...' in field_def
                    return ASN1Field(name=field_name, type_name=type_name, optional=optional, 
                                   comment=comment, extensible=extensible)
                return ASN1Field(name=field_name, type_name=type_name, optional=optional, comment=comment)
        
        match = re.match(r'([a-zA-Z0-9_-]+)\s+(BIT\s+STRING|OCTET\s+STRING|[A-Z][a-zA-Z0-9_-]*|INTEGER|BOOLEAN|NULL|ENUMERATED)', field_def)
        if match:
            field_name = match.group(1)
            type_name = match.group(2)
            
            size_constraint = None
            range_constraint = None
            
            size_match = re.search(r'SIZE\s*\(([^)]+)\)', field_def)
            if size_match:
                size_constraint = size_match.group(1).strip()
            
            if 'INTEGER' in field_def:
                range_match = re.search(r'INTEGER\s*\(([^)]+)\)', field_def)
                if range_match:
                    range_constraint = range_match.group(1).strip()
            
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
    
    def _extract_inline_types_from_choice(self, name: str, definition: str):
        """Extract and register inline SEQUENCE/CHOICE types from CHOICE alternatives"""
        match = re.search(r'CHOICE\s*\{(.*)\}', definition, re.DOTALL)
        if not match:
            return
        
        content = match.group(1)
        field_defs = self._split_by_comma(content)
        
        for field_def in field_defs:
            field_def = field_def.strip()
            if not field_def or field_def == '...' or field_def.startswith('--'):
                continue
            
            # Check for inline SEQUENCE
            seq_match = re.match(r'([a-zA-Z0-9_-]+)\s+SEQUENCE\s*\{(.*)\}', field_def, re.DOTALL)
            if seq_match:
                alt_name = seq_match.group(1)
                seq_content = seq_match.group(2)
                inline_type_name = f"{name}-{alt_name}"
                
                # Parse the inline SEQUENCE
                inline_seq = self._parse_sequence_from_content(inline_type_name, seq_content)
                if inline_seq and inline_type_name not in self.types:
                    self.types[inline_type_name] = inline_seq
                    # Also extract nested inline types from this SEQUENCE
                    self._extract_inline_types_from_sequence(inline_type_name, f"SEQUENCE {{{seq_content}}}")
                continue
            
            # Check for inline CHOICE
            choice_match = re.match(r'([a-zA-Z0-9_-]+)\s+CHOICE\s*\{(.*)\}', field_def, re.DOTALL)
            if choice_match:
                alt_name = choice_match.group(1)
                choice_content = choice_match.group(2)
                inline_type_name = f"{name}-{alt_name}"
                
                # Parse the inline CHOICE
                inline_choice = self._parse_choice_from_content(inline_type_name, choice_content)
                if inline_choice and inline_type_name not in self.types:
                    self.types[inline_type_name] = inline_choice
                    # Recursively extract nested inline types
                    self._extract_inline_types_from_choice(inline_type_name, f"CHOICE {{{choice_content}}}")
                continue
    
    def _extract_inline_types_from_sequence(self, name: str, definition: str):
        """Extract and register inline SEQUENCE/CHOICE/ENUMERATED types from SEQUENCE fields"""
        match = re.search(r'SEQUENCE\s*\{(.*)\}', definition, re.DOTALL)
        if not match:
            return
        
        content = match.group(1)
        field_defs = self._split_by_comma(content)
        
        for field_def in field_defs:
            field_def = field_def.strip()
            if not field_def or field_def == '...' or field_def.startswith('--') or field_def.startswith('[['):
                continue
            
            # Check for inline SEQUENCE
            seq_match = re.match(r'([a-zA-Z0-9_-]+)\s+SEQUENCE\s*\{(.*)\}', field_def, re.DOTALL)
            if seq_match:
                field_name = seq_match.group(1)
                seq_content = seq_match.group(2)
                inline_type_name = f"{name}-{field_name}"
                
                inline_seq = self._parse_sequence_from_content(inline_type_name, seq_content)
                if inline_seq and inline_type_name not in self.types:
                    self.types[inline_type_name] = inline_seq
                    self._extract_inline_types_from_sequence(inline_type_name, f"SEQUENCE {{{seq_content}}}")
                continue
            
            # Check for inline CHOICE
            choice_match = re.match(r'([a-zA-Z0-9_-]+)\s+CHOICE\s*\{(.*)\}', field_def, re.DOTALL)
            if choice_match:
                field_name = choice_match.group(1)
                choice_content = choice_match.group(2)
                inline_type_name = f"{name}-{field_name}"
                
                inline_choice = self._parse_choice_from_content(inline_type_name, choice_content)
                if inline_choice and inline_type_name not in self.types:
                    self.types[inline_type_name] = inline_choice
                    self._extract_inline_types_from_choice(inline_type_name, f"CHOICE {{{choice_content}}}")
                continue
            
            # Check for inline ENUMERATED
            enum_match = re.match(r'([a-zA-Z0-9_-]+)\s+ENUMERATED\s*\{([^}]+)\}', field_def)
            if enum_match:
                field_name = enum_match.group(1)
                enum_content = enum_match.group(2)
                
                # Parse the inline ENUMERATED
                extensible = '...' in enum_content
                values = []
                parts = re.split(r',\s*', enum_content)
                for part in parts:
                    part = part.strip()
                    if part and part != '...' and not part.startswith('--'):
                        value_name = re.match(r'([a-zA-Z0-9_-]+)', part)
                        if value_name:
                            values.append(value_name.group(1))
                
                # Special case: ENUMERATED {true} or other single-value enums should not create a separate type
                # They will be handled as bool in field generation
                if len(values) > 1 or (len(values) == 1 and values[0].lower() != 'true'):
                    inline_type_name = f"{name}-{field_name}"
                    inline_enum = ASN1Enumerated(name=inline_type_name, values=values, extensible=extensible)
                    if inline_type_name not in self.types:
                        self.types[inline_type_name] = inline_enum
                continue
            
            # Check for inline SEQUENCE OF CHOICE
            seq_of_choice_match = re.match(r'([a-zA-Z0-9_-]+)\s+SEQUENCE\s*\(SIZE\s*\([^)]+\)\)\s*OF\s+CHOICE\s*\{', field_def, re.DOTALL)
            if seq_of_choice_match:
                field_name = seq_of_choice_match.group(1)
                # Extract the CHOICE definition
                choice_content_match = re.search(r'CHOICE\s*\{(.*)\}', field_def, re.DOTALL)
                if choice_content_match:
                    choice_content = choice_content_match.group(1)
                    inline_type_name = f"{name}-{field_name}-Item"
                    
                    # Parse the inline CHOICE
                    inline_choice = self._parse_choice_from_content(inline_type_name, choice_content)
                    if inline_choice and inline_type_name not in self.types:
                        self.types[inline_type_name] = inline_choice
                        # Recursively extract nested inline types
                        self._extract_inline_types_from_choice(inline_type_name, f"CHOICE {{{choice_content}}}")
                continue
    
    def _parse_sequence_from_content(self, name: str, content: str) -> Optional[ASN1Sequence]:
        """Parse SEQUENCE from its inner content"""
        fields = []
        extensible = '...' in content
        
        field_defs = self._split_by_comma(content)
        
        for field_def in field_defs:
            field_def = field_def.strip()
            if not field_def or field_def == '...' or field_def.startswith('--') or field_def.startswith('[['):
                continue
            
            parsed_field = self._parse_field(field_def)
            if parsed_field:
                fields.append(parsed_field)
        
        return ASN1Sequence(name=name, fields=fields, extensible=extensible)
    
    def _parse_choice_from_content(self, name: str, content: str) -> Optional[ASN1Choice]:
        """Parse CHOICE from its inner content"""
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
    """Generate Go code from ASN.1 types"""
    
    def __init__(self, types: Dict[str, Any], constants: Dict[str, ASN1Constant], output_dir: str = 'ies'):
        self.types = types
        self.constants = constants
        self.output_dir = Path(output_dir)
        self.package_name = 'ies'
        
    def generate(self):
        """Generate all Go files"""
        self.output_dir.mkdir(exist_ok=True)
        
        print(f"Generating Go code to {self.output_dir}/")
        print(f"Found {len(self.types)} types and {len(self.constants)} constants")
        
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
        """Generate Go struct for SEQUENCE"""
        go_name = self._to_go_name(seq.name)
        
        f.write(f"// {seq.name} ::= SEQUENCE\n")
        if seq.extensible:
            f.write("// Extensible\n")
        f.write(f"type {go_name} struct {{\n")
        
        for field in seq.fields:
            self._generate_field(f, field, seq.name)
        
        f.write("}\n")
    
    def _generate_choice(self, f, choice: ASN1Choice):
        """Generate Go struct for CHOICE with Choice discriminator and constants"""
        go_name = self._to_go_name(choice.name)
        
        f.write(f"// {choice.name} ::= CHOICE\n")
        if choice.extensible:
            f.write("// Extensible\n")
        
        # Generate constants for alternatives
        f.write("const (\n")
        f.write(f"\t{go_name}ChoiceNothing = iota\n")
        for alt in choice.alternatives:
            const_name = f"{go_name}Choice{self._to_go_name(alt.name)}"
            f.write(f"\t{const_name}\n")
        f.write(")\n\n")
        
        # Generate struct with Choice discriminator and pointer fields
        f.write(f"type {go_name} struct {{\n")
        f.write(f"\tChoice uint64\n")
        
        for alt in choice.alternatives:
            alt_field_name = self._to_go_name(alt.name)
            alt_type = self._choice_alt_type_to_go(alt, go_name)
            
            # Build constraint tag for primitive types and inline SEQUENCE OF
            tag_str = ""
            if alt.size_constraint or alt.range_constraint:
                constraint = alt.size_constraint or alt.range_constraint
                lb, ub = self._resolve_constraint_bounds(constraint)
                tag_str = f" `lb:{lb},ub:{ub}`"
            
            # For inline SEQUENCE OF in CHOICE, the type is already a slice
            # so we need pointer to slice: *[]Type
            f.write(f"\t{alt_field_name} *{alt_type}{tag_str}\n")
        
        f.write("}\n")
    
    def _generate_sequence_of(self, f, seq_of: ASN1SequenceOf):
        """Generate Go slice type for SEQUENCE OF"""
        go_name = self._to_go_name(seq_of.name)
        
        # Handle primitive types specially
        if seq_of.element_type == 'INTEGER':
            elem_go_type = 'utils.INTEGER'
        elif seq_of.element_type == 'BOOLEAN':
            elem_go_type = 'utils.BOOLEAN'
        elif seq_of.element_type == 'BIT STRING' or seq_of.element_type == 'BITSTRING':
            elem_go_type = 'utils.BITSTRING'
        elif seq_of.element_type == 'OCTET STRING' or seq_of.element_type == 'OCTETSTRING':
            elem_go_type = 'utils.OCTETSTRING'
        elif seq_of.element_type == 'ENUMERATED':
            elem_go_type = 'utils.ENUMERATED'
        else:
            elem_go_type = self._to_go_name(seq_of.element_type)
        
        f.write(f"// {seq_of.name} ::= SEQUENCE OF {seq_of.element_type}\n")
        if seq_of.size_constraint:
            f.write(f"// SIZE ({seq_of.size_constraint})\n")
        
        # Build constraint tag for the Value field
        tag_str = ""
        if seq_of.size_constraint:
            lb, ub = self._resolve_constraint_bounds(seq_of.size_constraint)
            tag_str = f" `lb:{lb},ub:{ub}`"
        
        f.write(f"type {go_name} struct {{\n")
        f.write(f"\tValue []{elem_go_type}{tag_str}\n")
        f.write("}\n")
    
    def _generate_enumerated(self, f, enum: ASN1Enumerated):
        """Generate Go type for ENUMERATED with constants"""
        go_name = self._to_go_name(enum.name)
        
        f.write(f"// {enum.name} ::= ENUMERATED\n")
        if enum.extensible:
            f.write("// Extensible\n")
        
        f.write(f"type {go_name} struct {{\n")
        f.write("\tValue utils.ENUMERATED\n")
        f.write("}\n\n")
        
        f.write("const (\n")
        f.write(f"\t{go_name}EnumeratedNothing = iota\n")
        for value in enum.values:
            # Convert hyphens to underscores in enum values to avoid conflicts
            # e.g., v-10 -> V_10, but keep other naming rules
            value_go = self._enum_value_to_go_name(value)
            const_name = f"{go_name}Enumerated{value_go}"
            f.write(f"\t{const_name}\n")
        f.write(")\n")
    
    def _generate_integer(self, f, integer: ASN1Integer):
        """Generate Go type for INTEGER"""
        f.write(f"// {integer.name} ::= INTEGER")
        if integer.range_constraint:
            f.write(f" ({integer.range_constraint})")
        f.write("\n")
        
        # Build constraint tag
        tag_str = ""
        if integer.range_constraint:
            lb, ub = self._resolve_constraint_bounds(integer.range_constraint)
            # For INTEGER type definitions, use lb:0 if lb is numeric
            if isinstance(lb, int):
                tag_str = f" `lb:0,ub:{ub}`"
            else:
                # Keep constant names
                tag_str = f" `lb:{lb},ub:{ub}`"
        
        f.write(f"type {self._to_go_name(integer.name)} struct {{\n")
        f.write(f"\tValue utils.INTEGER{tag_str}\n")
        f.write("}\n")
    
    def _generate_bit_string(self, f, bit_str: ASN1BitString):
        """Generate Go type for BIT STRING"""
        f.write(f"// {bit_str.name} ::= BIT STRING")
        if bit_str.size_constraint:
            f.write(f" (SIZE ({bit_str.size_constraint}))")
        f.write("\n")
        
        # Build constraint tag
        tag_str = ""
        if bit_str.size_constraint:
            lb, ub = self._resolve_constraint_bounds(bit_str.size_constraint)
            tag_str = f" `lb:{lb},ub:{ub}`"
        
        f.write(f"type {self._to_go_name(bit_str.name)} struct {{\n")
        f.write(f"\tValue utils.BITSTRING{tag_str}\n")
        f.write("}\n")
    
    def _generate_octet_string(self, f, octet_str: ASN1OctetString):
        """Generate Go type for OCTET STRING"""
        f.write(f"// {octet_str.name} ::= OCTET STRING")
        if octet_str.size_constraint:
            f.write(f" (SIZE ({octet_str.size_constraint}))")
        if octet_str.containing:
            f.write(f" (CONTAINING {octet_str.containing})")
        f.write("\n")
        
        # Build constraint tag
        tag_str = ""
        if octet_str.size_constraint:
            lb, ub = self._resolve_constraint_bounds(octet_str.size_constraint)
            tag_str = f" `lb:{lb},ub:{ub}`"
        
        f.write(f"type {self._to_go_name(octet_str.name)} struct {{\n")
        f.write(f"\tValue utils.OCTETSTRING{tag_str}\n")
        f.write("}\n")
    
    def _generate_type_alias(self, f, alias: ASN1TypeAlias):
        """Generate Go type alias"""
        f.write(f"// {alias.name} ::= {alias.target_type}\n")
        
        if alias.target_type == 'BOOLEAN':
            # BOOLEAN should be wrapped in struct like INTEGER, BITSTRING, etc.
            f.write(f"type {self._to_go_name(alias.name)} struct {{\n")
            f.write("\tValue utils.BOOLEAN\n")
            f.write("}\n")
        elif alias.target_type == 'NULL':
            f.write(f"type {self._to_go_name(alias.name)} struct{{}}\n")
        else:
            go_type = self._to_go_type(alias.target_type)
            f.write(f"type {self._to_go_name(alias.name)} {go_type}\n")
    
    def _generate_field(self, f, field: ASN1Field, parent_name: str = ''):
        """Generate a struct field with constraint tags"""
        go_field_name = self._to_go_name(field.name)
        go_type = self._field_type_to_go(field, parent_name)
        
        # For inline SEQUENCE OF, if optional, make it pointer to slice
        is_inline_seq_of = field.type_name.startswith('SEQUENCE_OF_')
        
        if field.optional:
            if is_inline_seq_of:
                # For inline SEQUENCE OF, make it pointer to slice: *[]Type
                go_type = f"*{go_type}"
            else:
                go_type = f"*{go_type}"
        
        # Build constraint tag
        tag_str = ""
        if field.size_constraint or field.range_constraint:
            constraint = field.size_constraint or field.range_constraint
            lb, ub = self._resolve_constraint_bounds(constraint)
            # Special case: For INTEGER range constraints, use lb:0
            # For SIZE constraints (including inline SEQUENCE OF), use actual lower bound
            if field.range_constraint and field.type_name == 'INTEGER':
                tag_str = f" `lb:0,ub:{ub}`"
            else:
                tag_str = f" `lb:{lb},ub:{ub}`"
        
        comment_str = ""
        if field.comment:
            comment_str = f" // {field.comment}"
        
        f.write(f"\t{go_field_name} {go_type}{tag_str}{comment_str}\n")
    
    def _field_type_to_go(self, field: ASN1Field, parent_name: str = '') -> str:
        """Convert field type to Go type"""
        type_name = field.type_name
        
        if type_name == 'INTEGER':
            return 'utils.INTEGER'
        elif type_name == 'BOOLEAN':
            return 'utils.BOOLEAN'
        elif type_name == 'ENUMERATED':
            # Check if this is an inline ENUMERATED (check if the inline type was registered)
            if parent_name:
                inline_type_name = f"{parent_name}-{field.name}"
                if inline_type_name in self.types:
                    # This is an inline ENUMERATED, use the generated type
                    return self._to_go_name(inline_type_name)
                # If type is ENUMERATED but no type was registered, it's ENUMERATED {true}
                # which should be represented as bool
                elif field.type_name == 'ENUMERATED':
                    # Only return bool if this field was explicitly marked as inline ENUMERATED
                    return 'bool'
            return 'utils.ENUMERATED'
        elif type_name == 'BIT STRING':
            return 'utils.BITSTRING'
        elif type_name == 'OCTET STRING':
            return 'utils.OCTETSTRING'
        elif type_name == 'NULL':
            return 'struct{}'
        elif type_name.startswith('SEQUENCE_OF_'):
            # Inline SEQUENCE OF - extract element type and return as slice
            element_type = type_name[len('SEQUENCE_OF_'):]
            if element_type == 'INLINE_CHOICE':
                # Special case: inline CHOICE within SEQUENCE OF
                # The type name will be ParentName-FieldName-Item
                if parent_name:
                    return f'[]{self._to_go_name(f"{parent_name}-{field.name}-Item")}'
                return '[]Choice'  # Fallback
            # Handle primitive types
            elif element_type == 'INTEGER':
                return '[]utils.INTEGER'
            elif element_type == 'BOOLEAN':
                return '[]utils.BOOLEAN'
            elif element_type == 'BIT STRING':
                return '[]utils.BITSTRING'
            elif element_type == 'OCTET STRING':
                return '[]utils.OCTETSTRING'
            else:
                return f'[]{self._to_go_name(element_type)}'
        elif type_name == 'SEQUENCE' or type_name == 'CHOICE':
            # For inline types, construct the type name from parent and field name
            if parent_name:
                return self._to_go_name(f"{parent_name}-{field.name}")
            return 'interface{}'
        else:
            return self._to_go_name(type_name)
    
    def _choice_alt_type_to_go(self, alt: ASN1Field, parent_name: str) -> str:
        """Convert CHOICE alternative type to Go type"""
        type_name = alt.type_name
        
        if type_name == 'NULL':
            return 'struct{}'
        elif type_name in ['INTEGER', 'ENUMERATED', 'BIT STRING', 'OCTET STRING', 'BOOLEAN']:
            return self._field_type_to_go(alt)
        elif type_name.startswith('SEQUENCE_OF_'):
            # Inline SEQUENCE OF in CHOICE alternative
            element_type = type_name[len('SEQUENCE_OF_'):]
            # Handle primitive types
            if element_type == 'INTEGER':
                return '[]utils.INTEGER'
            elif element_type == 'BOOLEAN':
                return '[]utils.BOOLEAN'
            elif element_type == 'BIT STRING':
                return '[]utils.BITSTRING'
            elif element_type == 'OCTET STRING':
                return '[]utils.OCTETSTRING'
            else:
                return f'[]{self._to_go_name(element_type)}'
        elif type_name == 'SEQUENCE':
            return f"{parent_name}{self._to_go_name(alt.name)}"
        elif type_name == 'CHOICE':
            return f"{parent_name}{self._to_go_name(alt.name)}"
        else:
            return self._to_go_name(type_name)
    
    def _to_go_name(self, asn1_name: str, capitalize_first: bool = True) -> str:
        """Convert ASN.1 name to Go name - removes IEs suffix and converts to CamelCase"""
        # Remove "IEs" suffix
        asn1_name = re.sub(r'-IEs$', '', asn1_name)
        
        parts = asn1_name.replace('-', '_').split('_')
        
        if capitalize_first:
            return ''.join(word.capitalize() for word in parts)
        else:
            return ''.join(word.capitalize() if i > 0 else word for i, word in enumerate(parts))
    
    def _enum_value_to_go_name(self, value: str) -> str:
        """Convert ENUMERATED value to Go constant name
        Special handling: Convert '-' to '_' to avoid naming conflicts
        e.g., v-10 -> V_10, v0 -> V0
        """
        # Replace hyphens with underscores first
        value = value.replace('-', '_')
        # Then capitalize first letter of each part, but keep underscores
        parts = value.split('_')
        result = []
        for part in parts:
            if part:
                result.append(part.capitalize())
        return '_'.join(result)
    
    def _to_go_type(self, asn1_type: str) -> str:
        """Convert ASN.1 type to Go type"""
        type_map = {
            'BOOLEAN': 'bool',
            'NULL': 'struct{}',
            'INTEGER': 'int',
            'BIT STRING': '[]byte',
            'OCTET STRING': '[]byte',
            'BITSTRING': '[]byte',  # Single-word version
            'OCTETSTRING': '[]byte',  # Single-word version
        }
        
        if asn1_type in type_map:
            return type_map[asn1_type]
        
        return self._to_go_name(asn1_type)
    
    def _resolve_constraint_bounds(self, constraint: Optional[str]) -> Tuple[any, any]:
        """Resolve constraint bounds, handling constant references
        Returns tuple of (lower_bound, upper_bound) which can be int or string (for constants)
        """
        if not constraint:
            return (0, 0)
        
        constraint = constraint.strip()
        
        # Match range: number..number or number..constant or constant..constant
        match = re.match(r'(-?\d+|[a-z][a-zA-Z0-9_-]*)\s*\.\.\s*(-?\d+|[a-z][a-zA-Z0-9_-]*)', constraint)
        if match:
            lb_str = match.group(1).strip()
            ub_str = match.group(2).strip()
            
            # Parse lower bound
            if lb_str.lstrip('-').isdigit():
                lb = int(lb_str)
            else:
                # Keep constant name for tag (convert hyphens to camelCase but preserve original case)
                lb = self._const_to_go_name(lb_str)
            
            # Parse upper bound
            if ub_str.lstrip('-').isdigit():
                ub = int(ub_str)
            else:
                # Keep constant name for tag
                ub = self._const_to_go_name(ub_str)
            
            return (lb, ub)
        
        # Single value
        if constraint.lstrip('-').isdigit():
            val = int(constraint)
            return (val, val)
        else:
            # Constant reference
            const_name = self._const_to_go_name(constraint)
            return (const_name, const_name)
    
    def _const_to_go_name(self, const_name: str) -> str:
        """Convert constant name to Go name, preserving lowercase prefix"""
        # For constants like maxDRB, maxFreqANR-NB-r16, etc.
        # Convert hyphens to camelCase but keep the original case pattern
        parts = const_name.split('-')
        if len(parts) == 1:
            # No hyphens, return as-is
            return const_name
        
        # First part keeps its case, rest are capitalized
        result = parts[0]
        for part in parts[1:]:
            result += part.capitalize()
        return result
    
    def _resolve_constant(self, const_name: str) -> int:
        """Resolve a constant name to its integer value"""
        if const_name in self.constants:
            return self.constants[const_name].value
        
        const_name_underscore = const_name.replace('-', '_')
        if const_name_underscore in self.constants:
            return self.constants[const_name_underscore].value
        
        return 0


def validate_and_fix_generated_files(output_dir: str):
    """Validate and fix common issues in generated Go files"""
    print(f"\nValidating generated files in {output_dir}/...")
    
    issues_found = 0
    files_fixed = 0
    
    # Patterns to check and fix
    # Note: Avoid matching if already prefixed with utils.
    incorrect_patterns = [
        (r'(?<!utils\.)(?<!\.)Integer\b', 'utils.INTEGER', 'Integer -> utils.INTEGER'),
        (r'(?<!utils\.)(?<!\.)Boolean\b', 'utils.BOOLEAN', 'Boolean -> utils.BOOLEAN'),
        (r'(?<!utils\.)(?<!\.)Bitstring\b', 'utils.BITSTRING', 'Bitstring -> utils.BITSTRING'),
        (r'(?<!utils\.)(?<!\.)Octetstring\b', 'utils.OCTETSTRING', 'Octetstring -> utils.OCTETSTRING'),
        (r'(?<!utils\.)(?<!\.)Enumerated\b', 'utils.ENUMERATED', 'Enumerated -> utils.ENUMERATED'),
        # Also check for bare types that should have utils. prefix (but not in type definitions)
        (r'(?<!utils\.)\bINTEGER\s+(?!struct)(?!\w)', 'utils.INTEGER ', 'bare INTEGER -> utils.INTEGER'),
        (r'(?<!utils\.)\bBOOLEAN\s+(?!struct)(?!\w)', 'utils.BOOLEAN ', 'bare BOOLEAN -> utils.BOOLEAN'),
        (r'(?<!utils\.)\bBITSTRING\s+(?!struct)(?!\w)', 'utils.BITSTRING ', 'bare BITSTRING -> utils.BITSTRING'),
        (r'(?<!utils\.)\bOCTETSTRING\s+(?!struct)(?!\w)', 'utils.OCTETSTRING ', 'bare OCTETSTRING -> utils.OCTETSTRING'),
        (r'(?<!utils\.)\bENUMERATED\s+(?!struct)(?!\w)', 'utils.ENUMERATED ', 'bare ENUMERATED -> utils.ENUMERATED'),
    ]
    
    go_files = [f for f in os.listdir(output_dir) if f.endswith('.go')]
    
    for filename in go_files:
        filepath = os.path.join(output_dir, filename)
        
        try:
            with open(filepath, 'r', encoding='utf-8') as f:
                content = f.read()
            
            original_content = content
            file_issues = []
            
            # Check and fix each pattern
            for pattern, replacement, description in incorrect_patterns:
                matches = re.findall(pattern, content)
                if matches:
                    content = re.sub(pattern, replacement, content)
                    file_issues.append(f"{description} ({len(matches)} occurrences)")
            
            # If changes were made, write back
            if content != original_content:
                with open(filepath, 'w', encoding='utf-8') as f:
                    f.write(content)
                files_fixed += 1
                issues_found += len(file_issues)
                print(f"  Fixed {filename}: {', '.join(file_issues)}")
        
        except Exception as e:
            print(f"  Error processing {filename}: {e}")
    
    if files_fixed > 0:
        print(f"\n✓ Fixed {issues_found} issues in {files_fixed} files")
    else:
        print(f"\n✓ No issues found in {len(go_files)} files")
    
    return files_fixed


def main():
    """Main function"""
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
    
    # Validate and fix generated files
    files_fixed = validate_and_fix_generated_files(output_dir)
    
    # Format all Go files (after validation/fixes)
    print(f"\nFormatting Go files in {output_dir}/...")
    os.system("make format")
    
    if files_fixed > 0:
        print("\n⚠ Warning: Some files had incorrect type names and were fixed.")
        print("  Please review the changes and consider updating gen.py to prevent these issues.")

if __name__ == '__main__':
    main()