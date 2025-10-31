package ies

import "rrc/utils"

// SearchSpace-searchSpaceType-common-dci-Format2-3-dummy2 ::= ENUMERATED
type SearchspaceSearchspacetypeCommonDciFormat23Dummy2 struct {
	Value utils.ENUMERATED
}

const (
	SearchspaceSearchspacetypeCommonDciFormat23Dummy2EnumeratedNothing = iota
	SearchspaceSearchspacetypeCommonDciFormat23Dummy2EnumeratedN1
	SearchspaceSearchspacetypeCommonDciFormat23Dummy2EnumeratedN2
)
