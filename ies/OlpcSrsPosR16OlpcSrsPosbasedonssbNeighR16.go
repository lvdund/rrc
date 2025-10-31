package ies

import "rrc/utils"

// OLPC-SRS-Pos-r16-olpc-SRS-PosBasedOnSSB-Neigh-r16 ::= ENUMERATED
type OlpcSrsPosR16OlpcSrsPosbasedonssbNeighR16 struct {
	Value utils.ENUMERATED
}

const (
	OlpcSrsPosR16OlpcSrsPosbasedonssbNeighR16EnumeratedNothing = iota
	OlpcSrsPosR16OlpcSrsPosbasedonssbNeighR16EnumeratedSupported
)
