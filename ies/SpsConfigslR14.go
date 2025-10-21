package ies

import "rrc/utils"

// SPS-ConfigSL-r14 ::= SEQUENCE
type SpsConfigslR14 struct {
	SpsConfigindexR14             SpsConfigindexR14
	SemipersistschedintervalslR14 utils.ENUMERATED
}
