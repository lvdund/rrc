package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_intraBandENDC_Support_Enum_non_contiguous aper.Enumerated = 0
	MRDC_Parameters_intraBandENDC_Support_Enum_both           aper.Enumerated = 1
)

type MRDC_Parameters_intraBandENDC_Support struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MRDC_Parameters_intraBandENDC_Support) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_intraBandENDC_Support", err)
	}
	return nil
}

func (ie *MRDC_Parameters_intraBandENDC_Support) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_intraBandENDC_Support", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
