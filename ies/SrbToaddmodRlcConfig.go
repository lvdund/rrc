package ies

// SRB-ToAddMod-rlc-Config ::= CHOICE
const (
	SrbToaddmodRlcConfigChoiceNothing = iota
	SrbToaddmodRlcConfigChoiceExplicitvalue
	SrbToaddmodRlcConfigChoiceDefaultvalue
)

type SrbToaddmodRlcConfig struct {
	Choice        uint64
	Explicitvalue *RlcConfig
	Defaultvalue  *struct{}
}
