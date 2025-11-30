package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ProcessingParameters_fallback_Enum_sc        aper.Enumerated = 0
	ProcessingParameters_fallback_Enum_cap1_only aper.Enumerated = 1
)

type ProcessingParameters_fallback struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ProcessingParameters_fallback) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ProcessingParameters_fallback", err)
	}
	return nil
}

func (ie *ProcessingParameters_fallback) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ProcessingParameters_fallback", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
