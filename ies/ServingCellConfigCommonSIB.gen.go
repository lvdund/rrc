// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ServingCellConfigCommonSIB (line 14958).

var servingCellConfigCommonSIBConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "downlinkConfigCommon"},
		{Name: "uplinkConfigCommon", Optional: true},
		{Name: "supplementaryUplink", Optional: true},
		{Name: "n-TimingAdvanceOffset", Optional: true},
		{Name: "ssb-PositionsInBurst"},
		{Name: "ssb-PeriodicityServingCell"},
		{Name: "tdd-UL-DL-ConfigurationCommon", Optional: true},
		{Name: "ss-PBCH-BlockPower"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	ServingCellConfigCommonSIB_N_TimingAdvanceOffset_N0     = 0
	ServingCellConfigCommonSIB_N_TimingAdvanceOffset_N25600 = 1
	ServingCellConfigCommonSIB_N_TimingAdvanceOffset_N39936 = 2
)

var servingCellConfigCommonSIBNTimingAdvanceOffsetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommonSIB_N_TimingAdvanceOffset_N0, ServingCellConfigCommonSIB_N_TimingAdvanceOffset_N25600, ServingCellConfigCommonSIB_N_TimingAdvanceOffset_N39936},
}

const (
	ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms5   = 0
	ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms10  = 1
	ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms20  = 2
	ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms40  = 3
	ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms80  = 4
	ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms160 = 5
)

var servingCellConfigCommonSIBSsbPeriodicityServingCellConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms5, ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms10, ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms20, ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms40, ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms80, ServingCellConfigCommonSIB_Ssb_PeriodicityServingCell_Ms160},
}

var servingCellConfigCommonSIBSsPBCHBlockPowerConstraints = per.Constrained(-60, 50)

var servingCellConfigCommonSIBExtChannelAccessModeR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dynamic"},
		{Name: "semiStatic"},
	},
}

const (
	ServingCellConfigCommonSIB_Ext_ChannelAccessMode_r16_Dynamic    = 0
	ServingCellConfigCommonSIB_Ext_ChannelAccessMode_r16_SemiStatic = 1
)

const (
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms0dot5 = 0
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms1     = 1
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms2     = 2
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms3     = 3
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms4     = 4
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms5     = 5
)

var servingCellConfigCommonSIBExtDiscoveryBurstWindowLengthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms0dot5, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms1, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms2, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms3, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms4, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_r16_Ms5},
}

const (
	ServingCellConfigCommonSIB_Ext_ChannelAccessMode2_r17_Enabled = 0
)

var servingCellConfigCommonSIBExtChannelAccessMode2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommonSIB_Ext_ChannelAccessMode2_r17_Enabled},
}

const (
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot125 = 0
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot25  = 1
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot5   = 2
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot75  = 3
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms1       = 4
	ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms1dot25  = 5
)

var servingCellConfigCommonSIBExtDiscoveryBurstWindowLengthV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot125, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot25, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot5, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms0dot75, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms1, ServingCellConfigCommonSIB_Ext_DiscoveryBurstWindowLength_v1700_Ms1dot25},
}

const (
	ServingCellConfigCommonSIB_Ext_EnhancedMeasurementNGSO_r17_True = 0
)

var servingCellConfigCommonSIBExtEnhancedMeasurementNGSOR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommonSIB_Ext_EnhancedMeasurementNGSO_r17_True},
}

const (
	ServingCellConfigCommonSIB_Ext_Ra_ChannelAccess_r17_Enabled = 0
)

var servingCellConfigCommonSIBExtRaChannelAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommonSIB_Ext_Ra_ChannelAccess_r17_Enabled},
}

var servingCellConfigCommonSIBSsbPositionsInBurstConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inOneGroup"},
		{Name: "groupPresence", Optional: true},
	},
}

