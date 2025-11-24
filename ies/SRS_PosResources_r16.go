package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResources_r16 struct {
	MaxNumberSRS_PosResourceSetPerBWP_r16               SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16               `madatory`
	MaxNumberSRS_PosResourcesPerBWP_r16                 SRS_PosResources_r16_maxNumberSRS_PosResourcesPerBWP_r16                 `madatory`
	MaxNumberSRS_ResourcesPerBWP_PerSlot_r16            SRS_PosResources_r16_maxNumberSRS_ResourcesPerBWP_PerSlot_r16            `madatory`
	MaxNumberPeriodicSRS_PosResourcesPerBWP_r16         SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_r16         `madatory`
	MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16 SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16 `madatory`
}

func (ie *SRS_PosResources_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberSRS_PosResourceSetPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	if err = ie.MaxNumberSRS_PosResourcesPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberSRS_ResourcesPerBWP_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_ResourcesPerBWP_PerSlot_r16", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPeriodicSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}

func (ie *SRS_PosResources_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberSRS_PosResourceSetPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	if err = ie.MaxNumberSRS_PosResourcesPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberSRS_ResourcesPerBWP_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_ResourcesPerBWP_PerSlot_r16", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPeriodicSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}
