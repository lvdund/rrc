package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GeneralParametersMRDC_XDD_Diff_srb3_Enum_supported aper.Enumerated = 0
)

type GeneralParametersMRDC_XDD_Diff_srb3 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *GeneralParametersMRDC_XDD_Diff_srb3) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode GeneralParametersMRDC_XDD_Diff_srb3", err)
	}
	return nil
}

func (ie *GeneralParametersMRDC_XDD_Diff_srb3) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode GeneralParametersMRDC_XDD_Diff_srb3", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
