package ies

import "rrc/utils"

// RMTC-ConfigNR-r16 ::= SEQUENCE
// Extensible
type RmtcConfignrR16 struct {
	RmtcPeriodicitynrR16    RmtcConfignrR16RmtcPeriodicitynrR16
	RmtcSubframeoffsetnrR16 *utils.INTEGER `lb:0,ub:639`
	MeasdurationnrR16       RmtcConfignrR16MeasdurationnrR16
	RmtcFrequencynrR16      ArfcnValuenrR15
	RefscsCpNrR16           RmtcConfignrR16RefscsCpNrR16
}
