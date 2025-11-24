package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationSidelink_v1700_IEs struct {
	Sl_DRX_ConfigUC_PC5_r17            *SL_DRX_ConfigUC_r17           `optional,setuprelease`
	Sl_LatencyBoundIUC_Report_r17      *SL_LatencyBoundIUC_Report_r17 `optional,setuprelease`
	Sl_RLC_ChannelToReleaseListPC5_r17 []SL_RLC_ChannelID_r17         `lb:1,ub:maxSL_LCID_r16,optional`
	Sl_RLC_ChannelToAddModListPC5_r17  []SL_RLC_ChannelConfigPC5_r17  `lb:1,ub:maxSL_LCID_r16,optional`
	NonCriticalExtension               interface{}                    `optional`
}

func (ie *RRCReconfigurationSidelink_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_DRX_ConfigUC_PC5_r17 != nil, ie.Sl_LatencyBoundIUC_Report_r17 != nil, len(ie.Sl_RLC_ChannelToReleaseListPC5_r17) > 0, len(ie.Sl_RLC_ChannelToAddModListPC5_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_DRX_ConfigUC_PC5_r17 != nil {
		tmp_Sl_DRX_ConfigUC_PC5_r17 := utils.SetupRelease[*SL_DRX_ConfigUC_r17]{
			Setup: ie.Sl_DRX_ConfigUC_PC5_r17,
		}
		if err = tmp_Sl_DRX_ConfigUC_PC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_ConfigUC_PC5_r17", err)
		}
	}
	if ie.Sl_LatencyBoundIUC_Report_r17 != nil {
		tmp_Sl_LatencyBoundIUC_Report_r17 := utils.SetupRelease[*SL_LatencyBoundIUC_Report_r17]{
			Setup: ie.Sl_LatencyBoundIUC_Report_r17,
		}
		if err = tmp_Sl_LatencyBoundIUC_Report_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_LatencyBoundIUC_Report_r17", err)
		}
	}
	if len(ie.Sl_RLC_ChannelToReleaseListPC5_r17) > 0 {
		tmp_Sl_RLC_ChannelToReleaseListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelID_r17]([]*SL_RLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_ChannelToReleaseListPC5_r17 {
			tmp_Sl_RLC_ChannelToReleaseListPC5_r17.Value = append(tmp_Sl_RLC_ChannelToReleaseListPC5_r17.Value, &i)
		}
		if err = tmp_Sl_RLC_ChannelToReleaseListPC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_ChannelToReleaseListPC5_r17", err)
		}
	}
	if len(ie.Sl_RLC_ChannelToAddModListPC5_r17) > 0 {
		tmp_Sl_RLC_ChannelToAddModListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelConfigPC5_r17]([]*SL_RLC_ChannelConfigPC5_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_ChannelToAddModListPC5_r17 {
			tmp_Sl_RLC_ChannelToAddModListPC5_r17.Value = append(tmp_Sl_RLC_ChannelToAddModListPC5_r17.Value, &i)
		}
		if err = tmp_Sl_RLC_ChannelToAddModListPC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_ChannelToAddModListPC5_r17", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationSidelink_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DRX_ConfigUC_PC5_r17Present bool
	if Sl_DRX_ConfigUC_PC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_LatencyBoundIUC_Report_r17Present bool
	if Sl_LatencyBoundIUC_Report_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_ChannelToReleaseListPC5_r17Present bool
	if Sl_RLC_ChannelToReleaseListPC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_ChannelToAddModListPC5_r17Present bool
	if Sl_RLC_ChannelToAddModListPC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DRX_ConfigUC_PC5_r17Present {
		tmp_Sl_DRX_ConfigUC_PC5_r17 := utils.SetupRelease[*SL_DRX_ConfigUC_r17]{}
		if err = tmp_Sl_DRX_ConfigUC_PC5_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DRX_ConfigUC_PC5_r17", err)
		}
		ie.Sl_DRX_ConfigUC_PC5_r17 = tmp_Sl_DRX_ConfigUC_PC5_r17.Setup
	}
	if Sl_LatencyBoundIUC_Report_r17Present {
		tmp_Sl_LatencyBoundIUC_Report_r17 := utils.SetupRelease[*SL_LatencyBoundIUC_Report_r17]{}
		if err = tmp_Sl_LatencyBoundIUC_Report_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_LatencyBoundIUC_Report_r17", err)
		}
		ie.Sl_LatencyBoundIUC_Report_r17 = tmp_Sl_LatencyBoundIUC_Report_r17.Setup
	}
	if Sl_RLC_ChannelToReleaseListPC5_r17Present {
		tmp_Sl_RLC_ChannelToReleaseListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelID_r17]([]*SL_RLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_ChannelToReleaseListPC5_r17 := func() *SL_RLC_ChannelID_r17 {
			return new(SL_RLC_ChannelID_r17)
		}
		if err = tmp_Sl_RLC_ChannelToReleaseListPC5_r17.Decode(r, fn_Sl_RLC_ChannelToReleaseListPC5_r17); err != nil {
			return utils.WrapError("Decode Sl_RLC_ChannelToReleaseListPC5_r17", err)
		}
		ie.Sl_RLC_ChannelToReleaseListPC5_r17 = []SL_RLC_ChannelID_r17{}
		for _, i := range tmp_Sl_RLC_ChannelToReleaseListPC5_r17.Value {
			ie.Sl_RLC_ChannelToReleaseListPC5_r17 = append(ie.Sl_RLC_ChannelToReleaseListPC5_r17, *i)
		}
	}
	if Sl_RLC_ChannelToAddModListPC5_r17Present {
		tmp_Sl_RLC_ChannelToAddModListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelConfigPC5_r17]([]*SL_RLC_ChannelConfigPC5_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_ChannelToAddModListPC5_r17 := func() *SL_RLC_ChannelConfigPC5_r17 {
			return new(SL_RLC_ChannelConfigPC5_r17)
		}
		if err = tmp_Sl_RLC_ChannelToAddModListPC5_r17.Decode(r, fn_Sl_RLC_ChannelToAddModListPC5_r17); err != nil {
			return utils.WrapError("Decode Sl_RLC_ChannelToAddModListPC5_r17", err)
		}
		ie.Sl_RLC_ChannelToAddModListPC5_r17 = []SL_RLC_ChannelConfigPC5_r17{}
		for _, i := range tmp_Sl_RLC_ChannelToAddModListPC5_r17.Value {
			ie.Sl_RLC_ChannelToAddModListPC5_r17 = append(ie.Sl_RLC_ChannelToAddModListPC5_r17, *i)
		}
	}
	return nil
}
