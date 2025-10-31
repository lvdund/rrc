package ies

import "rrc/utils"

// TDD-PUSCH-UpPTS-r14-setup-symPUSCH-UpPTS-r14 ::= ENUMERATED
type TddPuschUpptsR14SetupSympuschUpptsR14 struct {
	Value utils.ENUMERATED
}

const (
	TddPuschUpptsR14SetupSympuschUpptsR14EnumeratedNothing = iota
	TddPuschUpptsR14SetupSympuschUpptsR14EnumeratedSym1
	TddPuschUpptsR14SetupSympuschUpptsR14EnumeratedSym2
	TddPuschUpptsR14SetupSympuschUpptsR14EnumeratedSym3
	TddPuschUpptsR14SetupSympuschUpptsR14EnumeratedSym4
	TddPuschUpptsR14SetupSympuschUpptsR14EnumeratedSym5
	TddPuschUpptsR14SetupSympuschUpptsR14EnumeratedSym6
)
