package ies

import "rrc/utils"

// MBMS-SupportedBandInfo-r16-subcarrierSpacingMBMS-khz0dot37-r16-timeSeparationSlot2-r16 ::= ENUMERATED
type MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz0dot37R16Timeseparationslot2R16 struct {
	Value utils.ENUMERATED
}

const (
	MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz0dot37R16Timeseparationslot2R16EnumeratedNothing = iota
	MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz0dot37R16Timeseparationslot2R16EnumeratedSupported
)
