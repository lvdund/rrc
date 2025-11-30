package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB_Enum_supported aper.Enumerated = 0
)

type NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB", err)
	}
	return nil
}

func (ie *NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
