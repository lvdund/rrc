package ies

import "rrc/utils"

// SIB-TypeInfo-type ::= utils.ENUMERATED // Extensible
type SibTypeinfoType struct {
	Value utils.ENUMERATED
}

const (
	SibTypeinfoTypeEnumeratedNothing = iota
	SibTypeinfoTypeEnumeratedSibtype2
	SibTypeinfoTypeEnumeratedSibtype3
	SibTypeinfoTypeEnumeratedSibtype4
	SibTypeinfoTypeEnumeratedSibtype5
	SibTypeinfoTypeEnumeratedSibtype6
	SibTypeinfoTypeEnumeratedSibtype7
	SibTypeinfoTypeEnumeratedSibtype8
	SibTypeinfoTypeEnumeratedSibtype9
	SibTypeinfoTypeEnumeratedSibtype10_V1610
	SibTypeinfoTypeEnumeratedSibtype11_V1610
	SibTypeinfoTypeEnumeratedSibtype12_V1610
	SibTypeinfoTypeEnumeratedSibtype13_V1610
	SibTypeinfoTypeEnumeratedSibtype14_V1610
	SibTypeinfoTypeEnumeratedSpare3
	SibTypeinfoTypeEnumeratedSpare2
	SibTypeinfoTypeEnumeratedSpare1
)
