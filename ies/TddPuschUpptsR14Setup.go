package ies

// TDD-PUSCH-UpPTS-r14-setup ::= SEQUENCE
type TddPuschUpptsR14Setup struct {
	SympuschUpptsR14       *TddPuschUpptsR14SetupSympuschUpptsR14
	DmrsLessupptsConfigR14 *bool
}
