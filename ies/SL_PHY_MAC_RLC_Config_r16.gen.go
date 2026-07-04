// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PHY-MAC-RLC-Config-r16 (line 26977).

var sLPHYMACRLCConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ScheduledConfig-r16", Optional: true},
		{Name: "sl-UE-SelectedConfig-r16", Optional: true},
		{Name: "sl-FreqInfoToReleaseList-r16", Optional: true},
		{Name: "sl-FreqInfoToAddModList-r16", Optional: true},
		{Name: "sl-RLC-BearerToReleaseList-r16", Optional: true},
		{Name: "sl-RLC-BearerToAddModList-r16", Optional: true},
		{Name: "sl-MaxNumConsecutiveDTX-r16", Optional: true},
		{Name: "sl-CSI-Acquisition-r16", Optional: true},
		{Name: "sl-CSI-SchedulingRequestId-r16", Optional: true},
		{Name: "sl-SSB-PriorityNR-r16", Optional: true},
		{Name: "networkControlledSyncTx-r16", Optional: true},
	},
}

var sL_PHY_MAC_RLC_Config_r16SlScheduledConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_PHY_MAC_RLC_Config_r16_Sl_ScheduledConfig_r16_Release = 0
	SL_PHY_MAC_RLC_Config_r16_Sl_ScheduledConfig_r16_Setup   = 1
)

var sL_PHY_MAC_RLC_Config_r16SlUESelectedConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_PHY_MAC_RLC_Config_r16_Sl_UE_SelectedConfig_r16_Release = 0
	SL_PHY_MAC_RLC_Config_r16_Sl_UE_SelectedConfig_r16_Setup   = 1
)

var sLPHYMACRLCConfigR16SlFreqInfoToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLPHYMACRLCConfigR16SlFreqInfoToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLPHYMACRLCConfigR16SlRLCBearerToReleaseListR16Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

var sLPHYMACRLCConfigR16SlRLCBearerToAddModListR16Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

const (
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N1  = 0
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N2  = 1
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N3  = 2
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N4  = 3
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N6  = 4
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N8  = 5
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N16 = 6
	SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N32 = 7
)

var sLPHYMACRLCConfigR16SlMaxNumConsecutiveDTXR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N1, SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N2, SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N3, SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N4, SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N6, SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N8, SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N16, SL_PHY_MAC_RLC_Config_r16_Sl_MaxNumConsecutiveDTX_r16_N32},
}

const (
	SL_PHY_MAC_RLC_Config_r16_Sl_CSI_Acquisition_r16_Enabled = 0
)

var sLPHYMACRLCConfigR16SlCSIAcquisitionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PHY_MAC_RLC_Config_r16_Sl_CSI_Acquisition_r16_Enabled},
}

var sL_PHY_MAC_RLC_Config_r16SlCSISchedulingRequestIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_PHY_MAC_RLC_Config_r16_Sl_CSI_SchedulingRequestId_r16_Release = 0
	SL_PHY_MAC_RLC_Config_r16_Sl_CSI_SchedulingRequestId_r16_Setup   = 1
)

var sLPHYMACRLCConfigR16SlSSBPriorityNRR16Constraints = per.Constrained(1, 8)

const (
	SL_PHY_MAC_RLC_Config_r16_NetworkControlledSyncTx_r16_On  = 0
	SL_PHY_MAC_RLC_Config_r16_NetworkControlledSyncTx_r16_Off = 1
)

var sLPHYMACRLCConfigR16NetworkControlledSyncTxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PHY_MAC_RLC_Config_r16_NetworkControlledSyncTx_r16_On, SL_PHY_MAC_RLC_Config_r16_NetworkControlledSyncTx_r16_Off},
}

type SL_PHY_MAC_RLC_Config_r16 struct {
	Sl_ScheduledConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_ScheduledConfig_r16
	}
	Sl_UE_SelectedConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_UE_SelectedConfig_r16
	}
	Sl_FreqInfoToReleaseList_r16   []SL_Freq_Id_r16
	Sl_FreqInfoToAddModList_r16    []SL_FreqConfig_r16
	Sl_RLC_BearerToReleaseList_r16 []SL_RLC_BearerConfigIndex_r16
	Sl_RLC_BearerToAddModList_r16  []SL_RLC_BearerConfig_r16
	Sl_MaxNumConsecutiveDTX_r16    *int64
	Sl_CSI_Acquisition_r16         *int64
	Sl_CSI_SchedulingRequestId_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SchedulingRequestId
	}
	Sl_SSB_PriorityNR_r16       *int64
	NetworkControlledSyncTx_r16 *int64
}

