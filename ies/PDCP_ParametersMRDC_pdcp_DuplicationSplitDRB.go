package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB_Enum_supported aper.Enumerated = 0
)

type PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB", err)
	}
	return nil
}

func (ie *PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
