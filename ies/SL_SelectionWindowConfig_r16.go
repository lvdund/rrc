package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_SelectionWindowConfig_r16 struct {
	Sl_Priority_r16        int64                                               `lb:1,ub:8,madatory`
	Sl_SelectionWindow_r16 SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16 `madatory`
}

func (ie *SL_SelectionWindowConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Sl_Priority_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_Priority_r16", err)
	}
	if err = ie.Sl_SelectionWindow_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_SelectionWindow_r16", err)
	}
	return nil
}

func (ie *SL_SelectionWindowConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Sl_Priority_r16 int64
	if tmp_int_Sl_Priority_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_Priority_r16", err)
	}
	ie.Sl_Priority_r16 = tmp_int_Sl_Priority_r16
	if err = ie.Sl_SelectionWindow_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_SelectionWindow_r16", err)
	}
	return nil
}
