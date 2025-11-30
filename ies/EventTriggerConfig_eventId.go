package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfig_eventId_Choice_nothing uint64 = iota
	EventTriggerConfig_eventId_Choice_EventA1
	EventTriggerConfig_eventId_Choice_EventA2
	EventTriggerConfig_eventId_Choice_EventA3
	EventTriggerConfig_eventId_Choice_EventA4
	EventTriggerConfig_eventId_Choice_EventA5
	EventTriggerConfig_eventId_Choice_EventA6
	EventTriggerConfig_eventId_Choice_EventX1_r17
	EventTriggerConfig_eventId_Choice_EventX2_r17
	EventTriggerConfig_eventId_Choice_EventD1_r17
)

type EventTriggerConfig_eventId struct {
	Choice      uint64
	EventA1     *EventTriggerConfig_eventId_eventA1
	EventA2     *EventTriggerConfig_eventId_eventA2
	EventA3     *EventTriggerConfig_eventId_eventA3
	EventA4     *EventTriggerConfig_eventId_eventA4
	EventA5     *EventTriggerConfig_eventId_eventA5
	EventA6     *EventTriggerConfig_eventId_eventA6
	EventX1_r17 *EventTriggerConfig_eventId_eventX1_r17
	EventX2_r17 *EventTriggerConfig_eventId_eventX2_r17
	EventD1_r17 *EventTriggerConfig_eventId_eventD1_r17
}

func (ie *EventTriggerConfig_eventId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfig_eventId_Choice_EventA1:
		if err = ie.EventA1.Encode(w); err != nil {
			err = utils.WrapError("Encode EventA1", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA2:
		if err = ie.EventA2.Encode(w); err != nil {
			err = utils.WrapError("Encode EventA2", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA3:
		if err = ie.EventA3.Encode(w); err != nil {
			err = utils.WrapError("Encode EventA3", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA4:
		if err = ie.EventA4.Encode(w); err != nil {
			err = utils.WrapError("Encode EventA4", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA5:
		if err = ie.EventA5.Encode(w); err != nil {
			err = utils.WrapError("Encode EventA5", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA6:
		if err = ie.EventA6.Encode(w); err != nil {
			err = utils.WrapError("Encode EventA6", err)
		}
	case EventTriggerConfig_eventId_Choice_EventX1_r17:
		if err = ie.EventX1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode EventX1_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_EventX2_r17:
		if err = ie.EventX2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode EventX2_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_EventD1_r17:
		if err = ie.EventD1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode EventD1_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventTriggerConfig_eventId) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfig_eventId_Choice_EventA1:
		ie.EventA1 = new(EventTriggerConfig_eventId_eventA1)
		if err = ie.EventA1.Decode(r); err != nil {
			return utils.WrapError("Decode EventA1", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA2:
		ie.EventA2 = new(EventTriggerConfig_eventId_eventA2)
		if err = ie.EventA2.Decode(r); err != nil {
			return utils.WrapError("Decode EventA2", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA3:
		ie.EventA3 = new(EventTriggerConfig_eventId_eventA3)
		if err = ie.EventA3.Decode(r); err != nil {
			return utils.WrapError("Decode EventA3", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA4:
		ie.EventA4 = new(EventTriggerConfig_eventId_eventA4)
		if err = ie.EventA4.Decode(r); err != nil {
			return utils.WrapError("Decode EventA4", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA5:
		ie.EventA5 = new(EventTriggerConfig_eventId_eventA5)
		if err = ie.EventA5.Decode(r); err != nil {
			return utils.WrapError("Decode EventA5", err)
		}
	case EventTriggerConfig_eventId_Choice_EventA6:
		ie.EventA6 = new(EventTriggerConfig_eventId_eventA6)
		if err = ie.EventA6.Decode(r); err != nil {
			return utils.WrapError("Decode EventA6", err)
		}
	case EventTriggerConfig_eventId_Choice_EventX1_r17:
		ie.EventX1_r17 = new(EventTriggerConfig_eventId_eventX1_r17)
		if err = ie.EventX1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EventX1_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_EventX2_r17:
		ie.EventX2_r17 = new(EventTriggerConfig_eventId_eventX2_r17)
		if err = ie.EventX2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EventX2_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_EventD1_r17:
		ie.EventD1_r17 = new(EventTriggerConfig_eventId_eventD1_r17)
		if err = ie.EventD1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EventD1_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
