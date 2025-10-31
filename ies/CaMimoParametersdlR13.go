package ies

// CA-MIMO-ParametersDL-r13 ::= SEQUENCE
type CaMimoParametersdlR13 struct {
	CaBandwidthclassdlR13            CaBandwidthclassR10
	SupportedmimoCapabilitydlR13     *MimoCapabilitydlR10
	Fourlayertm3Tm4R13               *CaMimoParametersdlR13Fourlayertm3Tm4R13
	IntrabandcontiguousccInfolistR13 []IntrabandcontiguousccInfoR12 `lb:1,ub:maxServCellR13`
}
