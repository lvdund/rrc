package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-logicalChannelSR-DelayTimer ::= ENUMERATED
type MacParametersxddDiffLogicalchannelsrDelaytimer struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffLogicalchannelsrDelaytimerEnumeratedNothing = iota
	MacParametersxddDiffLogicalchannelsrDelaytimerEnumeratedSupported
)
