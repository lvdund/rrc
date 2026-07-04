// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB12-IEs-r16 (line 4278).

var sIB12IEsR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfigCommonNR-r16"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var sIB12IEsR16LateNonCriticalExtensionConstraints = per.SizeConstraints{}

const (
	SIB12_IEs_r16_Ext_Sl_L2U2N_Relay_r17_Enabled = 0
)

var sIB12IEsR16ExtSlL2U2NRelayR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_Sl_L2U2N_Relay_r17_Enabled},
}

const (
	SIB12_IEs_r16_Ext_Sl_NonRelayDiscovery_r17_Enabled = 0
)

var sIB12IEsR16ExtSlNonRelayDiscoveryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_Sl_NonRelayDiscovery_r17_Enabled},
}

const (
	SIB12_IEs_r16_Ext_Sl_L3U2N_RelayDiscovery_r17_Enabled = 0
)

var sIB12IEsR16ExtSlL3U2NRelayDiscoveryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_Sl_L3U2N_RelayDiscovery_r17_Enabled},
}

var sIB12IEsR16ExtSlFreqInfoListSizeExtV1800Constraints = per.SizeRange(1, common.MaxNrofFreqSL_1_r18)

var sIB12IEsR16ExtSlRLCBearerConfigListSizeExtV1800Constraints = per.SizeRange(1, common.MaxSL_LCID_r16)

var sIB12IEsR16ExtSlSyncFreqListR18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

const (
	SIB12_IEs_r16_Ext_Sl_SyncTxMultiFreq_r18_True = 0
)

var sIB12IEsR16ExtSlSyncTxMultiFreqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_Sl_SyncTxMultiFreq_r18_True},
}

const (
	SIB12_IEs_r16_Ext_Sl_L2_U2U_Relay_r18_Enabled = 0
)

var sIB12IEsR16ExtSlL2U2URelayR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_Sl_L2_U2U_Relay_r18_Enabled},
}

const (
	SIB12_IEs_r16_Ext_Sl_L3_U2U_RelayDiscovery_r18_Enabled = 0
)

var sIB12IEsR16ExtSlL3U2URelayDiscoveryR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_Sl_L3_U2U_RelayDiscovery_r18_Enabled},
}

const (
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms200  = 0
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms400  = 1
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms600  = 2
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms800  = 3
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms1200 = 4
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms2000 = 5
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms3000 = 6
	SIB12_IEs_r16_Ext_T400_U2U_r18_Ms4000 = 7
)

var sIB12IEsR16ExtT400U2UR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_T400_U2U_r18_Ms200, SIB12_IEs_r16_Ext_T400_U2U_r18_Ms400, SIB12_IEs_r16_Ext_T400_U2U_r18_Ms600, SIB12_IEs_r16_Ext_T400_U2U_r18_Ms800, SIB12_IEs_r16_Ext_T400_U2U_r18_Ms1200, SIB12_IEs_r16_Ext_T400_U2U_r18_Ms2000, SIB12_IEs_r16_Ext_T400_U2U_r18_Ms3000, SIB12_IEs_r16_Ext_T400_U2U_r18_Ms4000},
}

const (
	SIB12_IEs_r16_Ext_Sl_L2U2N_MH_Relay_r19_Enabled = 0
)

var sIB12IEsR16ExtSlL2U2NMHRelayR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_IEs_r16_Ext_Sl_L2U2N_MH_Relay_r19_Enabled},
}

