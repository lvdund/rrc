package ies

// V2X-SupportedBandCombinationEUTRA-NR-v1630 ::= SEQUENCE OF V2X-BandCombinationParametersEUTRA-NR-v1630
// SIZE (1..maxBandCombSidelinkNR-r16)
type V2xSupportedbandcombinationeutraNrV1630 struct {
	Value []V2xBandcombinationparameterseutraNrV1630 `lb:1,ub:maxBandCombSidelinkNRR16`
}
