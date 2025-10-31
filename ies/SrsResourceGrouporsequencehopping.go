package ies

import "rrc/utils"

// SRS-Resource-groupOrSequenceHopping ::= ENUMERATED
type SrsResourceGrouporsequencehopping struct {
	Value utils.ENUMERATED
}

const (
	SrsResourceGrouporsequencehoppingEnumeratedNothing = iota
	SrsResourceGrouporsequencehoppingEnumeratedNeither
	SrsResourceGrouporsequencehoppingEnumeratedGrouphopping
	SrsResourceGrouporsequencehoppingEnumeratedSequencehopping
)
