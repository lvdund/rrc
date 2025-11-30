package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SubcarrierSpacing_Enum_kHz15        aper.Enumerated = 0
	SubcarrierSpacing_Enum_kHz30        aper.Enumerated = 1
	SubcarrierSpacing_Enum_kHz60        aper.Enumerated = 2
	SubcarrierSpacing_Enum_kHz120       aper.Enumerated = 3
	SubcarrierSpacing_Enum_kHz240       aper.Enumerated = 4
	SubcarrierSpacing_Enum_kHz480_v1700 aper.Enumerated = 5
	SubcarrierSpacing_Enum_kHz960_v1700 aper.Enumerated = 6
	SubcarrierSpacing_Enum_spare1       aper.Enumerated = 7
)

type SubcarrierSpacing struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SubcarrierSpacing) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing", err)
	}
	return nil
}

func (ie *SubcarrierSpacing) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
