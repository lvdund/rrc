package ies

// EventTriggerConfig-eventId ::= CHOICE
// Extensible
const (
	EventtriggerconfigEventidChoiceNothing = iota
	EventtriggerconfigEventidChoiceEventa1
	EventtriggerconfigEventidChoiceEventa2
	EventtriggerconfigEventidChoiceEventa3
	EventtriggerconfigEventidChoiceEventa4
	EventtriggerconfigEventidChoiceEventa5
	EventtriggerconfigEventidChoiceEventa6
	EventtriggerconfigEventidChoiceEventx1R17
	EventtriggerconfigEventidChoiceEventx2R17
	EventtriggerconfigEventidChoiceEventd1R17
)

type EventtriggerconfigEventid struct {
	Choice     uint64
	Eventa1    *EventtriggerconfigEventidEventa1
	Eventa2    *EventtriggerconfigEventidEventa2
	Eventa3    *EventtriggerconfigEventidEventa3
	Eventa4    *EventtriggerconfigEventidEventa4
	Eventa5    *EventtriggerconfigEventidEventa5
	Eventa6    *EventtriggerconfigEventidEventa6
	Eventx1R17 *EventtriggerconfigEventidEventx1R17
	Eventx2R17 *EventtriggerconfigEventidEventx2R17
	Eventd1R17 *EventtriggerconfigEventidEventd1R17
}
