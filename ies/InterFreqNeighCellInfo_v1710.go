package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqNeighCellInfo_v1710 struct {
	Ssb_PositionQCL_r17 *SSB_PositionQCL_Relation_r17 `optional`
}

func (ie *InterFreqNeighCellInfo_v1710) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_PositionQCL_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ssb_PositionQCL_r17 != nil {
		if err = ie.Ssb_PositionQCL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionQCL_r17", err)
		}
	}
	return nil
}

func (ie *InterFreqNeighCellInfo_v1710) Decode(r *aper.AperReader) error {
	var err error
	var Ssb_PositionQCL_r17Present bool
	if Ssb_PositionQCL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ssb_PositionQCL_r17Present {
		ie.Ssb_PositionQCL_r17 = new(SSB_PositionQCL_Relation_r17)
		if err = ie.Ssb_PositionQCL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionQCL_r17", err)
		}
	}
	return nil
}
