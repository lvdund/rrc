package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SpatialRelationInfoExt_r16 struct {
	Pucch_SpatialRelationInfoId_v1610  *PUCCH_SpatialRelationInfoId_v1610  `optional`
	Pucch_PathlossReferenceRS_Id_v1610 *PUCCH_PathlossReferenceRS_Id_v1610 `optional`
}

func (ie *PUCCH_SpatialRelationInfoExt_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pucch_SpatialRelationInfoId_v1610 != nil, ie.Pucch_PathlossReferenceRS_Id_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pucch_SpatialRelationInfoId_v1610 != nil {
		if err = ie.Pucch_SpatialRelationInfoId_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_SpatialRelationInfoId_v1610", err)
		}
	}
	if ie.Pucch_PathlossReferenceRS_Id_v1610 != nil {
		if err = ie.Pucch_PathlossReferenceRS_Id_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_PathlossReferenceRS_Id_v1610", err)
		}
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfoExt_r16) Decode(r *aper.AperReader) error {
	var err error
	var Pucch_SpatialRelationInfoId_v1610Present bool
	if Pucch_SpatialRelationInfoId_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_PathlossReferenceRS_Id_v1610Present bool
	if Pucch_PathlossReferenceRS_Id_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pucch_SpatialRelationInfoId_v1610Present {
		ie.Pucch_SpatialRelationInfoId_v1610 = new(PUCCH_SpatialRelationInfoId_v1610)
		if err = ie.Pucch_SpatialRelationInfoId_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_SpatialRelationInfoId_v1610", err)
		}
	}
	if Pucch_PathlossReferenceRS_Id_v1610Present {
		ie.Pucch_PathlossReferenceRS_Id_v1610 = new(PUCCH_PathlossReferenceRS_Id_v1610)
		if err = ie.Pucch_PathlossReferenceRS_Id_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_PathlossReferenceRS_Id_v1610", err)
		}
	}
	return nil
}
