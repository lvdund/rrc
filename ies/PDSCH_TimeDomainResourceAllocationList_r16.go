package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_TimeDomainResourceAllocationList_r16 struct {
	Value []PDSCH_TimeDomainResourceAllocation_r16 `lb:1,ub:maxNrofDL_Allocations,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocationList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PDSCH_TimeDomainResourceAllocation_r16]([]*PDSCH_TimeDomainResourceAllocation_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofDL_Allocations}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PDSCH_TimeDomainResourceAllocationList_r16", err)
	}
	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocationList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PDSCH_TimeDomainResourceAllocation_r16]([]*PDSCH_TimeDomainResourceAllocation_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofDL_Allocations}, false)
	fn := func() *PDSCH_TimeDomainResourceAllocation_r16 {
		return new(PDSCH_TimeDomainResourceAllocation_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PDSCH_TimeDomainResourceAllocationList_r16", err)
	}
	ie.Value = []PDSCH_TimeDomainResourceAllocation_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
