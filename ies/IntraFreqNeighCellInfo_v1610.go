package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraFreqNeighCellInfo_v1610 struct {
	Ssb_PositionQCL_r16 *SSB_PositionQCL_Relation_r16 `optional`
}

func (ie *IntraFreqNeighCellInfo_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_PositionQCL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ssb_PositionQCL_r16 != nil {
		if err = ie.Ssb_PositionQCL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionQCL_r16", err)
		}
	}
	return nil
}

func (ie *IntraFreqNeighCellInfo_v1610) Decode(r *uper.UperReader) error {
	var err error
	var Ssb_PositionQCL_r16Present bool
	if Ssb_PositionQCL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ssb_PositionQCL_r16Present {
		ie.Ssb_PositionQCL_r16 = new(SSB_PositionQCL_Relation_r16)
		if err = ie.Ssb_PositionQCL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionQCL_r16", err)
		}
	}
	return nil
}
