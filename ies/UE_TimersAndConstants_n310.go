package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TimersAndConstants_n310_Enum_n1  aper.Enumerated = 0
	UE_TimersAndConstants_n310_Enum_n2  aper.Enumerated = 1
	UE_TimersAndConstants_n310_Enum_n3  aper.Enumerated = 2
	UE_TimersAndConstants_n310_Enum_n4  aper.Enumerated = 3
	UE_TimersAndConstants_n310_Enum_n6  aper.Enumerated = 4
	UE_TimersAndConstants_n310_Enum_n8  aper.Enumerated = 5
	UE_TimersAndConstants_n310_Enum_n10 aper.Enumerated = 6
	UE_TimersAndConstants_n310_Enum_n20 aper.Enumerated = 7
)

type UE_TimersAndConstants_n310 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UE_TimersAndConstants_n310) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UE_TimersAndConstants_n310", err)
	}
	return nil
}

func (ie *UE_TimersAndConstants_n310) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UE_TimersAndConstants_n310", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
