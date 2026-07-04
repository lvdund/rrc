// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ConfigCommonNR-r16 (line 4316).

var sLConfigCommonNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-FreqInfoList-r16", Optional: true},
		{Name: "sl-UE-SelectedConfig-r16", Optional: true},
		{Name: "sl-NR-AnchorCarrierFreqList-r16", Optional: true},
		{Name: "sl-EUTRA-AnchorCarrierFreqList-r16", Optional: true},
		{Name: "sl-RadioBearerConfigList-r16", Optional: true},
		{Name: "sl-RLC-BearerConfigList-r16", Optional: true},
		{Name: "sl-MeasConfigCommon-r16", Optional: true},
		{Name: "sl-CSI-Acquisition-r16", Optional: true},
		{Name: "sl-OffsetDFN-r16", Optional: true},
		{Name: "t400-r16", Optional: true},
		{Name: "sl-MaxNumConsecutiveDTX-r16", Optional: true},
		{Name: "sl-SSB-PriorityNR-r16", Optional: true},
	},
}

var sLConfigCommonNRR16SlFreqInfoListR16Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLConfigCommonNRR16SlRadioBearerConfigListR16Constraints = per.SizeRange(1, common.MaxNrofSLRB_r16)

var sLConfigCommonNRR16SlRLCBearerConfigListR16Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

const (
	SL_ConfigCommonNR_r16_Sl_CSI_Acquisition_r16_Enabled = 0
)

var sLConfigCommonNRR16SlCSIAcquisitionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ConfigCommonNR_r16_Sl_CSI_Acquisition_r16_Enabled},
}

var sLConfigCommonNRR16SlOffsetDFNR16Constraints = per.Constrained(1, 1000)

const (
	SL_ConfigCommonNR_r16_T400_r16_Ms100  = 0
	SL_ConfigCommonNR_r16_T400_r16_Ms200  = 1
	SL_ConfigCommonNR_r16_T400_r16_Ms300  = 2
	SL_ConfigCommonNR_r16_T400_r16_Ms400  = 3
	SL_ConfigCommonNR_r16_T400_r16_Ms600  = 4
	SL_ConfigCommonNR_r16_T400_r16_Ms1000 = 5
	SL_ConfigCommonNR_r16_T400_r16_Ms1500 = 6
	SL_ConfigCommonNR_r16_T400_r16_Ms2000 = 7
)

var sLConfigCommonNRR16T400R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ConfigCommonNR_r16_T400_r16_Ms100, SL_ConfigCommonNR_r16_T400_r16_Ms200, SL_ConfigCommonNR_r16_T400_r16_Ms300, SL_ConfigCommonNR_r16_T400_r16_Ms400, SL_ConfigCommonNR_r16_T400_r16_Ms600, SL_ConfigCommonNR_r16_T400_r16_Ms1000, SL_ConfigCommonNR_r16_T400_r16_Ms1500, SL_ConfigCommonNR_r16_T400_r16_Ms2000},
}

const (
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N1  = 0
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N2  = 1
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N3  = 2
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N4  = 3
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N6  = 4
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N8  = 5
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N16 = 6
	SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N32 = 7
)

var sLConfigCommonNRR16SlMaxNumConsecutiveDTXR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N1, SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N2, SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N3, SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N4, SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N6, SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N8, SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N16, SL_ConfigCommonNR_r16_Sl_MaxNumConsecutiveDTX_r16_N32},
}

var sLConfigCommonNRR16SlSSBPriorityNRR16Constraints = per.Constrained(1, 8)

type SL_ConfigCommonNR_r16 struct {
	Sl_FreqInfoList_r16                []SL_FreqConfigCommon_r16
	Sl_UE_SelectedConfig_r16           *SL_UE_SelectedConfig_r16
	Sl_NR_AnchorCarrierFreqList_r16    *SL_NR_AnchorCarrierFreqList_r16
	Sl_EUTRA_AnchorCarrierFreqList_r16 *SL_EUTRA_AnchorCarrierFreqList_r16
	Sl_RadioBearerConfigList_r16       []SL_RadioBearerConfig_r16
	Sl_RLC_BearerConfigList_r16        []SL_RLC_BearerConfig_r16
	Sl_MeasConfigCommon_r16            *SL_MeasConfigCommon_r16
	Sl_CSI_Acquisition_r16             *int64
	Sl_OffsetDFN_r16                   *int64
	T400_r16                           *int64
	Sl_MaxNumConsecutiveDTX_r16        *int64
	Sl_SSB_PriorityNR_r16              *int64
}

