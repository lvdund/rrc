// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-LogicalChannelConfig-r16 (line 27408).

var sLLogicalChannelConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Priority-r16"},
		{Name: "sl-PrioritisedBitRate-r16"},
		{Name: "sl-BucketSizeDuration-r16"},
		{Name: "sl-ConfiguredGrantType1Allowed-r16", Optional: true},
		{Name: "sl-HARQ-FeedbackEnabled-r16", Optional: true},
		{Name: "sl-AllowedCG-List-r16", Optional: true},
		{Name: "sl-AllowedSCS-List-r16", Optional: true},
		{Name: "sl-MaxPUSCH-Duration-r16", Optional: true},
		{Name: "sl-LogicalChannelGroup-r16", Optional: true},
		{Name: "sl-SchedulingRequestId-r16", Optional: true},
		{Name: "sl-LogicalChannelSR-DelayTimerApplied-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLLogicalChannelConfigR16SlPriorityR16Constraints = per.Constrained(1, 8)

const (
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps0     = 0
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps8     = 1
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps16    = 2
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps32    = 3
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps64    = 4
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps128   = 5
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps256   = 6
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps512   = 7
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps1024  = 8
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps2048  = 9
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps4096  = 10
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps8192  = 11
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps16384 = 12
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps32768 = 13
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps65536 = 14
	SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_Infinity  = 15
)

var sLLogicalChannelConfigR16SlPrioritisedBitRateR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps0, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps8, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps16, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps32, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps64, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps128, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps256, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps512, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps1024, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps2048, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps4096, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps8192, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps16384, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps32768, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_KBps65536, SL_LogicalChannelConfig_r16_Sl_PrioritisedBitRate_r16_Infinity},
}

const (
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms5    = 0
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms10   = 1
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms20   = 2
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms50   = 3
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms100  = 4
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms150  = 5
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms300  = 6
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms500  = 7
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms1000 = 8
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare7 = 9
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare6 = 10
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare5 = 11
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare4 = 12
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare3 = 13
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare2 = 14
	SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare1 = 15
)

var sLLogicalChannelConfigR16SlBucketSizeDurationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms5, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms10, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms20, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms50, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms100, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms150, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms300, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms500, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Ms1000, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare7, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare6, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare5, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare4, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare3, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare2, SL_LogicalChannelConfig_r16_Sl_BucketSizeDuration_r16_Spare1},
}

const (
	SL_LogicalChannelConfig_r16_Sl_ConfiguredGrantType1Allowed_r16_True = 0
)

var sLLogicalChannelConfigR16SlConfiguredGrantType1AllowedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LogicalChannelConfig_r16_Sl_ConfiguredGrantType1Allowed_r16_True},
}

const (
	SL_LogicalChannelConfig_r16_Sl_HARQ_FeedbackEnabled_r16_Enabled  = 0
	SL_LogicalChannelConfig_r16_Sl_HARQ_FeedbackEnabled_r16_Disabled = 1
)

var sLLogicalChannelConfigR16SlHARQFeedbackEnabledR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LogicalChannelConfig_r16_Sl_HARQ_FeedbackEnabled_r16_Enabled, SL_LogicalChannelConfig_r16_Sl_HARQ_FeedbackEnabled_r16_Disabled},
}

var sLLogicalChannelConfigR16SlAllowedCGListR16Constraints = per.SizeRange(0, common.MaxNrofCG_SL_1_r16)

var sLLogicalChannelConfigR16SlAllowedSCSListR16Constraints = per.SizeRange(1, common.MaxSCSs)

const (
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p02   = 0
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p04   = 1
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p0625 = 2
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p125  = 3
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p25   = 4
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p5    = 5
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Spare2   = 6
	SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Spare1   = 7
)

var sLLogicalChannelConfigR16SlMaxPUSCHDurationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p02, SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p04, SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p0625, SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p125, SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p25, SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Ms0p5, SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Spare2, SL_LogicalChannelConfig_r16_Sl_MaxPUSCH_Duration_r16_Spare1},
}

