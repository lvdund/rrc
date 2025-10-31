package ies

import "rrc/utils"

// AC-BarringConfig-ac-BarringTime ::= ENUMERATED
type AcBarringconfigAcBarringtime struct {
	Value utils.ENUMERATED
}

const (
	AcBarringconfigAcBarringtimeEnumeratedNothing = iota
	AcBarringconfigAcBarringtimeEnumeratedS4
	AcBarringconfigAcBarringtimeEnumeratedS8
	AcBarringconfigAcBarringtimeEnumeratedS16
	AcBarringconfigAcBarringtimeEnumeratedS32
	AcBarringconfigAcBarringtimeEnumeratedS64
	AcBarringconfigAcBarringtimeEnumeratedS128
	AcBarringconfigAcBarringtimeEnumeratedS256
	AcBarringconfigAcBarringtimeEnumeratedS512
)
