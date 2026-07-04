<!-- Parent: ../AGENTS.md -->
<!-- Generated: 2026-06-29 | Updated: 2026-06-29 -->

# ies

## Purpose
Target package for generated Go code representing 3GPP NR-RRC Information Elements (IEs). Each IE maps to a Go struct with PER encode/decode methods produced from the ASN.1 definitions in `../docs/`. Currently a placeholder awaiting the code generator.

## Key Files
| File | Description |
|------|-------------|
| `common.go` | Package declaration (`package ies`) — placeholder until generated IE structs are added |

## Subdirectories
None.

## For AI Agents

### Working In This Directory
- This package is **generated** — do not hand-write IE structs; they will be overwritten by the code generator
- Generated types should implement the `IE`-style encode/decode interface pattern (see `../test/c.go`) using `github.com/lvdund/asn1go/per`
- Size/range constants referenced in constraints live in `../common/const.go` (e.g. `MaxCellIntra`) — reuse them rather than redefining
- Follow the release/setup choice pattern from `../common/` for `release`/`setup` CHOICE fields

### Testing Requirements
- Generated structs must round-trip through PER encode then decode
- Validate against known RRC message captures when available
- Ensure all `Max*` constraint references resolve to constants in `../common/const.go`

### Common Patterns
- Each IE: one struct + `Encode(*per.Encoder) error` + `Decode(*per.Decoder) error`
- OPTIONAL fields tracked via presence flags / pointers
- Extension groups encoded as open types with internal OPTIONAL bitmaps
- CHOICE types use `per.ChoiceConstraints`

## Dependencies

### Internal
- `../common/` - Shared constants and the `Release_Setup_Choice` pattern
- `../asn1go/per` (via module `github.com/lvdund/asn1go`) - PER encode/decode

### External
- Go standard library

<!-- MANUAL: Custom project notes can be added below -->
