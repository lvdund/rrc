package ies

import "rrc/utils"

// SIB-Type-r17 ::= ENUMERATED
type SibTypeR17 struct {
	Value utils.ENUMERATED
}

const (
	SibTypeR17EnumeratedNothing = iota
	SibTypeR17EnumeratedSibtype2
	SibTypeR17EnumeratedSibtype3
	SibTypeR17EnumeratedSibtype4
	SibTypeR17EnumeratedSibtype5
	SibTypeR17EnumeratedSibtype9
	SibTypeR17EnumeratedSibtype10_V1610
	SibTypeR17EnumeratedSibtype11_V1610
	SibTypeR17EnumeratedSibtype12_V1610
	SibTypeR17EnumeratedSibtype13_V1610
	SibTypeR17EnumeratedSibtype14_V1610
	SibTypeR17EnumeratedSpare6
	SibTypeR17EnumeratedSpare5
	SibTypeR17EnumeratedSpare4
	SibTypeR17EnumeratedSpare3
	SibTypeR17EnumeratedSpare2
	SibTypeR17EnumeratedSpare1
)
