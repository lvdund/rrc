package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddMod struct {
	PhysCellId           PhysCellId        `madatory`
	CellIndividualOffset Q_OffsetRangeList `madatory`
}

func (ie *CellsToAddMod) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	if err = ie.CellIndividualOffset.Encode(w); err != nil {
		return utils.WrapError("Encode CellIndividualOffset", err)
	}
	return nil
}

func (ie *CellsToAddMod) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	if err = ie.CellIndividualOffset.Decode(r); err != nil {
		return utils.WrapError("Decode CellIndividualOffset", err)
	}
	return nil
}
