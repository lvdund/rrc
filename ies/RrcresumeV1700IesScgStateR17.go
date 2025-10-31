package ies

import "rrc/utils"

// RRCResume-v1700-IEs-scg-State-r17 ::= ENUMERATED
type RrcresumeV1700IesScgStateR17 struct {
	Value utils.ENUMERATED
}

const (
	RrcresumeV1700IesScgStateR17EnumeratedNothing = iota
	RrcresumeV1700IesScgStateR17EnumeratedDeactivated
)
