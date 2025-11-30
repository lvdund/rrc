package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_v1610 struct {
	Idc_AssistanceConfig_r16                *IDC_AssistanceConfig_r16                               `optional,setuprelease`
	Drx_PreferenceConfig_r16                *DRX_PreferenceConfig_r16                               `optional,setuprelease`
	MaxBW_PreferenceConfig_r16              *MaxBW_PreferenceConfig_r16                             `optional,setuprelease`
	MaxCC_PreferenceConfig_r16              *MaxCC_PreferenceConfig_r16                             `optional,setuprelease`
	MaxMIMO_LayerPreferenceConfig_r16       *MaxMIMO_LayerPreferenceConfig_r16                      `optional,setuprelease`
	MinSchedulingOffsetPreferenceConfig_r16 *MinSchedulingOffsetPreferenceConfig_r16                `optional,setuprelease`
	ReleasePreferenceConfig_r16             *ReleasePreferenceConfig_r16                            `optional,setuprelease`
	ReferenceTimePreferenceReporting_r16    *OtherConfig_v1610_referenceTimePreferenceReporting_r16 `optional`
	BtNameList_r16                          *BT_NameList_r16                                        `optional,setuprelease`
	WlanNameList_r16                        *WLAN_NameList_r16                                      `optional,setuprelease`
	SensorNameList_r16                      *Sensor_NameList_r16                                    `optional,setuprelease`
	ObtainCommonLocation_r16                *OtherConfig_v1610_obtainCommonLocation_r16             `optional`
	Sl_AssistanceConfigNR_r16               *OtherConfig_v1610_sl_AssistanceConfigNR_r16            `optional`
}

func (ie *OtherConfig_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Idc_AssistanceConfig_r16 != nil, ie.Drx_PreferenceConfig_r16 != nil, ie.MaxBW_PreferenceConfig_r16 != nil, ie.MaxCC_PreferenceConfig_r16 != nil, ie.MaxMIMO_LayerPreferenceConfig_r16 != nil, ie.MinSchedulingOffsetPreferenceConfig_r16 != nil, ie.ReleasePreferenceConfig_r16 != nil, ie.ReferenceTimePreferenceReporting_r16 != nil, ie.BtNameList_r16 != nil, ie.WlanNameList_r16 != nil, ie.SensorNameList_r16 != nil, ie.ObtainCommonLocation_r16 != nil, ie.Sl_AssistanceConfigNR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Idc_AssistanceConfig_r16 != nil {
		tmp_Idc_AssistanceConfig_r16 := utils.SetupRelease[*IDC_AssistanceConfig_r16]{
			Setup: ie.Idc_AssistanceConfig_r16,
		}
		if err = tmp_Idc_AssistanceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Idc_AssistanceConfig_r16", err)
		}
	}
	if ie.Drx_PreferenceConfig_r16 != nil {
		tmp_Drx_PreferenceConfig_r16 := utils.SetupRelease[*DRX_PreferenceConfig_r16]{
			Setup: ie.Drx_PreferenceConfig_r16,
		}
		if err = tmp_Drx_PreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_PreferenceConfig_r16", err)
		}
	}
	if ie.MaxBW_PreferenceConfig_r16 != nil {
		tmp_MaxBW_PreferenceConfig_r16 := utils.SetupRelease[*MaxBW_PreferenceConfig_r16]{
			Setup: ie.MaxBW_PreferenceConfig_r16,
		}
		if err = tmp_MaxBW_PreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBW_PreferenceConfig_r16", err)
		}
	}
	if ie.MaxCC_PreferenceConfig_r16 != nil {
		tmp_MaxCC_PreferenceConfig_r16 := utils.SetupRelease[*MaxCC_PreferenceConfig_r16]{
			Setup: ie.MaxCC_PreferenceConfig_r16,
		}
		if err = tmp_MaxCC_PreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxCC_PreferenceConfig_r16", err)
		}
	}
	if ie.MaxMIMO_LayerPreferenceConfig_r16 != nil {
		tmp_MaxMIMO_LayerPreferenceConfig_r16 := utils.SetupRelease[*MaxMIMO_LayerPreferenceConfig_r16]{
			Setup: ie.MaxMIMO_LayerPreferenceConfig_r16,
		}
		if err = tmp_MaxMIMO_LayerPreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxMIMO_LayerPreferenceConfig_r16", err)
		}
	}
	if ie.MinSchedulingOffsetPreferenceConfig_r16 != nil {
		tmp_MinSchedulingOffsetPreferenceConfig_r16 := utils.SetupRelease[*MinSchedulingOffsetPreferenceConfig_r16]{
			Setup: ie.MinSchedulingOffsetPreferenceConfig_r16,
		}
		if err = tmp_MinSchedulingOffsetPreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MinSchedulingOffsetPreferenceConfig_r16", err)
		}
	}
	if ie.ReleasePreferenceConfig_r16 != nil {
		tmp_ReleasePreferenceConfig_r16 := utils.SetupRelease[*ReleasePreferenceConfig_r16]{
			Setup: ie.ReleasePreferenceConfig_r16,
		}
		if err = tmp_ReleasePreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReleasePreferenceConfig_r16", err)
		}
	}
	if ie.ReferenceTimePreferenceReporting_r16 != nil {
		if err = ie.ReferenceTimePreferenceReporting_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReferenceTimePreferenceReporting_r16", err)
		}
	}
	if ie.BtNameList_r16 != nil {
		tmp_BtNameList_r16 := utils.SetupRelease[*BT_NameList_r16]{
			Setup: ie.BtNameList_r16,
		}
		if err = tmp_BtNameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode BtNameList_r16", err)
		}
	}
	if ie.WlanNameList_r16 != nil {
		tmp_WlanNameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{
			Setup: ie.WlanNameList_r16,
		}
		if err = tmp_WlanNameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode WlanNameList_r16", err)
		}
	}
	if ie.SensorNameList_r16 != nil {
		tmp_SensorNameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{
			Setup: ie.SensorNameList_r16,
		}
		if err = tmp_SensorNameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SensorNameList_r16", err)
		}
	}
	if ie.ObtainCommonLocation_r16 != nil {
		if err = ie.ObtainCommonLocation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ObtainCommonLocation_r16", err)
		}
	}
	if ie.Sl_AssistanceConfigNR_r16 != nil {
		if err = ie.Sl_AssistanceConfigNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_AssistanceConfigNR_r16", err)
		}
	}
	return nil
}

