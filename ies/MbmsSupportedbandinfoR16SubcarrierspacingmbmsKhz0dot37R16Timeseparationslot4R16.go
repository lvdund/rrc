package ies

import "rrc/utils"

// MBMS-SupportedBandInfo-r16-subcarrierSpacingMBMS-khz0dot37-r16-timeSeparationSlot4-r16 ::= ENUMERATED
type MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz0dot37R16Timeseparationslot4R16 struct {
	Value utils.ENUMERATED
}

const (
	MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz0dot37R16Timeseparationslot4R16EnumeratedNothing = iota
	MbmsSupportedbandinfoR16SubcarrierspacingmbmsKhz0dot37R16Timeseparationslot4R16EnumeratedSupported
)
