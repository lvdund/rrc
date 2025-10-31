package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-gnss-Location-r16 ::= ENUMERATED
type UeBasedperfmeasParametersR16GnssLocationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16GnssLocationR16EnumeratedNothing = iota
	UeBasedperfmeasParametersR16GnssLocationR16EnumeratedSupported
)
