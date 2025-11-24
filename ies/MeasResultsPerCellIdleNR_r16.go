package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCellIdleNR_r16 struct {
	PhysCellId_r16       PhysCellId                                         `madatory`
	MeasIdleResultNR_r16 *MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16 `optional`
}

func (ie *MeasResultsPerCellIdleNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasIdleResultNR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r16", err)
	}
	if ie.MeasIdleResultNR_r16 != nil {
		if err = ie.MeasIdleResultNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdleResultNR_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultsPerCellIdleNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var MeasIdleResultNR_r16Present bool
	if MeasIdleResultNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r16", err)
	}
	if MeasIdleResultNR_r16Present {
		ie.MeasIdleResultNR_r16 = new(MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16)
		if err = ie.MeasIdleResultNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasIdleResultNR_r16", err)
		}
	}
	return nil
}
