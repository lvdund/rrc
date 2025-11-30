package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_TimeDomainResourceAllocationList_r16 struct {
	Value []PUSCH_TimeDomainResourceAllocation_r16 `lb:1,ub:maxNrofUL_Allocations_r16,madatory`
}

func (ie *PUSCH_TimeDomainResourceAllocationList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PUSCH_TimeDomainResourceAllocation_r16]([]*PUSCH_TimeDomainResourceAllocation_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PUSCH_TimeDomainResourceAllocationList_r16", err)
	}
	return nil
}

func (ie *PUSCH_TimeDomainResourceAllocationList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PUSCH_TimeDomainResourceAllocation_r16]([]*PUSCH_TimeDomainResourceAllocation_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
	fn := func() *PUSCH_TimeDomainResourceAllocation_r16 {
		return new(PUSCH_TimeDomainResourceAllocation_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PUSCH_TimeDomainResourceAllocationList_r16", err)
	}
	ie.Value = []PUSCH_TimeDomainResourceAllocation_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
