// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PosSI-SchedulingInfo-r16 (line 4877).

var posSISchedulingInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "posSchedulingInfoList-r16"},
		{Name: "posSI-RequestConfig-r16", Optional: true},
		{Name: "posSI-RequestConfigSUL-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var posSISchedulingInfoR16PosSchedulingInfoListR16Constraints = per.SizeRange(1, common.MaxSI_Message)

type PosSI_SchedulingInfo_r16 struct {
	PosSchedulingInfoList_r16                     []PosSchedulingInfo_r16
	PosSI_RequestConfig_r16                       *SI_RequestConfig
	PosSI_RequestConfigSUL_r16                    *SI_RequestConfig
	PosSI_RequestConfigRedCap_r17                 *SI_RequestConfig
	PosSI_RequestConfigMSG1_Repetition_r18        *SI_RequestConfigRepetition_r18
	PosSI_RequestConfigSUL_MSG1_Repetition_r18    *SI_RequestConfigRepetition_r18
	PosSI_RequestConfigRedCap_MSG1_Repetition_r18 *SI_RequestConfigRepetition_r18
}

func (ie *PosSI_SchedulingInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSISchedulingInfoR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.PosSI_RequestConfigRedCap_r17 != nil
	hasExtGroup1 := ie.PosSI_RequestConfigMSG1_Repetition_r18 != nil || ie.PosSI_RequestConfigSUL_MSG1_Repetition_r18 != nil || ie.PosSI_RequestConfigRedCap_MSG1_Repetition_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PosSI_RequestConfig_r16 != nil, ie.PosSI_RequestConfigSUL_r16 != nil}); err != nil {
		return err
	}

	// 3. posSchedulingInfoList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(posSISchedulingInfoR16PosSchedulingInfoListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.PosSchedulingInfoList_r16))); err != nil {
			return err
		}
		for i := range ie.PosSchedulingInfoList_r16 {
			if err := ie.PosSchedulingInfoList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. posSI-RequestConfig-r16: ref
	{
		if ie.PosSI_RequestConfig_r16 != nil {
			if err := ie.PosSI_RequestConfig_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. posSI-RequestConfigSUL-r16: ref
	{
		if ie.PosSI_RequestConfigSUL_r16 != nil {
			if err := ie.PosSI_RequestConfigSUL_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "posSI-RequestConfigRedCap-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PosSI_RequestConfigRedCap_r17 != nil}); err != nil {
				return err
			}

			if ie.PosSI_RequestConfigRedCap_r17 != nil {
				if err := ie.PosSI_RequestConfigRedCap_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "posSI-RequestConfigMSG1-Repetition-r18", Optional: true},
					{Name: "posSI-RequestConfigSUL-MSG1-Repetition-r18", Optional: true},
					{Name: "posSI-RequestConfigRedCap-MSG1-Repetition-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PosSI_RequestConfigMSG1_Repetition_r18 != nil, ie.PosSI_RequestConfigSUL_MSG1_Repetition_r18 != nil, ie.PosSI_RequestConfigRedCap_MSG1_Repetition_r18 != nil}); err != nil {
				return err
			}

			if ie.PosSI_RequestConfigMSG1_Repetition_r18 != nil {
				if err := ie.PosSI_RequestConfigMSG1_Repetition_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosSI_RequestConfigSUL_MSG1_Repetition_r18 != nil {
				if err := ie.PosSI_RequestConfigSUL_MSG1_Repetition_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosSI_RequestConfigRedCap_MSG1_Repetition_r18 != nil {
				if err := ie.PosSI_RequestConfigRedCap_MSG1_Repetition_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PosSI_SchedulingInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSISchedulingInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. posSchedulingInfoList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(posSISchedulingInfoR16PosSchedulingInfoListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.PosSchedulingInfoList_r16 = make([]PosSchedulingInfo_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.PosSchedulingInfoList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. posSI-RequestConfig-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.PosSI_RequestConfig_r16 = new(SI_RequestConfig)
			if err := ie.PosSI_RequestConfig_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. posSI-RequestConfigSUL-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.PosSI_RequestConfigSUL_r16 = new(SI_RequestConfig)
			if err := ie.PosSI_RequestConfigSUL_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "posSI-RequestConfigRedCap-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PosSI_RequestConfigRedCap_r17 = new(SI_RequestConfig)
			if err := ie.PosSI_RequestConfigRedCap_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "posSI-RequestConfigMSG1-Repetition-r18", Optional: true},
				{Name: "posSI-RequestConfigSUL-MSG1-Repetition-r18", Optional: true},
				{Name: "posSI-RequestConfigRedCap-MSG1-Repetition-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PosSI_RequestConfigMSG1_Repetition_r18 = new(SI_RequestConfigRepetition_r18)
			if err := ie.PosSI_RequestConfigMSG1_Repetition_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.PosSI_RequestConfigSUL_MSG1_Repetition_r18 = new(SI_RequestConfigRepetition_r18)
			if err := ie.PosSI_RequestConfigSUL_MSG1_Repetition_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.PosSI_RequestConfigRedCap_MSG1_Repetition_r18 = new(SI_RequestConfigRepetition_r18)
			if err := ie.PosSI_RequestConfigRedCap_MSG1_Repetition_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
