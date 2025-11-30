package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB0 aper.Enumerated = 0
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB2 aper.Enumerated = 1
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB4 aper.Enumerated = 2
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB6 aper.Enumerated = 3
)

type RA_Prioritization_powerRampingStepHighPriority struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RA_Prioritization_powerRampingStepHighPriority) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RA_Prioritization_powerRampingStepHighPriority", err)
	}
	return nil
}

func (ie *RA_Prioritization_powerRampingStepHighPriority) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RA_Prioritization_powerRampingStepHighPriority", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
