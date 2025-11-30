package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl1  aper.Enumerated = 0
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl2  aper.Enumerated = 1
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl4  aper.Enumerated = 2
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl8  aper.Enumerated = 3
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl10 aper.Enumerated = 4
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl20 aper.Enumerated = 5
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl40 aper.Enumerated = 6
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl80 aper.Enumerated = 7
)

type RACH_ConfigGeneric_ra_ResponseWindow struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_ra_ResponseWindow", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_ra_ResponseWindow", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
