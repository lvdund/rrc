package ies

import "rrc/utils"

// RMTC-ConfigNR-r16-measDurationNR-r16 ::= ENUMERATED
type RmtcConfignrR16MeasdurationnrR16 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfignrR16MeasdurationnrR16EnumeratedNothing = iota
	RmtcConfignrR16MeasdurationnrR16EnumeratedSym1
	RmtcConfignrR16MeasdurationnrR16EnumeratedSym14or12
	RmtcConfignrR16MeasdurationnrR16EnumeratedSym28or24
	RmtcConfignrR16MeasdurationnrR16EnumeratedSym42or36
	RmtcConfignrR16MeasdurationnrR16EnumeratedSym70or60
)