func (ie *SL_ConfigCommonNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfigCommonNRR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_FreqInfoList_r16 != nil, ie.Sl_UE_SelectedConfig_r16 != nil, ie.Sl_NR_AnchorCarrierFreqList_r16 != nil, ie.Sl_EUTRA_AnchorCarrierFreqList_r16 != nil, ie.Sl_RadioBearerConfigList_r16 != nil, ie.Sl_RLC_BearerConfigList_r16 != nil, ie.Sl_MeasConfigCommon_r16 != nil, ie.Sl_CSI_Acquisition_r16 != nil, ie.Sl_OffsetDFN_r16 != nil, ie.T400_r16 != nil, ie.Sl_MaxNumConsecutiveDTX_r16 != nil, ie.Sl_SSB_PriorityNR_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-FreqInfoList-r16: sequence-of
	{
		if ie.Sl_FreqInfoList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigCommonNRR16SlFreqInfoListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_FreqInfoList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_FreqInfoList_r16 {
				if err := ie.Sl_FreqInfoList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-UE-SelectedConfig-r16: ref
	{
		if ie.Sl_UE_SelectedConfig_r16 != nil {
			if err := ie.Sl_UE_SelectedConfig_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-NR-AnchorCarrierFreqList-r16: ref
	{
		if ie.Sl_NR_AnchorCarrierFreqList_r16 != nil {
			if err := ie.Sl_NR_AnchorCarrierFreqList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-EUTRA-AnchorCarrierFreqList-r16: ref
	{
		if ie.Sl_EUTRA_AnchorCarrierFreqList_r16 != nil {
			if err := ie.Sl_EUTRA_AnchorCarrierFreqList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-RadioBearerConfigList-r16: sequence-of
	{
		if ie.Sl_RadioBearerConfigList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigCommonNRR16SlRadioBearerConfigListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RadioBearerConfigList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RadioBearerConfigList_r16 {
				if err := ie.Sl_RadioBearerConfigList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-RLC-BearerConfigList-r16: sequence-of
	{
		if ie.Sl_RLC_BearerConfigList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigCommonNRR16SlRLCBearerConfigListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_BearerConfigList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RLC_BearerConfigList_r16 {
				if err := ie.Sl_RLC_BearerConfigList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-MeasConfigCommon-r16: ref
	{
		if ie.Sl_MeasConfigCommon_r16 != nil {
			if err := ie.Sl_MeasConfigCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. sl-CSI-Acquisition-r16: enumerated
	{
		if ie.Sl_CSI_Acquisition_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_CSI_Acquisition_r16, sLConfigCommonNRR16SlCSIAcquisitionR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-OffsetDFN-r16: integer
	{
		if ie.Sl_OffsetDFN_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_OffsetDFN_r16, sLConfigCommonNRR16SlOffsetDFNR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. t400-r16: enumerated
	{
		if ie.T400_r16 != nil {
			if err := e.EncodeEnumerated(*ie.T400_r16, sLConfigCommonNRR16T400R16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-MaxNumConsecutiveDTX-r16: enumerated
	{
		if ie.Sl_MaxNumConsecutiveDTX_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_MaxNumConsecutiveDTX_r16, sLConfigCommonNRR16SlMaxNumConsecutiveDTXR16Constraints); err != nil {
				return err
			}
		}
	}

	// 13. sl-SSB-PriorityNR-r16: integer
	{
		if ie.Sl_SSB_PriorityNR_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_SSB_PriorityNR_r16, sLConfigCommonNRR16SlSSBPriorityNRR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_ConfigCommonNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfigCommonNRR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-FreqInfoList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLConfigCommonNRR16SlFreqInfoListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_FreqInfoList_r16 = make([]SL_FreqConfigCommon_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_FreqInfoList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-UE-SelectedConfig-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_UE_SelectedConfig_r16 = new(SL_UE_SelectedConfig_r16)
			if err := ie.Sl_UE_SelectedConfig_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-NR-AnchorCarrierFreqList-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_NR_AnchorCarrierFreqList_r16 = new(SL_NR_AnchorCarrierFreqList_r16)
			if err := ie.Sl_NR_AnchorCarrierFreqList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-EUTRA-AnchorCarrierFreqList-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_EUTRA_AnchorCarrierFreqList_r16 = new(SL_EUTRA_AnchorCarrierFreqList_r16)
			if err := ie.Sl_EUTRA_AnchorCarrierFreqList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-RadioBearerConfigList-r16: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sLConfigCommonNRR16SlRadioBearerConfigListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RadioBearerConfigList_r16 = make([]SL_RadioBearerConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RadioBearerConfigList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-RLC-BearerConfigList-r16: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(sLConfigCommonNRR16SlRLCBearerConfigListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_BearerConfigList_r16 = make([]SL_RLC_BearerConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_BearerConfigList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-MeasConfigCommon-r16: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Sl_MeasConfigCommon_r16 = new(SL_MeasConfigCommon_r16)
			if err := ie.Sl_MeasConfigCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. sl-CSI-Acquisition-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sLConfigCommonNRR16SlCSIAcquisitionR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CSI_Acquisition_r16 = &idx
		}
	}

	// 10. sl-OffsetDFN-r16: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(sLConfigCommonNRR16SlOffsetDFNR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_OffsetDFN_r16 = &v
		}
	}

	// 11. t400-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sLConfigCommonNRR16T400R16Constraints)
			if err != nil {
				return err
			}
			ie.T400_r16 = &idx
		}
	}

	// 12. sl-MaxNumConsecutiveDTX-r16: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sLConfigCommonNRR16SlMaxNumConsecutiveDTXR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MaxNumConsecutiveDTX_r16 = &idx
		}
	}

	// 13. sl-SSB-PriorityNR-r16: integer
	{
		if seq.IsComponentPresent(11) {
			v, err := d.DecodeInteger(sLConfigCommonNRR16SlSSBPriorityNRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SSB_PriorityNR_r16 = &v
		}
	}

	return nil
}
