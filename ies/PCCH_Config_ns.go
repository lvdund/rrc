package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCCH_Config_ns_Enum_four aper.Enumerated = 0
	PCCH_Config_ns_Enum_two  aper.Enumerated = 1
	PCCH_Config_ns_Enum_one  aper.Enumerated = 2
)

type PCCH_Config_ns struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PCCH_Config_ns) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PCCH_Config_ns", err)
	}
	return nil
}

func (ie *PCCH_Config_ns) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PCCH_Config_ns", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
