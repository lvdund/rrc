package ies

// PUCCH-ConfigDedicated-v13c0-channelSelection-v13c0-n1PUCCH-AN-CS-v13c0 ::= CHOICE
const (
	PucchConfigdedicatedV13c0ChannelselectionV13c0N1pucchAnCsV13c0ChoiceNothing = iota
	PucchConfigdedicatedV13c0ChannelselectionV13c0N1pucchAnCsV13c0ChoiceRelease
	PucchConfigdedicatedV13c0ChannelselectionV13c0N1pucchAnCsV13c0ChoiceSetup
)

type PucchConfigdedicatedV13c0ChannelselectionV13c0N1pucchAnCsV13c0 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedV13c0ChannelselectionV13c0N1pucchAnCsV13c0Setup
}
