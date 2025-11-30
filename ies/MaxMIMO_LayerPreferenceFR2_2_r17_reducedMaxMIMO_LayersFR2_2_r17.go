package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17 struct {
	ReducedMIMO_LayersFR2_2_DL_r17 int64 `lb:1,ub:8,madatory`
	ReducedMIMO_LayersFR2_2_UL_r17 int64 `lb:1,ub:4,madatory`
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.ReducedMIMO_LayersFR2_2_DL_r17, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger ReducedMIMO_LayersFR2_2_DL_r17", err)
	}
	if err = w.WriteInteger(ie.ReducedMIMO_LayersFR2_2_UL_r17, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger ReducedMIMO_LayersFR2_2_UL_r17", err)
	}
	return nil
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_ReducedMIMO_LayersFR2_2_DL_r17 int64
	if tmp_int_ReducedMIMO_LayersFR2_2_DL_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger ReducedMIMO_LayersFR2_2_DL_r17", err)
	}
	ie.ReducedMIMO_LayersFR2_2_DL_r17 = tmp_int_ReducedMIMO_LayersFR2_2_DL_r17
	var tmp_int_ReducedMIMO_LayersFR2_2_UL_r17 int64
	if tmp_int_ReducedMIMO_LayersFR2_2_UL_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger ReducedMIMO_LayersFR2_2_UL_r17", err)
	}
	ie.ReducedMIMO_LayersFR2_2_UL_r17 = tmp_int_ReducedMIMO_LayersFR2_2_UL_r17
	return nil
}
