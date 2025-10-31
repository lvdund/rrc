package ies

import "rrc/utils"

// TRS-ResourceSet-r17-nrofResources-r17 ::= ENUMERATED
type TrsResourcesetR17NrofresourcesR17 struct {
	Value utils.ENUMERATED
}

const (
	TrsResourcesetR17NrofresourcesR17EnumeratedNothing = iota
	TrsResourcesetR17NrofresourcesR17EnumeratedN2
	TrsResourcesetR17NrofresourcesR17EnumeratedN4
)