func (ie *SL_PHY_MAC_RLC_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPHYMACRLCConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ScheduledConfig_r16 != nil, ie.Sl_UE_SelectedConfig_r16 != nil, ie.Sl_FreqInfoToReleaseList_r16 != nil, ie.Sl_FreqInfoToAddModList_r16 != nil, ie.Sl_RLC_BearerToReleaseList_r16 != nil, ie.Sl_RLC_BearerToAddModList_r16 != nil, ie.Sl_MaxNumConsecutiveDTX_r16 != nil, ie.Sl_CSI_Acquisition_r16 != nil, ie.Sl_CSI_SchedulingRequestId_r16 != nil, ie.Sl_SSB_PriorityNR_r16 != nil, ie.NetworkControlledSyncTx_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-ScheduledConfig-r16: choice
	{
		if ie.Sl_ScheduledConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_PHY_MAC_RLC_Config_r16SlScheduledConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_ScheduledConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_ScheduledConfig_r16).Choice {
			case SL_PHY_MAC_RLC_Config_r16_Sl_ScheduledConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_r16_Sl_ScheduledConfig_r16_Setup:
				if err := (*ie.Sl_ScheduledConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_ScheduledConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. sl-UE-SelectedConfig-r16: choice
	{
		if ie.Sl_UE_SelectedConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_PHY_MAC_RLC_Config_r16SlUESelectedConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_UE_SelectedConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_UE_SelectedConfig_r16).Choice {
			case SL_PHY_MAC_RLC_Config_r16_Sl_UE_SelectedConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_r16_Sl_UE_SelectedConfig_r16_Setup:
				if err := (*ie.Sl_UE_SelectedConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_UE_SelectedConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. sl-FreqInfoToReleaseList-r16: sequence-of
	{
		if ie.Sl_FreqInfoToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPHYMACRLCConfigR16SlFreqInfoToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_FreqInfoToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_FreqInfoToReleaseList_r16 {
				if err := ie.Sl_FreqInfoToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-FreqInfoToAddModList-r16: sequence-of
	{
		if ie.Sl_FreqInfoToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPHYMACRLCConfigR16SlFreqInfoToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_FreqInfoToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_FreqInfoToAddModList_r16 {
				if err := ie.Sl_FreqInfoToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-RLC-BearerToReleaseList-r16: sequence-of
	{
		if ie.Sl_RLC_BearerToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPHYMACRLCConfigR16SlRLCBearerToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_BearerToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RLC_BearerToReleaseList_r16 {
				if err := ie.Sl_RLC_BearerToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-RLC-BearerToAddModList-r16: sequence-of
	{
		if ie.Sl_RLC_BearerToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPHYMACRLCConfigR16SlRLCBearerToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_BearerToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RLC_BearerToAddModList_r16 {
				if err := ie.Sl_RLC_BearerToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-MaxNumConsecutiveDTX-r16: enumerated
	{
		if ie.Sl_MaxNumConsecutiveDTX_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_MaxNumConsecutiveDTX_r16, sLPHYMACRLCConfigR16SlMaxNumConsecutiveDTXR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-CSI-Acquisition-r16: enumerated
	{
		if ie.Sl_CSI_Acquisition_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_CSI_Acquisition_r16, sLPHYMACRLCConfigR16SlCSIAcquisitionR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-CSI-SchedulingRequestId-r16: choice
	{
		if ie.Sl_CSI_SchedulingRequestId_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_PHY_MAC_RLC_Config_r16SlCSISchedulingRequestIdR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_CSI_SchedulingRequestId_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_CSI_SchedulingRequestId_r16).Choice {
			case SL_PHY_MAC_RLC_Config_r16_Sl_CSI_SchedulingRequestId_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_r16_Sl_CSI_SchedulingRequestId_r16_Setup:
				if err := (*ie.Sl_CSI_SchedulingRequestId_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_CSI_SchedulingRequestId_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. sl-SSB-PriorityNR-r16: integer
	{
		if ie.Sl_SSB_PriorityNR_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_SSB_PriorityNR_r16, sLPHYMACRLCConfigR16SlSSBPriorityNRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. networkControlledSyncTx-r16: enumerated
	{
		if ie.NetworkControlledSyncTx_r16 != nil {
			if err := e.EncodeEnumerated(*ie.NetworkControlledSyncTx_r16, sLPHYMACRLCConfigR16NetworkControlledSyncTxR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PHY_MAC_RLC_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPHYMACRLCConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ScheduledConfig-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_ScheduledConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_ScheduledConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sL_PHY_MAC_RLC_Config_r16SlScheduledConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_ScheduledConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_PHY_MAC_RLC_Config_r16_Sl_ScheduledConfig_r16_Release:
				(*ie.Sl_ScheduledConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_r16_Sl_ScheduledConfig_r16_Setup:
				(*ie.Sl_ScheduledConfig_r16).Setup = new(SL_ScheduledConfig_r16)
				if err := (*ie.Sl_ScheduledConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-UE-SelectedConfig-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_UE_SelectedConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_UE_SelectedConfig_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sL_PHY_MAC_RLC_Config_r16SlUESelectedConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_UE_SelectedConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_PHY_MAC_RLC_Config_r16_Sl_UE_SelectedConfig_r16_Release:
				(*ie.Sl_UE_SelectedConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_r16_Sl_UE_SelectedConfig_r16_Setup:
				(*ie.Sl_UE_SelectedConfig_r16).Setup = new(SL_UE_SelectedConfig_r16)
				if err := (*ie.Sl_UE_SelectedConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-FreqInfoToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLPHYMACRLCConfigR16SlFreqInfoToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_FreqInfoToReleaseList_r16 = make([]SL_Freq_Id_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_FreqInfoToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-FreqInfoToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLPHYMACRLCConfigR16SlFreqInfoToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_FreqInfoToAddModList_r16 = make([]SL_FreqConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_FreqInfoToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-RLC-BearerToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sLPHYMACRLCConfigR16SlRLCBearerToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_BearerToReleaseList_r16 = make([]SL_RLC_BearerConfigIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_BearerToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-RLC-BearerToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(sLPHYMACRLCConfigR16SlRLCBearerToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_BearerToAddModList_r16 = make([]SL_RLC_BearerConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_BearerToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-MaxNumConsecutiveDTX-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sLPHYMACRLCConfigR16SlMaxNumConsecutiveDTXR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MaxNumConsecutiveDTX_r16 = &idx
		}
	}

	// 9. sl-CSI-Acquisition-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sLPHYMACRLCConfigR16SlCSIAcquisitionR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CSI_Acquisition_r16 = &idx
		}
	}

	// 10. sl-CSI-SchedulingRequestId-r16: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Sl_CSI_SchedulingRequestId_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SchedulingRequestId
			}{}
			choiceDec := d.NewChoiceDecoder(sL_PHY_MAC_RLC_Config_r16SlCSISchedulingRequestIdR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_CSI_SchedulingRequestId_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_PHY_MAC_RLC_Config_r16_Sl_CSI_SchedulingRequestId_r16_Release:
				(*ie.Sl_CSI_SchedulingRequestId_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_r16_Sl_CSI_SchedulingRequestId_r16_Setup:
				(*ie.Sl_CSI_SchedulingRequestId_r16).Setup = new(SchedulingRequestId)
				if err := (*ie.Sl_CSI_SchedulingRequestId_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. sl-SSB-PriorityNR-r16: integer
	{
		if seq.IsComponentPresent(9) {
			v, err := d.DecodeInteger(sLPHYMACRLCConfigR16SlSSBPriorityNRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SSB_PriorityNR_r16 = &v
		}
	}

	// 12. networkControlledSyncTx-r16: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sLPHYMACRLCConfigR16NetworkControlledSyncTxR16Constraints)
			if err != nil {
				return err
			}
			ie.NetworkControlledSyncTx_r16 = &idx
		}
	}

	return nil
}
