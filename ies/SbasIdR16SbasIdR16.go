package ies

import "rrc/utils"

// SBAS-ID-r16-sbas-id-r16 ::= utils.ENUMERATED // Extensible
type SbasIdR16SbasIdR16 struct {
	Value utils.ENUMERATED
}

const (
	SbasIdR16SbasIdR16EnumeratedNothing = iota
	SbasIdR16SbasIdR16EnumeratedWaas
	SbasIdR16SbasIdR16EnumeratedEgnos
	SbasIdR16SbasIdR16EnumeratedMsas
	SbasIdR16SbasIdR16EnumeratedGagan
)
