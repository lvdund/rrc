package ies

import "rrc/utils"

// PhysicalCellGroupConfig-ackNackFeedbackMode-r16 ::= ENUMERATED
type PhysicalcellgroupconfigAcknackfeedbackmodeR16 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigAcknackfeedbackmodeR16EnumeratedNothing = iota
	PhysicalcellgroupconfigAcknackfeedbackmodeR16EnumeratedJoint
	PhysicalcellgroupconfigAcknackfeedbackmodeR16EnumeratedSeparate
)
