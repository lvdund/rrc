package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfigNR_SL_r16_eventId_r16_Choice_nothing uint64 = iota
	EventTriggerConfigNR_SL_r16_eventId_r16_Choice_EventC1
	EventTriggerConfigNR_SL_r16_eventId_r16_Choice_EventC2_r16
)

type EventTriggerConfigNR_SL_r16_eventId_r16 struct {
	Choice      uint64
	EventC1     *EventTriggerConfigNR_SL_r16_eventId_r16_eventC1
	EventC2_r16 *EventTriggerConfigNR_SL_r16_eventId_r16_eventC2_r16
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_EventC1:
		if err = ie.EventC1.Encode(w); err != nil {
			err = utils.WrapError("Encode EventC1", err)
		}
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_EventC2_r16:
		if err = ie.EventC2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode EventC2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_EventC1:
		ie.EventC1 = new(EventTriggerConfigNR_SL_r16_eventId_r16_eventC1)
		if err = ie.EventC1.Decode(r); err != nil {
			return utils.WrapError("Decode EventC1", err)
		}
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_EventC2_r16:
		ie.EventC2_r16 = new(EventTriggerConfigNR_SL_r16_eventId_r16_eventC2_r16)
		if err = ie.EventC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode EventC2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
