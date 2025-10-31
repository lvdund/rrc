package ies

import "rrc/utils"

// SL-PSFCH-Config-r16-sl-PSFCH-Period-r16 ::= ENUMERATED
type SlPsfchConfigR16SlPsfchPeriodR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPsfchConfigR16SlPsfchPeriodR16EnumeratedNothing = iota
	SlPsfchConfigR16SlPsfchPeriodR16EnumeratedSl0
	SlPsfchConfigR16SlPsfchPeriodR16EnumeratedSl1
	SlPsfchConfigR16SlPsfchPeriodR16EnumeratedSl2
	SlPsfchConfigR16SlPsfchPeriodR16EnumeratedSl4
)
