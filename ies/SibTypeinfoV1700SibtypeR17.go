package ies

import "rrc/utils"

// SIB-TypeInfo-v1700-sibType-r17 ::= CHOICE
// Extensible
const (
	SibTypeinfoV1700SibtypeR17ChoiceNothing = iota
	SibTypeinfoV1700SibtypeR17ChoiceType1R17
	SibTypeinfoV1700SibtypeR17ChoiceType2R17
)

type SibTypeinfoV1700SibtypeR17 struct {
	Choice   uint64
	Type1R17 *utils.ENUMERATED
	Type2R17 *SibTypeinfoV1700SibtypeR17Type2R17
}
