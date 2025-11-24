package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2EUTRA struct {
	CarrierFreq             ARFCN_ValueEUTRA `madatory`
	MeasResultServingCell   *MeasResultEUTRA `optional`
	MeasResultBestNeighCell *MeasResultEUTRA `optional`
}

func (ie *MeasResult2EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultServingCell != nil, ie.MeasResultBestNeighCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq", err)
	}
	if ie.MeasResultServingCell != nil {
		if err = ie.MeasResultServingCell.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultServingCell", err)
		}
	}
	if ie.MeasResultBestNeighCell != nil {
		if err = ie.MeasResultBestNeighCell.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultBestNeighCell", err)
		}
	}
	return nil
}

func (ie *MeasResult2EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var MeasResultServingCellPresent bool
	if MeasResultServingCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultBestNeighCellPresent bool
	if MeasResultBestNeighCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq", err)
	}
	if MeasResultServingCellPresent {
		ie.MeasResultServingCell = new(MeasResultEUTRA)
		if err = ie.MeasResultServingCell.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultServingCell", err)
		}
	}
	if MeasResultBestNeighCellPresent {
		ie.MeasResultBestNeighCell = new(MeasResultEUTRA)
		if err = ie.MeasResultBestNeighCell.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultBestNeighCell", err)
		}
	}
	return nil
}
