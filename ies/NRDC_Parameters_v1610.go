package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters_v1610 struct {
	MeasAndMobParametersNRDC_v1610 *MeasAndMobParametersMRDC_v1610 `optional`
}

func (ie *NRDC_Parameters_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersNRDC_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersNRDC_v1610 != nil {
		if err = ie.MeasAndMobParametersNRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersNRDC_v1610", err)
		}
	}
	return nil
}

func (ie *NRDC_Parameters_v1610) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersNRDC_v1610Present bool
	if MeasAndMobParametersNRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersNRDC_v1610Present {
		ie.MeasAndMobParametersNRDC_v1610 = new(MeasAndMobParametersMRDC_v1610)
		if err = ie.MeasAndMobParametersNRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersNRDC_v1610", err)
		}
	}
	return nil
}
