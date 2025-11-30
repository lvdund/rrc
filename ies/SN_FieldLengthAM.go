package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SN_FieldLengthAM_Enum_size12 aper.Enumerated = 0
	SN_FieldLengthAM_Enum_size18 aper.Enumerated = 1
)

type SN_FieldLengthAM struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SN_FieldLengthAM) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SN_FieldLengthAM", err)
	}
	return nil
}

func (ie *SN_FieldLengthAM) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SN_FieldLengthAM", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
