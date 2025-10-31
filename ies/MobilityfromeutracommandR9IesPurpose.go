package ies

// MobilityFromEUTRACommand-r9-IEs-purpose ::= CHOICE
// Extensible
const (
	MobilityfromeutracommandR9IesPurposeChoiceNothing = iota
	MobilityfromeutracommandR9IesPurposeChoiceHandover
	MobilityfromeutracommandR9IesPurposeChoiceCellchangeorder
	MobilityfromeutracommandR9IesPurposeChoiceECsfbR9
)

type MobilityfromeutracommandR9IesPurpose struct {
	Choice          uint64
	Handover        *Handover
	Cellchangeorder *Cellchangeorder
	ECsfbR9         *ECsfbR9
}
