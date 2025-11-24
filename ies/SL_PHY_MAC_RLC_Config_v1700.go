package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PHY_MAC_RLC_Config_v1700 struct {
	Sl_DRX_Config_r17               *SL_DRX_Config_r17         `optional`
	Sl_RLC_ChannelToReleaseList_r17 []SL_RLC_ChannelID_r17     `lb:1,ub:maxSL_LCID_r16,optional`
	Sl_RLC_ChannelToAddModList_r17  []SL_RLC_ChannelConfig_r17 `lb:1,ub:maxSL_LCID_r16,optional`
}

func (ie *SL_PHY_MAC_RLC_Config_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_DRX_Config_r17 != nil, len(ie.Sl_RLC_ChannelToReleaseList_r17) > 0, len(ie.Sl_RLC_ChannelToAddModList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_DRX_Config_r17 != nil {
		if err = ie.Sl_DRX_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_Config_r17", err)
		}
	}
	if len(ie.Sl_RLC_ChannelToReleaseList_r17) > 0 {
		tmp_Sl_RLC_ChannelToReleaseList_r17 := utils.NewSequence[*SL_RLC_ChannelID_r17]([]*SL_RLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_ChannelToReleaseList_r17 {
			tmp_Sl_RLC_ChannelToReleaseList_r17.Value = append(tmp_Sl_RLC_ChannelToReleaseList_r17.Value, &i)
		}
		if err = tmp_Sl_RLC_ChannelToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_ChannelToReleaseList_r17", err)
		}
	}
	if len(ie.Sl_RLC_ChannelToAddModList_r17) > 0 {
		tmp_Sl_RLC_ChannelToAddModList_r17 := utils.NewSequence[*SL_RLC_ChannelConfig_r17]([]*SL_RLC_ChannelConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_ChannelToAddModList_r17 {
			tmp_Sl_RLC_ChannelToAddModList_r17.Value = append(tmp_Sl_RLC_ChannelToAddModList_r17.Value, &i)
		}
		if err = tmp_Sl_RLC_ChannelToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_ChannelToAddModList_r17", err)
		}
	}
	return nil
}

func (ie *SL_PHY_MAC_RLC_Config_v1700) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DRX_Config_r17Present bool
	if Sl_DRX_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_ChannelToReleaseList_r17Present bool
	if Sl_RLC_ChannelToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_ChannelToAddModList_r17Present bool
	if Sl_RLC_ChannelToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DRX_Config_r17Present {
		ie.Sl_DRX_Config_r17 = new(SL_DRX_Config_r17)
		if err = ie.Sl_DRX_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DRX_Config_r17", err)
		}
	}
	if Sl_RLC_ChannelToReleaseList_r17Present {
		tmp_Sl_RLC_ChannelToReleaseList_r17 := utils.NewSequence[*SL_RLC_ChannelID_r17]([]*SL_RLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_ChannelToReleaseList_r17 := func() *SL_RLC_ChannelID_r17 {
			return new(SL_RLC_ChannelID_r17)
		}
		if err = tmp_Sl_RLC_ChannelToReleaseList_r17.Decode(r, fn_Sl_RLC_ChannelToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Sl_RLC_ChannelToReleaseList_r17", err)
		}
		ie.Sl_RLC_ChannelToReleaseList_r17 = []SL_RLC_ChannelID_r17{}
		for _, i := range tmp_Sl_RLC_ChannelToReleaseList_r17.Value {
			ie.Sl_RLC_ChannelToReleaseList_r17 = append(ie.Sl_RLC_ChannelToReleaseList_r17, *i)
		}
	}
	if Sl_RLC_ChannelToAddModList_r17Present {
		tmp_Sl_RLC_ChannelToAddModList_r17 := utils.NewSequence[*SL_RLC_ChannelConfig_r17]([]*SL_RLC_ChannelConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_ChannelToAddModList_r17 := func() *SL_RLC_ChannelConfig_r17 {
			return new(SL_RLC_ChannelConfig_r17)
		}
		if err = tmp_Sl_RLC_ChannelToAddModList_r17.Decode(r, fn_Sl_RLC_ChannelToAddModList_r17); err != nil {
			return utils.WrapError("Decode Sl_RLC_ChannelToAddModList_r17", err)
		}
		ie.Sl_RLC_ChannelToAddModList_r17 = []SL_RLC_ChannelConfig_r17{}
		for _, i := range tmp_Sl_RLC_ChannelToAddModList_r17.Value {
			ie.Sl_RLC_ChannelToAddModList_r17 = append(ie.Sl_RLC_ChannelToAddModList_r17, *i)
		}
	}
	return nil
}
