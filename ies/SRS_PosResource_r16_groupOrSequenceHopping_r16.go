package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResource_r16_groupOrSequenceHopping_r16_Enum_neither         aper.Enumerated = 0
	SRS_PosResource_r16_groupOrSequenceHopping_r16_Enum_groupHopping    aper.Enumerated = 1
	SRS_PosResource_r16_groupOrSequenceHopping_r16_Enum_sequenceHopping aper.Enumerated = 2
)

type SRS_PosResource_r16_groupOrSequenceHopping_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SRS_PosResource_r16_groupOrSequenceHopping_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SRS_PosResource_r16_groupOrSequenceHopping_r16", err)
	}
	return nil
}

func (ie *SRS_PosResource_r16_groupOrSequenceHopping_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SRS_PosResource_r16_groupOrSequenceHopping_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