var sLLogicalChannelConfigR16SlLogicalChannelGroupR16Constraints = per.Constrained(0, common.MaxLCG_ID)

var sLLogicalChannelConfigR16SlChannelAccessPriorityR18Constraints = per.Constrained(1, 4)

var sLLogicalChannelConfigR16ExtSlAllowedCarriersR18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

type SL_LogicalChannelConfig_r16 struct {
	Sl_Priority_r16                           int64
	Sl_PrioritisedBitRate_r16                 int64
	Sl_BucketSizeDuration_r16                 int64
	Sl_ConfiguredGrantType1Allowed_r16        *int64
	Sl_HARQ_FeedbackEnabled_r16               *int64
	Sl_AllowedCG_List_r16                     []SL_ConfigIndexCG_r16
	Sl_AllowedSCS_List_r16                    []SubcarrierSpacing
	Sl_MaxPUSCH_Duration_r16                  *int64
	Sl_LogicalChannelGroup_r16                *int64
	Sl_SchedulingRequestId_r16                *SchedulingRequestId
	Sl_LogicalChannelSR_DelayTimerApplied_r16 *bool
	Sl_ChannelAccessPriority_r18              *int64
	Sl_AllowedCarriers_r18                    []int64
}

func (ie *SL_LogicalChannelConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLLogicalChannelConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_ChannelAccessPriority_r18 != nil || ie.Sl_AllowedCarriers_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfiguredGrantType1Allowed_r16 != nil, ie.Sl_HARQ_FeedbackEnabled_r16 != nil, ie.Sl_AllowedCG_List_r16 != nil, ie.Sl_AllowedSCS_List_r16 != nil, ie.Sl_MaxPUSCH_Duration_r16 != nil, ie.Sl_LogicalChannelGroup_r16 != nil, ie.Sl_SchedulingRequestId_r16 != nil, ie.Sl_LogicalChannelSR_DelayTimerApplied_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-Priority-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_Priority_r16, sLLogicalChannelConfigR16SlPriorityR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-PrioritisedBitRate-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_PrioritisedBitRate_r16, sLLogicalChannelConfigR16SlPrioritisedBitRateR16Constraints); err != nil {
			return err
		}
	}

	// 5. sl-BucketSizeDuration-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_BucketSizeDuration_r16, sLLogicalChannelConfigR16SlBucketSizeDurationR16Constraints); err != nil {
			return err
		}
	}

	// 6. sl-ConfiguredGrantType1Allowed-r16: enumerated
	{
		if ie.Sl_ConfiguredGrantType1Allowed_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_ConfiguredGrantType1Allowed_r16, sLLogicalChannelConfigR16SlConfiguredGrantType1AllowedR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-HARQ-FeedbackEnabled-r16: enumerated
	{
		if ie.Sl_HARQ_FeedbackEnabled_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_HARQ_FeedbackEnabled_r16, sLLogicalChannelConfigR16SlHARQFeedbackEnabledR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-AllowedCG-List-r16: sequence-of
	{
		if ie.Sl_AllowedCG_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLLogicalChannelConfigR16SlAllowedCGListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_AllowedCG_List_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_AllowedCG_List_r16 {
				if err := ie.Sl_AllowedCG_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-AllowedSCS-List-r16: sequence-of
	{
		if ie.Sl_AllowedSCS_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLLogicalChannelConfigR16SlAllowedSCSListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_AllowedSCS_List_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_AllowedSCS_List_r16 {
				if err := ie.Sl_AllowedSCS_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-MaxPUSCH-Duration-r16: enumerated
	{
		if ie.Sl_MaxPUSCH_Duration_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_MaxPUSCH_Duration_r16, sLLogicalChannelConfigR16SlMaxPUSCHDurationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sl-LogicalChannelGroup-r16: integer
	{
		if ie.Sl_LogicalChannelGroup_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_LogicalChannelGroup_r16, sLLogicalChannelConfigR16SlLogicalChannelGroupR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-SchedulingRequestId-r16: ref
	{
		if ie.Sl_SchedulingRequestId_r16 != nil {
			if err := ie.Sl_SchedulingRequestId_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 13. sl-LogicalChannelSR-DelayTimerApplied-r16: boolean
	{
		if ie.Sl_LogicalChannelSR_DelayTimerApplied_r16 != nil {
			if err := e.EncodeBoolean(*ie.Sl_LogicalChannelSR_DelayTimerApplied_r16); err != nil {
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
					{Name: "sl-ChannelAccessPriority-r18", Optional: true},
					{Name: "sl-AllowedCarriers-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_ChannelAccessPriority_r18 != nil, ie.Sl_AllowedCarriers_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_ChannelAccessPriority_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_ChannelAccessPriority_r18, sLLogicalChannelConfigR16SlChannelAccessPriorityR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_AllowedCarriers_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLLogicalChannelConfigR16ExtSlAllowedCarriersR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_AllowedCarriers_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_AllowedCarriers_r18 {
					if err := ex.EncodeInteger(ie.Sl_AllowedCarriers_r18[i], per.Constrained(1, common.MaxNrofFreqSL_r16)); err != nil {
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

func (ie *SL_LogicalChannelConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLLogicalChannelConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-Priority-r16: integer
	{
		v0, err := d.DecodeInteger(sLLogicalChannelConfigR16SlPriorityR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Priority_r16 = v0
	}

	// 4. sl-PrioritisedBitRate-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sLLogicalChannelConfigR16SlPrioritisedBitRateR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_PrioritisedBitRate_r16 = v1
	}

	// 5. sl-BucketSizeDuration-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(sLLogicalChannelConfigR16SlBucketSizeDurationR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_BucketSizeDuration_r16 = v2
	}

	// 6. sl-ConfiguredGrantType1Allowed-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLLogicalChannelConfigR16SlConfiguredGrantType1AllowedR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ConfiguredGrantType1Allowed_r16 = &idx
		}
	}

	// 7. sl-HARQ-FeedbackEnabled-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sLLogicalChannelConfigR16SlHARQFeedbackEnabledR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_HARQ_FeedbackEnabled_r16 = &idx
		}
	}

	// 8. sl-AllowedCG-List-r16: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(sLLogicalChannelConfigR16SlAllowedCGListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_AllowedCG_List_r16 = make([]SL_ConfigIndexCG_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_AllowedCG_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-AllowedSCS-List-r16: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(sLLogicalChannelConfigR16SlAllowedSCSListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_AllowedSCS_List_r16 = make([]SubcarrierSpacing, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_AllowedSCS_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-MaxPUSCH-Duration-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sLLogicalChannelConfigR16SlMaxPUSCHDurationR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MaxPUSCH_Duration_r16 = &idx
		}
	}

	// 11. sl-LogicalChannelGroup-r16: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(sLLogicalChannelConfigR16SlLogicalChannelGroupR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LogicalChannelGroup_r16 = &v
		}
	}

	// 12. sl-SchedulingRequestId-r16: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Sl_SchedulingRequestId_r16 = new(SchedulingRequestId)
			if err := ie.Sl_SchedulingRequestId_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 13. sl-LogicalChannelSR-DelayTimerApplied-r16: boolean
	{
		if seq.IsComponentPresent(10) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Sl_LogicalChannelSR_DelayTimerApplied_r16 = &v
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
				{Name: "sl-ChannelAccessPriority-r18", Optional: true},
				{Name: "sl-AllowedCarriers-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sLLogicalChannelConfigR16SlChannelAccessPriorityR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ChannelAccessPriority_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(sLLogicalChannelConfigR16ExtSlAllowedCarriersR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_AllowedCarriers_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofFreqSL_r16))
				if err != nil {
					return err
				}
				ie.Sl_AllowedCarriers_r18[i] = v
			}
		}
	}

	return nil
}
