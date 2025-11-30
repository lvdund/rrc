package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_density_dot5_Enum_evenPRBs aper.Enumerated = 0
	CSI_RS_ResourceMapping_density_dot5_Enum_oddPRBs  aper.Enumerated = 1
)

type CSI_RS_ResourceMapping_density_dot5 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CSI_RS_ResourceMapping_density_dot5) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_ResourceMapping_density_dot5", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping_density_dot5) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_ResourceMapping_density_dot5", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
