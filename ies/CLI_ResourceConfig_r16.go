package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CLI_ResourceConfig_r16 struct {
	Srs_ResourceConfig_r16  *SRS_ResourceListConfigCLI_r16  `optional,setuprelease`
	Rssi_ResourceConfig_r16 *RSSI_ResourceListConfigCLI_r16 `optional,setuprelease`
}

func (ie *CLI_ResourceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_ResourceConfig_r16 != nil, ie.Rssi_ResourceConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srs_ResourceConfig_r16 != nil {
		tmp_Srs_ResourceConfig_r16 := utils.SetupRelease[*SRS_ResourceListConfigCLI_r16]{
			Setup: ie.Srs_ResourceConfig_r16,
		}
		if err = tmp_Srs_ResourceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_ResourceConfig_r16", err)
		}
	}
	if ie.Rssi_ResourceConfig_r16 != nil {
		tmp_Rssi_ResourceConfig_r16 := utils.SetupRelease[*RSSI_ResourceListConfigCLI_r16]{
			Setup: ie.Rssi_ResourceConfig_r16,
		}
		if err = tmp_Rssi_ResourceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rssi_ResourceConfig_r16", err)
		}
	}
	return nil
}

func (ie *CLI_ResourceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var Srs_ResourceConfig_r16Present bool
	if Srs_ResourceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rssi_ResourceConfig_r16Present bool
	if Rssi_ResourceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_ResourceConfig_r16Present {
		tmp_Srs_ResourceConfig_r16 := utils.SetupRelease[*SRS_ResourceListConfigCLI_r16]{}
		if err = tmp_Srs_ResourceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_ResourceConfig_r16", err)
		}
		ie.Srs_ResourceConfig_r16 = tmp_Srs_ResourceConfig_r16.Setup
	}
	if Rssi_ResourceConfig_r16Present {
		tmp_Rssi_ResourceConfig_r16 := utils.SetupRelease[*RSSI_ResourceListConfigCLI_r16]{}
		if err = tmp_Rssi_ResourceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rssi_ResourceConfig_r16", err)
		}
		ie.Rssi_ResourceConfig_r16 = tmp_Rssi_ResourceConfig_r16.Setup
	}
	return nil
}
