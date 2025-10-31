package ies

import "rrc/utils"

// RMTC-Config-r16-measDurationSymbols-v1700 ::= ENUMERATED
type RmtcConfigR16MeasdurationsymbolsV1700 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR16MeasdurationsymbolsV1700EnumeratedNothing = iota
	RmtcConfigR16MeasdurationsymbolsV1700EnumeratedSym140
	RmtcConfigR16MeasdurationsymbolsV1700EnumeratedSym560
	RmtcConfigR16MeasdurationsymbolsV1700EnumeratedSym1120
)
