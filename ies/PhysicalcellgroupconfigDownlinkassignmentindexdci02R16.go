package ies

import "rrc/utils"

// PhysicalCellGroupConfig-downlinkAssignmentIndexDCI-0-2-r16 ::= ENUMERATED
type PhysicalcellgroupconfigDownlinkassignmentindexdci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigDownlinkassignmentindexdci02R16EnumeratedNothing = iota
	PhysicalcellgroupconfigDownlinkassignmentindexdci02R16EnumeratedEnabled
)
