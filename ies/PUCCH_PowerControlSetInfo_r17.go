package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PowerControlSetInfo_r17 struct {
	Pucch_PowerControlSetInfoId_r17  PUCCH_PowerControlSetInfoId_r17                         `madatory`
	P0_PUCCH_Id_r17                  P0_PUCCH_Id                                             `madatory`
	Pucch_ClosedLoopIndex_r17        PUCCH_PowerControlSetInfo_r17_pucch_ClosedLoopIndex_r17 `madatory`
	Pucch_PathlossReferenceRS_Id_r17 PUCCH_PathlossReferenceRS_Id_r17                        `madatory`
}

func (ie *PUCCH_PowerControlSetInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pucch_PowerControlSetInfoId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_PowerControlSetInfoId_r17", err)
	}
	if err = ie.P0_PUCCH_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode P0_PUCCH_Id_r17", err)
	}
	if err = ie.Pucch_ClosedLoopIndex_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_ClosedLoopIndex_r17", err)
	}
	if err = ie.Pucch_PathlossReferenceRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_PathlossReferenceRS_Id_r17", err)
	}
	return nil
}

func (ie *PUCCH_PowerControlSetInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pucch_PowerControlSetInfoId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_PowerControlSetInfoId_r17", err)
	}
	if err = ie.P0_PUCCH_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode P0_PUCCH_Id_r17", err)
	}
	if err = ie.Pucch_ClosedLoopIndex_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_ClosedLoopIndex_r17", err)
	}
	if err = ie.Pucch_PathlossReferenceRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_PathlossReferenceRS_Id_r17", err)
	}
	return nil
}
