package ies

// SystemInformationBlockType1-v1530-IEs-cellAccessRelatedInfo-5GC-r15 ::= SEQUENCE
type Systeminformationblocktype1V1530IesCellaccessrelatedinfo5gcR15 struct {
	Cellbarred5gcR15                Systeminformationblocktype1V1530IesCellaccessrelatedinfo5gcR15Cellbarred5gcR15
	Cellbarred5gcCrsR15             Systeminformationblocktype1V1530IesCellaccessrelatedinfo5gcR15Cellbarred5gcCrsR15
	Cellaccessrelatedinfolist5gcR15 []Cellaccessrelatedinfo5gcR15 `lb:1,ub:maxPLMNR11`
}
