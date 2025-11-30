package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610_interlaceAllocation_r16 struct {
	Rb_SetIndex_r16 int64                                                          `lb:0,ub:4,madatory`
	Interlace0_r16  PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16 `lb:0,ub:9,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Rb_SetIndex_r16, &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger Rb_SetIndex_r16", err)
	}
	if err = ie.Interlace0_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Interlace0_r16", err)
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Rb_SetIndex_r16 int64
	if tmp_int_Rb_SetIndex_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger Rb_SetIndex_r16", err)
	}
	ie.Rb_SetIndex_r16 = tmp_int_Rb_SetIndex_r16
	if err = ie.Interlace0_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Interlace0_r16", err)
	}
	return nil
}
