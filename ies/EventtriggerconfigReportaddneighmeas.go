package ies

import "rrc/utils"

// EventTriggerConfig-reportAddNeighMeas ::= ENUMERATED
type EventtriggerconfigReportaddneighmeas struct {
	Value utils.ENUMERATED
}

const (
	EventtriggerconfigReportaddneighmeasEnumeratedNothing = iota
	EventtriggerconfigReportaddneighmeasEnumeratedSetup
)
