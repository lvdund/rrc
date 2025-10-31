package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-ulPDCP-Delay-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16UlpdcpDelayR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16UlpdcpDelayR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16UlpdcpDelayR16EnumeratedSupported
)
