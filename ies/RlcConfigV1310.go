package ies

import "rrc/utils"

// RLC-Config-v1310 ::= SEQUENCE
type RlcConfigV1310 struct {
	UlExtendedRlcAmSnR13 utils.BOOLEAN
	DlExtendedRlcAmSnR13 utils.BOOLEAN
	PollpduV1310         *PollpduV1310
}
