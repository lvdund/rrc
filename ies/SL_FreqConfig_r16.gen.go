// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-FreqConfig-r16 (line 27253).

var sLFreqConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Freq-Id-r16"},
		{Name: "sl-SCS-SpecificCarrierList-r16"},
		{Name: "sl-AbsoluteFrequencyPointA-r16", Optional: true},
		{Name: "sl-AbsoluteFrequencySSB-r16", Optional: true},
		{Name: "frequencyShift7p5khzSL-r16", Optional: true},
		{Name: "valueN-r16"},
		{Name: "sl-BWP-ToReleaseList-r16", Optional: true},
		{Name: "sl-BWP-ToAddModList-r16", Optional: true},
		{Name: "sl-SyncConfigList-r16", Optional: true},
		{Name: "sl-SyncPriority-r16", Optional: true},
	},
}

var sLFreqConfigR16SlSCSSpecificCarrierListR16Constraints = per.SizeRange(1, common.MaxSCSs)

const (
	SL_FreqConfig_r16_FrequencyShift7p5khzSL_r16_True = 0
)

var sLFreqConfigR16FrequencyShift7p5khzSLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfig_r16_FrequencyShift7p5khzSL_r16_True},
}

var sLFreqConfigR16ValueNR16Constraints = per.Constrained(-1, 1)

var sLFreqConfigR16SlBWPToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofSL_BWPs_r16)

var sLFreqConfigR16SlBWPToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofSL_BWPs_r16)

const (
	SL_FreqConfig_r16_Sl_SyncPriority_r16_Gnss   = 0
	SL_FreqConfig_r16_Sl_SyncPriority_r16_GnbEnb = 1
)

var sLFreqConfigR16SlSyncPriorityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_FreqConfig_r16_Sl_SyncPriority_r16_Gnss, SL_FreqConfig_r16_Sl_SyncPriority_r16_GnbEnb},
}

type SL_FreqConfig_r16 struct {
	Sl_Freq_Id_r16                 SL_Freq_Id_r16
	Sl_SCS_SpecificCarrierList_r16 []SCS_SpecificCarrier
	Sl_AbsoluteFrequencyPointA_r16 *ARFCN_ValueNR
	Sl_AbsoluteFrequencySSB_r16    *ARFCN_ValueNR
	FrequencyShift7p5khzSL_r16     *int64
	ValueN_r16                     int64
	Sl_BWP_ToReleaseList_r16       []BWP_Id
	Sl_BWP_ToAddModList_r16        []SL_BWP_Config_r16
	Sl_SyncConfigList_r16          *SL_SyncConfigList_r16
	Sl_SyncPriority_r16            *int64
}

func (ie *SL_FreqConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLFreqConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_AbsoluteFrequencyPointA_r16 != nil, ie.Sl_AbsoluteFrequencySSB_r16 != nil, ie.FrequencyShift7p5khzSL_r16 != nil, ie.Sl_BWP_ToReleaseList_r16 != nil, ie.Sl_BWP_ToAddModList_r16 != nil, ie.Sl_SyncConfigList_r16 != nil, ie.Sl_SyncPriority_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-Freq-Id-r16: ref
	{
		if err := ie.Sl_Freq_Id_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-SCS-SpecificCarrierList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLFreqConfigR16SlSCSSpecificCarrierListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_SCS_SpecificCarrierList_r16))); err != nil {
			return err
		}
		for i := range ie.Sl_SCS_SpecificCarrierList_r16 {
			if err := ie.Sl_SCS_SpecificCarrierList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-AbsoluteFrequencyPointA-r16: ref
	{
		if ie.Sl_AbsoluteFrequencyPointA_r16 != nil {
			if err := ie.Sl_AbsoluteFrequencyPointA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-AbsoluteFrequencySSB-r16: ref
	{
		if ie.Sl_AbsoluteFrequencySSB_r16 != nil {
			if err := ie.Sl_AbsoluteFrequencySSB_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. frequencyShift7p5khzSL-r16: enumerated
	{
		if ie.FrequencyShift7p5khzSL_r16 != nil {
			if err := e.EncodeEnumerated(*ie.FrequencyShift7p5khzSL_r16, sLFreqConfigR16FrequencyShift7p5khzSLR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. valueN-r16: integer
	{
		if err := e.EncodeInteger(ie.ValueN_r16, sLFreqConfigR16ValueNR16Constraints); err != nil {
			return err
		}
	}

	// 8. sl-BWP-ToReleaseList-r16: sequence-of
	{
		if ie.Sl_BWP_ToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLFreqConfigR16SlBWPToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_BWP_ToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_BWP_ToReleaseList_r16 {
				if err := ie.Sl_BWP_ToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-BWP-ToAddModList-r16: sequence-of
	{
		if ie.Sl_BWP_ToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLFreqConfigR16SlBWPToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_BWP_ToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_BWP_ToAddModList_r16 {
				if err := ie.Sl_BWP_ToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-SyncConfigList-r16: ref
	{
		if ie.Sl_SyncConfigList_r16 != nil {
			if err := ie.Sl_SyncConfigList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. sl-SyncPriority-r16: enumerated
	{
		if ie.Sl_SyncPriority_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SyncPriority_r16, sLFreqConfigR16SlSyncPriorityR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_FreqConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLFreqConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-Freq-Id-r16: ref
	{
		if err := ie.Sl_Freq_Id_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-SCS-SpecificCarrierList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLFreqConfigR16SlSCSSpecificCarrierListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_SCS_SpecificCarrierList_r16 = make([]SCS_SpecificCarrier, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_SCS_SpecificCarrierList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-AbsoluteFrequencyPointA-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_AbsoluteFrequencyPointA_r16 = new(ARFCN_ValueNR)
			if err := ie.Sl_AbsoluteFrequencyPointA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-AbsoluteFrequencySSB-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_AbsoluteFrequencySSB_r16 = new(ARFCN_ValueNR)
			if err := ie.Sl_AbsoluteFrequencySSB_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. frequencyShift7p5khzSL-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sLFreqConfigR16FrequencyShift7p5khzSLR16Constraints)
			if err != nil {
				return err
			}
			ie.FrequencyShift7p5khzSL_r16 = &idx
		}
	}

	// 7. valueN-r16: integer
	{
		v5, err := d.DecodeInteger(sLFreqConfigR16ValueNR16Constraints)
		if err != nil {
			return err
		}
		ie.ValueN_r16 = v5
	}

	// 8. sl-BWP-ToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(sLFreqConfigR16SlBWPToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_BWP_ToReleaseList_r16 = make([]BWP_Id, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_BWP_ToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-BWP-ToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(sLFreqConfigR16SlBWPToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_BWP_ToAddModList_r16 = make([]SL_BWP_Config_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_BWP_ToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-SyncConfigList-r16: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Sl_SyncConfigList_r16 = new(SL_SyncConfigList_r16)
			if err := ie.Sl_SyncConfigList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. sl-SyncPriority-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sLFreqConfigR16SlSyncPriorityR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SyncPriority_r16 = &idx
		}
	}

	return nil
}
