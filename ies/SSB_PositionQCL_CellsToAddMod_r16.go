package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SSB_PositionQCL_CellsToAddMod_r16 struct {
	PhysCellId_r16      PhysCellId                   `madatory`
	Ssb_PositionQCL_r16 SSB_PositionQCL_Relation_r16 `madatory`
}

func (ie *SSB_PositionQCL_CellsToAddMod_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r16", err)
	}
	if err = ie.Ssb_PositionQCL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_PositionQCL_r16", err)
	}
	return nil
}

func (ie *SSB_PositionQCL_CellsToAddMod_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r16", err)
	}
	if err = ie.Ssb_PositionQCL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_PositionQCL_r16", err)
	}
	return nil
}
