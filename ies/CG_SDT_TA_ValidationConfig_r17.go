package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_SDT_TA_ValidationConfig_r17 struct {
	Cg_SDT_RSRP_ChangeThreshold_r17 CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17 `madatory`
}

func (ie *CG_SDT_TA_ValidationConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Cg_SDT_RSRP_ChangeThreshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Cg_SDT_RSRP_ChangeThreshold_r17", err)
	}
	return nil
}

func (ie *CG_SDT_TA_ValidationConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Cg_SDT_RSRP_ChangeThreshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Cg_SDT_RSRP_ChangeThreshold_r17", err)
	}
	return nil
}
