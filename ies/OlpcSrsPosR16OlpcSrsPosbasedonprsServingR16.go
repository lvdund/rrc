package ies

import "rrc/utils"

// OLPC-SRS-Pos-r16-olpc-SRS-PosBasedOnPRS-Serving-r16 ::= ENUMERATED
type OlpcSrsPosR16OlpcSrsPosbasedonprsServingR16 struct {
	Value utils.ENUMERATED
}

const (
	OlpcSrsPosR16OlpcSrsPosbasedonprsServingR16EnumeratedNothing = iota
	OlpcSrsPosR16OlpcSrsPosbasedonprsServingR16EnumeratedSupported
)
