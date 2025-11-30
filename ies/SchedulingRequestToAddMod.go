package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestToAddMod struct {
	SchedulingRequestId SchedulingRequestId                         `madatory`
	Sr_ProhibitTimer    *SchedulingRequestToAddMod_sr_ProhibitTimer `optional`
	Sr_TransMax         SchedulingRequestToAddMod_sr_TransMax       `madatory`
}

func (ie *SchedulingRequestToAddMod) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sr_ProhibitTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SchedulingRequestId.Encode(w); err != nil {
		return utils.WrapError("Encode SchedulingRequestId", err)
	}
	if ie.Sr_ProhibitTimer != nil {
		if err = ie.Sr_ProhibitTimer.Encode(w); err != nil {
			return utils.WrapError("Encode Sr_ProhibitTimer", err)
		}
	}
	if err = ie.Sr_TransMax.Encode(w); err != nil {
		return utils.WrapError("Encode Sr_TransMax", err)
	}
	return nil
}

func (ie *SchedulingRequestToAddMod) Decode(r *aper.AperReader) error {
	var err error
	var Sr_ProhibitTimerPresent bool
	if Sr_ProhibitTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SchedulingRequestId.Decode(r); err != nil {
		return utils.WrapError("Decode SchedulingRequestId", err)
	}
	if Sr_ProhibitTimerPresent {
		ie.Sr_ProhibitTimer = new(SchedulingRequestToAddMod_sr_ProhibitTimer)
		if err = ie.Sr_ProhibitTimer.Decode(r); err != nil {
			return utils.WrapError("Decode Sr_ProhibitTimer", err)
		}
	}
	if err = ie.Sr_TransMax.Decode(r); err != nil {
		return utils.WrapError("Decode Sr_TransMax", err)
	}
	return nil
}
