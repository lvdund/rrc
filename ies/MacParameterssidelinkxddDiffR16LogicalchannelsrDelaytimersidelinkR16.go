package ies

import "rrc/utils"

// MAC-ParametersSidelinkXDD-Diff-r16-logicalChannelSR-DelayTimerSidelink-r16 ::= ENUMERATED
type MacParameterssidelinkxddDiffR16LogicalchannelsrDelaytimersidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterssidelinkxddDiffR16LogicalchannelsrDelaytimersidelinkR16EnumeratedNothing = iota
	MacParameterssidelinkxddDiffR16LogicalchannelsrDelaytimersidelinkR16EnumeratedSupported
)
