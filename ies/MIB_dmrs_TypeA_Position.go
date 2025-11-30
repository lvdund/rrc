package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIB_dmrs_TypeA_Position_Enum_pos2 aper.Enumerated = 0
	MIB_dmrs_TypeA_Position_Enum_pos3 aper.Enumerated = 1
)

type MIB_dmrs_TypeA_Position struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MIB_dmrs_TypeA_Position) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MIB_dmrs_TypeA_Position", err)
	}
	return nil
}

func (ie *MIB_dmrs_TypeA_Position) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MIB_dmrs_TypeA_Position", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
