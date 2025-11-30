package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIB_intraFreqReselection_Enum_allowed    aper.Enumerated = 0
	MIB_intraFreqReselection_Enum_notAllowed aper.Enumerated = 1
)

type MIB_intraFreqReselection struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MIB_intraFreqReselection) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MIB_intraFreqReselection", err)
	}
	return nil
}

func (ie *MIB_intraFreqReselection) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MIB_intraFreqReselection", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
