package ies

import "rrc/utils"

// RMTC-Config-r16-measDurationSymbols-r16 ::= ENUMERATED
type RmtcConfigR16MeasdurationsymbolsR16 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR16MeasdurationsymbolsR16EnumeratedNothing = iota
	RmtcConfigR16MeasdurationsymbolsR16EnumeratedSym1
	RmtcConfigR16MeasdurationsymbolsR16EnumeratedSym14or12
	RmtcConfigR16MeasdurationsymbolsR16EnumeratedSym28or24
	RmtcConfigR16MeasdurationsymbolsR16EnumeratedSym42or36
	RmtcConfigR16MeasdurationsymbolsR16EnumeratedSym70or60
)
