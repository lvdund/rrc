package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BWP_DownlinkDedicatedSDT_r17 struct {
	Pdcch_Config_r17 *PDCCH_Config `optional,setuprelease`
	Pdsch_Config_r17 *PDSCH_Config `optional,setuprelease`
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcch_Config_r17 != nil, ie.Pdsch_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcch_Config_r17 != nil {
		tmp_Pdcch_Config_r17 := utils.SetupRelease[*PDCCH_Config]{
			Setup: ie.Pdcch_Config_r17,
		}
		if err = tmp_Pdcch_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_Config_r17", err)
		}
	}
	if ie.Pdsch_Config_r17 != nil {
		tmp_Pdsch_Config_r17 := utils.SetupRelease[*PDSCH_Config]{
			Setup: ie.Pdsch_Config_r17,
		}
		if err = tmp_Pdsch_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_Config_r17", err)
		}
	}
	return nil
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pdcch_Config_r17Present bool
	if Pdcch_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_Config_r17Present bool
	if Pdsch_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcch_Config_r17Present {
		tmp_Pdcch_Config_r17 := utils.SetupRelease[*PDCCH_Config]{}
		if err = tmp_Pdcch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_Config_r17", err)
		}
		ie.Pdcch_Config_r17 = tmp_Pdcch_Config_r17.Setup
	}
	if Pdsch_Config_r17Present {
		tmp_Pdsch_Config_r17 := utils.SetupRelease[*PDSCH_Config]{}
		if err = tmp_Pdsch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_Config_r17", err)
		}
		ie.Pdsch_Config_r17 = tmp_Pdsch_Config_r17.Setup
	}
	return nil
}
