package ies

import "rrc/utils"

// PUR-MPDCCH-Config-r16-mpdcch-StartSF-UESS-r16 ::= CHOICE
const (
	PurMpdcchConfigR16MpdcchStartsfUessR16ChoiceNothing = iota
	PurMpdcchConfigR16MpdcchStartsfUessR16ChoiceFdd
	PurMpdcchConfigR16MpdcchStartsfUessR16ChoiceTdd
)

type PurMpdcchConfigR16MpdcchStartsfUessR16 struct {
	Choice uint64
	Fdd    *utils.ENUMERATED
	Tdd    *utils.ENUMERATED
}
