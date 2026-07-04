// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PHY-MAC-RLC-Config-v1700 (line 26998).

var sLPHYMACRLCConfigV1700Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DRX-Config-r17", Optional: true},
		{Name: "sl-RLC-ChannelToReleaseList-r17", Optional: true},
		{Name: "sl-RLC-ChannelToAddModList-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLPHYMACRLCConfigV1700SlRLCChannelToReleaseListR17Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

var sLPHYMACRLCConfigV1700ExtSlRLCBearerToAddModListSizeExtV1800Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

var sLPHYMACRLCConfigV1700ExtSlRLCBearerToReleaseListSizeExtV1800Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

var sLPHYMACRLCConfigV1700ExtSlFreqInfoToAddModListExtV1800Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLPHYMACRLCConfigV1700ExtSlLBTSchedulingRequestIdR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_LBT_SchedulingRequestId_r18_Release = 0
	SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_LBT_SchedulingRequestId_r18_Setup   = 1
)

var sLPHYMACRLCConfigV1700ExtSlSyncFreqListR18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

const (
	SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SyncTxMultiFreq_r18_True = 0
)

var sLPHYMACRLCConfigV1700ExtSlSyncTxMultiFreqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SyncTxMultiFreq_r18_True},
}

var sLPHYMACRLCConfigV1700ExtSlSCCHCarrierSetConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SCCH_CarrierSetConfig_r18_Release = 0
	SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SCCH_CarrierSetConfig_r18_Setup   = 1
)

var sLPHYMACRLCConfigV1700ExtSlPRSSchedulingRequestIdR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_PRS_SchedulingRequestId_r18_Release = 0
	SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_PRS_SchedulingRequestId_r18_Setup   = 1
)

type SL_PHY_MAC_RLC_Config_v1700 struct {
	Sl_DRX_Config_r17                       *SL_DRX_Config_r17
	Sl_RLC_ChannelToReleaseList_r17         []SL_RLC_ChannelID_r17
	Sl_RLC_ChannelToAddModList_r17          *SL_RLC_ChannelToAddModList_r17
	Sl_RLC_BearerToAddModListSizeExt_v1800  []SL_RLC_BearerConfig_r16
	Sl_RLC_BearerToReleaseListSizeExt_v1800 []SL_RLC_BearerConfigIndex_v1800
	Sl_FreqInfoToAddModListExt_v1800        []SL_FreqConfigExt_v1800
	Sl_LBT_SchedulingRequestId_r18          *struct {
		Choice  int
		Release *struct{}
		Setup   *SchedulingRequestId
	}
	Sl_SyncFreqList_r18          []SL_Freq_Id_r16
	Sl_SyncTxMultiFreq_r18       *int64
	Sl_MaxTransPowerCA_r18       *P_Max
	Sl_SCCH_CarrierSetConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_SCCH_CarrierSetConfigList_r18
	}
	Sl_PRS_SchedulingRequestId_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SchedulingRequestId
	}
}

