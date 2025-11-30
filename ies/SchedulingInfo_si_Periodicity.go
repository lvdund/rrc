package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingInfo_si_Periodicity_Enum_rf8   aper.Enumerated = 0
	SchedulingInfo_si_Periodicity_Enum_rf16  aper.Enumerated = 1
	SchedulingInfo_si_Periodicity_Enum_rf32  aper.Enumerated = 2
	SchedulingInfo_si_Periodicity_Enum_rf64  aper.Enumerated = 3
	SchedulingInfo_si_Periodicity_Enum_rf128 aper.Enumerated = 4
	SchedulingInfo_si_Periodicity_Enum_rf256 aper.Enumerated = 5
	SchedulingInfo_si_Periodicity_Enum_rf512 aper.Enumerated = 6
)

type SchedulingInfo_si_Periodicity struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SchedulingInfo_si_Periodicity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SchedulingInfo_si_Periodicity", err)
	}
	return nil
}

func (ie *SchedulingInfo_si_Periodicity) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SchedulingInfo_si_Periodicity", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
