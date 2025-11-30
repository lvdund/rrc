package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SN_FieldLengthUM_Enum_size6  aper.Enumerated = 0
	SN_FieldLengthUM_Enum_size12 aper.Enumerated = 1
)

type SN_FieldLengthUM struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SN_FieldLengthUM) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SN_FieldLengthUM", err)
	}
	return nil
}

func (ie *SN_FieldLengthUM) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SN_FieldLengthUM", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
