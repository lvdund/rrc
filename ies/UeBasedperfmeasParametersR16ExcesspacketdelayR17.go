package ies

import "rrc/utils"

// UE-BasedPerfMeas-Parameters-r16-excessPacketDelay-r17 ::= ENUMERATED
type UeBasedperfmeasParametersR16ExcesspacketdelayR17 struct {
	Value utils.ENUMERATED
}

const (
	UeBasedperfmeasParametersR16ExcesspacketdelayR17EnumeratedNothing = iota
	UeBasedperfmeasParametersR16ExcesspacketdelayR17EnumeratedSupported
)
