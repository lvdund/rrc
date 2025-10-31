package ies

// MobilityFromEUTRACommand-r8-IEs-purpose ::= CHOICE
const (
	MobilityfromeutracommandR8IesPurposeChoiceNothing = iota
	MobilityfromeutracommandR8IesPurposeChoiceHandover
	MobilityfromeutracommandR8IesPurposeChoiceCellchangeorder
)

type MobilityfromeutracommandR8IesPurpose struct {
	Choice          uint64
	Handover        *Handover
	Cellchangeorder *Cellchangeorder
}
