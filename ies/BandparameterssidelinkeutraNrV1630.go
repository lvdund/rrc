package ies

// BandParametersSidelinkEUTRA-NR-v1630 ::= CHOICE
const (
	BandparameterssidelinkeutraNrV1630ChoiceNothing = iota
	BandparameterssidelinkeutraNrV1630ChoiceEutra
	BandparameterssidelinkeutraNrV1630ChoiceNr
)

type BandparameterssidelinkeutraNrV1630 struct {
	Choice uint64
	Eutra  *struct{}
	Nr     *BandparameterssidelinkeutraNrV1630Nr
}