type SIB12_IEs_r16 struct {
	Sl_ConfigCommonNR_r16                SL_ConfigCommonNR_r16
	LateNonCriticalExtension             []byte
	Sl_DRX_ConfigCommonGC_BC_r17         *SL_DRX_ConfigGC_BC_r17
	Sl_DiscConfigCommon_r17              *SL_DiscConfigCommon_r17
	Sl_L2U2N_Relay_r17                   *int64
	Sl_NonRelayDiscovery_r17             *int64
	Sl_L3U2N_RelayDiscovery_r17          *int64
	Sl_TimersAndConstantsRemoteUE_r17    *UE_TimersAndConstantsRemoteUE_r17
	Sl_FreqInfoListSizeExt_v1800         []SL_FreqConfigCommon_r16
	Sl_RLC_BearerConfigListSizeExt_v1800 []SL_RLC_BearerConfig_r16
	Sl_SyncFreqList_r18                  []SL_Freq_Id_r16
	Sl_SyncTxMultiFreq_r18               *int64
	Sl_MaxTransPowerCA_r18               *P_Max
	Sl_DiscConfigCommon_v1800            *SL_DiscConfigCommon_v1800
	Sl_L2_U2U_Relay_r18                  *int64
	Sl_L3_U2U_RelayDiscovery_r18         *int64
	T400_U2U_r18                         *int64
	Sl_DiscConfigCommon_v1840            *SL_DiscConfigCommon_v1840
	Sl_L2U2N_MH_Relay_r19                *int64
	Sl_DiscConfigCommon_v1900            *SL_DiscConfigCommon_v1900
}

