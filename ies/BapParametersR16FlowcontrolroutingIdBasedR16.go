package ies

import "rrc/utils"

// BAP-Parameters-r16-flowControlRouting-ID-Based-r16 ::= ENUMERATED
type BapParametersR16FlowcontrolroutingIdBasedR16 struct {
	Value utils.ENUMERATED
}

const (
	BapParametersR16FlowcontrolroutingIdBasedR16EnumeratedNothing = iota
	BapParametersR16FlowcontrolroutingIdBasedR16EnumeratedSupported
)
