package ies

// PUCCH-ConfigDedicated-v1020-pucch-Format-r10-channelSelection-r10-n1PUCCH-AN-CS-r10 ::= CHOICE
const (
	PucchConfigdedicatedV1020PucchFormatR10ChannelselectionR10N1pucchAnCsR10ChoiceNothing = iota
	PucchConfigdedicatedV1020PucchFormatR10ChannelselectionR10N1pucchAnCsR10ChoiceRelease
	PucchConfigdedicatedV1020PucchFormatR10ChannelselectionR10N1pucchAnCsR10ChoiceSetup
)

type PucchConfigdedicatedV1020PucchFormatR10ChannelselectionR10N1pucchAnCsR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedV1020PucchFormatR10ChannelselectionR10N1pucchAnCsR10Setup
}
