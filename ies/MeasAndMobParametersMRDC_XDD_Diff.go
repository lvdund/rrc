package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_XDD_Diff struct {
	Sftd_MeasPSCell  *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell  `optional`
	Sftd_MeasNR_Cell *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell `optional`
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sftd_MeasPSCell != nil, ie.Sftd_MeasNR_Cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sftd_MeasPSCell != nil {
		if err = ie.Sftd_MeasPSCell.Encode(w); err != nil {
			return utils.WrapError("Encode Sftd_MeasPSCell", err)
		}
	}
	if ie.Sftd_MeasNR_Cell != nil {
		if err = ie.Sftd_MeasNR_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode Sftd_MeasNR_Cell", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff) Decode(r *aper.AperReader) error {
	var err error
	var Sftd_MeasPSCellPresent bool
	if Sftd_MeasPSCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sftd_MeasNR_CellPresent bool
	if Sftd_MeasNR_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sftd_MeasPSCellPresent {
		ie.Sftd_MeasPSCell = new(MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell)
		if err = ie.Sftd_MeasPSCell.Decode(r); err != nil {
			return utils.WrapError("Decode Sftd_MeasPSCell", err)
		}
	}
	if Sftd_MeasNR_CellPresent {
		ie.Sftd_MeasNR_Cell = new(MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell)
		if err = ie.Sftd_MeasNR_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode Sftd_MeasNR_Cell", err)
		}
	}
	return nil
}
