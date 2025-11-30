package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	VictimSystemType_galileo_Enum_true aper.Enumerated = 0
)

type VictimSystemType_galileo struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *VictimSystemType_galileo) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode VictimSystemType_galileo", err)
	}
	return nil
}

func (ie *VictimSystemType_galileo) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode VictimSystemType_galileo", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
