package ies

import "rrc/utils"

// ControlResourceSet-followUnifiedTCI-State-r17 ::= ENUMERATED
type ControlresourcesetFollowunifiedtciStateR17 struct {
	Value utils.ENUMERATED
}

const (
	ControlresourcesetFollowunifiedtciStateR17EnumeratedNothing = iota
	ControlresourcesetFollowunifiedtciStateR17EnumeratedEnabled
)
