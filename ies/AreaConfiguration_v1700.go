package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AreaConfiguration_v1700 struct {
	AreaConfig_r17          *AreaConfig_r16           `optional`
	InterFreqTargetList_r17 []InterFreqTargetInfo_r16 `lb:1,ub:maxFreq,optional`
}

func (ie *AreaConfiguration_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AreaConfig_r17 != nil, len(ie.InterFreqTargetList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AreaConfig_r17 != nil {
		if err = ie.AreaConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AreaConfig_r17", err)
		}
	}
	if len(ie.InterFreqTargetList_r17) > 0 {
		tmp_InterFreqTargetList_r17 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		for _, i := range ie.InterFreqTargetList_r17 {
			tmp_InterFreqTargetList_r17.Value = append(tmp_InterFreqTargetList_r17.Value, &i)
		}
		if err = tmp_InterFreqTargetList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqTargetList_r17", err)
		}
	}
	return nil
}

func (ie *AreaConfiguration_v1700) Decode(r *uper.UperReader) error {
	var err error
	var AreaConfig_r17Present bool
	if AreaConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqTargetList_r17Present bool
	if InterFreqTargetList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AreaConfig_r17Present {
		ie.AreaConfig_r17 = new(AreaConfig_r16)
		if err = ie.AreaConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AreaConfig_r17", err)
		}
	}
	if InterFreqTargetList_r17Present {
		tmp_InterFreqTargetList_r17 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		fn_InterFreqTargetList_r17 := func() *InterFreqTargetInfo_r16 {
			return new(InterFreqTargetInfo_r16)
		}
		if err = tmp_InterFreqTargetList_r17.Decode(r, fn_InterFreqTargetList_r17); err != nil {
			return utils.WrapError("Decode InterFreqTargetList_r17", err)
		}
		ie.InterFreqTargetList_r17 = []InterFreqTargetInfo_r16{}
		for _, i := range tmp_InterFreqTargetList_r17.Value {
			ie.InterFreqTargetList_r17 = append(ie.InterFreqTargetList_r17, *i)
		}
	}
	return nil
}
