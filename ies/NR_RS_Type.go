package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_RS_Type_Enum_ssb    aper.Enumerated = 0
	NR_RS_Type_Enum_csi_rs aper.Enumerated = 1
)

type NR_RS_Type struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *NR_RS_Type) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode NR_RS_Type", err)
	}
	return nil
}

func (ie *NR_RS_Type) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode NR_RS_Type", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
