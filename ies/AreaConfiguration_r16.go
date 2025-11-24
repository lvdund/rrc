package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AreaConfiguration_r16 struct {
	AreaConfig_r16          AreaConfig_r16            `madatory`
	InterFreqTargetList_r16 []InterFreqTargetInfo_r16 `lb:1,ub:maxFreq,optional`
}

func (ie *AreaConfiguration_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.InterFreqTargetList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.AreaConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AreaConfig_r16", err)
	}
	if len(ie.InterFreqTargetList_r16) > 0 {
		tmp_InterFreqTargetList_r16 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		for _, i := range ie.InterFreqTargetList_r16 {
			tmp_InterFreqTargetList_r16.Value = append(tmp_InterFreqTargetList_r16.Value, &i)
		}
		if err = tmp_InterFreqTargetList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqTargetList_r16", err)
		}
	}
	return nil
}

func (ie *AreaConfiguration_r16) Decode(r *uper.UperReader) error {
	var err error
	var InterFreqTargetList_r16Present bool
	if InterFreqTargetList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.AreaConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AreaConfig_r16", err)
	}
	if InterFreqTargetList_r16Present {
		tmp_InterFreqTargetList_r16 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		fn_InterFreqTargetList_r16 := func() *InterFreqTargetInfo_r16 {
			return new(InterFreqTargetInfo_r16)
		}
		if err = tmp_InterFreqTargetList_r16.Decode(r, fn_InterFreqTargetList_r16); err != nil {
			return utils.WrapError("Decode InterFreqTargetList_r16", err)
		}
		ie.InterFreqTargetList_r16 = []InterFreqTargetInfo_r16{}
		for _, i := range tmp_InterFreqTargetList_r16.Value {
			ie.InterFreqTargetList_r16 = append(ie.InterFreqTargetList_r16, *i)
		}
	}
	return nil
}
