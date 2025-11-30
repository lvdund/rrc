package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NumberOfMsg3_Repetitions_r17_Enum_n1  aper.Enumerated = 0
	NumberOfMsg3_Repetitions_r17_Enum_n2  aper.Enumerated = 1
	NumberOfMsg3_Repetitions_r17_Enum_n3  aper.Enumerated = 2
	NumberOfMsg3_Repetitions_r17_Enum_n4  aper.Enumerated = 3
	NumberOfMsg3_Repetitions_r17_Enum_n7  aper.Enumerated = 4
	NumberOfMsg3_Repetitions_r17_Enum_n8  aper.Enumerated = 5
	NumberOfMsg3_Repetitions_r17_Enum_n12 aper.Enumerated = 6
	NumberOfMsg3_Repetitions_r17_Enum_n16 aper.Enumerated = 7
)

type NumberOfMsg3_Repetitions_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *NumberOfMsg3_Repetitions_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode NumberOfMsg3_Repetitions_r17", err)
	}
	return nil
}

func (ie *NumberOfMsg3_Repetitions_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode NumberOfMsg3_Repetitions_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
