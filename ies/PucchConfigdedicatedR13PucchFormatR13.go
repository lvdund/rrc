package ies

// PUCCH-ConfigDedicated-r13-pucch-Format-r13 ::= CHOICE
const (
	PucchConfigdedicatedR13PucchFormatR13ChoiceNothing = iota
	PucchConfigdedicatedR13PucchFormatR13ChoiceFormat3R13
	PucchConfigdedicatedR13PucchFormatR13ChoiceChannelselectionR13
	PucchConfigdedicatedR13PucchFormatR13ChoiceFormat4R13
	PucchConfigdedicatedR13PucchFormatR13ChoiceFormat5R13
)

type PucchConfigdedicatedR13PucchFormatR13 struct {
	Choice              uint64
	Format3R13          *PucchConfigdedicatedR13PucchFormatR13Format3R13
	ChannelselectionR13 *PucchConfigdedicatedR13PucchFormatR13ChannelselectionR13
	Format4R13          *PucchConfigdedicatedR13PucchFormatR13Format4R13
	Format5R13          *PucchConfigdedicatedR13PucchFormatR13Format5R13
}
