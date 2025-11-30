package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n2     aper.Enumerated = 0
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n4     aper.Enumerated = 1
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n6     aper.Enumerated = 2
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n8     aper.Enumerated = 3
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n16    aper.Enumerated = 4
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n32    aper.Enumerated = 5
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_spare2 aper.Enumerated = 6
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_spare1 aper.Enumerated = 7
)

type RepFactorAndTimeGap_r17_repetitionFactor_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RepFactorAndTimeGap_r17_repetitionFactor_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RepFactorAndTimeGap_r17_repetitionFactor_r17", err)
	}
	return nil
}

func (ie *RepFactorAndTimeGap_r17_repetitionFactor_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RepFactorAndTimeGap_r17_repetitionFactor_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