func (ie *SIB12_IEs_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB12IEsR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_DRX_ConfigCommonGC_BC_r17 != nil || ie.Sl_DiscConfigCommon_r17 != nil || ie.Sl_L2U2N_Relay_r17 != nil || ie.Sl_NonRelayDiscovery_r17 != nil || ie.Sl_L3U2N_RelayDiscovery_r17 != nil || ie.Sl_TimersAndConstantsRemoteUE_r17 != nil
	hasExtGroup1 := ie.Sl_FreqInfoListSizeExt_v1800 != nil || ie.Sl_RLC_BearerConfigListSizeExt_v1800 != nil || ie.Sl_SyncFreqList_r18 != nil || ie.Sl_SyncTxMultiFreq_r18 != nil || ie.Sl_MaxTransPowerCA_r18 != nil || ie.Sl_DiscConfigCommon_v1800 != nil || ie.Sl_L2_U2U_Relay_r18 != nil || ie.Sl_L3_U2U_RelayDiscovery_r18 != nil || ie.T400_U2U_r18 != nil
	hasExtGroup2 := ie.Sl_DiscConfigCommon_v1840 != nil
	hasExtGroup3 := ie.Sl_L2U2N_MH_Relay_r19 != nil || ie.Sl_DiscConfigCommon_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. sl-ConfigCommonNR-r16: ref
	{
		if err := ie.Sl_ConfigCommonNR_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB12IEsR16LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-DRX-ConfigCommonGC-BC-r17", Optional: true},
					{Name: "sl-DiscConfigCommon-r17", Optional: true},
					{Name: "sl-L2U2N-Relay-r17", Optional: true},
					{Name: "sl-NonRelayDiscovery-r17", Optional: true},
					{Name: "sl-L3U2N-RelayDiscovery-r17", Optional: true},
					{Name: "sl-TimersAndConstantsRemoteUE-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DRX_ConfigCommonGC_BC_r17 != nil, ie.Sl_DiscConfigCommon_r17 != nil, ie.Sl_L2U2N_Relay_r17 != nil, ie.Sl_NonRelayDiscovery_r17 != nil, ie.Sl_L3U2N_RelayDiscovery_r17 != nil, ie.Sl_TimersAndConstantsRemoteUE_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_DRX_ConfigCommonGC_BC_r17 != nil {
				if err := ie.Sl_DRX_ConfigCommonGC_BC_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_DiscConfigCommon_r17 != nil {
				if err := ie.Sl_DiscConfigCommon_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_L2U2N_Relay_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_L2U2N_Relay_r17, sIB12IEsR16ExtSlL2U2NRelayR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_NonRelayDiscovery_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_NonRelayDiscovery_r17, sIB12IEsR16ExtSlNonRelayDiscoveryR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_L3U2N_RelayDiscovery_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_L3U2N_RelayDiscovery_r17, sIB12IEsR16ExtSlL3U2NRelayDiscoveryR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_TimersAndConstantsRemoteUE_r17 != nil {
				if err := ie.Sl_TimersAndConstantsRemoteUE_r17.Encode(ex); err != nil {
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
					{Name: "sl-FreqInfoListSizeExt-v1800", Optional: true},
					{Name: "sl-RLC-BearerConfigListSizeExt-v1800", Optional: true},
					{Name: "sl-SyncFreqList-r18", Optional: true},
					{Name: "sl-SyncTxMultiFreq-r18", Optional: true},
					{Name: "sl-MaxTransPowerCA-r18", Optional: true},
					{Name: "sl-DiscConfigCommon-v1800", Optional: true},
					{Name: "sl-L2-U2U-Relay-r18", Optional: true},
					{Name: "sl-L3-U2U-RelayDiscovery-r18", Optional: true},
					{Name: "t400-U2U-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_FreqInfoListSizeExt_v1800 != nil, ie.Sl_RLC_BearerConfigListSizeExt_v1800 != nil, ie.Sl_SyncFreqList_r18 != nil, ie.Sl_SyncTxMultiFreq_r18 != nil, ie.Sl_MaxTransPowerCA_r18 != nil, ie.Sl_DiscConfigCommon_v1800 != nil, ie.Sl_L2_U2U_Relay_r18 != nil, ie.Sl_L3_U2U_RelayDiscovery_r18 != nil, ie.T400_U2U_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_FreqInfoListSizeExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(sIB12IEsR16ExtSlFreqInfoListSizeExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_FreqInfoListSizeExt_v1800))); err != nil {
					return err
				}
				for i := range ie.Sl_FreqInfoListSizeExt_v1800 {
					if err := ie.Sl_FreqInfoListSizeExt_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_RLC_BearerConfigListSizeExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(sIB12IEsR16ExtSlRLCBearerConfigListSizeExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_BearerConfigListSizeExt_v1800))); err != nil {
					return err
				}
				for i := range ie.Sl_RLC_BearerConfigListSizeExt_v1800 {
					if err := ie.Sl_RLC_BearerConfigListSizeExt_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_SyncFreqList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sIB12IEsR16ExtSlSyncFreqListR18Constraints)
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
				if err := ex.EncodeEnumerated(*ie.Sl_SyncTxMultiFreq_r18, sIB12IEsR16ExtSlSyncTxMultiFreqR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_MaxTransPowerCA_r18 != nil {
				if err := ie.Sl_MaxTransPowerCA_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_DiscConfigCommon_v1800 != nil {
				if err := ie.Sl_DiscConfigCommon_v1800.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_L2_U2U_Relay_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_L2_U2U_Relay_r18, sIB12IEsR16ExtSlL2U2URelayR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_L3_U2U_RelayDiscovery_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_L3_U2U_RelayDiscovery_r18, sIB12IEsR16ExtSlL3U2URelayDiscoveryR18Constraints); err != nil {
					return err
				}
			}

			if ie.T400_U2U_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.T400_U2U_r18, sIB12IEsR16ExtT400U2UR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-DiscConfigCommon-v1840", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DiscConfigCommon_v1840 != nil}); err != nil {
				return err
			}

			if ie.Sl_DiscConfigCommon_v1840 != nil {
				if err := ie.Sl_DiscConfigCommon_v1840.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-L2U2N-MH-Relay-r19", Optional: true},
					{Name: "sl-DiscConfigCommon-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_L2U2N_MH_Relay_r19 != nil, ie.Sl_DiscConfigCommon_v1900 != nil}); err != nil {
				return err
			}

			if ie.Sl_L2U2N_MH_Relay_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_L2U2N_MH_Relay_r19, sIB12IEsR16ExtSlL2U2NMHRelayR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_DiscConfigCommon_v1900 != nil {
				if err := ie.Sl_DiscConfigCommon_v1900.Encode(ex); err != nil {
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

func (ie *SIB12_IEs_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB12IEsR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-ConfigCommonNR-r16: ref
	{
		if err := ie.Sl_ConfigCommonNR_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB12IEsR16LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
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
				{Name: "sl-DRX-ConfigCommonGC-BC-r17", Optional: true},
				{Name: "sl-DiscConfigCommon-r17", Optional: true},
				{Name: "sl-L2U2N-Relay-r17", Optional: true},
				{Name: "sl-NonRelayDiscovery-r17", Optional: true},
				{Name: "sl-L3U2N-RelayDiscovery-r17", Optional: true},
				{Name: "sl-TimersAndConstantsRemoteUE-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_DRX_ConfigCommonGC_BC_r17 = new(SL_DRX_ConfigGC_BC_r17)
			if err := ie.Sl_DRX_ConfigCommonGC_BC_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_DiscConfigCommon_r17 = new(SL_DiscConfigCommon_r17)
			if err := ie.Sl_DiscConfigCommon_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtSlL2U2NRelayR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_L2U2N_Relay_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtSlNonRelayDiscoveryR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NonRelayDiscovery_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtSlL3U2NRelayDiscoveryR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_L3U2N_RelayDiscovery_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Sl_TimersAndConstantsRemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17)
			if err := ie.Sl_TimersAndConstantsRemoteUE_r17.Decode(dx); err != nil {
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
				{Name: "sl-FreqInfoListSizeExt-v1800", Optional: true},
				{Name: "sl-RLC-BearerConfigListSizeExt-v1800", Optional: true},
				{Name: "sl-SyncFreqList-r18", Optional: true},
				{Name: "sl-SyncTxMultiFreq-r18", Optional: true},
				{Name: "sl-MaxTransPowerCA-r18", Optional: true},
				{Name: "sl-DiscConfigCommon-v1800", Optional: true},
				{Name: "sl-L2-U2U-Relay-r18", Optional: true},
				{Name: "sl-L3-U2U-RelayDiscovery-r18", Optional: true},
				{Name: "t400-U2U-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sIB12IEsR16ExtSlFreqInfoListSizeExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_FreqInfoListSizeExt_v1800 = make([]SL_FreqConfigCommon_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_FreqInfoListSizeExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(sIB12IEsR16ExtSlRLCBearerConfigListSizeExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_BearerConfigListSizeExt_v1800 = make([]SL_RLC_BearerConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_BearerConfigListSizeExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(sIB12IEsR16ExtSlSyncFreqListR18Constraints)
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

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtSlSyncTxMultiFreqR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SyncTxMultiFreq_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Sl_MaxTransPowerCA_r18 = new(P_Max)
			if err := ie.Sl_MaxTransPowerCA_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Sl_DiscConfigCommon_v1800 = new(SL_DiscConfigCommon_v1800)
			if err := ie.Sl_DiscConfigCommon_v1800.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtSlL2U2URelayR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_L2_U2U_Relay_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtSlL3U2URelayDiscoveryR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_L3_U2U_RelayDiscovery_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtT400U2UR18Constraints)
			if err != nil {
				return err
			}
			ie.T400_U2U_r18 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-DiscConfigCommon-v1840", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_DiscConfigCommon_v1840 = new(SL_DiscConfigCommon_v1840)
			if err := ie.Sl_DiscConfigCommon_v1840.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-L2U2N-MH-Relay-r19", Optional: true},
				{Name: "sl-DiscConfigCommon-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sIB12IEsR16ExtSlL2U2NMHRelayR19Constraints)
			if err != nil {
				return err
			}
			ie.Sl_L2U2N_MH_Relay_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_DiscConfigCommon_v1900 = new(SL_DiscConfigCommon_v1900)
			if err := ie.Sl_DiscConfigCommon_v1900.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
