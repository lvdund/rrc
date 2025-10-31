package ies

import "rrc/utils"

// IntraBandContiguousCC-Info-r12-fourLayerTM3-TM4-perCC-r12 ::= ENUMERATED
type IntrabandcontiguousccInfoR12Fourlayertm3Tm4PerccR12 struct {
	Value utils.ENUMERATED
}

const (
	IntrabandcontiguousccInfoR12Fourlayertm3Tm4PerccR12EnumeratedNothing = iota
	IntrabandcontiguousccInfoR12Fourlayertm3Tm4PerccR12EnumeratedSupported
)
