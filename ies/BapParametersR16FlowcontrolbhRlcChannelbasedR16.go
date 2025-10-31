package ies

import "rrc/utils"

// BAP-Parameters-r16-flowControlBH-RLC-ChannelBased-r16 ::= ENUMERATED
type BapParametersR16FlowcontrolbhRlcChannelbasedR16 struct {
	Value utils.ENUMERATED
}

const (
	BapParametersR16FlowcontrolbhRlcChannelbasedR16EnumeratedNothing = iota
	BapParametersR16FlowcontrolbhRlcChannelbasedR16EnumeratedSupported
)
