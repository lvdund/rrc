package ies

// CA-MIMO-ParametersDL-r15 ::= SEQUENCE
type CaMimoParametersdlR15 struct {
	SupportedmimoCapabilitydlR15     *MimoCapabilitydlR10
	Fourlayertm3Tm4R15               *CaMimoParametersdlR15Fourlayertm3Tm4R15
	IntrabandcontiguousccInfolistR15 *[]IntrabandcontiguousccInfoR12 `lb:1,ub:maxServCellR13`
}
