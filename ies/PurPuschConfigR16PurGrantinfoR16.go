package ies

// PUR-PUSCH-Config-r16-pur-GrantInfo-r16 ::= CHOICE
const (
	PurPuschConfigR16PurGrantinfoR16ChoiceNothing = iota
	PurPuschConfigR16PurGrantinfoR16ChoiceCeModea
	PurPuschConfigR16PurGrantinfoR16ChoiceCeModeb
)

type PurPuschConfigR16PurGrantinfoR16 struct {
	Choice  uint64
	CeModea *PurPuschConfigR16PurGrantinfoR16CeModea
	CeModeb *PurPuschConfigR16PurGrantinfoR16CeModeb
}
