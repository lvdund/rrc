package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResume_v1700_IEs struct {
	Sl_ConfigDedicatedNR_r17       *SL_ConfigDedicatedNR_r16          `optional,setuprelease`
	Sl_L2RemoteUE_Config_r17       *SL_L2RemoteUE_Config_r17          `optional,setuprelease`
	NeedForGapNCSG_ConfigNR_r17    *NeedForGapNCSG_ConfigNR_r17       `optional,setuprelease`
	NeedForGapNCSG_ConfigEUTRA_r17 *NeedForGapNCSG_ConfigEUTRA_r17    `optional,setuprelease`
	Scg_State_r17                  *RRCResume_v1700_IEs_scg_State_r17 `optional`
	AppLayerMeasConfig_r17         *AppLayerMeasConfig_r17            `optional`
	NonCriticalExtension           interface{}                        `optional`
}

func (ie *RRCResume_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ConfigDedicatedNR_r17 != nil, ie.Sl_L2RemoteUE_Config_r17 != nil, ie.NeedForGapNCSG_ConfigNR_r17 != nil, ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil, ie.Scg_State_r17 != nil, ie.AppLayerMeasConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_ConfigDedicatedNR_r17 != nil {
		tmp_Sl_ConfigDedicatedNR_r17 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{
			Setup: ie.Sl_ConfigDedicatedNR_r17,
		}
		if err = tmp_Sl_ConfigDedicatedNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ConfigDedicatedNR_r17", err)
		}
	}
	if ie.Sl_L2RemoteUE_Config_r17 != nil {
		tmp_Sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{
			Setup: ie.Sl_L2RemoteUE_Config_r17,
		}
		if err = tmp_Sl_L2RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_L2RemoteUE_Config_r17", err)
		}
	}
	if ie.NeedForGapNCSG_ConfigNR_r17 != nil {
		tmp_NeedForGapNCSG_ConfigNR_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigNR_r17]{
			Setup: ie.NeedForGapNCSG_ConfigNR_r17,
		}
		if err = tmp_NeedForGapNCSG_ConfigNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapNCSG_ConfigNR_r17", err)
		}
	}
	if ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil {
		tmp_NeedForGapNCSG_ConfigEUTRA_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigEUTRA_r17]{
			Setup: ie.NeedForGapNCSG_ConfigEUTRA_r17,
		}
		if err = tmp_NeedForGapNCSG_ConfigEUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapNCSG_ConfigEUTRA_r17", err)
		}
	}
	if ie.Scg_State_r17 != nil {
		if err = ie.Scg_State_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_State_r17", err)
		}
	}
	if ie.AppLayerMeasConfig_r17 != nil {
		if err = ie.AppLayerMeasConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AppLayerMeasConfig_r17", err)
		}
	}
	return nil
}

func (ie *RRCResume_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Sl_ConfigDedicatedNR_r17Present bool
	if Sl_ConfigDedicatedNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_L2RemoteUE_Config_r17Present bool
	if Sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapNCSG_ConfigNR_r17Present bool
	if NeedForGapNCSG_ConfigNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapNCSG_ConfigEUTRA_r17Present bool
	if NeedForGapNCSG_ConfigEUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_State_r17Present bool
	if Scg_State_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AppLayerMeasConfig_r17Present bool
	if AppLayerMeasConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_ConfigDedicatedNR_r17Present {
		tmp_Sl_ConfigDedicatedNR_r17 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{}
		if err = tmp_Sl_ConfigDedicatedNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ConfigDedicatedNR_r17", err)
		}
		ie.Sl_ConfigDedicatedNR_r17 = tmp_Sl_ConfigDedicatedNR_r17.Setup
	}
	if Sl_L2RemoteUE_Config_r17Present {
		tmp_Sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{}
		if err = tmp_Sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_L2RemoteUE_Config_r17", err)
		}
		ie.Sl_L2RemoteUE_Config_r17 = tmp_Sl_L2RemoteUE_Config_r17.Setup
	}
	if NeedForGapNCSG_ConfigNR_r17Present {
		tmp_NeedForGapNCSG_ConfigNR_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigNR_r17]{}
		if err = tmp_NeedForGapNCSG_ConfigNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapNCSG_ConfigNR_r17", err)
		}
		ie.NeedForGapNCSG_ConfigNR_r17 = tmp_NeedForGapNCSG_ConfigNR_r17.Setup
	}
	if NeedForGapNCSG_ConfigEUTRA_r17Present {
		tmp_NeedForGapNCSG_ConfigEUTRA_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigEUTRA_r17]{}
		if err = tmp_NeedForGapNCSG_ConfigEUTRA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapNCSG_ConfigEUTRA_r17", err)
		}
		ie.NeedForGapNCSG_ConfigEUTRA_r17 = tmp_NeedForGapNCSG_ConfigEUTRA_r17.Setup
	}
	if Scg_State_r17Present {
		ie.Scg_State_r17 = new(RRCResume_v1700_IEs_scg_State_r17)
		if err = ie.Scg_State_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_State_r17", err)
		}
	}
	if AppLayerMeasConfig_r17Present {
		ie.AppLayerMeasConfig_r17 = new(AppLayerMeasConfig_r17)
		if err = ie.AppLayerMeasConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AppLayerMeasConfig_r17", err)
		}
	}
	return nil
}
