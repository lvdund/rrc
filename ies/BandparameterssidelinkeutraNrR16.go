package ies

// BandParametersSidelinkEUTRA-NR-r16 ::= CHOICE
const (
	BandparameterssidelinkeutraNrR16ChoiceNothing = iota
	BandparameterssidelinkeutraNrR16ChoiceEutra
	BandparameterssidelinkeutraNrR16ChoiceNr
)

type BandparameterssidelinkeutraNrR16 struct {
	Choice uint64
	Eutra  *BandparameterssidelinkeutraNrR16Eutra
	Nr     *BandparameterssidelinkeutraNrR16Nr
}
