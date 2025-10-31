package ies

import "rrc/utils"

// SR-NPRACH-Resource-NB-r15-alpha-r15 ::= ENUMERATED
type SrNprachResourceNbR15AlphaR15 struct {
	Value utils.ENUMERATED
}

const (
	SrNprachResourceNbR15AlphaR15EnumeratedNothing = iota
	SrNprachResourceNbR15AlphaR15EnumeratedAl0
	SrNprachResourceNbR15AlphaR15EnumeratedAl04
	SrNprachResourceNbR15AlphaR15EnumeratedAl05
	SrNprachResourceNbR15AlphaR15EnumeratedAl06
	SrNprachResourceNbR15AlphaR15EnumeratedAl07
	SrNprachResourceNbR15AlphaR15EnumeratedAl08
	SrNprachResourceNbR15AlphaR15EnumeratedAl09
	SrNprachResourceNbR15AlphaR15EnumeratedAl1
)
