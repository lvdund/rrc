package ies

import "rrc/utils"

// MAC-ParametersCommon-jointPrioritizationCG-Retx-Timer-r17 ::= ENUMERATED
type MacParameterscommonJointprioritizationcgRetxTimerR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonJointprioritizationcgRetxTimerR17EnumeratedNothing = iota
	MacParameterscommonJointprioritizationcgRetxTimerR17EnumeratedSupported
)
