package ies

import "rrc/utils"

// NPRACH-ConfigSIB-NB-r13-nprach-CP-Length-r13 ::= ENUMERATED
type NprachConfigsibNbR13NprachCpLengthR13 struct {
	Value utils.ENUMERATED
}

const (
	NprachConfigsibNbR13NprachCpLengthR13EnumeratedNothing = iota
	NprachConfigsibNbR13NprachCpLengthR13EnumeratedUs66dot7
	NprachConfigsibNbR13NprachCpLengthR13EnumeratedUs266dot7
)
