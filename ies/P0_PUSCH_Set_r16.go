package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUSCH_Set_r16 struct {
	P0_PUSCH_SetId_r16 P0_PUSCH_SetId_r16 `madatory`
	P0_List_r16        []P0_PUSCH_r16     `lb:1,ub:maxNrofP0_PUSCH_Set_r16,optional`
}

func (ie *P0_PUSCH_Set_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.P0_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.P0_PUSCH_SetId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode P0_PUSCH_SetId_r16", err)
	}
	if len(ie.P0_List_r16) > 0 {
		tmp_P0_List_r16 := utils.NewSequence[*P0_PUSCH_r16]([]*P0_PUSCH_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_Set_r16}, false)
		for _, i := range ie.P0_List_r16 {
			tmp_P0_List_r16.Value = append(tmp_P0_List_r16.Value, &i)
		}
		if err = tmp_P0_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode P0_List_r16", err)
		}
	}
	return nil
}

func (ie *P0_PUSCH_Set_r16) Decode(r *uper.UperReader) error {
	var err error
	var P0_List_r16Present bool
	if P0_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.P0_PUSCH_SetId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode P0_PUSCH_SetId_r16", err)
	}
	if P0_List_r16Present {
		tmp_P0_List_r16 := utils.NewSequence[*P0_PUSCH_r16]([]*P0_PUSCH_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_Set_r16}, false)
		fn_P0_List_r16 := func() *P0_PUSCH_r16 {
			return new(P0_PUSCH_r16)
		}
		if err = tmp_P0_List_r16.Decode(r, fn_P0_List_r16); err != nil {
			return utils.WrapError("Decode P0_List_r16", err)
		}
		ie.P0_List_r16 = []P0_PUSCH_r16{}
		for _, i := range tmp_P0_List_r16.Value {
			ie.P0_List_r16 = append(ie.P0_List_r16, *i)
		}
	}
	return nil
}
