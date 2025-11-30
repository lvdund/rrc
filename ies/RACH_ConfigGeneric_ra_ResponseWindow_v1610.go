package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_ra_ResponseWindow_v1610_Enum_sl60  aper.Enumerated = 0
	RACH_ConfigGeneric_ra_ResponseWindow_v1610_Enum_sl160 aper.Enumerated = 1
)

type RACH_ConfigGeneric_ra_ResponseWindow_v1610 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_ra_ResponseWindow_v1610", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow_v1610) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_ra_ResponseWindow_v1610", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
