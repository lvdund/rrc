package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_UplinkConfig_transformPrecoderDisabled struct {
	FrequencyDensity      []int64                                                            `lb:2,ub:2,e_lb:0,e_ub:0,optional`
	TimeDensity           []int64                                                            `lb:3,ub:3,e_lb:0,e_ub:0,optional`
	MaxNrofPorts          PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts           `madatory`
	ResourceElementOffset *PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset `optional`
	Ptrs_Power            PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power             `madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.FrequencyDensity) > 0, len(ie.TimeDensity) > 0, ie.ResourceElementOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.FrequencyDensity) > 0 {
		tmp_FrequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.FrequencyDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_FrequencyDensity.Value = append(tmp_FrequencyDensity.Value, &tmp_ie)
		}
		if err = tmp_FrequencyDensity.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyDensity", err)
		}
	}
	if len(ie.TimeDensity) > 0 {
		tmp_TimeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.TimeDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_TimeDensity.Value = append(tmp_TimeDensity.Value, &tmp_ie)
		}
		if err = tmp_TimeDensity.Encode(w); err != nil {
			return utils.WrapError("Encode TimeDensity", err)
		}
	}
	if err = ie.MaxNrofPorts.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNrofPorts", err)
	}
	if ie.ResourceElementOffset != nil {
		if err = ie.ResourceElementOffset.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceElementOffset", err)
		}
	}
	if err = ie.Ptrs_Power.Encode(w); err != nil {
		return utils.WrapError("Encode Ptrs_Power", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled) Decode(r *uper.UperReader) error {
	var err error
	var FrequencyDensityPresent bool
	if FrequencyDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeDensityPresent bool
	if TimeDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceElementOffsetPresent bool
	if ResourceElementOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyDensityPresent {
		tmp_FrequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_FrequencyDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_FrequencyDensity.Decode(r, fn_FrequencyDensity); err != nil {
			return utils.WrapError("Decode FrequencyDensity", err)
		}
		ie.FrequencyDensity = []int64{}
		for _, i := range tmp_FrequencyDensity.Value {
			ie.FrequencyDensity = append(ie.FrequencyDensity, int64(i.Value))
		}
	}
	if TimeDensityPresent {
		tmp_TimeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_TimeDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_TimeDensity.Decode(r, fn_TimeDensity); err != nil {
			return utils.WrapError("Decode TimeDensity", err)
		}
		ie.TimeDensity = []int64{}
		for _, i := range tmp_TimeDensity.Value {
			ie.TimeDensity = append(ie.TimeDensity, int64(i.Value))
		}
	}
	if err = ie.MaxNrofPorts.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNrofPorts", err)
	}
	if ResourceElementOffsetPresent {
		ie.ResourceElementOffset = new(PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset)
		if err = ie.ResourceElementOffset.Decode(r); err != nil {
			return utils.WrapError("Decode ResourceElementOffset", err)
		}
	}
	if err = ie.Ptrs_Power.Decode(r); err != nil {
		return utils.WrapError("Decode Ptrs_Power", err)
	}
	return nil
}
