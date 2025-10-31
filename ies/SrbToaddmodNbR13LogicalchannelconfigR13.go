package ies

// SRB-ToAddMod-NB-r13-logicalChannelConfig-r13 ::= CHOICE
const (
	SrbToaddmodNbR13LogicalchannelconfigR13ChoiceNothing = iota
	SrbToaddmodNbR13LogicalchannelconfigR13ChoiceExplicitvalue
	SrbToaddmodNbR13LogicalchannelconfigR13ChoiceDefaultvalue
)

type SrbToaddmodNbR13LogicalchannelconfigR13 struct {
	Choice        uint64
	Explicitvalue *LogicalchannelconfigNbR13
	Defaultvalue  *struct{}
}
