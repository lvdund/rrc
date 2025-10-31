package ies

import "rrc/utils"

// MAC-ParametersCommon-lastTransmissionUL-r17 ::= ENUMERATED
type MacParameterscommonLasttransmissionulR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLasttransmissionulR17EnumeratedNothing = iota
	MacParameterscommonLasttransmissionulR17EnumeratedSupported
)
