package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NeedForNCSG_NR_r17_gapIndication_r17_Enum_gap          aper.Enumerated = 0
	NeedForNCSG_NR_r17_gapIndication_r17_Enum_ncsg         aper.Enumerated = 1
	NeedForNCSG_NR_r17_gapIndication_r17_Enum_nogap_noncsg aper.Enumerated = 2
)

type NeedForNCSG_NR_r17_gapIndication_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *NeedForNCSG_NR_r17_gapIndication_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode NeedForNCSG_NR_r17_gapIndication_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_NR_r17_gapIndication_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode NeedForNCSG_NR_r17_gapIndication_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
