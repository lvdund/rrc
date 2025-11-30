package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_Cell struct {
	CellIndexEUTRA       EUTRA_CellIndex     `madatory`
	PhysCellId           EUTRA_PhysCellId    `madatory`
	CellIndividualOffset EUTRA_Q_OffsetRange `madatory`
}

func (ie *EUTRA_Cell) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CellIndexEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode CellIndexEUTRA", err)
	}
	if err = ie.PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	if err = ie.CellIndividualOffset.Encode(w); err != nil {
		return utils.WrapError("Encode CellIndividualOffset", err)
	}
	return nil
}

func (ie *EUTRA_Cell) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CellIndexEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode CellIndexEUTRA", err)
	}
	if err = ie.PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	if err = ie.CellIndividualOffset.Decode(r); err != nil {
		return utils.WrapError("Decode CellIndividualOffset", err)
	}
	return nil
}
