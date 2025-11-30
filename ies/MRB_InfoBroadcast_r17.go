package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRB_InfoBroadcast_r17 struct {
	Pdcp_Config_r17 MRB_PDCP_ConfigBroadcast_r17 `madatory`
	Rlc_Config_r17  MRB_RLC_ConfigBroadcast_r17  `madatory`
}

func (ie *MRB_InfoBroadcast_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pdcp_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcp_Config_r17", err)
	}
	if err = ie.Rlc_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Rlc_Config_r17", err)
	}
	return nil
}

func (ie *MRB_InfoBroadcast_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pdcp_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcp_Config_r17", err)
	}
	if err = ie.Rlc_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Rlc_Config_r17", err)
	}
	return nil
}
