package ies

import "rrc/utils"

// MAC-ParametersFR2-2-r17-directSCG-SCellActivationResume-r17 ::= ENUMERATED
type MacParametersfr22R17DirectscgScellactivationresumeR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersfr22R17DirectscgScellactivationresumeR17EnumeratedNothing = iota
	MacParametersfr22R17DirectscgScellactivationresumeR17EnumeratedSupported
)
