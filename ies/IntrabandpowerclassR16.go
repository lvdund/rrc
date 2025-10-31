package ies

import "rrc/utils"

// IntraBandPowerClass-r16 ::= ENUMERATED
type IntrabandpowerclassR16 struct {
	Value utils.ENUMERATED
}

const (
	IntrabandpowerclassR16EnumeratedNothing = iota
	IntrabandpowerclassR16EnumeratedPc2
	IntrabandpowerclassR16EnumeratedPc3
	IntrabandpowerclassR16EnumeratedSpare6
	IntrabandpowerclassR16EnumeratedSpare5
	IntrabandpowerclassR16EnumeratedSpare4
	IntrabandpowerclassR16EnumeratedSpare3
	IntrabandpowerclassR16EnumeratedSpare2
	IntrabandpowerclassR16EnumeratedSpare1
)
