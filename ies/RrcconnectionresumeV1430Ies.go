package ies

import "rrc/utils"

// RRCConnectionResume-v1430-IEs ::= SEQUENCE
type RrcconnectionresumeV1430Ies struct {
	OtherconfigR14              *OtherconfigR9
	RrcconnectionresumeV1510Ies *RrcconnectionresumeV1510Ies
}
