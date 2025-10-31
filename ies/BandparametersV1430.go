package ies

// BandParameters-v1430 ::= SEQUENCE
type BandparametersV1430 struct {
	BandparametersdlV1430           *MimoCaParametersperbobcV1430
	Ul256qamR14                     *BandparametersV1430Ul256qamR14
	Ul256qamPerccInfolistR14        *[]Ul256qamPerccInfoR14        `lb:2,ub:maxServCellR13`
	SrsCapabilityperbandpairlistR14 *[]SrsCapabilityperbandpairR14 `lb:1,ub:maxSimultaneousBandsR10`
}
