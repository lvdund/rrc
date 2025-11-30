package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TimersAndConstants_t311_Enum_ms1000  aper.Enumerated = 0
	UE_TimersAndConstants_t311_Enum_ms3000  aper.Enumerated = 1
	UE_TimersAndConstants_t311_Enum_ms5000  aper.Enumerated = 2
	UE_TimersAndConstants_t311_Enum_ms10000 aper.Enumerated = 3
	UE_TimersAndConstants_t311_Enum_ms15000 aper.Enumerated = 4
	UE_TimersAndConstants_t311_Enum_ms20000 aper.Enumerated = 5
	UE_TimersAndConstants_t311_Enum_ms30000 aper.Enumerated = 6
)

type UE_TimersAndConstants_t311 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *UE_TimersAndConstants_t311) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode UE_TimersAndConstants_t311", err)
	}
	return nil
}

func (ie *UE_TimersAndConstants_t311) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode UE_TimersAndConstants_t311", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
