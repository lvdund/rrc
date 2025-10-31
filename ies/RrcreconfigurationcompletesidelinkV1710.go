package ies

import "rrc/utils"

// RRCReconfigurationCompleteSidelink-v1710-IEs ::= SEQUENCE
type RrcreconfigurationcompletesidelinkV1710 struct {
	Dummy                utils.BOOLEAN
	Noncriticalextension *RrcreconfigurationcompletesidelinkV1720
}
