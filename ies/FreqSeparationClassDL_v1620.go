package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqSeparationClassDL_v1620_Enum_mhz1000 aper.Enumerated = 0
	FreqSeparationClassDL_v1620_Enum_mhz1600 aper.Enumerated = 1
	FreqSeparationClassDL_v1620_Enum_mhz1800 aper.Enumerated = 2
	FreqSeparationClassDL_v1620_Enum_mhz2000 aper.Enumerated = 3
	FreqSeparationClassDL_v1620_Enum_mhz2200 aper.Enumerated = 4
	FreqSeparationClassDL_v1620_Enum_mhz2400 aper.Enumerated = 5
)

type FreqSeparationClassDL_v1620 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *FreqSeparationClassDL_v1620) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode FreqSeparationClassDL_v1620", err)
	}
	return nil
}

func (ie *FreqSeparationClassDL_v1620) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode FreqSeparationClassDL_v1620", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
