package ies

import "rrc/utils"

// RMTC-ConfigNR-r16-refSCS-CP-NR-r16 ::= ENUMERATED
type RmtcConfignrR16RefscsCpNrR16 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfignrR16RefscsCpNrR16EnumeratedNothing = iota
	RmtcConfignrR16RefscsCpNrR16EnumeratedKhz15
	RmtcConfignrR16RefscsCpNrR16EnumeratedKhz30
	RmtcConfignrR16RefscsCpNrR16EnumeratedKhz60_Ncp
	RmtcConfignrR16RefscsCpNrR16EnumeratedKhz60_Ecp
)
