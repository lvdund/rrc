package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	IAB_IP_Usage_r16_Enum_f1_C   aper.Enumerated = 0
	IAB_IP_Usage_r16_Enum_f1_U   aper.Enumerated = 1
	IAB_IP_Usage_r16_Enum_non_F1 aper.Enumerated = 2
	IAB_IP_Usage_r16_Enum_spare  aper.Enumerated = 3
)

type IAB_IP_Usage_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *IAB_IP_Usage_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode IAB_IP_Usage_r16", err)
	}
	return nil
}

func (ie *IAB_IP_Usage_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode IAB_IP_Usage_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
