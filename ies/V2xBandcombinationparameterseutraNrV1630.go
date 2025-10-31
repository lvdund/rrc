package ies

// V2X-BandCombinationParametersEUTRA-NR-v1630 ::= SEQUENCE
type V2xBandcombinationparameterseutraNrV1630 struct {
	BandlistsidelinkeutraNrR16   []V2xBandparameterseutraNrR16   `lb:1,ub:maxSimultaneousBandsR10`
	BandlistsidelinkeutraNrV1630 []V2xBandparameterseutraNrV1630 `lb:1,ub:maxSimultaneousBandsR10`
}
