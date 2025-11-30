package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_PositionQCL_Relation_r16_Enum_n1 aper.Enumerated = 0
	SSB_PositionQCL_Relation_r16_Enum_n2 aper.Enumerated = 1
	SSB_PositionQCL_Relation_r16_Enum_n4 aper.Enumerated = 2
	SSB_PositionQCL_Relation_r16_Enum_n8 aper.Enumerated = 3
)

type SSB_PositionQCL_Relation_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SSB_PositionQCL_Relation_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SSB_PositionQCL_Relation_r16", err)
	}
	return nil
}

func (ie *SSB_PositionQCL_Relation_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SSB_PositionQCL_Relation_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
