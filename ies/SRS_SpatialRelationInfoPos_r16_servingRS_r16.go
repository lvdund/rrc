package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SpatialRelationInfoPos_r16_servingRS_r16 struct {
	ServingCellId       *ServCellIndex                                                   `optional`
	ReferenceSignal_r16 SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16 `madatory`
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ServingCellId != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ServingCellId != nil {
		if err = ie.ServingCellId.Encode(w); err != nil {
			return utils.WrapError("Encode ServingCellId", err)
		}
	}
	if err = ie.ReferenceSignal_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal_r16", err)
	}
	return nil
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16) Decode(r *uper.UperReader) error {
	var err error
	var ServingCellIdPresent bool
	if ServingCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ServingCellIdPresent {
		ie.ServingCellId = new(ServCellIndex)
		if err = ie.ServingCellId.Decode(r); err != nil {
			return utils.WrapError("Decode ServingCellId", err)
		}
	}
	if err = ie.ReferenceSignal_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal_r16", err)
	}
	return nil
}
