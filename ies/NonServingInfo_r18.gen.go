// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NonServingInfo-r18 (line 28466).

var nonServingInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "freqInfoMBS-r18", Optional: true},
		{Name: "cfr-InfoMBS-r18", Optional: true},
		{Name: "subcarrierSpacing-r18", Optional: true},
	},
}

var nonServingInfo_r18CfrInfoMBSR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cfr-Bandwidth-r18"},
		{Name: "cfr-LocationAndBW-r18"},
	},
}

const (
	NonServingInfo_r18_Cfr_InfoMBS_r18_Cfr_Bandwidth_r18     = 0
	NonServingInfo_r18_Cfr_InfoMBS_r18_Cfr_LocationAndBW_r18 = 1
)

type NonServingInfo_r18 struct {
	FreqInfoMBS_r18 *FreqInfoMBS_r18
	Cfr_InfoMBS_r18 *struct {
		Choice                int
		Cfr_Bandwidth_r18     *int64
		Cfr_LocationAndBW_r18 *CFR_LocationAndBW_r18
	}
	SubcarrierSpacing_r18 *SubcarrierSpacing
}

func (ie *NonServingInfo_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nonServingInfoR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FreqInfoMBS_r18 != nil, ie.Cfr_InfoMBS_r18 != nil, ie.SubcarrierSpacing_r18 != nil}); err != nil {
		return err
	}

	// 2. freqInfoMBS-r18: ref
	{
		if ie.FreqInfoMBS_r18 != nil {
			if err := ie.FreqInfoMBS_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. cfr-InfoMBS-r18: choice
	{
		if ie.Cfr_InfoMBS_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(nonServingInfo_r18CfrInfoMBSR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Cfr_InfoMBS_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Cfr_InfoMBS_r18).Choice {
			case NonServingInfo_r18_Cfr_InfoMBS_r18_Cfr_Bandwidth_r18:
				if err := e.EncodeInteger((*(*ie.Cfr_InfoMBS_r18).Cfr_Bandwidth_r18), per.Constrained(1, common.MaxNrofPhysicalResourceBlocks)); err != nil {
					return err
				}
			case NonServingInfo_r18_Cfr_InfoMBS_r18_Cfr_LocationAndBW_r18:
				if err := (*ie.Cfr_InfoMBS_r18).Cfr_LocationAndBW_r18.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Cfr_InfoMBS_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. subcarrierSpacing-r18: ref
	{
		if ie.SubcarrierSpacing_r18 != nil {
			if err := ie.SubcarrierSpacing_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NonServingInfo_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nonServingInfoR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. freqInfoMBS-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FreqInfoMBS_r18 = new(FreqInfoMBS_r18)
			if err := ie.FreqInfoMBS_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. cfr-InfoMBS-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Cfr_InfoMBS_r18 = &struct {
				Choice                int
				Cfr_Bandwidth_r18     *int64
				Cfr_LocationAndBW_r18 *CFR_LocationAndBW_r18
			}{}
			choiceDec := d.NewChoiceDecoder(nonServingInfo_r18CfrInfoMBSR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cfr_InfoMBS_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case NonServingInfo_r18_Cfr_InfoMBS_r18_Cfr_Bandwidth_r18:
				(*ie.Cfr_InfoMBS_r18).Cfr_Bandwidth_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofPhysicalResourceBlocks))
				if err != nil {
					return err
				}
				(*(*ie.Cfr_InfoMBS_r18).Cfr_Bandwidth_r18) = v
			case NonServingInfo_r18_Cfr_InfoMBS_r18_Cfr_LocationAndBW_r18:
				(*ie.Cfr_InfoMBS_r18).Cfr_LocationAndBW_r18 = new(CFR_LocationAndBW_r18)
				if err := (*ie.Cfr_InfoMBS_r18).Cfr_LocationAndBW_r18.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. subcarrierSpacing-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SubcarrierSpacing_r18 = new(SubcarrierSpacing)
			if err := ie.SubcarrierSpacing_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
