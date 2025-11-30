package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRI_PUSCH_PowerControl struct {
	Sri_PUSCH_PowerControlId         SRI_PUSCH_PowerControlId                         `madatory`
	Sri_PUSCH_PathlossReferenceRS_Id PUSCH_PathlossReferenceRS_Id                     `madatory`
	Sri_P0_PUSCH_AlphaSetId          P0_PUSCH_AlphaSetId                              `madatory`
	Sri_PUSCH_ClosedLoopIndex        SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex `madatory`
}

func (ie *SRI_PUSCH_PowerControl) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sri_PUSCH_PowerControlId.Encode(w); err != nil {
		return utils.WrapError("Encode Sri_PUSCH_PowerControlId", err)
	}
	if err = ie.Sri_PUSCH_PathlossReferenceRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Sri_PUSCH_PathlossReferenceRS_Id", err)
	}
	if err = ie.Sri_P0_PUSCH_AlphaSetId.Encode(w); err != nil {
		return utils.WrapError("Encode Sri_P0_PUSCH_AlphaSetId", err)
	}
	if err = ie.Sri_PUSCH_ClosedLoopIndex.Encode(w); err != nil {
		return utils.WrapError("Encode Sri_PUSCH_ClosedLoopIndex", err)
	}
	return nil
}

func (ie *SRI_PUSCH_PowerControl) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sri_PUSCH_PowerControlId.Decode(r); err != nil {
		return utils.WrapError("Decode Sri_PUSCH_PowerControlId", err)
	}
	if err = ie.Sri_PUSCH_PathlossReferenceRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Sri_PUSCH_PathlossReferenceRS_Id", err)
	}
	if err = ie.Sri_P0_PUSCH_AlphaSetId.Decode(r); err != nil {
		return utils.WrapError("Decode Sri_P0_PUSCH_AlphaSetId", err)
	}
	if err = ie.Sri_PUSCH_ClosedLoopIndex.Decode(r); err != nil {
		return utils.WrapError("Decode Sri_PUSCH_ClosedLoopIndex", err)
	}
	return nil
}
