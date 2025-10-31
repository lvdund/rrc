package ies

import "rrc/utils"

// PUR-Config-NB-r16-pur-ImplicitReleaseAfter-r16 ::= ENUMERATED
type PurConfigNbR16PurImplicitreleaseafterR16 struct {
	Value utils.ENUMERATED
}

const (
	PurConfigNbR16PurImplicitreleaseafterR16EnumeratedNothing = iota
	PurConfigNbR16PurImplicitreleaseafterR16EnumeratedN2
	PurConfigNbR16PurImplicitreleaseafterR16EnumeratedN4
	PurConfigNbR16PurImplicitreleaseafterR16EnumeratedN8
	PurConfigNbR16PurImplicitreleaseafterR16EnumeratedSpare
)
