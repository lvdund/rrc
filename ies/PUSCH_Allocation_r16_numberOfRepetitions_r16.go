package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n1  aper.Enumerated = 0
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n2  aper.Enumerated = 1
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n3  aper.Enumerated = 2
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n4  aper.Enumerated = 3
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n7  aper.Enumerated = 4
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n8  aper.Enumerated = 5
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n12 aper.Enumerated = 6
	PUSCH_Allocation_r16_numberOfRepetitions_r16_Enum_n16 aper.Enumerated = 7
)

type PUSCH_Allocation_r16_numberOfRepetitions_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PUSCH_Allocation_r16_numberOfRepetitions_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Allocation_r16_numberOfRepetitions_r16", err)
	}
	return nil
}

func (ie *PUSCH_Allocation_r16_numberOfRepetitions_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Allocation_r16_numberOfRepetitions_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
