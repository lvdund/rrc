package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventType_r16_Choice_nothing uint64 = iota
	EventType_r16_Choice_OutOfCoverage
	EventType_r16_Choice_EventL1
)

type EventType_r16 struct {
	Choice        uint64
	OutOfCoverage aper.NULL `madatory`
	EventL1       *EventType_r16_eventL1
}

func (ie *EventType_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventType_r16_Choice_OutOfCoverage:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode OutOfCoverage", err)
		}
	case EventType_r16_Choice_EventL1:
		if err = ie.EventL1.Encode(w); err != nil {
			err = utils.WrapError("Encode EventL1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventType_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventType_r16_Choice_OutOfCoverage:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode OutOfCoverage", err)
		}
	case EventType_r16_Choice_EventL1:
		ie.EventL1 = new(EventType_r16_eventL1)
		if err = ie.EventL1.Decode(r); err != nil {
			return utils.WrapError("Decode EventL1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
