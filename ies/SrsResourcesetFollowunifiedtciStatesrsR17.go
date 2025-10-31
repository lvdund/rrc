package ies

import "rrc/utils"

// SRS-ResourceSet-followUnifiedTCI-StateSRS-r17 ::= ENUMERATED
type SrsResourcesetFollowunifiedtciStatesrsR17 struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcesetFollowunifiedtciStatesrsR17EnumeratedNothing = iota
	SrsResourcesetFollowunifiedtciStatesrsR17EnumeratedEnabled
)
