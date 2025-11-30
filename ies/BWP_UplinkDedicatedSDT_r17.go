package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkDedicatedSDT_r17 struct {
	Pusch_Config_r17                       *PUSCH_Config                           `optional,setuprelease`
	ConfiguredGrantConfigToAddModList_r17  *ConfiguredGrantConfigToAddModList_r16  `optional`
	ConfiguredGrantConfigToReleaseList_r17 *ConfiguredGrantConfigToReleaseList_r16 `optional`
}

func (ie *BWP_UplinkDedicatedSDT_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pusch_Config_r17 != nil, ie.ConfiguredGrantConfigToAddModList_r17 != nil, ie.ConfiguredGrantConfigToReleaseList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pusch_Config_r17 != nil {
		tmp_Pusch_Config_r17 := utils.SetupRelease[*PUSCH_Config]{
			Setup: ie.Pusch_Config_r17,
		}
		if err = tmp_Pusch_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_Config_r17", err)
		}
	}
	if ie.ConfiguredGrantConfigToAddModList_r17 != nil {
		if err = ie.ConfiguredGrantConfigToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredGrantConfigToAddModList_r17", err)
		}
	}
	if ie.ConfiguredGrantConfigToReleaseList_r17 != nil {
		if err = ie.ConfiguredGrantConfigToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredGrantConfigToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *BWP_UplinkDedicatedSDT_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pusch_Config_r17Present bool
	if Pusch_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredGrantConfigToAddModList_r17Present bool
	if ConfiguredGrantConfigToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredGrantConfigToReleaseList_r17Present bool
	if ConfiguredGrantConfigToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pusch_Config_r17Present {
		tmp_Pusch_Config_r17 := utils.SetupRelease[*PUSCH_Config]{}
		if err = tmp_Pusch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_Config_r17", err)
		}
		ie.Pusch_Config_r17 = tmp_Pusch_Config_r17.Setup
	}
	if ConfiguredGrantConfigToAddModList_r17Present {
		ie.ConfiguredGrantConfigToAddModList_r17 = new(ConfiguredGrantConfigToAddModList_r16)
		if err = ie.ConfiguredGrantConfigToAddModList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredGrantConfigToAddModList_r17", err)
		}
	}
	if ConfiguredGrantConfigToReleaseList_r17Present {
		ie.ConfiguredGrantConfigToReleaseList_r17 = new(ConfiguredGrantConfigToReleaseList_r16)
		if err = ie.ConfiguredGrantConfigToReleaseList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredGrantConfigToReleaseList_r17", err)
		}
	}
	return nil
}
