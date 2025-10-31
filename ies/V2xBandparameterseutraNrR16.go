package ies

// V2X-BandParametersEUTRA-NR-r16 ::= CHOICE
const (
	V2xBandparameterseutraNrR16ChoiceNothing = iota
	V2xBandparameterseutraNrR16ChoiceEutra
	V2xBandparameterseutraNrR16ChoiceNr
)

type V2xBandparameterseutraNrR16 struct {
	Choice uint64
	Eutra  *V2xBandparameterseutraNrR16Eutra
	Nr     *V2xBandparameterseutraNrR16Nr
}
