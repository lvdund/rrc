package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestToAddMod_sr_TransMax_Enum_n4     aper.Enumerated = 0
	SchedulingRequestToAddMod_sr_TransMax_Enum_n8     aper.Enumerated = 1
	SchedulingRequestToAddMod_sr_TransMax_Enum_n16    aper.Enumerated = 2
	SchedulingRequestToAddMod_sr_TransMax_Enum_n32    aper.Enumerated = 3
	SchedulingRequestToAddMod_sr_TransMax_Enum_n64    aper.Enumerated = 4
	SchedulingRequestToAddMod_sr_TransMax_Enum_spare3 aper.Enumerated = 5
	SchedulingRequestToAddMod_sr_TransMax_Enum_spare2 aper.Enumerated = 6
	SchedulingRequestToAddMod_sr_TransMax_Enum_spare1 aper.Enumerated = 7
)

type SchedulingRequestToAddMod_sr_TransMax struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SchedulingRequestToAddMod_sr_TransMax) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SchedulingRequestToAddMod_sr_TransMax", err)
	}
	return nil
}

func (ie *SchedulingRequestToAddMod_sr_TransMax) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SchedulingRequestToAddMod_sr_TransMax", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
