package ies

import "rrc/utils"

// SupportedUDC-r15 ::= SEQUENCE
type SupportedudcR15 struct {
	SupportedstandarddicR15 *utils.ENUMERATED
	SupportedoperatordicR15 *SupportedoperatordicR15
}
