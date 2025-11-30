package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedConfig_v1700 struct {
	HighSpeedMeasCA_Scell_r17  *HighSpeedConfig_v1700_highSpeedMeasCA_Scell_r17  `optional`
	HighSpeedMeasInterFreq_r17 *HighSpeedConfig_v1700_highSpeedMeasInterFreq_r17 `optional`
	HighSpeedDemodCA_Scell_r17 *HighSpeedConfig_v1700_highSpeedDemodCA_Scell_r17 `optional`
}

func (ie *HighSpeedConfig_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.HighSpeedMeasCA_Scell_r17 != nil, ie.HighSpeedMeasInterFreq_r17 != nil, ie.HighSpeedDemodCA_Scell_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.HighSpeedMeasCA_Scell_r17 != nil {
		if err = ie.HighSpeedMeasCA_Scell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedMeasCA_Scell_r17", err)
		}
	}
	if ie.HighSpeedMeasInterFreq_r17 != nil {
		if err = ie.HighSpeedMeasInterFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedMeasInterFreq_r17", err)
		}
	}
	if ie.HighSpeedDemodCA_Scell_r17 != nil {
		if err = ie.HighSpeedDemodCA_Scell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedDemodCA_Scell_r17", err)
		}
	}
	return nil
}

func (ie *HighSpeedConfig_v1700) Decode(r *aper.AperReader) error {
	var err error
	var HighSpeedMeasCA_Scell_r17Present bool
	if HighSpeedMeasCA_Scell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedMeasInterFreq_r17Present bool
	if HighSpeedMeasInterFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedDemodCA_Scell_r17Present bool
	if HighSpeedDemodCA_Scell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if HighSpeedMeasCA_Scell_r17Present {
		ie.HighSpeedMeasCA_Scell_r17 = new(HighSpeedConfig_v1700_highSpeedMeasCA_Scell_r17)
		if err = ie.HighSpeedMeasCA_Scell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedMeasCA_Scell_r17", err)
		}
	}
	if HighSpeedMeasInterFreq_r17Present {
		ie.HighSpeedMeasInterFreq_r17 = new(HighSpeedConfig_v1700_highSpeedMeasInterFreq_r17)
		if err = ie.HighSpeedMeasInterFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedMeasInterFreq_r17", err)
		}
	}
	if HighSpeedDemodCA_Scell_r17Present {
		ie.HighSpeedDemodCA_Scell_r17 = new(HighSpeedConfig_v1700_highSpeedDemodCA_Scell_r17)
		if err = ie.HighSpeedDemodCA_Scell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedDemodCA_Scell_r17", err)
		}
	}
	return nil
}
