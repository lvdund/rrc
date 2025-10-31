package ies

// CLI-TriggeredList-r16 ::= CHOICE
const (
	CliTriggeredlistR16ChoiceNothing = iota
	CliTriggeredlistR16ChoiceSrsRsrpTriggeredlistR16
	CliTriggeredlistR16ChoiceCliRssiTriggeredlistR16
)

type CliTriggeredlistR16 struct {
	Choice                  uint64
	SrsRsrpTriggeredlistR16 *SrsRsrpTriggeredlistR16
	CliRssiTriggeredlistR16 *CliRssiTriggeredlistR16
}
