package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityConfig_keyToUse_Enum_master    aper.Enumerated = 0
	SecurityConfig_keyToUse_Enum_secondary aper.Enumerated = 1
)

type SecurityConfig_keyToUse struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SecurityConfig_keyToUse) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SecurityConfig_keyToUse", err)
	}
	return nil
}

func (ie *SecurityConfig_keyToUse) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SecurityConfig_keyToUse", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
