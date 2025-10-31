package ies

import "rrc/utils"

// RRCReconfigurationCompleteSidelink-v1720-IEs ::= SEQUENCE
type RrcreconfigurationcompletesidelinkV1720 struct {
	SlDrxConfigrejectV1720 *utils.BOOLEAN
	Noncriticalextension   *RrcreconfigurationcompletesidelinkV1720IesNoncriticalextension
}
