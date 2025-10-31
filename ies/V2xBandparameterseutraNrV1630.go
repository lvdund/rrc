package ies

// V2X-BandParametersEUTRA-NR-v1630 ::= CHOICE
const (
	V2xBandparameterseutraNrV1630ChoiceNothing = iota
	V2xBandparameterseutraNrV1630ChoiceEutra
	V2xBandparameterseutraNrV1630ChoiceNr
)

type V2xBandparameterseutraNrV1630 struct {
	Choice uint64
	Eutra  *struct{}
	Nr     *V2xBandparameterseutraNrV1630Nr
}
