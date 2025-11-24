package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpatialRelationsSRS_Pos_r16 struct {
	SpatialRelation_SRS_PosBasedOnSSB_Serving_r16    *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Serving_r16    `optional`
	SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 `optional`
	SpatialRelation_SRS_PosBasedOnPRS_Serving_r16    *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnPRS_Serving_r16    `optional`
	SpatialRelation_SRS_PosBasedOnSRS_r16            *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSRS_r16            `optional`
	SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16      *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16      `optional`
	SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16      *SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnPRS_Neigh_r16      `optional`
}

func (ie *SpatialRelationsSRS_Pos_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnSRS_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16 != nil, ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16 != nil {
		if err = ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelation_SRS_PosBasedOnSSB_Serving_r16", err)
		}
	}
	if ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 != nil {
		if err = ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16", err)
		}
	}
	if ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16 != nil {
		if err = ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelation_SRS_PosBasedOnPRS_Serving_r16", err)
		}
	}
	if ie.SpatialRelation_SRS_PosBasedOnSRS_r16 != nil {
		if err = ie.SpatialRelation_SRS_PosBasedOnSRS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelation_SRS_PosBasedOnSRS_r16", err)
		}
	}
	if ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16 != nil {
		if err = ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16", err)
		}
	}
	if ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16 != nil {
		if err = ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16", err)
		}
	}
	return nil
}

func (ie *SpatialRelationsSRS_Pos_r16) Decode(r *uper.UperReader) error {
	var err error
	var SpatialRelation_SRS_PosBasedOnSSB_Serving_r16Present bool
	if SpatialRelation_SRS_PosBasedOnSSB_Serving_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16Present bool
	if SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialRelation_SRS_PosBasedOnPRS_Serving_r16Present bool
	if SpatialRelation_SRS_PosBasedOnPRS_Serving_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialRelation_SRS_PosBasedOnSRS_r16Present bool
	if SpatialRelation_SRS_PosBasedOnSRS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16Present bool
	if SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16Present bool
	if SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SpatialRelation_SRS_PosBasedOnSSB_Serving_r16Present {
		ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16 = new(SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Serving_r16)
		if err = ie.SpatialRelation_SRS_PosBasedOnSSB_Serving_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SpatialRelation_SRS_PosBasedOnSSB_Serving_r16", err)
		}
	}
	if SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16Present {
		ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16 = new(SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16)
		if err = ie.SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SpatialRelation_SRS_PosBasedOnCSI_RS_Serving_r16", err)
		}
	}
	if SpatialRelation_SRS_PosBasedOnPRS_Serving_r16Present {
		ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16 = new(SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnPRS_Serving_r16)
		if err = ie.SpatialRelation_SRS_PosBasedOnPRS_Serving_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SpatialRelation_SRS_PosBasedOnPRS_Serving_r16", err)
		}
	}
	if SpatialRelation_SRS_PosBasedOnSRS_r16Present {
		ie.SpatialRelation_SRS_PosBasedOnSRS_r16 = new(SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSRS_r16)
		if err = ie.SpatialRelation_SRS_PosBasedOnSRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SpatialRelation_SRS_PosBasedOnSRS_r16", err)
		}
	}
	if SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16Present {
		ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16 = new(SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnSSB_Neigh_r16)
		if err = ie.SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SpatialRelation_SRS_PosBasedOnSSB_Neigh_r16", err)
		}
	}
	if SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16Present {
		ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16 = new(SpatialRelationsSRS_Pos_r16_spatialRelation_SRS_PosBasedOnPRS_Neigh_r16)
		if err = ie.SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SpatialRelation_SRS_PosBasedOnPRS_Neigh_r16", err)
		}
	}
	return nil
}
