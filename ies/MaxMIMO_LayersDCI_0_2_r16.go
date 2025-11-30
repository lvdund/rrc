package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayersDCI_0_2_r16 struct {
	Value uint64 `lb:1,ub:4,madatory`
}

func (ie *MaxMIMO_LayersDCI_0_2_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MaxMIMO_LayersDCI_0_2_r16", err)
	}
	return nil
}

func (ie *MaxMIMO_LayersDCI_0_2_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MaxMIMO_LayersDCI_0_2_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
