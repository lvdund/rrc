package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SpatialRelationInfo struct {
	Pucch_SpatialRelationInfoId  PUCCH_SpatialRelationInfoId               `madatory`
	ServingCellId                *ServCellIndex                            `optional`
	ReferenceSignal              PUCCH_SpatialRelationInfo_referenceSignal `madatory`
	Pucch_PathlossReferenceRS_Id PUCCH_PathlossReferenceRS_Id              `madatory`
	P0_PUCCH_Id                  P0_PUCCH_Id                               `madatory`
	ClosedLoopIndex              PUCCH_SpatialRelationInfo_closedLoopIndex `madatory`
}

func (ie *PUCCH_SpatialRelationInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ServingCellId != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Pucch_SpatialRelationInfoId.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_SpatialRelationInfoId", err)
	}
	if ie.ServingCellId != nil {
		if err = ie.ServingCellId.Encode(w); err != nil {
			return utils.WrapError("Encode ServingCellId", err)
		}
	}
	if err = ie.ReferenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal", err)
	}
	if err = ie.Pucch_PathlossReferenceRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.P0_PUCCH_Id.Encode(w); err != nil {
		return utils.WrapError("Encode P0_PUCCH_Id", err)
	}
	if err = ie.ClosedLoopIndex.Encode(w); err != nil {
		return utils.WrapError("Encode ClosedLoopIndex", err)
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfo) Decode(r *uper.UperReader) error {
	var err error
	var ServingCellIdPresent bool
	if ServingCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Pucch_SpatialRelationInfoId.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_SpatialRelationInfoId", err)
	}
	if ServingCellIdPresent {
		ie.ServingCellId = new(ServCellIndex)
		if err = ie.ServingCellId.Decode(r); err != nil {
			return utils.WrapError("Decode ServingCellId", err)
		}
	}
	if err = ie.ReferenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal", err)
	}
	if err = ie.Pucch_PathlossReferenceRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.P0_PUCCH_Id.Decode(r); err != nil {
		return utils.WrapError("Decode P0_PUCCH_Id", err)
	}
	if err = ie.ClosedLoopIndex.Decode(r); err != nil {
		return utils.WrapError("Decode ClosedLoopIndex", err)
	}
	return nil
}
