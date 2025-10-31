package ies

import "rrc/utils"

// SearchSpace-searchSpaceType-ue-Specific-dci-FormatsExt-r16 ::= ENUMERATED
type SearchspaceSearchspacetypeUeSpecificDciFormatsextR16 struct {
	Value utils.ENUMERATED
}

const (
	SearchspaceSearchspacetypeUeSpecificDciFormatsextR16EnumeratedNothing = iota
	SearchspaceSearchspacetypeUeSpecificDciFormatsextR16EnumeratedFormats0_2_And_1_2
	SearchspaceSearchspacetypeUeSpecificDciFormatsextR16EnumeratedFormats0_1_And_1_1and_0_2_And_1_2
)
