package ies

import "rrc/utils"

// SL-PSBCH-Config-r16-dl-Alpha-PSBCH-r16 ::= ENUMERATED
type SlPsbchConfigR16DlAlphaPsbchR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedNothing = iota
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha0
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha04
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha05
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha06
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha07
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha08
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha09
	SlPsbchConfigR16DlAlphaPsbchR16EnumeratedAlpha1
)
