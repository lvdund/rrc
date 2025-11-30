package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms0p5   aper.Enumerated = 0
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms0p625 aper.Enumerated = 1
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms1     aper.Enumerated = 2
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms1p25  aper.Enumerated = 3
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms2     aper.Enumerated = 4
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms2p5   aper.Enumerated = 5
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms5     aper.Enumerated = 6
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms10    aper.Enumerated = 7
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms20    aper.Enumerated = 8
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms40    aper.Enumerated = 9
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms80    aper.Enumerated = 10
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms160   aper.Enumerated = 11
)

type IAB_ResourceConfig_r17_periodicitySlotList_r17 struct {
	Value aper.Enumerated `lb:0,ub:11,madatory`
}

func (ie *IAB_ResourceConfig_r17_periodicitySlotList_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Encode IAB_ResourceConfig_r17_periodicitySlotList_r17", err)
	}
	return nil
}

func (ie *IAB_ResourceConfig_r17_periodicitySlotList_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Decode IAB_ResourceConfig_r17_periodicitySlotList_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
