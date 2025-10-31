package ies

// PUCCH-ConfigDedicated-r13-pucch-Format-r13-channelSelection-r13-n1PUCCH-AN-CS-r13 ::= CHOICE
const (
	PucchConfigdedicatedR13PucchFormatR13ChannelselectionR13N1pucchAnCsR13ChoiceNothing = iota
	PucchConfigdedicatedR13PucchFormatR13ChannelselectionR13N1pucchAnCsR13ChoiceRelease
	PucchConfigdedicatedR13PucchFormatR13ChannelselectionR13N1pucchAnCsR13ChoiceSetup
)

type PucchConfigdedicatedR13PucchFormatR13ChannelselectionR13N1pucchAnCsR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedR13PucchFormatR13ChannelselectionR13N1pucchAnCsR13Setup
}
