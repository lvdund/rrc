package ies

import "rrc/utils"

// NPUSCH-ConfigDedicated-NB-v1610-npusch-MultiTB-Config-r16 ::= ENUMERATED
type NpuschConfigdedicatedNbV1610NpuschMultitbConfigR16 struct {
	Value utils.ENUMERATED
}

const (
	NpuschConfigdedicatedNbV1610NpuschMultitbConfigR16EnumeratedNothing = iota
	NpuschConfigdedicatedNbV1610NpuschMultitbConfigR16EnumeratedInterleaved
	NpuschConfigdedicatedNbV1610NpuschMultitbConfigR16EnumeratedNoninterleaved
)
