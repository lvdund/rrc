package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SSB_PositionQCL_Cell_r17 struct {
	PhysCellId_r17      PhysCellId                   `madatory`
	Ssb_PositionQCL_r17 SSB_PositionQCL_Relation_r17 `madatory`
}

func (ie *SSB_PositionQCL_Cell_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PhysCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r17", err)
	}
	if err = ie.Ssb_PositionQCL_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_PositionQCL_r17", err)
	}
	return nil
}

func (ie *SSB_PositionQCL_Cell_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PhysCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r17", err)
	}
	if err = ie.Ssb_PositionQCL_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_PositionQCL_r17", err)
	}
	return nil
}
