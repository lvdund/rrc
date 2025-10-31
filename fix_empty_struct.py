#!/usr/bin/env python3
"""
Replace empty struct type definitions with direct struct{} usage
1. Parse note file to get all empty struct types
2. Find all usages of these types in Go files
3. Replace type references with struct{}
4. Delete the type definition files
"""

import re
import os
from pathlib import Path
from typing import Set, Dict, List

# Directory containing Go files
IES_DIR = Path('ies')

def parse_note_file(note_file: str) -> Dict[str, str]:
    """
    Parse note file to extract empty struct types and their file paths
    Returns: Dict mapping type_name -> file_path
    """
    empty_structs = {}
    
    with open(note_file, 'r') as f:
        lines = f.readlines()
    
    i = 0
    while i < len(lines):
        line = lines[i].strip()
        if line.startswith('/') and line.endswith('.go'):
            filepath = line
            i += 1
            if i < len(lines):
                type_line = lines[i].strip()
                # Match: "  4,1: type TypeName struct{}"
                match = re.match(r'\s*\d+,\d+:\s*type\s+([A-Za-z0-9_]+)\s+struct\{\}', type_line)
                if match:
                    type_name = match.group(1)
                    empty_structs[type_name] = filepath
        i += 1
    
    return empty_structs

def find_usages_in_file(filepath: Path, type_name: str) -> List[tuple]:
    """
    Find all usages of type_name in a Go file
    Returns: List of (line_number, line_content, needs_replacement) tuples
    """
    usages = []
    
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            lines = f.readlines()
    except Exception as e:
        print(f"  Warning: Could not read {filepath}: {e}")
        return usages
    
    for i, line in enumerate(lines, 1):
        # Check if this line uses the type
        # Pattern: *TypeName or []TypeName or TypeName (without struct{}, type, or import)
        # But not in type definition: "type TypeName struct{}"
        if re.search(r'\b' + re.escape(type_name) + r'\b', line):
            # Skip if it's the type definition itself
            if re.match(r'^\s*type\s+' + re.escape(type_name) + r'\s+struct\{\}', line):
                continue
            # Skip if it's in a comment
            if '//' in line and line.index('//') < line.find(type_name):
                comment_start = line.index('//')
                if type_name not in line[:comment_start]:
                    continue
            usages.append((i, line, True))
    
    return usages

def replace_type_in_content(content: str, type_name: str) -> str:
    """
    Replace all usages of type_name with struct{} in content
    Handles: *TypeName, []TypeName, TypeName
    """
    # Skip if this is the type definition file itself
    if re.search(r'^\s*type\s+' + re.escape(type_name) + r'\s+struct\{\}', content, re.MULTILINE):
        return content
    
    # Pattern 1: *TypeName (pointer to type) - most common case
    content = re.sub(r'\*\b' + re.escape(type_name) + r'\b', '*struct{}', content)
    
    # Pattern 2: []TypeName (slice of type)
    content = re.sub(r'\[\]\b' + re.escape(type_name) + r'\b', '[]struct{}', content)
    
    # Pattern 3: Direct TypeName in field declarations (not pointer, not slice)
    # Match: whitespace + field_name + whitespace + TypeName + optional tags
    # But avoid matching if it's part of a type definition or import
    def replace_field_type(match):
        full_match = match.group(0)
        before = content[:match.start()]
        
        # Skip if it's in a type definition
        if re.search(r'type\s+' + re.escape(type_name) + r'\s+', before):
            return full_match
        
        # Skip if it's in an import statement
        # Check if we're inside an import block
        import_match = re.search(r'import\s*(?:\()?', before)
        if import_match:
            # Check if there's a closing parenthesis after the import
            after_import = before[import_match.end():]
            if ')' not in after_import[:200]:  # Simple check for import block
                return full_match
        
        # Replace the type name
        return full_match.replace(type_name, 'struct{}')
    
    # Match field declarations: tabs/spaces + identifier + tabs/spaces + TypeName + optional tags/newline
    content = re.sub(
        r'(\t+[A-Za-z0-9_]+\s+)\b' + re.escape(type_name) + r'\b(\s*(?:`[^`]*`)?\s*$)',
        lambda m: m.group(1) + 'struct{}' + m.group(2),
        content,
        flags=re.MULTILINE
    )
    
    return content

def process_all_files(empty_structs: Dict[str, str]):
    """
    Process all Go files to replace empty struct types with struct{}
    """
    # Get all Go files in ies directory
    go_files = list(IES_DIR.glob('*.go'))
    
    files_to_delete = set()
    files_modified = set()
    
    print(f"Found {len(empty_structs)} empty struct types to replace")
    print(f"Scanning {len(go_files)} Go files...\n")
    
    for type_name, definition_file in empty_structs.items():
        print(f"Processing {type_name}...")
        
        # Mark definition file for deletion
        def_file_path = Path(definition_file)
        if def_file_path.exists():
            files_to_delete.add(def_file_path)
        else:
            print(f"  Warning: Definition file not found: {definition_file}")
        
        # Find and replace in all Go files
        for go_file in go_files:
            # Skip the definition file itself
            if go_file == def_file_path:
                continue
            
            try:
                with open(go_file, 'r', encoding='utf-8') as f:
                    original_content = f.read()
                
                # Check if file contains this type
                if type_name in original_content:
                    # Replace the type
                    new_content = replace_type_in_content(original_content, type_name)
                    
                    if new_content != original_content:
                        # Write back
                        with open(go_file, 'w', encoding='utf-8') as f:
                            f.write(new_content)
                        files_modified.add(go_file)
                        print(f"  Updated: {go_file.name}")
            
            except Exception as e:
                print(f"  Error processing {go_file}: {e}")
    
    # Delete definition files
    print(f"\nDeleting {len(files_to_delete)} empty struct definition files...")
    for file_path in files_to_delete:
        try:
            file_path.unlink()
            print(f"  Deleted: {file_path.name}")
        except Exception as e:
            print(f"  Error deleting {file_path}: {e}")
    
    print(f"\n✓ Done!")
    print(f"  Files modified: {len(files_modified)}")
    print(f"  Files deleted: {len(files_to_delete)}")

def main():
    note_file = 'note'
    
    if not os.path.exists(note_file):
        print(f"Error: {note_file} not found")
        return
    
    if not IES_DIR.exists():
        print(f"Error: {IES_DIR} directory not found")
        return
    
    print("Parsing note file...")
    empty_structs = parse_note_file(note_file)
    
    if not empty_structs:
        print("No empty struct types found in note file")
        return
    
    print(f"Found {len(empty_structs)} empty struct types")
    print("\nEmpty struct types:")
    for type_name, filepath in sorted(empty_structs.items()):
        print(f"  {type_name} -> {filepath}")
    
    print("\nProceeding with replacement...")
    process_all_files(empty_structs)

if __name__ == '__main__':
    main()

