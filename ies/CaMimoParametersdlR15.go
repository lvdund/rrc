package ies

import "rrc/utils"

// CA-MIMO-ParametersDL-r15 ::= SEQUENCE
type CaMimoParametersdlR15 struct {
	SupportedmimoCapabilitydlR15     *MimoCapabilitydlR10
	Fourlayertm3Tm4R15               *utils.ENUMERATED
	IntrabandcontiguousccInfolistR15 *interface{}
}
