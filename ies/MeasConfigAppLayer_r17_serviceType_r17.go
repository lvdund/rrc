package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasConfigAppLayer_r17_serviceType_r17_Enum_streaming aper.Enumerated = 0
	MeasConfigAppLayer_r17_serviceType_r17_Enum_mtsi      aper.Enumerated = 1
	MeasConfigAppLayer_r17_serviceType_r17_Enum_vr        aper.Enumerated = 2
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare5    aper.Enumerated = 3
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare4    aper.Enumerated = 4
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare3    aper.Enumerated = 5
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare2    aper.Enumerated = 6
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare1    aper.Enumerated = 7
)

type MeasConfigAppLayer_r17_serviceType_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MeasConfigAppLayer_r17_serviceType_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MeasConfigAppLayer_r17_serviceType_r17", err)
	}
	return nil
}

func (ie *MeasConfigAppLayer_r17_serviceType_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MeasConfigAppLayer_r17_serviceType_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
