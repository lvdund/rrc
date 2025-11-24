package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_RequestConfig_rach_OccasionsSI struct {
	Rach_ConfigSI        RACH_ConfigGeneric                                     `madatory`
	Ssb_perRACH_Occasion SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion `madatory`
}

func (ie *SI_RequestConfig_rach_OccasionsSI) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Rach_ConfigSI.Encode(w); err != nil {
		return utils.WrapError("Encode Rach_ConfigSI", err)
	}
	if err = ie.Ssb_perRACH_Occasion.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_perRACH_Occasion", err)
	}
	return nil
}

func (ie *SI_RequestConfig_rach_OccasionsSI) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Rach_ConfigSI.Decode(r); err != nil {
		return utils.WrapError("Decode Rach_ConfigSI", err)
	}
	if err = ie.Ssb_perRACH_Occasion.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_perRACH_Occasion", err)
	}
	return nil
}
