package ies

import "rrc/utils"

// SupportedBandEUTRA-v1320-ue-PowerClass-N-r13 ::= ENUMERATED
type SupportedbandeutraV1320UePowerclassNR13 struct {
	Value utils.ENUMERATED
}

const (
	SupportedbandeutraV1320UePowerclassNR13EnumeratedNothing = iota
	SupportedbandeutraV1320UePowerclassNR13EnumeratedClass1
	SupportedbandeutraV1320UePowerclassNR13EnumeratedClass2
	SupportedbandeutraV1320UePowerclassNR13EnumeratedClass4
)
