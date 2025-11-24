package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ZoneConfigMCR_r16 struct {
	Sl_ZoneConfigMCR_Index_r16 int64                                   `lb:0,ub:15,madatory`
	Sl_TransRange_r16          *SL_ZoneConfigMCR_r16_sl_TransRange_r16 `optional`
	Sl_ZoneConfig_r16          *SL_ZoneConfig_r16                      `optional`
}

func (ie *SL_ZoneConfigMCR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_TransRange_r16 != nil, ie.Sl_ZoneConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Sl_ZoneConfigMCR_Index_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_ZoneConfigMCR_Index_r16", err)
	}
	if ie.Sl_TransRange_r16 != nil {
		if err = ie.Sl_TransRange_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TransRange_r16", err)
		}
	}
	if ie.Sl_ZoneConfig_r16 != nil {
		if err = ie.Sl_ZoneConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ZoneConfig_r16", err)
		}
	}
	return nil
}

func (ie *SL_ZoneConfigMCR_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_TransRange_r16Present bool
	if Sl_TransRange_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ZoneConfig_r16Present bool
	if Sl_ZoneConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Sl_ZoneConfigMCR_Index_r16 int64
	if tmp_int_Sl_ZoneConfigMCR_Index_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_ZoneConfigMCR_Index_r16", err)
	}
	ie.Sl_ZoneConfigMCR_Index_r16 = tmp_int_Sl_ZoneConfigMCR_Index_r16
	if Sl_TransRange_r16Present {
		ie.Sl_TransRange_r16 = new(SL_ZoneConfigMCR_r16_sl_TransRange_r16)
		if err = ie.Sl_TransRange_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TransRange_r16", err)
		}
	}
	if Sl_ZoneConfig_r16Present {
		ie.Sl_ZoneConfig_r16 = new(SL_ZoneConfig_r16)
		if err = ie.Sl_ZoneConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ZoneConfig_r16", err)
		}
	}
	return nil
}
