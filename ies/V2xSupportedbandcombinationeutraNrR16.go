package ies

// V2X-SupportedBandCombinationEUTRA-NR-r16 ::= SEQUENCE OF V2X-BandParametersEUTRA-NR-r16
// SIZE (1..maxBandCombSidelinkNR-r16)
type V2xSupportedbandcombinationeutraNrR16 struct {
	Value []V2xBandparameterseutraNrR16 `lb:1,ub:maxBandCombSidelinkNRR16`
}
