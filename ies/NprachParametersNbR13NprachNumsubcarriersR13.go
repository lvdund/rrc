package ies

import "rrc/utils"

// NPRACH-Parameters-NB-r13-nprach-NumSubcarriers-r13 ::= ENUMERATED
type NprachParametersNbR13NprachNumsubcarriersR13 struct {
	Value utils.ENUMERATED
}

const (
	NprachParametersNbR13NprachNumsubcarriersR13EnumeratedNothing = iota
	NprachParametersNbR13NprachNumsubcarriersR13EnumeratedN12
	NprachParametersNbR13NprachNumsubcarriersR13EnumeratedN24
	NprachParametersNbR13NprachNumsubcarriersR13EnumeratedN36
	NprachParametersNbR13NprachNumsubcarriersR13EnumeratedN48
)
