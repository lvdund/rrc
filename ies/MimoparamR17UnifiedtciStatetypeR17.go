package ies

import "rrc/utils"

// MIMOParam-r17-unifiedTCI-StateType-r17 ::= ENUMERATED
type MimoparamR17UnifiedtciStatetypeR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoparamR17UnifiedtciStatetypeR17EnumeratedNothing = iota
	MimoparamR17UnifiedtciStatetypeR17EnumeratedSeparate
	MimoparamR17UnifiedtciStatetypeR17EnumeratedJoint
)
