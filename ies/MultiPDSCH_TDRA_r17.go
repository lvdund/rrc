package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MultiPDSCH_TDRA_r17 struct {
	Pdsch_TDRA_List_r17 []PDSCH_TimeDomainResourceAllocation_r16 `lb:1,ub:maxNrofMultiplePDSCHs_r17,madatory`
}

func (ie *MultiPDSCH_TDRA_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Pdsch_TDRA_List_r17 := utils.NewSequence[*PDSCH_TimeDomainResourceAllocation_r16]([]*PDSCH_TimeDomainResourceAllocation_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofMultiplePDSCHs_r17}, false)
	for _, i := range ie.Pdsch_TDRA_List_r17 {
		tmp_Pdsch_TDRA_List_r17.Value = append(tmp_Pdsch_TDRA_List_r17.Value, &i)
	}
	if err = tmp_Pdsch_TDRA_List_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pdsch_TDRA_List_r17", err)
	}
	return nil
}

func (ie *MultiPDSCH_TDRA_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_Pdsch_TDRA_List_r17 := utils.NewSequence[*PDSCH_TimeDomainResourceAllocation_r16]([]*PDSCH_TimeDomainResourceAllocation_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofMultiplePDSCHs_r17}, false)
	fn_Pdsch_TDRA_List_r17 := func() *PDSCH_TimeDomainResourceAllocation_r16 {
		return new(PDSCH_TimeDomainResourceAllocation_r16)
	}
	if err = tmp_Pdsch_TDRA_List_r17.Decode(r, fn_Pdsch_TDRA_List_r17); err != nil {
		return utils.WrapError("Decode Pdsch_TDRA_List_r17", err)
	}
	ie.Pdsch_TDRA_List_r17 = []PDSCH_TimeDomainResourceAllocation_r16{}
	for _, i := range tmp_Pdsch_TDRA_List_r17.Value {
		ie.Pdsch_TDRA_List_r17 = append(ie.Pdsch_TDRA_List_r17, *i)
	}
	return nil
}
