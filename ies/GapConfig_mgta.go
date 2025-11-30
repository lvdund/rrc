package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_mgta_Enum_ms0      aper.Enumerated = 0
	GapConfig_mgta_Enum_ms0dot25 aper.Enumerated = 1
	GapConfig_mgta_Enum_ms0dot5  aper.Enumerated = 2
)

type GapConfig_mgta struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *GapConfig_mgta) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode GapConfig_mgta", err)
	}
	return nil
}

func (ie *GapConfig_mgta) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode GapConfig_mgta", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
