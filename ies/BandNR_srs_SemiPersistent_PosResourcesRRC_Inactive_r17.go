package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17 struct {
	MaxNumOfSemiPersistentSRSposResources_r17        BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResources_r17        `madatory`
	MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17 `madatory`
}

func (ie *BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumOfSemiPersistentSRSposResources_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumOfSemiPersistentSRSposResources_r17", err)
	}
	if err = ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
	}
	return nil
}

func (ie *BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumOfSemiPersistentSRSposResources_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumOfSemiPersistentSRSposResources_r17", err)
	}
	if err = ie.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
	}
	return nil
}
