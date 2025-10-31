package ies

// BandParametersSidelinkEUTRA-NR-v1710 ::= CHOICE
const (
	BandparameterssidelinkeutraNrV1710ChoiceNothing = iota
	BandparameterssidelinkeutraNrV1710ChoiceEutra
	BandparameterssidelinkeutraNrV1710ChoiceNr
)

type BandparameterssidelinkeutraNrV1710 struct {
	Choice uint64
	Eutra  *struct{}
	Nr     *BandparameterssidelinkeutraNrV1710Nr
}
