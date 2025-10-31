package ies

// MeasResultsSL-r16-measResultsListSL-r16 ::= CHOICE
// Extensible
const (
	MeasresultsslR16MeasresultslistslR16ChoiceNothing = iota
	MeasresultsslR16MeasresultslistslR16ChoiceMeasresultnrSlR16
)

type MeasresultsslR16MeasresultslistslR16 struct {
	Choice            uint64
	MeasresultnrSlR16 *MeasresultnrSlR16
}
