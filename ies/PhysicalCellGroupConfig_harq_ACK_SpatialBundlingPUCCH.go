package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH_Enum_true aper.Enumerated = 0
)

type PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
