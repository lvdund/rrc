package ies

// SRB-ToAddMod-logicalChannelConfig ::= CHOICE
const (
	SrbToaddmodLogicalchannelconfigChoiceNothing = iota
	SrbToaddmodLogicalchannelconfigChoiceExplicitvalue
	SrbToaddmodLogicalchannelconfigChoiceDefaultvalue
)

type SrbToaddmodLogicalchannelconfig struct {
	Choice        uint64
	Explicitvalue *Logicalchannelconfig
	Defaultvalue  *struct{}
}
