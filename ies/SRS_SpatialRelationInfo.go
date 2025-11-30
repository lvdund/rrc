package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SpatialRelationInfo struct {
	ServingCellId   *ServCellIndex                          `optional`
	ReferenceSignal SRS_SpatialRelationInfo_referenceSignal `madatory`
}

func (ie *SRS_SpatialRelationInfo) Encode(w *aper.AperWriter) error {
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
	if err = ie.ReferenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal", err)
	}
	return nil
}

func (ie *SRS_SpatialRelationInfo) Decode(r *aper.AperReader) error {
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
	if err = ie.ReferenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal", err)
	}
	return nil
}
