package ies

import "rrc/utils"

// RLC-Config-v1310 ::= SEQUENCE
type RlcConfigV1310 struct {
	UlExtendedRlcAmSnR13 bool
	DlExtendedRlcAmSnR13 bool
	PollpduV1310         *PollpduV1310
}
