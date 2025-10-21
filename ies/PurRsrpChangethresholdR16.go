package ies

import "rrc/utils"

// PUR-RSRP-ChangeThreshold-r16 ::= SEQUENCE
type PurRsrpChangethresholdR16 struct {
	IncreasethreshR16 RsrpChangethreshR16
	DecreasethreshR16 *RsrpChangethreshR16
}
