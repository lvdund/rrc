package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_COT_Sharing_r17_cot_Sharing_r17 struct {
	Duration_r17 int64 `lb:1,ub:319,madatory`
	Offset_r17   int64 `lb:1,ub:319,madatory`
}

func (ie *CG_COT_Sharing_r17_cot_Sharing_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Duration_r17, &aper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("WriteInteger Duration_r17", err)
	}
	if err = w.WriteInteger(ie.Offset_r17, &aper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("WriteInteger Offset_r17", err)
	}
	return nil
}

func (ie *CG_COT_Sharing_r17_cot_Sharing_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Duration_r17 int64
	if tmp_int_Duration_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("ReadInteger Duration_r17", err)
	}
	ie.Duration_r17 = tmp_int_Duration_r17
	var tmp_int_Offset_r17 int64
	if tmp_int_Offset_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("ReadInteger Offset_r17", err)
	}
	ie.Offset_r17 = tmp_int_Offset_r17
	return nil
}
