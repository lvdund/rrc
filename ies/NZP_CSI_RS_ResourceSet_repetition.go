package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NZP_CSI_RS_ResourceSet_repetition_Enum_on  aper.Enumerated = 0
	NZP_CSI_RS_ResourceSet_repetition_Enum_off aper.Enumerated = 1
)

type NZP_CSI_RS_ResourceSet_repetition struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *NZP_CSI_RS_ResourceSet_repetition) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode NZP_CSI_RS_ResourceSet_repetition", err)
	}
	return nil
}

func (ie *NZP_CSI_RS_ResourceSet_repetition) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode NZP_CSI_RS_ResourceSet_repetition", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
