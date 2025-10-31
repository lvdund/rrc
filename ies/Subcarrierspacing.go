package ies

import "rrc/utils"

// SubcarrierSpacing ::= ENUMERATED
type Subcarrierspacing struct {
	Value utils.ENUMERATED
}

const (
	SubcarrierspacingEnumeratedNothing = iota
	SubcarrierspacingEnumeratedKhz15
	SubcarrierspacingEnumeratedKhz30
	SubcarrierspacingEnumeratedKhz60
	SubcarrierspacingEnumeratedKhz120
	SubcarrierspacingEnumeratedKhz240
	SubcarrierspacingEnumeratedKhz480_V1700
	SubcarrierspacingEnumeratedKhz960_V1700
	SubcarrierspacingEnumeratedSpare1
)
