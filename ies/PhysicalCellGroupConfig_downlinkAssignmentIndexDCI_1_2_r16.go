package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16_Enum_n1 aper.Enumerated = 0
	PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16_Enum_n2 aper.Enumerated = 1
	PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16_Enum_n4 aper.Enumerated = 2
)

type PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
