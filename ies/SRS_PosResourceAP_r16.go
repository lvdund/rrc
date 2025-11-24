package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResourceAP_r16 struct {
	MaxNumberAP_SRS_PosResourcesPerBWP_r16         SRS_PosResourceAP_r16_maxNumberAP_SRS_PosResourcesPerBWP_r16         `madatory`
	MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16 SRS_PosResourceAP_r16_maxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16 `madatory`
}

func (ie *SRS_PosResourceAP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberAP_SRS_PosResourcesPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberAP_SRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}

func (ie *SRS_PosResourceAP_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberAP_SRS_PosResourcesPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberAP_SRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberAP_SRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}
