package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasGapSharingScheme_Enum_scheme00 aper.Enumerated = 0
	MeasGapSharingScheme_Enum_scheme01 aper.Enumerated = 1
	MeasGapSharingScheme_Enum_scheme10 aper.Enumerated = 2
	MeasGapSharingScheme_Enum_scheme11 aper.Enumerated = 3
)

type MeasGapSharingScheme struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MeasGapSharingScheme) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MeasGapSharingScheme", err)
	}
	return nil
}

func (ie *MeasGapSharingScheme) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MeasGapSharingScheme", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
