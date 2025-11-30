package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PeriodicRNAU_TimerValue_Enum_min5   aper.Enumerated = 0
	PeriodicRNAU_TimerValue_Enum_min10  aper.Enumerated = 1
	PeriodicRNAU_TimerValue_Enum_min20  aper.Enumerated = 2
	PeriodicRNAU_TimerValue_Enum_min30  aper.Enumerated = 3
	PeriodicRNAU_TimerValue_Enum_min60  aper.Enumerated = 4
	PeriodicRNAU_TimerValue_Enum_min120 aper.Enumerated = 5
	PeriodicRNAU_TimerValue_Enum_min360 aper.Enumerated = 6
	PeriodicRNAU_TimerValue_Enum_min720 aper.Enumerated = 7
)

type PeriodicRNAU_TimerValue struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PeriodicRNAU_TimerValue) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PeriodicRNAU_TimerValue", err)
	}
	return nil
}

func (ie *PeriodicRNAU_TimerValue) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PeriodicRNAU_TimerValue", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
