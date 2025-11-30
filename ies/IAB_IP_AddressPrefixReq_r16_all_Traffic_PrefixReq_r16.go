package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16_Enum_true aper.Enumerated = 0
)

type IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16", err)
	}
	return nil
}

func (ie *IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
