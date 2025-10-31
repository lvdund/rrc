package ies

import "rrc/utils"

// RMTC-ConfigNR-r16-rmtc-PeriodicityNR-r16 ::= ENUMERATED
type RmtcConfignrR16RmtcPeriodicitynrR16 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfignrR16RmtcPeriodicitynrR16EnumeratedNothing = iota
	RmtcConfignrR16RmtcPeriodicitynrR16EnumeratedMs40
	RmtcConfignrR16RmtcPeriodicitynrR16EnumeratedMs80
	RmtcConfignrR16RmtcPeriodicitynrR16EnumeratedMs160
	RmtcConfignrR16RmtcPeriodicitynrR16EnumeratedMs320
	RmtcConfignrR16RmtcPeriodicitynrR16EnumeratedMs640
)
