package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1730 struct {
	Prs_AsSpatialRelationRS_For_SRS_r17 *FeatureSetDownlink_v1730_prs_AsSpatialRelationRS_For_SRS_r17 `optional`
}

func (ie *FeatureSetDownlink_v1730) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Prs_AsSpatialRelationRS_For_SRS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Prs_AsSpatialRelationRS_For_SRS_r17 != nil {
		if err = ie.Prs_AsSpatialRelationRS_For_SRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Prs_AsSpatialRelationRS_For_SRS_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1730) Decode(r *aper.AperReader) error {
	var err error
	var Prs_AsSpatialRelationRS_For_SRS_r17Present bool
	if Prs_AsSpatialRelationRS_For_SRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Prs_AsSpatialRelationRS_For_SRS_r17Present {
		ie.Prs_AsSpatialRelationRS_For_SRS_r17 = new(FeatureSetDownlink_v1730_prs_AsSpatialRelationRS_For_SRS_r17)
		if err = ie.Prs_AsSpatialRelationRS_For_SRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Prs_AsSpatialRelationRS_For_SRS_r17", err)
		}
	}
	return nil
}
