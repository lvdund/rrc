package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasObjectNR_measCyclePSCell_r17_Enum_ms160  aper.Enumerated = 0
	MeasObjectNR_measCyclePSCell_r17_Enum_ms256  aper.Enumerated = 1
	MeasObjectNR_measCyclePSCell_r17_Enum_ms320  aper.Enumerated = 2
	MeasObjectNR_measCyclePSCell_r17_Enum_ms512  aper.Enumerated = 3
	MeasObjectNR_measCyclePSCell_r17_Enum_ms640  aper.Enumerated = 4
	MeasObjectNR_measCyclePSCell_r17_Enum_ms1024 aper.Enumerated = 5
	MeasObjectNR_measCyclePSCell_r17_Enum_ms1280 aper.Enumerated = 6
	MeasObjectNR_measCyclePSCell_r17_Enum_spare1 aper.Enumerated = 7
)

type MeasObjectNR_measCyclePSCell_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MeasObjectNR_measCyclePSCell_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MeasObjectNR_measCyclePSCell_r17", err)
	}
	return nil
}

func (ie *MeasObjectNR_measCyclePSCell_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MeasObjectNR_measCyclePSCell_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
