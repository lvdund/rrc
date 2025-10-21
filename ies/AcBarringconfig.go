package ies

import "rrc/utils"

// AC-BarringConfig ::= SEQUENCE
type AcBarringconfig struct {
	AcBarringfactor       utils.ENUMERATED
	AcBarringtime         utils.ENUMERATED
	AcBarringforspecialac utils.BITSTRING
}
