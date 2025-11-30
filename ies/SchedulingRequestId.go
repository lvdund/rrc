package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestId struct {
	Value uint64 `lb:0,ub:7,madatory`
}

func (ie *SchedulingRequestId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SchedulingRequestId", err)
	}
	return nil
}

func (ie *SchedulingRequestId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SchedulingRequestId", err)
	}
	ie.Value = uint64(v)
	return nil
}
