package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_TimersAndConstants_n311_Enum_n1  aper.Enumerated = 0
	RLF_TimersAndConstants_n311_Enum_n2  aper.Enumerated = 1
	RLF_TimersAndConstants_n311_Enum_n3  aper.Enumerated = 2
	RLF_TimersAndConstants_n311_Enum_n4  aper.Enumerated = 3
	RLF_TimersAndConstants_n311_Enum_n5  aper.Enumerated = 4
	RLF_TimersAndConstants_n311_Enum_n6  aper.Enumerated = 5
	RLF_TimersAndConstants_n311_Enum_n8  aper.Enumerated = 6
	RLF_TimersAndConstants_n311_Enum_n10 aper.Enumerated = 7
)

type RLF_TimersAndConstants_n311 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RLF_TimersAndConstants_n311) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RLF_TimersAndConstants_n311", err)
	}
	return nil
}

func (ie *RLF_TimersAndConstants_n311) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RLF_TimersAndConstants_n311", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
