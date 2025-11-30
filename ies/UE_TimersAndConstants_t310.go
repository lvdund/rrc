package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TimersAndConstants_t310_Enum_ms0    aper.Enumerated = 0
	UE_TimersAndConstants_t310_Enum_ms50   aper.Enumerated = 1
	UE_TimersAndConstants_t310_Enum_ms100  aper.Enumerated = 2
	UE_TimersAndConstants_t310_Enum_ms200  aper.Enumerated = 3
	UE_TimersAndConstants_t310_Enum_ms500  aper.Enumerated = 4
	UE_TimersAndConstants_t310_Enum_ms1000 aper.Enumerated = 5
	UE_TimersAndConstants_t310_Enum_ms2000 aper.Enumerated = 6
)

type UE_TimersAndConstants_t310 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *UE_TimersAndConstants_t310) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode UE_TimersAndConstants_t310", err)
	}
	return nil
}

func (ie *UE_TimersAndConstants_t310) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode UE_TimersAndConstants_t310", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
