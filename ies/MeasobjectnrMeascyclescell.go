package ies

import "rrc/utils"

// MeasObjectNR-measCycleSCell ::= ENUMERATED
type MeasobjectnrMeascyclescell struct {
	Value utils.ENUMERATED
}

const (
	MeasobjectnrMeascyclescellEnumeratedNothing = iota
	MeasobjectnrMeascyclescellEnumeratedSf160
	MeasobjectnrMeascyclescellEnumeratedSf256
	MeasobjectnrMeascyclescellEnumeratedSf320
	MeasobjectnrMeascyclescellEnumeratedSf512
	MeasobjectnrMeascyclescellEnumeratedSf640
	MeasobjectnrMeascyclescellEnumeratedSf1024
	MeasobjectnrMeascyclescellEnumeratedSf1280
)
