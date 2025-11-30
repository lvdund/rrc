package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCellIdleEUTRA_r16 struct {
	Eutra_PhysCellId_r16    EUTRA_PhysCellId                                         `madatory`
	MeasIdleResultEUTRA_r16 *MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16 `optional`
}

func (ie *MeasResultsPerCellIdleEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasIdleResultEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Eutra_PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Eutra_PhysCellId_r16", err)
	}
	if ie.MeasIdleResultEUTRA_r16 != nil {
		if err = ie.MeasIdleResultEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdleResultEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultsPerCellIdleEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasIdleResultEUTRA_r16Present bool
	if MeasIdleResultEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Eutra_PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Eutra_PhysCellId_r16", err)
	}
	if MeasIdleResultEUTRA_r16Present {
		ie.MeasIdleResultEUTRA_r16 = new(MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16)
		if err = ie.MeasIdleResultEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasIdleResultEUTRA_r16", err)
		}
	}
	return nil
}
