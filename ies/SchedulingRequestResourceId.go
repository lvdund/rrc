package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestResourceId struct {
	Value uint64 `lb:1,ub:maxNrofSR_Resources,madatory`
}

func (ie *SchedulingRequestResourceId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false); err != nil {
		return utils.WrapError("Encode SchedulingRequestResourceId", err)
	}
	return nil
}

func (ie *SchedulingRequestResourceId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false); err != nil {
		return utils.WrapError("Decode SchedulingRequestResourceId", err)
	}
	ie.Value = uint64(v)
	return nil
}
