package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Sensor_NameList_r16 struct {
	MeasUncomBarPre_r16 *Sensor_NameList_r16_measUncomBarPre_r16 `optional`
	MeasUeSpeed         *Sensor_NameList_r16_measUeSpeed         `optional`
	MeasUeOrientation   *Sensor_NameList_r16_measUeOrientation   `optional`
}

func (ie *Sensor_NameList_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasUncomBarPre_r16 != nil, ie.MeasUeSpeed != nil, ie.MeasUeOrientation != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasUncomBarPre_r16 != nil {
		if err = ie.MeasUncomBarPre_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasUncomBarPre_r16", err)
		}
	}
	if ie.MeasUeSpeed != nil {
		if err = ie.MeasUeSpeed.Encode(w); err != nil {
			return utils.WrapError("Encode MeasUeSpeed", err)
		}
	}
	if ie.MeasUeOrientation != nil {
		if err = ie.MeasUeOrientation.Encode(w); err != nil {
			return utils.WrapError("Encode MeasUeOrientation", err)
		}
	}
	return nil
}

func (ie *Sensor_NameList_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasUncomBarPre_r16Present bool
	if MeasUncomBarPre_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasUeSpeedPresent bool
	if MeasUeSpeedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasUeOrientationPresent bool
	if MeasUeOrientationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasUncomBarPre_r16Present {
		ie.MeasUncomBarPre_r16 = new(Sensor_NameList_r16_measUncomBarPre_r16)
		if err = ie.MeasUncomBarPre_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasUncomBarPre_r16", err)
		}
	}
	if MeasUeSpeedPresent {
		ie.MeasUeSpeed = new(Sensor_NameList_r16_measUeSpeed)
		if err = ie.MeasUeSpeed.Decode(r); err != nil {
			return utils.WrapError("Decode MeasUeSpeed", err)
		}
	}
	if MeasUeOrientationPresent {
		ie.MeasUeOrientation = new(Sensor_NameList_r16_measUeOrientation)
		if err = ie.MeasUeOrientation.Decode(r); err != nil {
			return utils.WrapError("Decode MeasUeOrientation", err)
		}
	}
	return nil
}
