package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17 struct {
	MaxNumberSRS_PosResourceSetPerBWP_r17               SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberSRS_PosResourceSetPerBWP_r17               `madatory`
	MaxNumberSRS_PosResourcesPerBWP_r17                 SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberSRS_PosResourcesPerBWP_r17                 `madatory`
	MaxNumberSRS_ResourcesPerBWP_PerSlot_r17            SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberSRS_ResourcesPerBWP_PerSlot_r17            `madatory`
	MaxNumberPeriodicSRS_PosResourcesPerBWP_r17         SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberPeriodicSRS_PosResourcesPerBWP_r17         `madatory`
	MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17 SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17 `madatory`
	Dummy1                                              SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1                                              `madatory`
	Dummy2                                              SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy2                                              `madatory`
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberSRS_PosResourceSetPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_PosResourceSetPerBWP_r17", err)
	}
	if err = ie.MaxNumberSRS_PosResourcesPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.MaxNumberSRS_ResourcesPerBWP_PerSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSRS_ResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPeriodicSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.Dummy1.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy1", err)
	}
	if err = ie.Dummy2.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy2", err)
	}
	return nil
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberSRS_PosResourceSetPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_PosResourceSetPerBWP_r17", err)
	}
	if err = ie.MaxNumberSRS_PosResourcesPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.MaxNumberSRS_ResourcesPerBWP_PerSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSRS_ResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPeriodicSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.Dummy1.Decode(r); err != nil {
		return utils.WrapError("Decode Dummy1", err)
	}
	if err = ie.Dummy2.Decode(r); err != nil {
		return utils.WrapError("Decode Dummy2", err)
	}
	return nil
}
