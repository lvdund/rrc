package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_xScale_Enum_dB0    aper.Enumerated = 0
	PhysicalCellGroupConfig_xScale_Enum_dB6    aper.Enumerated = 1
	PhysicalCellGroupConfig_xScale_Enum_spare2 aper.Enumerated = 2
	PhysicalCellGroupConfig_xScale_Enum_spare1 aper.Enumerated = 3
)

type PhysicalCellGroupConfig_xScale struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PhysicalCellGroupConfig_xScale) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_xScale", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_xScale) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_xScale", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
