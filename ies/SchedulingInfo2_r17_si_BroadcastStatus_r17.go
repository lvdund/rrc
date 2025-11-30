package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingInfo2_r17_si_BroadcastStatus_r17_Enum_broadcasting    aper.Enumerated = 0
	SchedulingInfo2_r17_si_BroadcastStatus_r17_Enum_notBroadcasting aper.Enumerated = 1
)

type SchedulingInfo2_r17_si_BroadcastStatus_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SchedulingInfo2_r17_si_BroadcastStatus_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SchedulingInfo2_r17_si_BroadcastStatus_r17", err)
	}
	return nil
}

func (ie *SchedulingInfo2_r17_si_BroadcastStatus_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SchedulingInfo2_r17_si_BroadcastStatus_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
