package ies

import "rrc/utils"

// SL-PSCCH-Config-r16-sl-FreqResourcePSCCH-r16 ::= ENUMERATED
type SlPscchConfigR16SlFreqresourcepscchR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPscchConfigR16SlFreqresourcepscchR16EnumeratedNothing = iota
	SlPscchConfigR16SlFreqresourcepscchR16EnumeratedN10
	SlPscchConfigR16SlFreqresourcepscchR16EnumeratedN12
	SlPscchConfigR16SlFreqresourcepscchR16EnumeratedN15
	SlPscchConfigR16SlFreqresourcepscchR16EnumeratedN20
	SlPscchConfigR16SlFreqresourcepscchR16EnumeratedN25
)
