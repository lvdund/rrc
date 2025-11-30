package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResourceSP_r16 struct {
	MaxNumberSP_SRS_PosResourcesPerBWP_r16         SRS_PosResourceSP_r16_maxNumberSP_SRS_PosResourcesPerBWP_r16         `madatory`
	MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16 SRS_PosResourceSP_r16_maxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16 `madatory`
}

func (ie *SRS_PosResourceSP_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxNumberSP_SRS_PosResourcesPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSP_SRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}

func (ie *SRS_PosResourceSP_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxNumberSP_SRS_PosResourcesPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSP_SRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}
