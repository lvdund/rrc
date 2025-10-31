package ies

// IntraBandContiguousCC-Info-r12 ::= SEQUENCE
type IntrabandcontiguousccInfoR12 struct {
	Fourlayertm3Tm4PerccR12      *IntrabandcontiguousccInfoR12Fourlayertm3Tm4PerccR12
	SupportedmimoCapabilitydlR12 *MimoCapabilitydlR10
	SupportedcsiProcR12          *IntrabandcontiguousccInfoR12SupportedcsiProcR12
}
