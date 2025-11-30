package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedParameters_v1700 struct {
	MeasurementEnhancementCA_r17        *HighSpeedParameters_v1700_measurementEnhancementCA_r17        `optional`
	MeasurementEnhancementInterFreq_r17 *HighSpeedParameters_v1700_measurementEnhancementInterFreq_r17 `optional`
}

func (ie *HighSpeedParameters_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasurementEnhancementCA_r17 != nil, ie.MeasurementEnhancementInterFreq_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasurementEnhancementCA_r17 != nil {
		if err = ie.MeasurementEnhancementCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasurementEnhancementCA_r17", err)
		}
	}
	if ie.MeasurementEnhancementInterFreq_r17 != nil {
		if err = ie.MeasurementEnhancementInterFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasurementEnhancementInterFreq_r17", err)
		}
	}
	return nil
}

func (ie *HighSpeedParameters_v1700) Decode(r *aper.AperReader) error {
	var err error
	var MeasurementEnhancementCA_r17Present bool
	if MeasurementEnhancementCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasurementEnhancementInterFreq_r17Present bool
	if MeasurementEnhancementInterFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasurementEnhancementCA_r17Present {
		ie.MeasurementEnhancementCA_r17 = new(HighSpeedParameters_v1700_measurementEnhancementCA_r17)
		if err = ie.MeasurementEnhancementCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementEnhancementCA_r17", err)
		}
	}
	if MeasurementEnhancementInterFreq_r17Present {
		ie.MeasurementEnhancementInterFreq_r17 = new(HighSpeedParameters_v1700_measurementEnhancementInterFreq_r17)
		if err = ie.MeasurementEnhancementInterFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementEnhancementInterFreq_r17", err)
		}
	}
	return nil
}
