package ies

import "rrc/utils"

// OLPC-SRS-Pos-r16-olpc-SRS-PosBasedOnPRS-Neigh-r16 ::= ENUMERATED
type OlpcSrsPosR16OlpcSrsPosbasedonprsNeighR16 struct {
	Value utils.ENUMERATED
}

const (
	OlpcSrsPosR16OlpcSrsPosbasedonprsNeighR16EnumeratedNothing = iota
	OlpcSrsPosR16OlpcSrsPosbasedonprsNeighR16EnumeratedSupported
)
