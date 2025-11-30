package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_COT_Sharing_r16_cot_Sharing_r16 struct {
	Duration_r16              int64 `lb:1,ub:39,madatory`
	Offset_r16                int64 `lb:1,ub:39,madatory`
	ChannelAccessPriority_r16 int64 `lb:1,ub:4,madatory`
}

func (ie *CG_COT_Sharing_r16_cot_Sharing_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Duration_r16, &aper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("WriteInteger Duration_r16", err)
	}
	if err = w.WriteInteger(ie.Offset_r16, &aper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("WriteInteger Offset_r16", err)
	}
	if err = w.WriteInteger(ie.ChannelAccessPriority_r16, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger ChannelAccessPriority_r16", err)
	}
	return nil
}

func (ie *CG_COT_Sharing_r16_cot_Sharing_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Duration_r16 int64
	if tmp_int_Duration_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("ReadInteger Duration_r16", err)
	}
	ie.Duration_r16 = tmp_int_Duration_r16
	var tmp_int_Offset_r16 int64
	if tmp_int_Offset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("ReadInteger Offset_r16", err)
	}
	ie.Offset_r16 = tmp_int_Offset_r16
	var tmp_int_ChannelAccessPriority_r16 int64
	if tmp_int_ChannelAccessPriority_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger ChannelAccessPriority_r16", err)
	}
	ie.ChannelAccessPriority_r16 = tmp_int_ChannelAccessPriority_r16
	return nil
}
