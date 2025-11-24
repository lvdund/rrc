package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestResourceConfig struct {
	SchedulingRequestResourceId SchedulingRequestResourceId                           `madatory`
	SchedulingRequestID         SchedulingRequestId                                   `madatory`
	PeriodicityAndOffset        *SchedulingRequestResourceConfig_periodicityAndOffset `lb:0,ub:1,optional`
	Resource                    *PUCCH_ResourceId                                     `optional`
}

func (ie *SchedulingRequestResourceConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PeriodicityAndOffset != nil, ie.Resource != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SchedulingRequestResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode SchedulingRequestResourceId", err)
	}
	if err = ie.SchedulingRequestID.Encode(w); err != nil {
		return utils.WrapError("Encode SchedulingRequestID", err)
	}
	if ie.PeriodicityAndOffset != nil {
		if err = ie.PeriodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicityAndOffset", err)
		}
	}
	if ie.Resource != nil {
		if err = ie.Resource.Encode(w); err != nil {
			return utils.WrapError("Encode Resource", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestResourceConfig) Decode(r *uper.UperReader) error {
	var err error
	var PeriodicityAndOffsetPresent bool
	if PeriodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourcePresent bool
	if ResourcePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SchedulingRequestResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode SchedulingRequestResourceId", err)
	}
	if err = ie.SchedulingRequestID.Decode(r); err != nil {
		return utils.WrapError("Decode SchedulingRequestID", err)
	}
	if PeriodicityAndOffsetPresent {
		ie.PeriodicityAndOffset = new(SchedulingRequestResourceConfig_periodicityAndOffset)
		if err = ie.PeriodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicityAndOffset", err)
		}
	}
	if ResourcePresent {
		ie.Resource = new(PUCCH_ResourceId)
		if err = ie.Resource.Decode(r); err != nil {
			return utils.WrapError("Decode Resource", err)
		}
	}
	return nil
}
