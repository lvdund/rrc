package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_LayersUL_Enum_oneLayer   aper.Enumerated = 0
	MIMO_LayersUL_Enum_twoLayers  aper.Enumerated = 1
	MIMO_LayersUL_Enum_fourLayers aper.Enumerated = 2
)

type MIMO_LayersUL struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MIMO_LayersUL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MIMO_LayersUL", err)
	}
	return nil
}

func (ie *MIMO_LayersUL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MIMO_LayersUL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
