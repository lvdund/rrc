<!-- Parent: ../AGENTS.md -->
<!-- Generated: 2026-06-28 | Updated: 2026-06-29 -->

# common

## Purpose
Common utilities, types, and base structures used across the RRC code generation project. Provides generic encoders/decoders and shared types for handling release/setup choice patterns common in 3GPP specifications.

## Key Files
| File | Description |
|------|-------------|
| `const.go` | All `Max*` numeric constants extracted from the NR-RRC ASN.1 specification (size/range bounds used in constraints, e.g. `MaxCellIntra`, `MaxMultiBands`) |
| `release_setup.go` | Generic `Release_Setup_Choice` type for the release/setup (NULL/concrete) pattern pervasive in 3GPP messages |

## Subdirectories
None.

## For AI Agents

### Working In This Directory
- This package provides type-safe generic structures for ASN.1 encoding/decoding
- The Release_Setup_Choice pattern is used throughout 3GPP specifications for backward compatibility
- Uses generics (Go 1.18+) to provide type-safe encoding/decoding operations
- All types implement the GenericSetuprelease interface for PER operations

### Testing Requirements
- Run `go test .` from this directory
- Verify release/setup choice encoding matches 3GPP specification
- Test generic type instantiation with concrete types

### Common Patterns
- Generic type parameter T must implement GenericSetuprelease interface
- Choice values: 0 = release (NULL), 1 = setup (concrete type)
- Always check err after Encode/Decode operations
- Use per.ChoiceConstraints for choice type constraints

## Dependencies

### Internal
- None

### External
- `github.com/lvdund/asn1go/per` - PER encoding/decoding library

<!-- MANUAL: Custom project notes can be added below -->