func (ie *OtherConfig_v1610) Decode(r *aper.AperReader) error {
	var err error
	var Idc_AssistanceConfig_r16Present bool
	if Idc_AssistanceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_PreferenceConfig_r16Present bool
	if Drx_PreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxBW_PreferenceConfig_r16Present bool
	if MaxBW_PreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCC_PreferenceConfig_r16Present bool
	if MaxCC_PreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxMIMO_LayerPreferenceConfig_r16Present bool
	if MaxMIMO_LayerPreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MinSchedulingOffsetPreferenceConfig_r16Present bool
	if MinSchedulingOffsetPreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReleasePreferenceConfig_r16Present bool
	if ReleasePreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReferenceTimePreferenceReporting_r16Present bool
	if ReferenceTimePreferenceReporting_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BtNameList_r16Present bool
	if BtNameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var WlanNameList_r16Present bool
	if WlanNameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SensorNameList_r16Present bool
	if SensorNameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ObtainCommonLocation_r16Present bool
	if ObtainCommonLocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_AssistanceConfigNR_r16Present bool
	if Sl_AssistanceConfigNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Idc_AssistanceConfig_r16Present {
		tmp_Idc_AssistanceConfig_r16 := utils.SetupRelease[*IDC_AssistanceConfig_r16]{}
		if err = tmp_Idc_AssistanceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Idc_AssistanceConfig_r16", err)
		}
		ie.Idc_AssistanceConfig_r16 = tmp_Idc_AssistanceConfig_r16.Setup
	}
	if Drx_PreferenceConfig_r16Present {
		tmp_Drx_PreferenceConfig_r16 := utils.SetupRelease[*DRX_PreferenceConfig_r16]{}
		if err = tmp_Drx_PreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_PreferenceConfig_r16", err)
		}
		ie.Drx_PreferenceConfig_r16 = tmp_Drx_PreferenceConfig_r16.Setup
	}
	if MaxBW_PreferenceConfig_r16Present {
		tmp_MaxBW_PreferenceConfig_r16 := utils.SetupRelease[*MaxBW_PreferenceConfig_r16]{}
		if err = tmp_MaxBW_PreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBW_PreferenceConfig_r16", err)
		}
		ie.MaxBW_PreferenceConfig_r16 = tmp_MaxBW_PreferenceConfig_r16.Setup
	}
	if MaxCC_PreferenceConfig_r16Present {
		tmp_MaxCC_PreferenceConfig_r16 := utils.SetupRelease[*MaxCC_PreferenceConfig_r16]{}
		if err = tmp_MaxCC_PreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxCC_PreferenceConfig_r16", err)
		}
		ie.MaxCC_PreferenceConfig_r16 = tmp_MaxCC_PreferenceConfig_r16.Setup
	}
	if MaxMIMO_LayerPreferenceConfig_r16Present {
		tmp_MaxMIMO_LayerPreferenceConfig_r16 := utils.SetupRelease[*MaxMIMO_LayerPreferenceConfig_r16]{}
		if err = tmp_MaxMIMO_LayerPreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxMIMO_LayerPreferenceConfig_r16", err)
		}
		ie.MaxMIMO_LayerPreferenceConfig_r16 = tmp_MaxMIMO_LayerPreferenceConfig_r16.Setup
	}
	if MinSchedulingOffsetPreferenceConfig_r16Present {
		tmp_MinSchedulingOffsetPreferenceConfig_r16 := utils.SetupRelease[*MinSchedulingOffsetPreferenceConfig_r16]{}
		if err = tmp_MinSchedulingOffsetPreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MinSchedulingOffsetPreferenceConfig_r16", err)
		}
		ie.MinSchedulingOffsetPreferenceConfig_r16 = tmp_MinSchedulingOffsetPreferenceConfig_r16.Setup
	}
	if ReleasePreferenceConfig_r16Present {
		tmp_ReleasePreferenceConfig_r16 := utils.SetupRelease[*ReleasePreferenceConfig_r16]{}
		if err = tmp_ReleasePreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReleasePreferenceConfig_r16", err)
		}
		ie.ReleasePreferenceConfig_r16 = tmp_ReleasePreferenceConfig_r16.Setup
	}
	if ReferenceTimePreferenceReporting_r16Present {
		ie.ReferenceTimePreferenceReporting_r16 = new(OtherConfig_v1610_referenceTimePreferenceReporting_r16)
		if err = ie.ReferenceTimePreferenceReporting_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReferenceTimePreferenceReporting_r16", err)
		}
	}
	if BtNameList_r16Present {
		tmp_BtNameList_r16 := utils.SetupRelease[*BT_NameList_r16]{}
		if err = tmp_BtNameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode BtNameList_r16", err)
		}
		ie.BtNameList_r16 = tmp_BtNameList_r16.Setup
	}
	if WlanNameList_r16Present {
		tmp_WlanNameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{}
		if err = tmp_WlanNameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode WlanNameList_r16", err)
		}
		ie.WlanNameList_r16 = tmp_WlanNameList_r16.Setup
	}
	if SensorNameList_r16Present {
		tmp_SensorNameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{}
		if err = tmp_SensorNameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SensorNameList_r16", err)
		}
		ie.SensorNameList_r16 = tmp_SensorNameList_r16.Setup
	}
	if ObtainCommonLocation_r16Present {
		ie.ObtainCommonLocation_r16 = new(OtherConfig_v1610_obtainCommonLocation_r16)
		if err = ie.ObtainCommonLocation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ObtainCommonLocation_r16", err)
		}
	}
	if Sl_AssistanceConfigNR_r16Present {
		ie.Sl_AssistanceConfigNR_r16 = new(OtherConfig_v1610_sl_AssistanceConfigNR_r16)
		if err = ie.Sl_AssistanceConfigNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_AssistanceConfigNR_r16", err)
		}
	}
	return nil
}
