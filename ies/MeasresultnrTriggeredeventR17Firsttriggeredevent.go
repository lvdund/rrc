package ies

import "rrc/utils"

// MeasResultNR-triggeredEvent-r17-firstTriggeredEvent ::= ENUMERATED
type MeasresultnrTriggeredeventR17Firsttriggeredevent struct {
	Value utils.ENUMERATED
}

const (
	MeasresultnrTriggeredeventR17FirsttriggeredeventEnumeratedNothing = iota
	MeasresultnrTriggeredeventR17FirsttriggeredeventEnumeratedCondfirstevent
	MeasresultnrTriggeredeventR17FirsttriggeredeventEnumeratedCondsecondevent
)
