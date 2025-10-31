package ies

// SRB-ToAddMod-NB-r13-rlc-Config-r13 ::= CHOICE
const (
	SrbToaddmodNbR13RlcConfigR13ChoiceNothing = iota
	SrbToaddmodNbR13RlcConfigR13ChoiceExplicitvalue
	SrbToaddmodNbR13RlcConfigR13ChoiceDefaultvalue
)

type SrbToaddmodNbR13RlcConfigR13 struct {
	Choice        uint64
	Explicitvalue *RlcConfigNbR13
	Defaultvalue  *struct{}
}
