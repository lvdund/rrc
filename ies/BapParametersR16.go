package ies

// BAP-Parameters-r16 ::= SEQUENCE
type BapParametersR16 struct {
	FlowcontrolbhRlcChannelbasedR16 *BapParametersR16FlowcontrolbhRlcChannelbasedR16
	FlowcontrolroutingIdBasedR16    *BapParametersR16FlowcontrolroutingIdBasedR16
}
