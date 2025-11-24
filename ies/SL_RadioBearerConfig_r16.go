package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RadioBearerConfig_r16 struct {
	Slrb_Uu_ConfigIndex_r16 SLRB_Uu_ConfigIndex_r16                     `madatory`
	Sl_SDAP_Config_r16      *SL_SDAP_Config_r16                         `optional`
	Sl_PDCP_Config_r16      *SL_PDCP_Config_r16                         `optional`
	Sl_TransRange_r16       *SL_RadioBearerConfig_r16_sl_TransRange_r16 `optional`
}

func (ie *SL_RadioBearerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_SDAP_Config_r16 != nil, ie.Sl_PDCP_Config_r16 != nil, ie.Sl_TransRange_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Slrb_Uu_ConfigIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Slrb_Uu_ConfigIndex_r16", err)
	}
	if ie.Sl_SDAP_Config_r16 != nil {
		if err = ie.Sl_SDAP_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SDAP_Config_r16", err)
		}
	}
	if ie.Sl_PDCP_Config_r16 != nil {
		if err = ie.Sl_PDCP_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PDCP_Config_r16", err)
		}
	}
	if ie.Sl_TransRange_r16 != nil {
		if err = ie.Sl_TransRange_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TransRange_r16", err)
		}
	}
	return nil
}

func (ie *SL_RadioBearerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_SDAP_Config_r16Present bool
	if Sl_SDAP_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PDCP_Config_r16Present bool
	if Sl_PDCP_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TransRange_r16Present bool
	if Sl_TransRange_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Slrb_Uu_ConfigIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Slrb_Uu_ConfigIndex_r16", err)
	}
	if Sl_SDAP_Config_r16Present {
		ie.Sl_SDAP_Config_r16 = new(SL_SDAP_Config_r16)
		if err = ie.Sl_SDAP_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SDAP_Config_r16", err)
		}
	}
	if Sl_PDCP_Config_r16Present {
		ie.Sl_PDCP_Config_r16 = new(SL_PDCP_Config_r16)
		if err = ie.Sl_PDCP_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PDCP_Config_r16", err)
		}
	}
	if Sl_TransRange_r16Present {
		ie.Sl_TransRange_r16 = new(SL_RadioBearerConfig_r16_sl_TransRange_r16)
		if err = ie.Sl_TransRange_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TransRange_r16", err)
		}
	}
	return nil
}
