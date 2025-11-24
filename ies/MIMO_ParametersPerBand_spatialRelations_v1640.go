package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_spatialRelations_v1640 struct {
	MaxNumberConfiguredSpatialRelations_v1640 MIMO_ParametersPerBand_spatialRelations_v1640_maxNumberConfiguredSpatialRelations_v1640 `madatory`
}

func (ie *MIMO_ParametersPerBand_spatialRelations_v1640) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberConfiguredSpatialRelations_v1640.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberConfiguredSpatialRelations_v1640", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_spatialRelations_v1640) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberConfiguredSpatialRelations_v1640.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberConfiguredSpatialRelations_v1640", err)
	}
	return nil
}
