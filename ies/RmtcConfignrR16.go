package ies

import "rrc/utils"

// RMTC-ConfigNR-r16 ::= SEQUENCE
// Extensible
type RmtcConfignrR16 struct {
	RmtcPeriodicitynrR16    utils.ENUMERATED
	RmtcSubframeoffsetnrR16 *utils.INTEGER
	MeasdurationnrR16       utils.ENUMERATED
	RmtcFrequencynrR16      ArfcnValuenrR15
	RefscsCpNrR16           utils.ENUMERATED
}
