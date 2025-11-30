package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type AdditionalRACH_Config_r17 struct {
	Rach_ConfigCommon_r17 *RACH_ConfigCommon     `optional`
	MsgA_ConfigCommon_r17 *MsgA_ConfigCommon_r16 `optional`
}

func (ie *AdditionalRACH_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rach_ConfigCommon_r17 != nil, ie.MsgA_ConfigCommon_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rach_ConfigCommon_r17 != nil {
		if err = ie.Rach_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rach_ConfigCommon_r17", err)
		}
	}
	if ie.MsgA_ConfigCommon_r17 != nil {
		if err = ie.MsgA_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_ConfigCommon_r17", err)
		}
	}
	return nil
}

func (ie *AdditionalRACH_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var Rach_ConfigCommon_r17Present bool
	if Rach_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_ConfigCommon_r17Present bool
	if MsgA_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rach_ConfigCommon_r17Present {
		ie.Rach_ConfigCommon_r17 = new(RACH_ConfigCommon)
		if err = ie.Rach_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rach_ConfigCommon_r17", err)
		}
	}
	if MsgA_ConfigCommon_r17Present {
		ie.MsgA_ConfigCommon_r17 = new(MsgA_ConfigCommon_r16)
		if err = ie.MsgA_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_ConfigCommon_r17", err)
		}
	}
	return nil
}