type ServingCellConfigCommonSIB struct {
	DownlinkConfigCommon  DownlinkConfigCommonSIB
	UplinkConfigCommon    *UplinkConfigCommonSIB
	SupplementaryUplink   *UplinkConfigCommonSIB
	N_TimingAdvanceOffset *int64
	Ssb_PositionsInBurst  struct {
		InOneGroup    per.BitString
		GroupPresence *per.BitString
	}
	Ssb_PeriodicityServingCell    int64
	Tdd_UL_DL_ConfigurationCommon *TDD_UL_DL_ConfigCommon
	Ss_PBCH_BlockPower            int64
	ChannelAccessMode_r16         *struct {
		Choice     int
		Dynamic    *struct{}
		SemiStatic *SemiStaticChannelAccessConfig_r16
	}
	DiscoveryBurstWindowLength_r16   *int64
	HighSpeedConfig_r16              *HighSpeedConfig_r16
	ChannelAccessMode2_r17           *int64
	DiscoveryBurstWindowLength_v1700 *int64
	HighSpeedConfigFR2_r17           *HighSpeedConfigFR2_r17
	UplinkConfigCommon_v1700         *UplinkConfigCommonSIB_v1700
	EnhancedMeasurementNGSO_r17      *int64
	Ra_ChannelAccess_r17             *int64
	DownlinkConfigCommon_v1760       *DownlinkConfigCommonSIB_v1760
	UplinkConfigCommon_v1760         *UplinkConfigCommonSIB_v1760
}

