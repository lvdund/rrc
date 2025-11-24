package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReferenceSignalConfig struct {
	Ssb_ConfigMobility            *SSB_ConfigMobility            `optional`
	Csi_rs_ResourceConfigMobility *CSI_RS_ResourceConfigMobility `optional,setuprelease`
}

func (ie *ReferenceSignalConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_ConfigMobility != nil, ie.Csi_rs_ResourceConfigMobility != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ssb_ConfigMobility != nil {
		if err = ie.Ssb_ConfigMobility.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_ConfigMobility", err)
		}
	}
	if ie.Csi_rs_ResourceConfigMobility != nil {
		tmp_Csi_rs_ResourceConfigMobility := utils.SetupRelease[*CSI_RS_ResourceConfigMobility]{
			Setup: ie.Csi_rs_ResourceConfigMobility,
		}
		if err = tmp_Csi_rs_ResourceConfigMobility.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_rs_ResourceConfigMobility", err)
		}
	}
	return nil
}

func (ie *ReferenceSignalConfig) Decode(r *uper.UperReader) error {
	var err error
	var Ssb_ConfigMobilityPresent bool
	if Ssb_ConfigMobilityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_rs_ResourceConfigMobilityPresent bool
	if Csi_rs_ResourceConfigMobilityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ssb_ConfigMobilityPresent {
		ie.Ssb_ConfigMobility = new(SSB_ConfigMobility)
		if err = ie.Ssb_ConfigMobility.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_ConfigMobility", err)
		}
	}
	if Csi_rs_ResourceConfigMobilityPresent {
		tmp_Csi_rs_ResourceConfigMobility := utils.SetupRelease[*CSI_RS_ResourceConfigMobility]{}
		if err = tmp_Csi_rs_ResourceConfigMobility.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_rs_ResourceConfigMobility", err)
		}
		ie.Csi_rs_ResourceConfigMobility = tmp_Csi_rs_ResourceConfigMobility.Setup
	}
	return nil
}
