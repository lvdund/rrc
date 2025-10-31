package ies

import "rrc/utils"

// SearchSpace-searchSpaceType-ue-Specific-dci-Formats-MT-r16 ::= ENUMERATED
type SearchspaceSearchspacetypeUeSpecificDciFormatsMtR16 struct {
	Value utils.ENUMERATED
}

const (
	SearchspaceSearchspacetypeUeSpecificDciFormatsMtR16EnumeratedNothing = iota
	SearchspaceSearchspacetypeUeSpecificDciFormatsMtR16EnumeratedFormats2_5
)
