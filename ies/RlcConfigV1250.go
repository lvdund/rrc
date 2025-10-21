package ies

import "rrc/utils"

// RLC-Config-v1250 ::= SEQUENCE
type RlcConfigV1250 struct {
	UlExtendedRlcLiFieldR12 bool
	DlExtendedRlcLiFieldR12 bool
}
