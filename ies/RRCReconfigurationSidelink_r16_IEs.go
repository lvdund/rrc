package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationSidelink_r16_IEs struct {
	Slrb_ConfigToAddModList_r16   []SLRB_Config_r16                                      `lb:1,ub:maxNrofSLRB_r16,optional`
	Slrb_ConfigToReleaseList_r16  []SLRB_PC5_ConfigIndex_r16                             `lb:1,ub:maxNrofSLRB_r16,optional`
	Sl_MeasConfig_r16             *SL_MeasConfig_r16                                     `optional,setuprelease`
	Sl_CSI_RS_Config_r16          *SL_CSI_RS_Config_r16                                  `optional,setuprelease`
	Sl_ResetConfig_r16            *RRCReconfigurationSidelink_r16_IEs_sl_ResetConfig_r16 `optional`
	Sl_LatencyBoundCSI_Report_r16 *int64                                                 `lb:3,ub:160,optional`
	LateNonCriticalExtension      *[]byte                                                `optional`
	NonCriticalExtension          *RRCReconfigurationSidelink_v1700_IEs                  `optional`
}

func (ie *RRCReconfigurationSidelink_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Slrb_ConfigToAddModList_r16) > 0, len(ie.Slrb_ConfigToReleaseList_r16) > 0, ie.Sl_MeasConfig_r16 != nil, ie.Sl_CSI_RS_Config_r16 != nil, ie.Sl_ResetConfig_r16 != nil, ie.Sl_LatencyBoundCSI_Report_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Slrb_ConfigToAddModList_r16) > 0 {
		tmp_Slrb_ConfigToAddModList_r16 := utils.NewSequence[*SLRB_Config_r16]([]*SLRB_Config_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.Slrb_ConfigToAddModList_r16 {
			tmp_Slrb_ConfigToAddModList_r16.Value = append(tmp_Slrb_ConfigToAddModList_r16.Value, &i)
		}
		if err = tmp_Slrb_ConfigToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Slrb_ConfigToAddModList_r16", err)
		}
	}
	if len(ie.Slrb_ConfigToReleaseList_r16) > 0 {
		tmp_Slrb_ConfigToReleaseList_r16 := utils.NewSequence[*SLRB_PC5_ConfigIndex_r16]([]*SLRB_PC5_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.Slrb_ConfigToReleaseList_r16 {
			tmp_Slrb_ConfigToReleaseList_r16.Value = append(tmp_Slrb_ConfigToReleaseList_r16.Value, &i)
		}
		if err = tmp_Slrb_ConfigToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Slrb_ConfigToReleaseList_r16", err)
		}
	}
	if ie.Sl_MeasConfig_r16 != nil {
		tmp_Sl_MeasConfig_r16 := utils.SetupRelease[*SL_MeasConfig_r16]{
			Setup: ie.Sl_MeasConfig_r16,
		}
		if err = tmp_Sl_MeasConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasConfig_r16", err)
		}
	}
	if ie.Sl_CSI_RS_Config_r16 != nil {
		tmp_Sl_CSI_RS_Config_r16 := utils.SetupRelease[*SL_CSI_RS_Config_r16]{
			Setup: ie.Sl_CSI_RS_Config_r16,
		}
		if err = tmp_Sl_CSI_RS_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CSI_RS_Config_r16", err)
		}
	}
	if ie.Sl_ResetConfig_r16 != nil {
		if err = ie.Sl_ResetConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ResetConfig_r16", err)
		}
	}
	if ie.Sl_LatencyBoundCSI_Report_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_LatencyBoundCSI_Report_r16, &uper.Constraint{Lb: 3, Ub: 160}, false); err != nil {
			return utils.WrapError("Encode Sl_LatencyBoundCSI_Report_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationSidelink_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Slrb_ConfigToAddModList_r16Present bool
	if Slrb_ConfigToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Slrb_ConfigToReleaseList_r16Present bool
	if Slrb_ConfigToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasConfig_r16Present bool
	if Sl_MeasConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CSI_RS_Config_r16Present bool
	if Sl_CSI_RS_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ResetConfig_r16Present bool
	if Sl_ResetConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_LatencyBoundCSI_Report_r16Present bool
	if Sl_LatencyBoundCSI_Report_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Slrb_ConfigToAddModList_r16Present {
		tmp_Slrb_ConfigToAddModList_r16 := utils.NewSequence[*SLRB_Config_r16]([]*SLRB_Config_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_Slrb_ConfigToAddModList_r16 := func() *SLRB_Config_r16 {
			return new(SLRB_Config_r16)
		}
		if err = tmp_Slrb_ConfigToAddModList_r16.Decode(r, fn_Slrb_ConfigToAddModList_r16); err != nil {
			return utils.WrapError("Decode Slrb_ConfigToAddModList_r16", err)
		}
		ie.Slrb_ConfigToAddModList_r16 = []SLRB_Config_r16{}
		for _, i := range tmp_Slrb_ConfigToAddModList_r16.Value {
			ie.Slrb_ConfigToAddModList_r16 = append(ie.Slrb_ConfigToAddModList_r16, *i)
		}
	}
	if Slrb_ConfigToReleaseList_r16Present {
		tmp_Slrb_ConfigToReleaseList_r16 := utils.NewSequence[*SLRB_PC5_ConfigIndex_r16]([]*SLRB_PC5_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_Slrb_ConfigToReleaseList_r16 := func() *SLRB_PC5_ConfigIndex_r16 {
			return new(SLRB_PC5_ConfigIndex_r16)
		}
		if err = tmp_Slrb_ConfigToReleaseList_r16.Decode(r, fn_Slrb_ConfigToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Slrb_ConfigToReleaseList_r16", err)
		}
		ie.Slrb_ConfigToReleaseList_r16 = []SLRB_PC5_ConfigIndex_r16{}
		for _, i := range tmp_Slrb_ConfigToReleaseList_r16.Value {
			ie.Slrb_ConfigToReleaseList_r16 = append(ie.Slrb_ConfigToReleaseList_r16, *i)
		}
	}
	if Sl_MeasConfig_r16Present {
		tmp_Sl_MeasConfig_r16 := utils.SetupRelease[*SL_MeasConfig_r16]{}
		if err = tmp_Sl_MeasConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasConfig_r16", err)
		}
		ie.Sl_MeasConfig_r16 = tmp_Sl_MeasConfig_r16.Setup
	}
	if Sl_CSI_RS_Config_r16Present {
		tmp_Sl_CSI_RS_Config_r16 := utils.SetupRelease[*SL_CSI_RS_Config_r16]{}
		if err = tmp_Sl_CSI_RS_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CSI_RS_Config_r16", err)
		}
		ie.Sl_CSI_RS_Config_r16 = tmp_Sl_CSI_RS_Config_r16.Setup
	}
	if Sl_ResetConfig_r16Present {
		ie.Sl_ResetConfig_r16 = new(RRCReconfigurationSidelink_r16_IEs_sl_ResetConfig_r16)
		if err = ie.Sl_ResetConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ResetConfig_r16", err)
		}
	}
	if Sl_LatencyBoundCSI_Report_r16Present {
		var tmp_int_Sl_LatencyBoundCSI_Report_r16 int64
		if tmp_int_Sl_LatencyBoundCSI_Report_r16, err = r.ReadInteger(&uper.Constraint{Lb: 3, Ub: 160}, false); err != nil {
			return utils.WrapError("Decode Sl_LatencyBoundCSI_Report_r16", err)
		}
		ie.Sl_LatencyBoundCSI_Report_r16 = &tmp_int_Sl_LatencyBoundCSI_Report_r16
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfigurationSidelink_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
