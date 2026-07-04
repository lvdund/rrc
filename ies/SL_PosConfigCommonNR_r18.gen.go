// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PosConfigCommonNR-r18 (line 4644).

var sLPosConfigCommonNRR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PosFreqInfoList-r18", Optional: true},
		{Name: "sl-PosUE-SelectedConfig-r18", Optional: true},
		{Name: "sl-PosNR-AnchorCarrierFreqList-r18", Optional: true},
		{Name: "sl-PosMeasConfigCommon-r18", Optional: true},
		{Name: "sl-PosOffsetDFN-r18", Optional: true},
		{Name: "sl-PosSSB-PriorityNR-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLPosConfigCommonNRR18SlPosFreqInfoListR18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLPosConfigCommonNRR18SlPosOffsetDFNR18Constraints = per.Constrained(1, 1000)

var sLPosConfigCommonNRR18SlPosSSBPriorityNRR18Constraints = per.Constrained(1, 8)

var sLPosConfigCommonNRR18ExtSlPosFreqInfoListExtV1880Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

type SL_PosConfigCommonNR_r18 struct {
	Sl_PosFreqInfoList_r18             []SL_FreqConfigCommon_r16
	Sl_PosUE_SelectedConfig_r18        *SL_UE_SelectedConfig_r16
	Sl_PosNR_AnchorCarrierFreqList_r18 *SL_NR_AnchorCarrierFreqList_r16
	Sl_PosMeasConfigCommon_r18         *SL_MeasConfigCommon_r16
	Sl_PosOffsetDFN_r18                *int64
	Sl_PosSSB_PriorityNR_r18           *int64
	Sl_PosFreqInfoListExt_v1880        []SL_FreqConfigCommonExt_V16k0
}

func (ie *SL_PosConfigCommonNR_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPosConfigCommonNRR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_PosFreqInfoListExt_v1880 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PosFreqInfoList_r18 != nil, ie.Sl_PosUE_SelectedConfig_r18 != nil, ie.Sl_PosNR_AnchorCarrierFreqList_r18 != nil, ie.Sl_PosMeasConfigCommon_r18 != nil, ie.Sl_PosOffsetDFN_r18 != nil, ie.Sl_PosSSB_PriorityNR_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-PosFreqInfoList-r18: sequence-of
	{
		if ie.Sl_PosFreqInfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPosConfigCommonNRR18SlPosFreqInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PosFreqInfoList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PosFreqInfoList_r18 {
				if err := ie.Sl_PosFreqInfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PosUE-SelectedConfig-r18: ref
	{
		if ie.Sl_PosUE_SelectedConfig_r18 != nil {
			if err := ie.Sl_PosUE_SelectedConfig_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-PosNR-AnchorCarrierFreqList-r18: ref
	{
		if ie.Sl_PosNR_AnchorCarrierFreqList_r18 != nil {
			if err := ie.Sl_PosNR_AnchorCarrierFreqList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-PosMeasConfigCommon-r18: ref
	{
		if ie.Sl_PosMeasConfigCommon_r18 != nil {
			if err := ie.Sl_PosMeasConfigCommon_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. sl-PosOffsetDFN-r18: integer
	{
		if ie.Sl_PosOffsetDFN_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PosOffsetDFN_r18, sLPosConfigCommonNRR18SlPosOffsetDFNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-PosSSB-PriorityNR-r18: integer
	{
		if ie.Sl_PosSSB_PriorityNR_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PosSSB_PriorityNR_r18, sLPosConfigCommonNRR18SlPosSSBPriorityNRR18Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-PosFreqInfoListExt-v1880", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PosFreqInfoListExt_v1880 != nil}); err != nil {
				return err
			}

			if ie.Sl_PosFreqInfoListExt_v1880 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLPosConfigCommonNRR18ExtSlPosFreqInfoListExtV1880Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_PosFreqInfoListExt_v1880))); err != nil {
					return err
				}
				for i := range ie.Sl_PosFreqInfoListExt_v1880 {
					if err := ie.Sl_PosFreqInfoListExt_v1880[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *SL_PosConfigCommonNR_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPosConfigCommonNRR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PosFreqInfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLPosConfigCommonNRR18SlPosFreqInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PosFreqInfoList_r18 = make([]SL_FreqConfigCommon_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PosFreqInfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PosUE-SelectedConfig-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_PosUE_SelectedConfig_r18 = new(SL_UE_SelectedConfig_r16)
			if err := ie.Sl_PosUE_SelectedConfig_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-PosNR-AnchorCarrierFreqList-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_PosNR_AnchorCarrierFreqList_r18 = new(SL_NR_AnchorCarrierFreqList_r16)
			if err := ie.Sl_PosNR_AnchorCarrierFreqList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-PosMeasConfigCommon-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_PosMeasConfigCommon_r18 = new(SL_MeasConfigCommon_r16)
			if err := ie.Sl_PosMeasConfigCommon_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. sl-PosOffsetDFN-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLPosConfigCommonNRR18SlPosOffsetDFNR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PosOffsetDFN_r18 = &v
		}
	}

	// 8. sl-PosSSB-PriorityNR-r18: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(sLPosConfigCommonNRR18SlPosSSBPriorityNRR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PosSSB_PriorityNR_r18 = &v
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
				{Name: "sl-PosFreqInfoListExt-v1880", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sLPosConfigCommonNRR18ExtSlPosFreqInfoListExtV1880Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PosFreqInfoListExt_v1880 = make([]SL_FreqConfigCommonExt_V16k0, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PosFreqInfoListExt_v1880[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
