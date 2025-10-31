package ies

import "rrc/utils"

// PUR-Config-r16-pur-ImplicitReleaseAfter-r16 ::= ENUMERATED
type PurConfigR16PurImplicitreleaseafterR16 struct {
	Value utils.ENUMERATED
}

const (
	PurConfigR16PurImplicitreleaseafterR16EnumeratedNothing = iota
	PurConfigR16PurImplicitreleaseafterR16EnumeratedN2
	PurConfigR16PurImplicitreleaseafterR16EnumeratedN4
	PurConfigR16PurImplicitreleaseafterR16EnumeratedN8
	PurConfigR16PurImplicitreleaseafterR16EnumeratedSpare
)
