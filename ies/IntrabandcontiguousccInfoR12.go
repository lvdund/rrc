package ies

import "rrc/utils"

// IntraBandContiguousCC-Info-r12 ::= SEQUENCE
type IntrabandcontiguousccInfoR12 struct {
	Fourlayertm3Tm4PerccR12      *utils.ENUMERATED
	SupportedmimoCapabilitydlR12 *MimoCapabilitydlR10
	SupportedcsiProcR12          *utils.ENUMERATED
}
