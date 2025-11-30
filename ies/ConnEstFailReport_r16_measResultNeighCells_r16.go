package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConnEstFailReport_r16_measResultNeighCells_r16 struct {
	MeasResultNeighCellListNR    *MeasResultList2NR_r16    `optional`
	MeasResultNeighCellListEUTRA *MeasResultList2EUTRA_r16 `optional`
}

func (ie *ConnEstFailReport_r16_measResultNeighCells_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultNeighCellListNR != nil, ie.MeasResultNeighCellListEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasResultNeighCellListNR != nil {
		if err = ie.MeasResultNeighCellListNR.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCellListNR", err)
		}
	}
	if ie.MeasResultNeighCellListEUTRA != nil {
		if err = ie.MeasResultNeighCellListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCellListEUTRA", err)
		}
	}
	return nil
}

func (ie *ConnEstFailReport_r16_measResultNeighCells_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasResultNeighCellListNRPresent bool
	if MeasResultNeighCellListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultNeighCellListEUTRAPresent bool
	if MeasResultNeighCellListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasResultNeighCellListNRPresent {
		ie.MeasResultNeighCellListNR = new(MeasResultList2NR_r16)
		if err = ie.MeasResultNeighCellListNR.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCellListNR", err)
		}
	}
	if MeasResultNeighCellListEUTRAPresent {
		ie.MeasResultNeighCellListEUTRA = new(MeasResultList2EUTRA_r16)
		if err = ie.MeasResultNeighCellListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCellListEUTRA", err)
		}
	}
	return nil
}
