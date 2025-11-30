package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BWP_cyclicPrefix_Enum_extended aper.Enumerated = 0
)

type BWP_cyclicPrefix struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BWP_cyclicPrefix) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BWP_cyclicPrefix", err)
	}
	return nil
}

func (ie *BWP_cyclicPrefix) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BWP_cyclicPrefix", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
