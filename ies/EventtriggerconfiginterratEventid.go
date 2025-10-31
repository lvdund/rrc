package ies

// EventTriggerConfigInterRAT-eventId ::= CHOICE
// Extensible
const (
	EventtriggerconfiginterratEventidChoiceNothing = iota
	EventtriggerconfiginterratEventidChoiceEventb1
	EventtriggerconfiginterratEventidChoiceEventb2
	EventtriggerconfiginterratEventidChoiceEventb1UtraFddR16
	EventtriggerconfiginterratEventidChoiceEventb2UtraFddR16
	EventtriggerconfiginterratEventidChoiceEventy1RelayR17
	EventtriggerconfiginterratEventidChoiceEventy2RelayR17
)

type EventtriggerconfiginterratEventid struct {
	Choice            uint64
	Eventb1           *EventtriggerconfiginterratEventidEventb1
	Eventb2           *EventtriggerconfiginterratEventidEventb2
	Eventb1UtraFddR16 *EventtriggerconfiginterratEventidEventb1UtraFddR16
	Eventb2UtraFddR16 *EventtriggerconfiginterratEventidEventb2UtraFddR16
	Eventy1RelayR17   *EventtriggerconfiginterratEventidEventy1RelayR17
	Eventy2RelayR17   *EventtriggerconfiginterratEventidEventy2RelayR17
}
