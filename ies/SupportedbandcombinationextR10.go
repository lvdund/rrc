package ies

// SupportedBandCombinationExt-r10 ::= SEQUENCE OF BandCombinationParametersExt-r10
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationextR10 struct {
	Value []BandcombinationparametersextR10 `lb:1,ub:maxBandCombR10`
}
