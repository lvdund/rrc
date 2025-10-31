package ies

import "rrc/utils"

// MAC-ParametersCommon-preEmptiveBSR-r16 ::= ENUMERATED
type MacParameterscommonPreemptivebsrR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonPreemptivebsrR16EnumeratedNothing = iota
	MacParameterscommonPreemptivebsrR16EnumeratedSupported
)
