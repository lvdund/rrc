package ies

import "rrc/utils"

// RMTC-Config-r16-rmtc-Bandwidth-r17 ::= ENUMERATED
type RmtcConfigR16RmtcBandwidthR17 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR16RmtcBandwidthR17EnumeratedNothing = iota
	RmtcConfigR16RmtcBandwidthR17EnumeratedMhz100
	RmtcConfigR16RmtcBandwidthR17EnumeratedMhz400
	RmtcConfigR16RmtcBandwidthR17EnumeratedMhz800
	RmtcConfigR16RmtcBandwidthR17EnumeratedMhz1600
	RmtcConfigR16RmtcBandwidthR17EnumeratedMhz2000
)
