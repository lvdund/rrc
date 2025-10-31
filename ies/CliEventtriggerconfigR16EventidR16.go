package ies

// CLI-EventTriggerConfig-r16-eventId-r16 ::= CHOICE
// Extensible
const (
	CliEventtriggerconfigR16EventidR16ChoiceNothing = iota
	CliEventtriggerconfigR16EventidR16ChoiceEventi1R16
)

type CliEventtriggerconfigR16EventidR16 struct {
	Choice     uint64
	Eventi1R16 *CliEventtriggerconfigR16EventidR16Eventi1R16
}
