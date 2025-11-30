package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingInfo_si_BroadcastStatus_Enum_broadcasting    aper.Enumerated = 0
	SchedulingInfo_si_BroadcastStatus_Enum_notBroadcasting aper.Enumerated = 1
)

type SchedulingInfo_si_BroadcastStatus struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SchedulingInfo_si_BroadcastStatus) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SchedulingInfo_si_BroadcastStatus", err)
	}
	return nil
}

func (ie *SchedulingInfo_si_BroadcastStatus) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SchedulingInfo_si_BroadcastStatus", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
