package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n1     aper.Enumerated = 0
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n2     aper.Enumerated = 1
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n3     aper.Enumerated = 2
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n4     aper.Enumerated = 3
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n7     aper.Enumerated = 4
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n8     aper.Enumerated = 5
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n12    aper.Enumerated = 6
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n16    aper.Enumerated = 7
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n20    aper.Enumerated = 8
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n24    aper.Enumerated = 9
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n28    aper.Enumerated = 10
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n32    aper.Enumerated = 11
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare4 aper.Enumerated = 12
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare3 aper.Enumerated = 13
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare2 aper.Enumerated = 14
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare1 aper.Enumerated = 15
)

type PUSCH_Allocation_r16_numberOfRepetitionsExt_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PUSCH_Allocation_r16_numberOfRepetitionsExt_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Allocation_r16_numberOfRepetitionsExt_r17", err)
	}
	return nil
}

func (ie *PUSCH_Allocation_r16_numberOfRepetitionsExt_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Allocation_r16_numberOfRepetitionsExt_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
