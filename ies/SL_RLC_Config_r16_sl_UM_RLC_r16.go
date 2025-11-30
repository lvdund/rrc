package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_Config_r16_sl_UM_RLC_r16 struct {
	Sl_SN_FieldLengthUM_r16 *SN_FieldLengthUM `optional`
}

func (ie *SL_RLC_Config_r16_sl_UM_RLC_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_SN_FieldLengthUM_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_SN_FieldLengthUM_r16 != nil {
		if err = ie.Sl_SN_FieldLengthUM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SN_FieldLengthUM_r16", err)
		}
	}
	return nil
}

func (ie *SL_RLC_Config_r16_sl_UM_RLC_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_SN_FieldLengthUM_r16Present bool
	if Sl_SN_FieldLengthUM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_SN_FieldLengthUM_r16Present {
		ie.Sl_SN_FieldLengthUM_r16 = new(SN_FieldLengthUM)
		if err = ie.Sl_SN_FieldLengthUM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SN_FieldLengthUM_r16", err)
		}
	}
	return nil
}
