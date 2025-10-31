package ies

// PUCCH-ConfigDedicated-v1020-pucch-Format-r10 ::= CHOICE
const (
	PucchConfigdedicatedV1020PucchFormatR10ChoiceNothing = iota
	PucchConfigdedicatedV1020PucchFormatR10ChoiceFormat3R10
	PucchConfigdedicatedV1020PucchFormatR10ChoiceChannelselectionR10
)

type PucchConfigdedicatedV1020PucchFormatR10 struct {
	Choice              uint64
	Format3R10          *PucchFormat3ConfR13
	ChannelselectionR10 *PucchConfigdedicatedV1020PucchFormatR10ChannelselectionR10
}
