package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16 struct {
	MeasResultListNR_r16    *MeasResultList2NR_r16    `optional`
	MeasResultListEUTRA_r16 *MeasResultList2EUTRA_r16 `optional`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultListNR_r16 != nil, ie.MeasResultListEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasResultListNR_r16 != nil {
		if err = ie.MeasResultListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultListNR_r16", err)
		}
	}
	if ie.MeasResultListEUTRA_r16 != nil {
		if err = ie.MeasResultListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultListEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16) Decode(r *uper.UperReader) error {
	var err error
	var MeasResultListNR_r16Present bool
	if MeasResultListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultListEUTRA_r16Present bool
	if MeasResultListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasResultListNR_r16Present {
		ie.MeasResultListNR_r16 = new(MeasResultList2NR_r16)
		if err = ie.MeasResultListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListNR_r16", err)
		}
	}
	if MeasResultListEUTRA_r16Present {
		ie.MeasResultListEUTRA_r16 = new(MeasResultList2EUTRA_r16)
		if err = ie.MeasResultListEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListEUTRA_r16", err)
		}
	}
	return nil
}
