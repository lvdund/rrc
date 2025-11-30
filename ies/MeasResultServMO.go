package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServMO struct {
	ServCellId              ServCellIndex `madatory`
	MeasResultServingCell   MeasResultNR  `madatory`
	MeasResultBestNeighCell *MeasResultNR `optional`
}

func (ie *MeasResultServMO) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultBestNeighCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ServCellId.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellId", err)
	}
	if err = ie.MeasResultServingCell.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultServingCell", err)
	}
	if ie.MeasResultBestNeighCell != nil {
		if err = ie.MeasResultBestNeighCell.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultBestNeighCell", err)
		}
	}
	return nil
}

func (ie *MeasResultServMO) Decode(r *aper.AperReader) error {
	var err error
	var MeasResultBestNeighCellPresent bool
	if MeasResultBestNeighCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ServCellId.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellId", err)
	}
	if err = ie.MeasResultServingCell.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultServingCell", err)
	}
	if MeasResultBestNeighCellPresent {
		ie.MeasResultBestNeighCell = new(MeasResultNR)
		if err = ie.MeasResultBestNeighCell.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultBestNeighCell", err)
		}
	}
	return nil
}