func (ie *ServingCellConfigCommonSIB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servingCellConfigCommonSIBConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ChannelAccessMode_r16 != nil || ie.DiscoveryBurstWindowLength_r16 != nil || ie.HighSpeedConfig_r16 != nil
	hasExtGroup1 := ie.ChannelAccessMode2_r17 != nil || ie.DiscoveryBurstWindowLength_v1700 != nil || ie.HighSpeedConfigFR2_r17 != nil || ie.UplinkConfigCommon_v1700 != nil
	hasExtGroup2 := ie.EnhancedMeasurementNGSO_r17 != nil
	hasExtGroup3 := ie.Ra_ChannelAccess_r17 != nil
	hasExtGroup4 := ie.DownlinkConfigCommon_v1760 != nil || ie.UplinkConfigCommon_v1760 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkConfigCommon != nil, ie.SupplementaryUplink != nil, ie.N_TimingAdvanceOffset != nil, ie.Tdd_UL_DL_ConfigurationCommon != nil}); err != nil {
		return err
	}

	// 3. downlinkConfigCommon: ref
	{
		if err := ie.DownlinkConfigCommon.Encode(e); err != nil {
			return err
		}
	}

	// 4. uplinkConfigCommon: ref
	{
		if ie.UplinkConfigCommon != nil {
			if err := ie.UplinkConfigCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. supplementaryUplink: ref
	{
		if ie.SupplementaryUplink != nil {
			if err := ie.SupplementaryUplink.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. n-TimingAdvanceOffset: enumerated
	{
		if ie.N_TimingAdvanceOffset != nil {
			if err := e.EncodeEnumerated(*ie.N_TimingAdvanceOffset, servingCellConfigCommonSIBNTimingAdvanceOffsetConstraints); err != nil {
				return err
			}
		}
	}

	// 7. ssb-PositionsInBurst: sequence
	{
		{
			c := &ie.Ssb_PositionsInBurst
			servingCellConfigCommonSIBSsbPositionsInBurstSeq := e.NewSequenceEncoder(servingCellConfigCommonSIBSsbPositionsInBurstConstraints)
			if err := servingCellConfigCommonSIBSsbPositionsInBurstSeq.EncodePreamble([]bool{c.GroupPresence != nil}); err != nil {
				return err
			}
			if err := e.EncodeBitString(c.InOneGroup, per.FixedSize(8)); err != nil {
				return err
			}
			if c.GroupPresence != nil {
				if err := e.EncodeBitString((*c.GroupPresence), per.FixedSize(8)); err != nil {
					return err
				}
			}
		}
	}

	// 8. ssb-PeriodicityServingCell: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ssb_PeriodicityServingCell, servingCellConfigCommonSIBSsbPeriodicityServingCellConstraints); err != nil {
			return err
		}
	}

	// 9. tdd-UL-DL-ConfigurationCommon: ref
	{
		if ie.Tdd_UL_DL_ConfigurationCommon != nil {
			if err := ie.Tdd_UL_DL_ConfigurationCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. ss-PBCH-BlockPower: integer
	{
		if err := e.EncodeInteger(ie.Ss_PBCH_BlockPower, servingCellConfigCommonSIBSsPBCHBlockPowerConstraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "channelAccessMode-r16", Optional: true},
					{Name: "discoveryBurstWindowLength-r16", Optional: true},
					{Name: "highSpeedConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ChannelAccessMode_r16 != nil, ie.DiscoveryBurstWindowLength_r16 != nil, ie.HighSpeedConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.ChannelAccessMode_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigCommonSIBExtChannelAccessModeR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelAccessMode_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelAccessMode_r16).Choice {
				case ServingCellConfigCommonSIB_Ext_ChannelAccessMode_r16_Dynamic:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfigCommonSIB_Ext_ChannelAccessMode_r16_SemiStatic:
					if err := (*ie.ChannelAccessMode_r16).SemiStatic.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.DiscoveryBurstWindowLength_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.DiscoveryBurstWindowLength_r16, servingCellConfigCommonSIBExtDiscoveryBurstWindowLengthR16Constraints); err != nil {
					return err
				}
			}

			if ie.HighSpeedConfig_r16 != nil {
				if err := ie.HighSpeedConfig_r16.Encode(ex); err != nil {
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
					{Name: "channelAccessMode2-r17", Optional: true},
					{Name: "discoveryBurstWindowLength-v1700", Optional: true},
					{Name: "highSpeedConfigFR2-r17", Optional: true},
					{Name: "uplinkConfigCommon-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ChannelAccessMode2_r17 != nil, ie.DiscoveryBurstWindowLength_v1700 != nil, ie.HighSpeedConfigFR2_r17 != nil, ie.UplinkConfigCommon_v1700 != nil}); err != nil {
				return err
			}

			if ie.ChannelAccessMode2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ChannelAccessMode2_r17, servingCellConfigCommonSIBExtChannelAccessMode2R17Constraints); err != nil {
					return err
				}
			}

			if ie.DiscoveryBurstWindowLength_v1700 != nil {
				if err := ex.EncodeEnumerated(*ie.DiscoveryBurstWindowLength_v1700, servingCellConfigCommonSIBExtDiscoveryBurstWindowLengthV1700Constraints); err != nil {
					return err
				}
			}

			if ie.HighSpeedConfigFR2_r17 != nil {
				if err := ie.HighSpeedConfigFR2_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.UplinkConfigCommon_v1700 != nil {
				if err := ie.UplinkConfigCommon_v1700.Encode(ex); err != nil {
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
					{Name: "enhancedMeasurementNGSO-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnhancedMeasurementNGSO_r17 != nil}); err != nil {
				return err
			}

			if ie.EnhancedMeasurementNGSO_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedMeasurementNGSO_r17, servingCellConfigCommonSIBExtEnhancedMeasurementNGSOR17Constraints); err != nil {
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
					{Name: "ra-ChannelAccess-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_ChannelAccess_r17 != nil}); err != nil {
				return err
			}

			if ie.Ra_ChannelAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ra_ChannelAccess_r17, servingCellConfigCommonSIBExtRaChannelAccessR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "downlinkConfigCommon-v1760", Optional: true},
					{Name: "uplinkConfigCommon-v1760", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DownlinkConfigCommon_v1760 != nil, ie.UplinkConfigCommon_v1760 != nil}); err != nil {
				return err
			}

			if ie.DownlinkConfigCommon_v1760 != nil {
				if err := ie.DownlinkConfigCommon_v1760.Encode(ex); err != nil {
					return err
				}
			}

			if ie.UplinkConfigCommon_v1760 != nil {
				if err := ie.UplinkConfigCommon_v1760.Encode(ex); err != nil {
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

func (ie *ServingCellConfigCommonSIB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servingCellConfigCommonSIBConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. downlinkConfigCommon: ref
	{
		if err := ie.DownlinkConfigCommon.Decode(d); err != nil {
			return err
		}
	}

	// 4. uplinkConfigCommon: ref
	{
		if seq.IsComponentPresent(1) {
			ie.UplinkConfigCommon = new(UplinkConfigCommonSIB)
			if err := ie.UplinkConfigCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. supplementaryUplink: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SupplementaryUplink = new(UplinkConfigCommonSIB)
			if err := ie.SupplementaryUplink.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. n-TimingAdvanceOffset: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(servingCellConfigCommonSIBNTimingAdvanceOffsetConstraints)
			if err != nil {
				return err
			}
			ie.N_TimingAdvanceOffset = &idx
		}
	}

	// 7. ssb-PositionsInBurst: sequence
	{
		{
			c := &ie.Ssb_PositionsInBurst
			servingCellConfigCommonSIBSsbPositionsInBurstSeq := d.NewSequenceDecoder(servingCellConfigCommonSIBSsbPositionsInBurstConstraints)
			if err := servingCellConfigCommonSIBSsbPositionsInBurstSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				c.InOneGroup = v
			}
			if servingCellConfigCommonSIBSsbPositionsInBurstSeq.IsComponentPresent(1) {
				c.GroupPresence = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*c.GroupPresence) = v
			}
		}
	}

	// 8. ssb-PeriodicityServingCell: enumerated
	{
		v5, err := d.DecodeEnumerated(servingCellConfigCommonSIBSsbPeriodicityServingCellConstraints)
		if err != nil {
			return err
		}
		ie.Ssb_PeriodicityServingCell = v5
	}

	// 9. tdd-UL-DL-ConfigurationCommon: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Tdd_UL_DL_ConfigurationCommon = new(TDD_UL_DL_ConfigCommon)
			if err := ie.Tdd_UL_DL_ConfigurationCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. ss-PBCH-BlockPower: integer
	{
		v7, err := d.DecodeInteger(servingCellConfigCommonSIBSsPBCHBlockPowerConstraints)
		if err != nil {
			return err
		}
		ie.Ss_PBCH_BlockPower = v7
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
				{Name: "channelAccessMode-r16", Optional: true},
				{Name: "discoveryBurstWindowLength-r16", Optional: true},
				{Name: "highSpeedConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ChannelAccessMode_r16 = &struct {
				Choice     int
				Dynamic    *struct{}
				SemiStatic *SemiStaticChannelAccessConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigCommonSIBExtChannelAccessModeR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelAccessMode_r16).Choice = int(index)
			switch index {
			case ServingCellConfigCommonSIB_Ext_ChannelAccessMode_r16_Dynamic:
				(*ie.ChannelAccessMode_r16).Dynamic = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfigCommonSIB_Ext_ChannelAccessMode_r16_SemiStatic:
				(*ie.ChannelAccessMode_r16).SemiStatic = new(SemiStaticChannelAccessConfig_r16)
				if err := (*ie.ChannelAccessMode_r16).SemiStatic.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonSIBExtDiscoveryBurstWindowLengthR16Constraints)
			if err != nil {
				return err
			}
			ie.DiscoveryBurstWindowLength_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.HighSpeedConfig_r16 = new(HighSpeedConfig_r16)
			if err := ie.HighSpeedConfig_r16.Decode(dx); err != nil {
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
				{Name: "channelAccessMode2-r17", Optional: true},
				{Name: "discoveryBurstWindowLength-v1700", Optional: true},
				{Name: "highSpeedConfigFR2-r17", Optional: true},
				{Name: "uplinkConfigCommon-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonSIBExtChannelAccessMode2R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelAccessMode2_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonSIBExtDiscoveryBurstWindowLengthV1700Constraints)
			if err != nil {
				return err
			}
			ie.DiscoveryBurstWindowLength_v1700 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.HighSpeedConfigFR2_r17 = new(HighSpeedConfigFR2_r17)
			if err := ie.HighSpeedConfigFR2_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.UplinkConfigCommon_v1700 = new(UplinkConfigCommonSIB_v1700)
			if err := ie.UplinkConfigCommon_v1700.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "enhancedMeasurementNGSO-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonSIBExtEnhancedMeasurementNGSOR17Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedMeasurementNGSO_r17 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-ChannelAccess-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonSIBExtRaChannelAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.Ra_ChannelAccess_r17 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "downlinkConfigCommon-v1760", Optional: true},
				{Name: "uplinkConfigCommon-v1760", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.DownlinkConfigCommon_v1760 = new(DownlinkConfigCommonSIB_v1760)
			if err := ie.DownlinkConfigCommon_v1760.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.UplinkConfigCommon_v1760 = new(UplinkConfigCommonSIB_v1760)
			if err := ie.UplinkConfigCommon_v1760.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
