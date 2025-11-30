package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_ExcludedCell struct {
	CellIndexEUTRA  EUTRA_CellIndex       `madatory`
	PhysCellIdRange EUTRA_PhysCellIdRange `madatory`
}

func (ie *EUTRA_ExcludedCell) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CellIndexEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode CellIndexEUTRA", err)
	}
	if err = ie.PhysCellIdRange.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellIdRange", err)
	}
	return nil
}

func (ie *EUTRA_ExcludedCell) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CellIndexEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode CellIndexEUTRA", err)
	}
	if err = ie.PhysCellIdRange.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellIdRange", err)
	}
	return nil
}
