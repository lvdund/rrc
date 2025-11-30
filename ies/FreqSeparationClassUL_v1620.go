package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqSeparationClassUL_v1620_Enum_mhz1000 aper.Enumerated = 0
)

type FreqSeparationClassUL_v1620 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FreqSeparationClassUL_v1620) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FreqSeparationClassUL_v1620", err)
	}
	return nil
}

func (ie *FreqSeparationClassUL_v1620) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FreqSeparationClassUL_v1620", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
