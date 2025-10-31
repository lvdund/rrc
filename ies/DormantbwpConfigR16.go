package ies

import "rrc/utils"

// DormantBWP-Config-r16 ::= SEQUENCE
type DormantbwpConfigR16 struct {
	DormantbwpIdR16            *BwpId
	WithinactivetimeconfigR16  *utils.Setuprelease[WithinactivetimeconfigR16]
	OutsideactivetimeconfigR16 *utils.Setuprelease[OutsideactivetimeconfigR16]
}
