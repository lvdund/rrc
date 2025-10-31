package ies

import "rrc/utils"

// SL-PSCCH-Config-r16-sl-TimeResourcePSCCH-r16 ::= ENUMERATED
type SlPscchConfigR16SlTimeresourcepscchR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPscchConfigR16SlTimeresourcepscchR16EnumeratedNothing = iota
	SlPscchConfigR16SlTimeresourcepscchR16EnumeratedN2
	SlPscchConfigR16SlTimeresourcepscchR16EnumeratedN3
)
