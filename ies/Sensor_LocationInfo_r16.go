package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Sensor_LocationInfo_r16 struct {
	Sensor_MeasurementInformation_r16 *[]byte `optional`
	Sensor_MotionInformation_r16      *[]byte `optional`
}

func (ie *Sensor_LocationInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sensor_MeasurementInformation_r16 != nil, ie.Sensor_MotionInformation_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sensor_MeasurementInformation_r16 != nil {
		if err = w.WriteOctetString(*ie.Sensor_MeasurementInformation_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Sensor_MeasurementInformation_r16", err)
		}
	}
	if ie.Sensor_MotionInformation_r16 != nil {
		if err = w.WriteOctetString(*ie.Sensor_MotionInformation_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Sensor_MotionInformation_r16", err)
		}
	}
	return nil
}

func (ie *Sensor_LocationInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sensor_MeasurementInformation_r16Present bool
	if Sensor_MeasurementInformation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sensor_MotionInformation_r16Present bool
	if Sensor_MotionInformation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sensor_MeasurementInformation_r16Present {
		var tmp_os_Sensor_MeasurementInformation_r16 []byte
		if tmp_os_Sensor_MeasurementInformation_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Sensor_MeasurementInformation_r16", err)
		}
		ie.Sensor_MeasurementInformation_r16 = &tmp_os_Sensor_MeasurementInformation_r16
	}
	if Sensor_MotionInformation_r16Present {
		var tmp_os_Sensor_MotionInformation_r16 []byte
		if tmp_os_Sensor_MotionInformation_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Sensor_MotionInformation_r16", err)
		}
		ie.Sensor_MotionInformation_r16 = &tmp_os_Sensor_MotionInformation_r16
	}
	return nil
}
