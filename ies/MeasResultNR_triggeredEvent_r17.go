package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_triggeredEvent_r17 struct {
	TimeBetweenEvents_r17 *TimeBetweenEvent_r17                                `optional`
	FirstTriggeredEvent   *MeasResultNR_triggeredEvent_r17_firstTriggeredEvent `optional`
}

func (ie *MeasResultNR_triggeredEvent_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TimeBetweenEvents_r17 != nil, ie.FirstTriggeredEvent != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.TimeBetweenEvents_r17 != nil {
		if err = ie.TimeBetweenEvents_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TimeBetweenEvents_r17", err)
		}
	}
	if ie.FirstTriggeredEvent != nil {
		if err = ie.FirstTriggeredEvent.Encode(w); err != nil {
			return utils.WrapError("Encode FirstTriggeredEvent", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_triggeredEvent_r17) Decode(r *uper.UperReader) error {
	var err error
	var TimeBetweenEvents_r17Present bool
	if TimeBetweenEvents_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FirstTriggeredEventPresent bool
	if FirstTriggeredEventPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if TimeBetweenEvents_r17Present {
		ie.TimeBetweenEvents_r17 = new(TimeBetweenEvent_r17)
		if err = ie.TimeBetweenEvents_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TimeBetweenEvents_r17", err)
		}
	}
	if FirstTriggeredEventPresent {
		ie.FirstTriggeredEvent = new(MeasResultNR_triggeredEvent_r17_firstTriggeredEvent)
		if err = ie.FirstTriggeredEvent.Decode(r); err != nil {
			return utils.WrapError("Decode FirstTriggeredEvent", err)
		}
	}
	return nil
}
