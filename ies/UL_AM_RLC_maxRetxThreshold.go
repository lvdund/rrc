package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_AM_RLC_maxRetxThreshold_Enum_t1  aper.Enumerated = 0
	UL_AM_RLC_maxRetxThreshold_Enum_t2  aper.Enumerated = 1
	UL_AM_RLC_maxRetxThreshold_Enum_t3  aper.Enumerated = 2
	UL_AM_RLC_maxRetxThreshold_Enum_t4  aper.Enumerated = 3
	UL_AM_RLC_maxRetxThreshold_Enum_t6  aper.Enumerated = 4
	UL_AM_RLC_maxRetxThreshold_Enum_t8  aper.Enumerated = 5
	UL_AM_RLC_maxRetxThreshold_Enum_t16 aper.Enumerated = 6
	UL_AM_RLC_maxRetxThreshold_Enum_t32 aper.Enumerated = 7
)

type UL_AM_RLC_maxRetxThreshold struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UL_AM_RLC_maxRetxThreshold) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UL_AM_RLC_maxRetxThreshold", err)
	}
	return nil
}

func (ie *UL_AM_RLC_maxRetxThreshold) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UL_AM_RLC_maxRetxThreshold", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
