package ies

import "rrc/utils"

// MAC-Parameters-r12-logicalChannelSR-ProhibitTimer-r12 ::= ENUMERATED
type MacParametersR12LogicalchannelsrProhibittimerR12 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersR12LogicalchannelsrProhibittimerR12EnumeratedNothing = iota
	MacParametersR12LogicalchannelsrProhibittimerR12EnumeratedSupported
)
