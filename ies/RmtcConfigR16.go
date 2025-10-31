package ies

import "rrc/utils"

// RMTC-Config-r16 ::= SEQUENCE
// Extensible
type RmtcConfigR16 struct {
	RmtcPeriodicityR16       RmtcConfigR16RmtcPeriodicityR16
	RmtcSubframeoffsetR16    *utils.INTEGER `lb:0,ub:639`
	MeasdurationsymbolsR16   RmtcConfigR16MeasdurationsymbolsR16
	RmtcFrequencyR16         ArfcnValuenr
	RefScsCpR16              RmtcConfigR16RefScsCpR16
	RmtcBandwidthR17         *RmtcConfigR16RmtcBandwidthR17         `ext`
	MeasdurationsymbolsV1700 *RmtcConfigR16MeasdurationsymbolsV1700 `ext`
	RefScsCpV1700            *RmtcConfigR16RefScsCpV1700            `ext`
	TciStateinfoR17          *RmtcConfigR16TciStateinfoR17          `ext`
	RefBwpidR17              *BwpId                                 `ext`
}
