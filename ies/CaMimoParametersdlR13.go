package ies

import "rrc/utils"

// CA-MIMO-ParametersDL-r13 ::= SEQUENCE
type CaMimoParametersdlR13 struct {
	CaBandwidthclassdlR13            CaBandwidthclassR10
	SupportedmimoCapabilitydlR13     *MimoCapabilitydlR10
	Fourlayertm3Tm4R13               *utils.ENUMERATED
	IntrabandcontiguousccInfolistR13 interface{}
}
