package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters_v1700 struct {
	F1c_OverNR_RRC_r17             *NRDC_Parameters_v1700_f1c_OverNR_RRC_r17 `optional`
	MeasAndMobParametersNRDC_v1700 MeasAndMobParametersMRDC_v1700            `madatory`
}

func (ie *NRDC_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.F1c_OverNR_RRC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.F1c_OverNR_RRC_r17 != nil {
		if err = ie.F1c_OverNR_RRC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode F1c_OverNR_RRC_r17", err)
		}
	}
	if err = ie.MeasAndMobParametersNRDC_v1700.Encode(w); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersNRDC_v1700", err)
	}
	return nil
}

func (ie *NRDC_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var F1c_OverNR_RRC_r17Present bool
	if F1c_OverNR_RRC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if F1c_OverNR_RRC_r17Present {
		ie.F1c_OverNR_RRC_r17 = new(NRDC_Parameters_v1700_f1c_OverNR_RRC_r17)
		if err = ie.F1c_OverNR_RRC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode F1c_OverNR_RRC_r17", err)
		}
	}
	if err = ie.MeasAndMobParametersNRDC_v1700.Decode(r); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersNRDC_v1700", err)
	}
	return nil
}
