package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpatialRelations struct {
	MaxNumberConfiguredSpatialRelations  SpatialRelations_maxNumberConfiguredSpatialRelations   `madatory`
	MaxNumberActiveSpatialRelations      SpatialRelations_maxNumberActiveSpatialRelations       `madatory`
	AdditionalActiveSpatialRelationPUCCH *SpatialRelations_additionalActiveSpatialRelationPUCCH `optional`
	MaxNumberDL_RS_QCL_TypeD             SpatialRelations_maxNumberDL_RS_QCL_TypeD              `madatory`
}

func (ie *SpatialRelations) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalActiveSpatialRelationPUCCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MaxNumberConfiguredSpatialRelations.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberConfiguredSpatialRelations", err)
	}
	if err = ie.MaxNumberActiveSpatialRelations.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberActiveSpatialRelations", err)
	}
	if ie.AdditionalActiveSpatialRelationPUCCH != nil {
		if err = ie.AdditionalActiveSpatialRelationPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalActiveSpatialRelationPUCCH", err)
		}
	}
	if err = ie.MaxNumberDL_RS_QCL_TypeD.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberDL_RS_QCL_TypeD", err)
	}
	return nil
}

func (ie *SpatialRelations) Decode(r *uper.UperReader) error {
	var err error
	var AdditionalActiveSpatialRelationPUCCHPresent bool
	if AdditionalActiveSpatialRelationPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MaxNumberConfiguredSpatialRelations.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberConfiguredSpatialRelations", err)
	}
	if err = ie.MaxNumberActiveSpatialRelations.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberActiveSpatialRelations", err)
	}
	if AdditionalActiveSpatialRelationPUCCHPresent {
		ie.AdditionalActiveSpatialRelationPUCCH = new(SpatialRelations_additionalActiveSpatialRelationPUCCH)
		if err = ie.AdditionalActiveSpatialRelationPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalActiveSpatialRelationPUCCH", err)
		}
	}
	if err = ie.MaxNumberDL_RS_QCL_TypeD.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberDL_RS_QCL_TypeD", err)
	}
	return nil
}
