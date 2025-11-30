package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfigIndexCG_r16 struct {
	Value uint64 `lb:0,ub:maxNrofCG_SL_1_r16,madatory`
}

func (ie *SL_ConfigIndexCG_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofCG_SL_1_r16}, false); err != nil {
		return utils.WrapError("Encode SL_ConfigIndexCG_r16", err)
	}
	return nil
}

func (ie *SL_ConfigIndexCG_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofCG_SL_1_r16}, false); err != nil {
		return utils.WrapError("Decode SL_ConfigIndexCG_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
