package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_v1540 struct {
	OverheatingAssistanceConfig *OverheatingAssistanceConfig `optional,setuprelease`
}

func (ie *OtherConfig_v1540) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OverheatingAssistanceConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OverheatingAssistanceConfig != nil {
		tmp_OverheatingAssistanceConfig := utils.SetupRelease[*OverheatingAssistanceConfig]{
			Setup: ie.OverheatingAssistanceConfig,
		}
		if err = tmp_OverheatingAssistanceConfig.Encode(w); err != nil {
			return utils.WrapError("Encode OverheatingAssistanceConfig", err)
		}
	}
	return nil
}

func (ie *OtherConfig_v1540) Decode(r *aper.AperReader) error {
	var err error
	var OverheatingAssistanceConfigPresent bool
	if OverheatingAssistanceConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OverheatingAssistanceConfigPresent {
		tmp_OverheatingAssistanceConfig := utils.SetupRelease[*OverheatingAssistanceConfig]{}
		if err = tmp_OverheatingAssistanceConfig.Decode(r); err != nil {
			return utils.WrapError("Decode OverheatingAssistanceConfig", err)
		}
		ie.OverheatingAssistanceConfig = tmp_OverheatingAssistanceConfig.Setup
	}
	return nil
}
