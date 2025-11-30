package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UCI_OnPUSCH_scaling_Enum_f0p5  aper.Enumerated = 0
	UCI_OnPUSCH_scaling_Enum_f0p65 aper.Enumerated = 1
	UCI_OnPUSCH_scaling_Enum_f0p8  aper.Enumerated = 2
	UCI_OnPUSCH_scaling_Enum_f1    aper.Enumerated = 3
)

type UCI_OnPUSCH_scaling struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *UCI_OnPUSCH_scaling) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode UCI_OnPUSCH_scaling", err)
	}
	return nil
}

func (ie *UCI_OnPUSCH_scaling) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode UCI_OnPUSCH_scaling", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
