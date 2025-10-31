package ies

import "rrc/utils"

// SearchSpace-searchSpaceType-ue-Specific-dci-Formats ::= ENUMERATED
type SearchspaceSearchspacetypeUeSpecificDciFormats struct {
	Value utils.ENUMERATED
}

const (
	SearchspaceSearchspacetypeUeSpecificDciFormatsEnumeratedNothing = iota
	SearchspaceSearchspacetypeUeSpecificDciFormatsEnumeratedFormats0_0_And_1_0
	SearchspaceSearchspacetypeUeSpecificDciFormatsEnumeratedFormats0_1_And_1_1
)
