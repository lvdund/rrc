package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	IntegrityProtAlgorithm_Enum_nia0   aper.Enumerated = 0
	IntegrityProtAlgorithm_Enum_nia1   aper.Enumerated = 1
	IntegrityProtAlgorithm_Enum_nia2   aper.Enumerated = 2
	IntegrityProtAlgorithm_Enum_nia3   aper.Enumerated = 3
	IntegrityProtAlgorithm_Enum_spare4 aper.Enumerated = 4
	IntegrityProtAlgorithm_Enum_spare3 aper.Enumerated = 5
	IntegrityProtAlgorithm_Enum_spare2 aper.Enumerated = 6
	IntegrityProtAlgorithm_Enum_spare1 aper.Enumerated = 7
)

type IntegrityProtAlgorithm struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *IntegrityProtAlgorithm) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode IntegrityProtAlgorithm", err)
	}
	return nil
}

func (ie *IntegrityProtAlgorithm) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode IntegrityProtAlgorithm", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
