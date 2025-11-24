package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_ParametersXDD_Diff struct {
	RsrqMeasWidebandEUTRA *EUTRA_ParametersXDD_Diff_rsrqMeasWidebandEUTRA `optional`
}

func (ie *EUTRA_ParametersXDD_Diff) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RsrqMeasWidebandEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RsrqMeasWidebandEUTRA != nil {
		if err = ie.RsrqMeasWidebandEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode RsrqMeasWidebandEUTRA", err)
		}
	}
	return nil
}

func (ie *EUTRA_ParametersXDD_Diff) Decode(r *uper.UperReader) error {
	var err error
	var RsrqMeasWidebandEUTRAPresent bool
	if RsrqMeasWidebandEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if RsrqMeasWidebandEUTRAPresent {
		ie.RsrqMeasWidebandEUTRA = new(EUTRA_ParametersXDD_Diff_rsrqMeasWidebandEUTRA)
		if err = ie.RsrqMeasWidebandEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode RsrqMeasWidebandEUTRA", err)
		}
	}
	return nil
}
