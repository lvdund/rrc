package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedParameters_r16 struct {
	MeasurementEnhancement_r16  *HighSpeedParameters_r16_measurementEnhancement_r16  `optional`
	DemodulationEnhancement_r16 *HighSpeedParameters_r16_demodulationEnhancement_r16 `optional`
}

func (ie *HighSpeedParameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasurementEnhancement_r16 != nil, ie.DemodulationEnhancement_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasurementEnhancement_r16 != nil {
		if err = ie.MeasurementEnhancement_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasurementEnhancement_r16", err)
		}
	}
	if ie.DemodulationEnhancement_r16 != nil {
		if err = ie.DemodulationEnhancement_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DemodulationEnhancement_r16", err)
		}
	}
	return nil
}

func (ie *HighSpeedParameters_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasurementEnhancement_r16Present bool
	if MeasurementEnhancement_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DemodulationEnhancement_r16Present bool
	if DemodulationEnhancement_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasurementEnhancement_r16Present {
		ie.MeasurementEnhancement_r16 = new(HighSpeedParameters_r16_measurementEnhancement_r16)
		if err = ie.MeasurementEnhancement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementEnhancement_r16", err)
		}
	}
	if DemodulationEnhancement_r16Present {
		ie.DemodulationEnhancement_r16 = new(HighSpeedParameters_r16_demodulationEnhancement_r16)
		if err = ie.DemodulationEnhancement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DemodulationEnhancement_r16", err)
		}
	}
	return nil
}
