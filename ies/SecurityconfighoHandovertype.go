package ies

// SecurityConfigHO-handoverType ::= CHOICE
const (
	SecurityconfighoHandovertypeChoiceNothing = iota
	SecurityconfighoHandovertypeChoiceIntralte
	SecurityconfighoHandovertypeChoiceInterrat
)

type SecurityconfighoHandovertype struct {
	Choice   uint64
	Intralte *SecurityconfighoHandovertypeIntralte
	Interrat *SecurityconfighoHandovertypeInterrat
}
