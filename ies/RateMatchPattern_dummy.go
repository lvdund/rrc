package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPattern_dummy_Enum_dynamic    aper.Enumerated = 0
	RateMatchPattern_dummy_Enum_semiStatic aper.Enumerated = 1
)

type RateMatchPattern_dummy struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RateMatchPattern_dummy) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RateMatchPattern_dummy", err)
	}
	return nil
}

func (ie *RateMatchPattern_dummy) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RateMatchPattern_dummy", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
