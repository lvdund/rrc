package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_TimersAndConstants_t310_Enum_ms0    aper.Enumerated = 0
	RLF_TimersAndConstants_t310_Enum_ms50   aper.Enumerated = 1
	RLF_TimersAndConstants_t310_Enum_ms100  aper.Enumerated = 2
	RLF_TimersAndConstants_t310_Enum_ms200  aper.Enumerated = 3
	RLF_TimersAndConstants_t310_Enum_ms500  aper.Enumerated = 4
	RLF_TimersAndConstants_t310_Enum_ms1000 aper.Enumerated = 5
	RLF_TimersAndConstants_t310_Enum_ms2000 aper.Enumerated = 6
	RLF_TimersAndConstants_t310_Enum_ms4000 aper.Enumerated = 7
	RLF_TimersAndConstants_t310_Enum_ms6000 aper.Enumerated = 8
)

type RLF_TimersAndConstants_t310 struct {
	Value aper.Enumerated `lb:0,ub:8,madatory`
}

func (ie *RLF_TimersAndConstants_t310) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Encode RLF_TimersAndConstants_t310", err)
	}
	return nil
}

func (ie *RLF_TimersAndConstants_t310) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Decode RLF_TimersAndConstants_t310", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
