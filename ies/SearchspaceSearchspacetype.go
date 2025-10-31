package ies

// SearchSpace-searchSpaceType ::= CHOICE
// Extensible
const (
	SearchspaceSearchspacetypeChoiceNothing = iota
	SearchspaceSearchspacetypeChoiceCommon
	SearchspaceSearchspacetypeChoiceUeSpecific
)

type SearchspaceSearchspacetype struct {
	Choice     uint64
	Common     *SearchspaceSearchspacetypeCommon
	UeSpecific *SearchspaceSearchspacetypeUeSpecific
}
