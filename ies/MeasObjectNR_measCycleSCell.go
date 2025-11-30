package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasObjectNR_measCycleSCell_Enum_sf160  aper.Enumerated = 0
	MeasObjectNR_measCycleSCell_Enum_sf256  aper.Enumerated = 1
	MeasObjectNR_measCycleSCell_Enum_sf320  aper.Enumerated = 2
	MeasObjectNR_measCycleSCell_Enum_sf512  aper.Enumerated = 3
	MeasObjectNR_measCycleSCell_Enum_sf640  aper.Enumerated = 4
	MeasObjectNR_measCycleSCell_Enum_sf1024 aper.Enumerated = 5
	MeasObjectNR_measCycleSCell_Enum_sf1280 aper.Enumerated = 6
)

type MeasObjectNR_measCycleSCell struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *MeasObjectNR_measCycleSCell) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode MeasObjectNR_measCycleSCell", err)
	}
	return nil
}

func (ie *MeasObjectNR_measCycleSCell) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode MeasObjectNR_measCycleSCell", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
