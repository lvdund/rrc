package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_LayersDL_Enum_twoLayers   aper.Enumerated = 0
	MIMO_LayersDL_Enum_fourLayers  aper.Enumerated = 1
	MIMO_LayersDL_Enum_eightLayers aper.Enumerated = 2
)

type MIMO_LayersDL struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MIMO_LayersDL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MIMO_LayersDL", err)
	}
	return nil
}

func (ie *MIMO_LayersDL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MIMO_LayersDL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
