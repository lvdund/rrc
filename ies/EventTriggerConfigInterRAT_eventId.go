package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfigInterRAT_eventId_Choice_nothing uint64 = iota
	EventTriggerConfigInterRAT_eventId_Choice_EventB1
	EventTriggerConfigInterRAT_eventId_Choice_EventB2
	EventTriggerConfigInterRAT_eventId_Choice_EventB1_UTRA_FDD_r16
	EventTriggerConfigInterRAT_eventId_Choice_EventB2_UTRA_FDD_r16
	EventTriggerConfigInterRAT_eventId_Choice_EventY1_Relay_r17
	EventTriggerConfigInterRAT_eventId_Choice_EventY2_Relay_r17
)

type EventTriggerConfigInterRAT_eventId struct {
	Choice               uint64
	EventB1              *EventTriggerConfigInterRAT_eventId_eventB1
	EventB2              *EventTriggerConfigInterRAT_eventId_eventB2
	EventB1_UTRA_FDD_r16 *EventTriggerConfigInterRAT_eventId_eventB1_UTRA_FDD_r16
	EventB2_UTRA_FDD_r16 *EventTriggerConfigInterRAT_eventId_eventB2_UTRA_FDD_r16
	EventY1_Relay_r17    *EventTriggerConfigInterRAT_eventId_eventY1_Relay_r17
	EventY2_Relay_r17    *EventTriggerConfigInterRAT_eventId_eventY2_Relay_r17
}

func (ie *EventTriggerConfigInterRAT_eventId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigInterRAT_eventId_Choice_EventB1:
		if err = ie.EventB1.Encode(w); err != nil {
			err = utils.WrapError("Encode EventB1", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventB2:
		if err = ie.EventB2.Encode(w); err != nil {
			err = utils.WrapError("Encode EventB2", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventB1_UTRA_FDD_r16:
		if err = ie.EventB1_UTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode EventB1_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventB2_UTRA_FDD_r16:
		if err = ie.EventB2_UTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode EventB2_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventY1_Relay_r17:
		if err = ie.EventY1_Relay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode EventY1_Relay_r17", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventY2_Relay_r17:
		if err = ie.EventY2_Relay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode EventY2_Relay_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventTriggerConfigInterRAT_eventId) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigInterRAT_eventId_Choice_EventB1:
		ie.EventB1 = new(EventTriggerConfigInterRAT_eventId_eventB1)
		if err = ie.EventB1.Decode(r); err != nil {
			return utils.WrapError("Decode EventB1", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventB2:
		ie.EventB2 = new(EventTriggerConfigInterRAT_eventId_eventB2)
		if err = ie.EventB2.Decode(r); err != nil {
			return utils.WrapError("Decode EventB2", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventB1_UTRA_FDD_r16:
		ie.EventB1_UTRA_FDD_r16 = new(EventTriggerConfigInterRAT_eventId_eventB1_UTRA_FDD_r16)
		if err = ie.EventB1_UTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode EventB1_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventB2_UTRA_FDD_r16:
		ie.EventB2_UTRA_FDD_r16 = new(EventTriggerConfigInterRAT_eventId_eventB2_UTRA_FDD_r16)
		if err = ie.EventB2_UTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode EventB2_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventY1_Relay_r17:
		ie.EventY1_Relay_r17 = new(EventTriggerConfigInterRAT_eventId_eventY1_Relay_r17)
		if err = ie.EventY1_Relay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EventY1_Relay_r17", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_EventY2_Relay_r17:
		ie.EventY2_Relay_r17 = new(EventTriggerConfigInterRAT_eventId_eventY2_Relay_r17)
		if err = ie.EventY2_Relay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EventY2_Relay_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
