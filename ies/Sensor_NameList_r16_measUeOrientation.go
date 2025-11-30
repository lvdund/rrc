package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Sensor_NameList_r16_measUeOrientation_Enum_true aper.Enumerated = 0
)

type Sensor_NameList_r16_measUeOrientation struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Sensor_NameList_r16_measUeOrientation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Sensor_NameList_r16_measUeOrientation", err)
	}
	return nil
}

func (ie *Sensor_NameList_r16_measUeOrientation) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Sensor_NameList_r16_measUeOrientation", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
