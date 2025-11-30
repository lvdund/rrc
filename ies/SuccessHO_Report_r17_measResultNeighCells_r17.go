package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17_measResultNeighCells_r17 struct {
	MeasResultListNR_r17    *MeasResultList2NR_r16    `optional`
	MeasResultListEUTRA_r17 *MeasResultList2EUTRA_r16 `optional`
}

func (ie *SuccessHO_Report_r17_measResultNeighCells_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultListNR_r17 != nil, ie.MeasResultListEUTRA_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasResultListNR_r17 != nil {
		if err = ie.MeasResultListNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultListNR_r17", err)
		}
	}
	if ie.MeasResultListEUTRA_r17 != nil {
		if err = ie.MeasResultListEUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultListEUTRA_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17_measResultNeighCells_r17) Decode(r *aper.AperReader) error {
	var err error
	var MeasResultListNR_r17Present bool
	if MeasResultListNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultListEUTRA_r17Present bool
	if MeasResultListEUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasResultListNR_r17Present {
		ie.MeasResultListNR_r17 = new(MeasResultList2NR_r16)
		if err = ie.MeasResultListNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListNR_r17", err)
		}
	}
	if MeasResultListEUTRA_r17Present {
		ie.MeasResultListEUTRA_r17 = new(MeasResultList2EUTRA_r16)
		if err = ie.MeasResultListEUTRA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListEUTRA_r17", err)
		}
	}
	return nil
}