func (ie *SL_PHY_MAC_RLC_Config_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPHYMACRLCConfigV1700Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_RLC_BearerToAddModListSizeExt_v1800 != nil || ie.Sl_RLC_BearerToReleaseListSizeExt_v1800 != nil || ie.Sl_FreqInfoToAddModListExt_v1800 != nil || ie.Sl_LBT_SchedulingRequestId_r18 != nil || ie.Sl_SyncFreqList_r18 != nil || ie.Sl_SyncTxMultiFreq_r18 != nil || ie.Sl_MaxTransPowerCA_r18 != nil || ie.Sl_SCCH_CarrierSetConfig_r18 != nil || ie.Sl_PRS_SchedulingRequestId_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DRX_Config_r17 != nil, ie.Sl_RLC_ChannelToReleaseList_r17 != nil, ie.Sl_RLC_ChannelToAddModList_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DRX-Config-r17: ref
	{
		if ie.Sl_DRX_Config_r17 != nil {
			if err := ie.Sl_DRX_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-RLC-ChannelToReleaseList-r17: sequence-of
	{
		if ie.Sl_RLC_ChannelToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPHYMACRLCConfigV1700SlRLCChannelToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_ChannelToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_RLC_ChannelToReleaseList_r17 {
				if err := ie.Sl_RLC_ChannelToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-RLC-ChannelToAddModList-r17: ref
	{
		if ie.Sl_RLC_ChannelToAddModList_r17 != nil {
			if err := ie.Sl_RLC_ChannelToAddModList_r17.Encode(e); err != nil {
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
					{Name: "sl-RLC-BearerToAddModListSizeExt-v1800", Optional: true},
					{Name: "sl-RLC-BearerToReleaseListSizeExt-v1800", Optional: true},
					{Name: "sl-FreqInfoToAddModListExt-v1800", Optional: true},
					{Name: "sl-LBT-SchedulingRequestId-r18", Optional: true},
					{Name: "sl-SyncFreqList-r18", Optional: true},
					{Name: "sl-SyncTxMultiFreq-r18", Optional: true},
					{Name: "sl-MaxTransPowerCA-r18", Optional: true},
					{Name: "sl-SCCH-CarrierSetConfig-r18", Optional: true},
					{Name: "sl-PRS-SchedulingRequestId-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_RLC_BearerToAddModListSizeExt_v1800 != nil, ie.Sl_RLC_BearerToReleaseListSizeExt_v1800 != nil, ie.Sl_FreqInfoToAddModListExt_v1800 != nil, ie.Sl_LBT_SchedulingRequestId_r18 != nil, ie.Sl_SyncFreqList_r18 != nil, ie.Sl_SyncTxMultiFreq_r18 != nil, ie.Sl_MaxTransPowerCA_r18 != nil, ie.Sl_SCCH_CarrierSetConfig_r18 != nil, ie.Sl_PRS_SchedulingRequestId_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_RLC_BearerToAddModListSizeExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLPHYMACRLCConfigV1700ExtSlRLCBearerToAddModListSizeExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_BearerToAddModListSizeExt_v1800))); err != nil {
					return err
				}
				for i := range ie.Sl_RLC_BearerToAddModListSizeExt_v1800 {
					if err := ie.Sl_RLC_BearerToAddModListSizeExt_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_RLC_BearerToReleaseListSizeExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLPHYMACRLCConfigV1700ExtSlRLCBearerToReleaseListSizeExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_BearerToReleaseListSizeExt_v1800))); err != nil {
					return err
				}
				for i := range ie.Sl_RLC_BearerToReleaseListSizeExt_v1800 {
					if err := ie.Sl_RLC_BearerToReleaseListSizeExt_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_FreqInfoToAddModListExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLPHYMACRLCConfigV1700ExtSlFreqInfoToAddModListExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_FreqInfoToAddModListExt_v1800))); err != nil {
					return err
				}
				for i := range ie.Sl_FreqInfoToAddModListExt_v1800 {
					if err := ie.Sl_FreqInfoToAddModListExt_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_LBT_SchedulingRequestId_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLPHYMACRLCConfigV1700ExtSlLBTSchedulingRequestIdR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_LBT_SchedulingRequestId_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_LBT_SchedulingRequestId_r18).Choice {
				case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_LBT_SchedulingRequestId_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_LBT_SchedulingRequestId_r18_Setup:
					if err := (*ie.Sl_LBT_SchedulingRequestId_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_SyncFreqList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLPHYMACRLCConfigV1700ExtSlSyncFreqListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_SyncFreqList_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_SyncFreqList_r18 {
					if err := ie.Sl_SyncFreqList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_SyncTxMultiFreq_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_SyncTxMultiFreq_r18, sLPHYMACRLCConfigV1700ExtSlSyncTxMultiFreqR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_MaxTransPowerCA_r18 != nil {
				if err := ie.Sl_MaxTransPowerCA_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_SCCH_CarrierSetConfig_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLPHYMACRLCConfigV1700ExtSlSCCHCarrierSetConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_SCCH_CarrierSetConfig_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_SCCH_CarrierSetConfig_r18).Choice {
				case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SCCH_CarrierSetConfig_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SCCH_CarrierSetConfig_r18_Setup:
					if err := (*ie.Sl_SCCH_CarrierSetConfig_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_PRS_SchedulingRequestId_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLPHYMACRLCConfigV1700ExtSlPRSSchedulingRequestIdR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PRS_SchedulingRequestId_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_PRS_SchedulingRequestId_r18).Choice {
				case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_PRS_SchedulingRequestId_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_PRS_SchedulingRequestId_r18_Setup:
					if err := (*ie.Sl_PRS_SchedulingRequestId_r18).Setup.Encode(ex); err != nil {
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

func (ie *SL_PHY_MAC_RLC_Config_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPHYMACRLCConfigV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DRX-Config-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_DRX_Config_r17 = new(SL_DRX_Config_r17)
			if err := ie.Sl_DRX_Config_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-RLC-ChannelToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLPHYMACRLCConfigV1700SlRLCChannelToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_ChannelToReleaseList_r17 = make([]SL_RLC_ChannelID_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_ChannelToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-RLC-ChannelToAddModList-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_RLC_ChannelToAddModList_r17 = new(SL_RLC_ChannelToAddModList_r17)
			if err := ie.Sl_RLC_ChannelToAddModList_r17.Decode(d); err != nil {
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
				{Name: "sl-RLC-BearerToAddModListSizeExt-v1800", Optional: true},
				{Name: "sl-RLC-BearerToReleaseListSizeExt-v1800", Optional: true},
				{Name: "sl-FreqInfoToAddModListExt-v1800", Optional: true},
				{Name: "sl-LBT-SchedulingRequestId-r18", Optional: true},
				{Name: "sl-SyncFreqList-r18", Optional: true},
				{Name: "sl-SyncTxMultiFreq-r18", Optional: true},
				{Name: "sl-MaxTransPowerCA-r18", Optional: true},
				{Name: "sl-SCCH-CarrierSetConfig-r18", Optional: true},
				{Name: "sl-PRS-SchedulingRequestId-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sLPHYMACRLCConfigV1700ExtSlRLCBearerToAddModListSizeExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_BearerToAddModListSizeExt_v1800 = make([]SL_RLC_BearerConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_BearerToAddModListSizeExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(sLPHYMACRLCConfigV1700ExtSlRLCBearerToReleaseListSizeExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_BearerToReleaseListSizeExt_v1800 = make([]SL_RLC_BearerConfigIndex_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_BearerToReleaseListSizeExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(sLPHYMACRLCConfigV1700ExtSlFreqInfoToAddModListExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_FreqInfoToAddModListExt_v1800 = make([]SL_FreqConfigExt_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_FreqInfoToAddModListExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Sl_LBT_SchedulingRequestId_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SchedulingRequestId
			}{}
			choiceDec := dx.NewChoiceDecoder(sLPHYMACRLCConfigV1700ExtSlLBTSchedulingRequestIdR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_LBT_SchedulingRequestId_r18).Choice = int(index)
			switch index {
			case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_LBT_SchedulingRequestId_r18_Release:
				(*ie.Sl_LBT_SchedulingRequestId_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_LBT_SchedulingRequestId_r18_Setup:
				(*ie.Sl_LBT_SchedulingRequestId_r18).Setup = new(SchedulingRequestId)
				if err := (*ie.Sl_LBT_SchedulingRequestId_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(sLPHYMACRLCConfigV1700ExtSlSyncFreqListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_SyncFreqList_r18 = make([]SL_Freq_Id_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_SyncFreqList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(sLPHYMACRLCConfigV1700ExtSlSyncTxMultiFreqR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SyncTxMultiFreq_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Sl_MaxTransPowerCA_r18 = new(P_Max)
			if err := ie.Sl_MaxTransPowerCA_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Sl_SCCH_CarrierSetConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_SCCH_CarrierSetConfigList_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sLPHYMACRLCConfigV1700ExtSlSCCHCarrierSetConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_SCCH_CarrierSetConfig_r18).Choice = int(index)
			switch index {
			case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SCCH_CarrierSetConfig_r18_Release:
				(*ie.Sl_SCCH_CarrierSetConfig_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_SCCH_CarrierSetConfig_r18_Setup:
				(*ie.Sl_SCCH_CarrierSetConfig_r18).Setup = new(SL_SCCH_CarrierSetConfigList_r18)
				if err := (*ie.Sl_SCCH_CarrierSetConfig_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			ie.Sl_PRS_SchedulingRequestId_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SchedulingRequestId
			}{}
			choiceDec := dx.NewChoiceDecoder(sLPHYMACRLCConfigV1700ExtSlPRSSchedulingRequestIdR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PRS_SchedulingRequestId_r18).Choice = int(index)
			switch index {
			case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_PRS_SchedulingRequestId_r18_Release:
				(*ie.Sl_PRS_SchedulingRequestId_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_PHY_MAC_RLC_Config_v1700_Ext_Sl_PRS_SchedulingRequestId_r18_Setup:
				(*ie.Sl_PRS_SchedulingRequestId_r18).Setup = new(SchedulingRequestId)
				if err := (*ie.Sl_PRS_SchedulingRequestId_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
