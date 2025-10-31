package ies

import "rrc/utils"

// RRCReconfiguration-v1700-IEs-scg-State-r17 ::= ENUMERATED
type RrcreconfigurationV1700IesScgStateR17 struct {
	Value utils.ENUMERATED
}

const (
	RrcreconfigurationV1700IesScgStateR17EnumeratedNothing = iota
	RrcreconfigurationV1700IesScgStateR17EnumeratedDeactivated
)
