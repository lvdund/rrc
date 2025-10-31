package ies

import "rrc/utils"

// SemiStaticChannelAccessConfigUE-r17 ::= SEQUENCE
type SemistaticchannelaccessconfigueR17 struct {
	PeriodueR17 SemistaticchannelaccessconfigueR17PeriodueR17
	OffsetueR17 utils.INTEGER `lb:0,ub:559`
}
