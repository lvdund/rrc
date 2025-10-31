package ies

// SL-EventTriggerConfig-r16-sl-EventId-r16 ::= CHOICE
// Extensible
const (
	SlEventtriggerconfigR16SlEventidR16ChoiceNothing = iota
	SlEventtriggerconfigR16SlEventidR16ChoiceEvents1R16
	SlEventtriggerconfigR16SlEventidR16ChoiceEvents2R16
)

type SlEventtriggerconfigR16SlEventidR16 struct {
	Choice     uint64
	Events1R16 *SlEventtriggerconfigR16SlEventidR16Events1R16
	Events2R16 *SlEventtriggerconfigR16SlEventidR16Events2R16
}
