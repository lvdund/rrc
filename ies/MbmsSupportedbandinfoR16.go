package ies

import "rrc/utils"

// MBMS-SupportedBandInfo-r16 ::= SEQUENCE
type MbmsSupportedbandinfoR16 struct {
	SubcarrierspacingmbmsKhz2dot5R16  *utils.ENUMERATED
	SubcarrierspacingmbmsKhz0dot37R16 *interface{}
}
