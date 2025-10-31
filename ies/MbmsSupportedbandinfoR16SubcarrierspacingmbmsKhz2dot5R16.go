package ies

import "rrc/utils"

// MBMS-SupportedBandInfo-r16-subcarrierSpacingMBMS-khz2dot5-r16 ::= ENUMERATED
type MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz2dot5R16 struct {
	Value utils.ENUMERATED
}

const (
	MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz2dot5R16EnumeratedNothing = iota
	MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz2dot5R16EnumeratedSupported
)
