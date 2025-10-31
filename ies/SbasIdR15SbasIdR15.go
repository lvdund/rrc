package ies

import "rrc/utils"

// SBAS-ID-r15-sbas-id-r15 ::= utils.ENUMERATED // Extensible
type SbasIdR15SbasIdR15 struct {
	Value utils.ENUMERATED
}

const (
	SbasIdR15SbasIdR15EnumeratedNothing = iota
	SbasIdR15SbasIdR15EnumeratedWaas
	SbasIdR15SbasIdR15EnumeratedEgnos
	SbasIdR15SbasIdR15EnumeratedMsas
	SbasIdR15SbasIdR15EnumeratedGagan
)
