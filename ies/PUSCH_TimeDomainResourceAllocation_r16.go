package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_TimeDomainResourceAllocation_r16 struct {
	K2_r16                  *int64                 `lb:0,ub:32,optional`
	PuschAllocationList_r16 []PUSCH_Allocation_r16 `lb:1,ub:maxNrofMultiplePUSCHs_r16,madatory`
}

func (ie *PUSCH_TimeDomainResourceAllocation_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.K2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.K2_r16 != nil {
		if err = w.WriteInteger(*ie.K2_r16, &aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode K2_r16", err)
		}
	}
	tmp_PuschAllocationList_r16 := utils.NewSequence[*PUSCH_Allocation_r16]([]*PUSCH_Allocation_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofMultiplePUSCHs_r16}, false)
	for _, i := range ie.PuschAllocationList_r16 {
		tmp_PuschAllocationList_r16.Value = append(tmp_PuschAllocationList_r16.Value, &i)
	}
	if err = tmp_PuschAllocationList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PuschAllocationList_r16", err)
	}
	return nil
}

func (ie *PUSCH_TimeDomainResourceAllocation_r16) Decode(r *aper.AperReader) error {
	var err error
	var K2_r16Present bool
	if K2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if K2_r16Present {
		var tmp_int_K2_r16 int64
		if tmp_int_K2_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode K2_r16", err)
		}
		ie.K2_r16 = &tmp_int_K2_r16
	}
	tmp_PuschAllocationList_r16 := utils.NewSequence[*PUSCH_Allocation_r16]([]*PUSCH_Allocation_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofMultiplePUSCHs_r16}, false)
	fn_PuschAllocationList_r16 := func() *PUSCH_Allocation_r16 {
		return new(PUSCH_Allocation_r16)
	}
	if err = tmp_PuschAllocationList_r16.Decode(r, fn_PuschAllocationList_r16); err != nil {
		return utils.WrapError("Decode PuschAllocationList_r16", err)
	}
	ie.PuschAllocationList_r16 = []PUSCH_Allocation_r16{}
	for _, i := range tmp_PuschAllocationList_r16.Value {
		ie.PuschAllocationList_r16 = append(ie.PuschAllocationList_r16, *i)
	}
	return nil
}
