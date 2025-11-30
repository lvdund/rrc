package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddModUTRA_FDD_r16 struct {
	CellIndexUTRA_FDD_r16 UTRA_FDD_CellIndex_r16 `madatory`
	PhysCellId_r16        PhysCellIdUTRA_FDD_r16 `madatory`
}

func (ie *CellsToAddModUTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CellIndexUTRA_FDD_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CellIndexUTRA_FDD_r16", err)
	}
	if err = ie.PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r16", err)
	}
	return nil
}

func (ie *CellsToAddModUTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CellIndexUTRA_FDD_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CellIndexUTRA_FDD_r16", err)
	}
	if err = ie.PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r16", err)
	}
	return nil
}
