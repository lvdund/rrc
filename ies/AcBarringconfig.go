package ies

import "rrc/utils"

// AC-BarringConfig ::= SEQUENCE
type AcBarringconfig struct {
	AcBarringfactor       AcBarringconfigAcBarringfactor
	AcBarringtime         AcBarringconfigAcBarringtime
	AcBarringforspecialac utils.BITSTRING `lb:5,ub:5`
}
