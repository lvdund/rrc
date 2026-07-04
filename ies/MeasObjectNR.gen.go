// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasObjectNR (line 9431).

var measObjectNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssbFrequency", Optional: true},
		{Name: "ssbSubcarrierSpacing", Optional: true},
		{Name: "smtc1", Optional: true},
		{Name: "smtc2", Optional: true},
		{Name: "refFreqCSI-RS", Optional: true},
		{Name: "referenceSignalConfig"},
		{Name: "absThreshSS-BlocksConsolidation", Optional: true},
		{Name: "absThreshCSI-RS-Consolidation", Optional: true},
		{Name: "nrofSS-BlocksToAverage", Optional: true},
		{Name: "nrofCSI-RS-ResourcesToAverage", Optional: true},
		{Name: "quantityConfigIndex"},
		{Name: "offsetMO"},
		{Name: "cellsToRemoveList", Optional: true},
		{Name: "cellsToAddModList", Optional: true},
		{Name: "excludedCellsToRemoveList", Optional: true},
		{Name: "excludedCellsToAddModList", Optional: true},
		{Name: "allowedCellsToRemoveList", Optional: true},
		{Name: "allowedCellsToAddModList", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
	},
}

var measObjectNRNrofSSBlocksToAverageConstraints = per.Constrained(2, common.MaxNrofSS_BlocksToAverage)

var measObjectNRNrofCSIRSResourcesToAverageConstraints = per.Constrained(2, common.MaxNrofCSI_RS_ResourcesToAverage)

var measObjectNRQuantityConfigIndexConstraints = per.Constrained(1, common.MaxNrofQuantityConfig)

var measObjectNRExcludedCellsToAddModListConstraints = per.SizeRange(1, common.MaxNrofPCI_Ranges)

var measObjectNRAllowedCellsToAddModListConstraints = per.SizeRange(1, common.MaxNrofPCI_Ranges)

const (
	MeasObjectNR_Ext_MeasCycleSCell_Sf160  = 0
	MeasObjectNR_Ext_MeasCycleSCell_Sf256  = 1
	MeasObjectNR_Ext_MeasCycleSCell_Sf320  = 2
	MeasObjectNR_Ext_MeasCycleSCell_Sf512  = 3
	MeasObjectNR_Ext_MeasCycleSCell_Sf640  = 4
	MeasObjectNR_Ext_MeasCycleSCell_Sf1024 = 5
	MeasObjectNR_Ext_MeasCycleSCell_Sf1280 = 6
)

var measObjectNRExtMeasCycleSCellConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasObjectNR_Ext_MeasCycleSCell_Sf160, MeasObjectNR_Ext_MeasCycleSCell_Sf256, MeasObjectNR_Ext_MeasCycleSCell_Sf320, MeasObjectNR_Ext_MeasCycleSCell_Sf512, MeasObjectNR_Ext_MeasCycleSCell_Sf640, MeasObjectNR_Ext_MeasCycleSCell_Sf1024, MeasObjectNR_Ext_MeasCycleSCell_Sf1280},
}

var measObjectNRExtRmtcConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasObjectNR_Ext_Rmtc_Config_r16_Release = 0
	MeasObjectNR_Ext_Rmtc_Config_r16_Setup   = 1
)

var measObjectNRExtT312R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MeasObjectNR_Ext_T312_r16_Release = 0
	MeasObjectNR_Ext_T312_r16_Setup   = 1
)

const (
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms160  = 0
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms256  = 1
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms320  = 2
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms512  = 3
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms640  = 4
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms1024 = 5
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms1280 = 6
	MeasObjectNR_Ext_MeasCyclePSCell_r17_Spare1 = 7
)

var measObjectNRExtMeasCyclePSCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms160, MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms256, MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms320, MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms512, MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms640, MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms1024, MeasObjectNR_Ext_MeasCyclePSCell_r17_Ms1280, MeasObjectNR_Ext_MeasCyclePSCell_r17_Spare1},
}

const (
	MeasObjectNR_Ext_NeighSCellMeasSkipping_r19_Enabled = 0
)

var measObjectNRExtNeighSCellMeasSkippingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasObjectNR_Ext_NeighSCellMeasSkipping_r19_Enabled},
}

type MeasObjectNR struct {
	SsbFrequency                    *ARFCN_ValueNR
	SsbSubcarrierSpacing            *SubcarrierSpacing
	Smtc1                           *SSB_MTC
	Smtc2                           *SSB_MTC2
	RefFreqCSI_RS                   *ARFCN_ValueNR
	ReferenceSignalConfig           ReferenceSignalConfig
	AbsThreshSS_BlocksConsolidation *ThresholdNR
	AbsThreshCSI_RS_Consolidation   *ThresholdNR
	NrofSS_BlocksToAverage          *int64
	NrofCSI_RS_ResourcesToAverage   *int64
	QuantityConfigIndex             int64
	OffsetMO                        Q_OffsetRangeList
	CellsToRemoveList               *PCI_List
	CellsToAddModList               *CellsToAddModList
	ExcludedCellsToRemoveList       *PCI_RangeIndexList
	ExcludedCellsToAddModList       []PCI_RangeElement
	AllowedCellsToRemoveList        *PCI_RangeIndexList
	AllowedCellsToAddModList        []PCI_RangeElement
	FreqBandIndicatorNR             *FreqBandIndicatorNR
	MeasCycleSCell                  *int64
	Smtc3list_r16                   *SSB_MTC3List_r16
	Rmtc_Config_r16                 *struct {
		Choice  int
		Release *struct{}
		Setup   *RMTC_Config_r16
	}
	T312_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *T312_r16
	}
	AssociatedMeasGapSSB_r17      *MeasGapId_r17
	AssociatedMeasGapCSIRS_r17    *MeasGapId_r17
	Smtc4list_r17                 *SSB_MTC4List_r17
	MeasCyclePSCell_r17           *int64
	CellsToAddModListExt_v1710    *CellsToAddModListExt_v1710
	AssociatedMeasGapSSB2_v1720   *MeasGapId_r17
	AssociatedMeasGapCSIRS2_v1720 *MeasGapId_r17
	MeasSequence_r18              *MeasSequence_r18
	CellsToAddModListExt_v1800    *CellsToAddModListExt_v1800
	Smtc5list_r19                 *SSB_MTC5List_r19
	Smtc6list_r19                 *SSB_MTC6List_r19
	Smtc7_SSBAdapt_r19            *SSB_MTC_SSBAdapt_r19
	NeighSCellMeasSkipping_r19    *int64
}

func (ie *MeasObjectNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectNRConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.FreqBandIndicatorNR != nil || ie.MeasCycleSCell != nil
	hasExtGroup1 := ie.Smtc3list_r16 != nil || ie.Rmtc_Config_r16 != nil || ie.T312_r16 != nil
	hasExtGroup2 := ie.AssociatedMeasGapSSB_r17 != nil || ie.AssociatedMeasGapCSIRS_r17 != nil || ie.Smtc4list_r17 != nil || ie.MeasCyclePSCell_r17 != nil || ie.CellsToAddModListExt_v1710 != nil
	hasExtGroup3 := ie.AssociatedMeasGapSSB2_v1720 != nil || ie.AssociatedMeasGapCSIRS2_v1720 != nil
	hasExtGroup4 := ie.MeasSequence_r18 != nil || ie.CellsToAddModListExt_v1800 != nil
	hasExtGroup5 := ie.Smtc5list_r19 != nil || ie.Smtc6list_r19 != nil || ie.Smtc7_SSBAdapt_r19 != nil || ie.NeighSCellMeasSkipping_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SsbFrequency != nil, ie.SsbSubcarrierSpacing != nil, ie.Smtc1 != nil, ie.Smtc2 != nil, ie.RefFreqCSI_RS != nil, ie.AbsThreshSS_BlocksConsolidation != nil, ie.AbsThreshCSI_RS_Consolidation != nil, ie.NrofSS_BlocksToAverage != nil, ie.NrofCSI_RS_ResourcesToAverage != nil, ie.CellsToRemoveList != nil, ie.CellsToAddModList != nil, ie.ExcludedCellsToRemoveList != nil, ie.ExcludedCellsToAddModList != nil, ie.AllowedCellsToRemoveList != nil, ie.AllowedCellsToAddModList != nil}); err != nil {
		return err
	}

	// 3. ssbFrequency: ref
	{
		if ie.SsbFrequency != nil {
			if err := ie.SsbFrequency.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ssbSubcarrierSpacing: ref
	{
		if ie.SsbSubcarrierSpacing != nil {
			if err := ie.SsbSubcarrierSpacing.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. smtc1: ref
	{
		if ie.Smtc1 != nil {
			if err := ie.Smtc1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. smtc2: ref
	{
		if ie.Smtc2 != nil {
			if err := ie.Smtc2.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. refFreqCSI-RS: ref
	{
		if ie.RefFreqCSI_RS != nil {
			if err := ie.RefFreqCSI_RS.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. referenceSignalConfig: ref
	{
		if err := ie.ReferenceSignalConfig.Encode(e); err != nil {
			return err
		}
	}

	// 9. absThreshSS-BlocksConsolidation: ref
	{
		if ie.AbsThreshSS_BlocksConsolidation != nil {
			if err := ie.AbsThreshSS_BlocksConsolidation.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. absThreshCSI-RS-Consolidation: ref
	{
		if ie.AbsThreshCSI_RS_Consolidation != nil {
			if err := ie.AbsThreshCSI_RS_Consolidation.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. nrofSS-BlocksToAverage: integer
	{
		if ie.NrofSS_BlocksToAverage != nil {
			if err := e.EncodeInteger(*ie.NrofSS_BlocksToAverage, measObjectNRNrofSSBlocksToAverageConstraints); err != nil {
				return err
			}
		}
	}

	// 12. nrofCSI-RS-ResourcesToAverage: integer
	{
		if ie.NrofCSI_RS_ResourcesToAverage != nil {
			if err := e.EncodeInteger(*ie.NrofCSI_RS_ResourcesToAverage, measObjectNRNrofCSIRSResourcesToAverageConstraints); err != nil {
				return err
			}
		}
	}

	// 13. quantityConfigIndex: integer
	{
		if err := e.EncodeInteger(ie.QuantityConfigIndex, measObjectNRQuantityConfigIndexConstraints); err != nil {
			return err
		}
	}

	// 14. offsetMO: ref
	{
		if err := ie.OffsetMO.Encode(e); err != nil {
			return err
		}
	}

	// 15. cellsToRemoveList: ref
	{
		if ie.CellsToRemoveList != nil {
			if err := ie.CellsToRemoveList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 16. cellsToAddModList: ref
	{
		if ie.CellsToAddModList != nil {
			if err := ie.CellsToAddModList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 17. excludedCellsToRemoveList: ref
	{
		if ie.ExcludedCellsToRemoveList != nil {
			if err := ie.ExcludedCellsToRemoveList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 18. excludedCellsToAddModList: sequence-of
	{
		if ie.ExcludedCellsToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(measObjectNRExcludedCellsToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ExcludedCellsToAddModList))); err != nil {
				return err
			}
			for i := range ie.ExcludedCellsToAddModList {
				if err := ie.ExcludedCellsToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 19. allowedCellsToRemoveList: ref
	{
		if ie.AllowedCellsToRemoveList != nil {
			if err := ie.AllowedCellsToRemoveList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 20. allowedCellsToAddModList: sequence-of
	{
		if ie.AllowedCellsToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(measObjectNRAllowedCellsToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.AllowedCellsToAddModList))); err != nil {
				return err
			}
			for i := range ie.AllowedCellsToAddModList {
				if err := ie.AllowedCellsToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "freqBandIndicatorNR", Optional: true},
					{Name: "measCycleSCell", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FreqBandIndicatorNR != nil, ie.MeasCycleSCell != nil}); err != nil {
				return err
			}

			if ie.FreqBandIndicatorNR != nil {
				if err := ie.FreqBandIndicatorNR.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasCycleSCell != nil {
				if err := ex.EncodeEnumerated(*ie.MeasCycleSCell, measObjectNRExtMeasCycleSCellConstraints); err != nil {
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
					{Name: "smtc3list-r16", Optional: true},
					{Name: "rmtc-Config-r16", Optional: true},
					{Name: "t312-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Smtc3list_r16 != nil, ie.Rmtc_Config_r16 != nil, ie.T312_r16 != nil}); err != nil {
				return err
			}

			if ie.Smtc3list_r16 != nil {
				if err := ie.Smtc3list_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Rmtc_Config_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(measObjectNRExtRmtcConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Rmtc_Config_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Rmtc_Config_r16).Choice {
				case MeasObjectNR_Ext_Rmtc_Config_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasObjectNR_Ext_Rmtc_Config_r16_Setup:
					if err := (*ie.Rmtc_Config_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.T312_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(measObjectNRExtT312R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.T312_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.T312_r16).Choice {
				case MeasObjectNR_Ext_T312_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MeasObjectNR_Ext_T312_r16_Setup:
					if err := (*ie.T312_r16).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "associatedMeasGapSSB-r17", Optional: true},
					{Name: "associatedMeasGapCSIRS-r17", Optional: true},
					{Name: "smtc4list-r17", Optional: true},
					{Name: "measCyclePSCell-r17", Optional: true},
					{Name: "cellsToAddModListExt-v1710", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AssociatedMeasGapSSB_r17 != nil, ie.AssociatedMeasGapCSIRS_r17 != nil, ie.Smtc4list_r17 != nil, ie.MeasCyclePSCell_r17 != nil, ie.CellsToAddModListExt_v1710 != nil}); err != nil {
				return err
			}

			if ie.AssociatedMeasGapSSB_r17 != nil {
				if err := ie.AssociatedMeasGapSSB_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.AssociatedMeasGapCSIRS_r17 != nil {
				if err := ie.AssociatedMeasGapCSIRS_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Smtc4list_r17 != nil {
				if err := ie.Smtc4list_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasCyclePSCell_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MeasCyclePSCell_r17, measObjectNRExtMeasCyclePSCellR17Constraints); err != nil {
					return err
				}
			}

			if ie.CellsToAddModListExt_v1710 != nil {
				if err := ie.CellsToAddModListExt_v1710.Encode(ex); err != nil {
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
					{Name: "associatedMeasGapSSB2-v1720", Optional: true},
					{Name: "associatedMeasGapCSIRS2-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AssociatedMeasGapSSB2_v1720 != nil, ie.AssociatedMeasGapCSIRS2_v1720 != nil}); err != nil {
				return err
			}

			if ie.AssociatedMeasGapSSB2_v1720 != nil {
				if err := ie.AssociatedMeasGapSSB2_v1720.Encode(ex); err != nil {
					return err
				}
			}

			if ie.AssociatedMeasGapCSIRS2_v1720 != nil {
				if err := ie.AssociatedMeasGapCSIRS2_v1720.Encode(ex); err != nil {
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
					{Name: "measSequence-r18", Optional: true},
					{Name: "cellsToAddModListExt-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasSequence_r18 != nil, ie.CellsToAddModListExt_v1800 != nil}); err != nil {
				return err
			}

			if ie.MeasSequence_r18 != nil {
				if err := ie.MeasSequence_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CellsToAddModListExt_v1800 != nil {
				if err := ie.CellsToAddModListExt_v1800.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "smtc5list-r19", Optional: true},
					{Name: "smtc6list-r19", Optional: true},
					{Name: "smtc7-SSBAdapt-r19", Optional: true},
					{Name: "neighSCellMeasSkipping-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Smtc5list_r19 != nil, ie.Smtc6list_r19 != nil, ie.Smtc7_SSBAdapt_r19 != nil, ie.NeighSCellMeasSkipping_r19 != nil}); err != nil {
				return err
			}

			if ie.Smtc5list_r19 != nil {
				if err := ie.Smtc5list_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Smtc6list_r19 != nil {
				if err := ie.Smtc6list_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Smtc7_SSBAdapt_r19 != nil {
				if err := ie.Smtc7_SSBAdapt_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.NeighSCellMeasSkipping_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.NeighSCellMeasSkipping_r19, measObjectNRExtNeighSCellMeasSkippingR19Constraints); err != nil {
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

func (ie *MeasObjectNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ssbFrequency: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SsbFrequency = new(ARFCN_ValueNR)
			if err := ie.SsbFrequency.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ssbSubcarrierSpacing: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SsbSubcarrierSpacing = new(SubcarrierSpacing)
			if err := ie.SsbSubcarrierSpacing.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. smtc1: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Smtc1 = new(SSB_MTC)
			if err := ie.Smtc1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. smtc2: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Smtc2 = new(SSB_MTC2)
			if err := ie.Smtc2.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. refFreqCSI-RS: ref
	{
		if seq.IsComponentPresent(4) {
			ie.RefFreqCSI_RS = new(ARFCN_ValueNR)
			if err := ie.RefFreqCSI_RS.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. referenceSignalConfig: ref
	{
		if err := ie.ReferenceSignalConfig.Decode(d); err != nil {
			return err
		}
	}

	// 9. absThreshSS-BlocksConsolidation: ref
	{
		if seq.IsComponentPresent(6) {
			ie.AbsThreshSS_BlocksConsolidation = new(ThresholdNR)
			if err := ie.AbsThreshSS_BlocksConsolidation.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. absThreshCSI-RS-Consolidation: ref
	{
		if seq.IsComponentPresent(7) {
			ie.AbsThreshCSI_RS_Consolidation = new(ThresholdNR)
			if err := ie.AbsThreshCSI_RS_Consolidation.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. nrofSS-BlocksToAverage: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(measObjectNRNrofSSBlocksToAverageConstraints)
			if err != nil {
				return err
			}
			ie.NrofSS_BlocksToAverage = &v
		}
	}

	// 12. nrofCSI-RS-ResourcesToAverage: integer
	{
		if seq.IsComponentPresent(9) {
			v, err := d.DecodeInteger(measObjectNRNrofCSIRSResourcesToAverageConstraints)
			if err != nil {
				return err
			}
			ie.NrofCSI_RS_ResourcesToAverage = &v
		}
	}

	// 13. quantityConfigIndex: integer
	{
		v10, err := d.DecodeInteger(measObjectNRQuantityConfigIndexConstraints)
		if err != nil {
			return err
		}
		ie.QuantityConfigIndex = v10
	}

	// 14. offsetMO: ref
	{
		if err := ie.OffsetMO.Decode(d); err != nil {
			return err
		}
	}

	// 15. cellsToRemoveList: ref
	{
		if seq.IsComponentPresent(12) {
			ie.CellsToRemoveList = new(PCI_List)
			if err := ie.CellsToRemoveList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 16. cellsToAddModList: ref
	{
		if seq.IsComponentPresent(13) {
			ie.CellsToAddModList = new(CellsToAddModList)
			if err := ie.CellsToAddModList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 17. excludedCellsToRemoveList: ref
	{
		if seq.IsComponentPresent(14) {
			ie.ExcludedCellsToRemoveList = new(PCI_RangeIndexList)
			if err := ie.ExcludedCellsToRemoveList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 18. excludedCellsToAddModList: sequence-of
	{
		if seq.IsComponentPresent(15) {
			seqOf := d.NewSequenceOfDecoder(measObjectNRExcludedCellsToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ExcludedCellsToAddModList = make([]PCI_RangeElement, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ExcludedCellsToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 19. allowedCellsToRemoveList: ref
	{
		if seq.IsComponentPresent(16) {
			ie.AllowedCellsToRemoveList = new(PCI_RangeIndexList)
			if err := ie.AllowedCellsToRemoveList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 20. allowedCellsToAddModList: sequence-of
	{
		if seq.IsComponentPresent(17) {
			seqOf := d.NewSequenceOfDecoder(measObjectNRAllowedCellsToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AllowedCellsToAddModList = make([]PCI_RangeElement, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AllowedCellsToAddModList[i].Decode(d); err != nil {
					return err
				}
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
				{Name: "freqBandIndicatorNR", Optional: true},
				{Name: "measCycleSCell", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.FreqBandIndicatorNR = new(FreqBandIndicatorNR)
			if err := ie.FreqBandIndicatorNR.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measObjectNRExtMeasCycleSCellConstraints)
			if err != nil {
				return err
			}
			ie.MeasCycleSCell = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "smtc3list-r16", Optional: true},
				{Name: "rmtc-Config-r16", Optional: true},
				{Name: "t312-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Smtc3list_r16 = new(SSB_MTC3List_r16)
			if err := ie.Smtc3list_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Rmtc_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RMTC_Config_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(measObjectNRExtRmtcConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rmtc_Config_r16).Choice = int(index)
			switch index {
			case MeasObjectNR_Ext_Rmtc_Config_r16_Release:
				(*ie.Rmtc_Config_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasObjectNR_Ext_Rmtc_Config_r16_Setup:
				(*ie.Rmtc_Config_r16).Setup = new(RMTC_Config_r16)
				if err := (*ie.Rmtc_Config_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.T312_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *T312_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(measObjectNRExtT312R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.T312_r16).Choice = int(index)
			switch index {
			case MeasObjectNR_Ext_T312_r16_Release:
				(*ie.T312_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MeasObjectNR_Ext_T312_r16_Setup:
				(*ie.T312_r16).Setup = new(T312_r16)
				if err := (*ie.T312_r16).Setup.Decode(dx); err != nil {
					return err
				}
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
				{Name: "associatedMeasGapSSB-r17", Optional: true},
				{Name: "associatedMeasGapCSIRS-r17", Optional: true},
				{Name: "smtc4list-r17", Optional: true},
				{Name: "measCyclePSCell-r17", Optional: true},
				{Name: "cellsToAddModListExt-v1710", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AssociatedMeasGapSSB_r17 = new(MeasGapId_r17)
			if err := ie.AssociatedMeasGapSSB_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AssociatedMeasGapCSIRS_r17 = new(MeasGapId_r17)
			if err := ie.AssociatedMeasGapCSIRS_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Smtc4list_r17 = new(SSB_MTC4List_r17)
			if err := ie.Smtc4list_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measObjectNRExtMeasCyclePSCellR17Constraints)
			if err != nil {
				return err
			}
			ie.MeasCyclePSCell_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.CellsToAddModListExt_v1710 = new(CellsToAddModListExt_v1710)
			if err := ie.CellsToAddModListExt_v1710.Decode(dx); err != nil {
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
				{Name: "associatedMeasGapSSB2-v1720", Optional: true},
				{Name: "associatedMeasGapCSIRS2-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AssociatedMeasGapSSB2_v1720 = new(MeasGapId_r17)
			if err := ie.AssociatedMeasGapSSB2_v1720.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AssociatedMeasGapCSIRS2_v1720 = new(MeasGapId_r17)
			if err := ie.AssociatedMeasGapCSIRS2_v1720.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "measSequence-r18", Optional: true},
				{Name: "cellsToAddModListExt-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MeasSequence_r18 = new(MeasSequence_r18)
			if err := ie.MeasSequence_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CellsToAddModListExt_v1800 = new(CellsToAddModListExt_v1800)
			if err := ie.CellsToAddModListExt_v1800.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "smtc5list-r19", Optional: true},
				{Name: "smtc6list-r19", Optional: true},
				{Name: "smtc7-SSBAdapt-r19", Optional: true},
				{Name: "neighSCellMeasSkipping-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Smtc5list_r19 = new(SSB_MTC5List_r19)
			if err := ie.Smtc5list_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Smtc6list_r19 = new(SSB_MTC6List_r19)
			if err := ie.Smtc6list_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Smtc7_SSBAdapt_r19 = new(SSB_MTC_SSBAdapt_r19)
			if err := ie.Smtc7_SSBAdapt_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measObjectNRExtNeighSCellMeasSkippingR19Constraints)
			if err != nil {
				return err
			}
			ie.NeighSCellMeasSkipping_r19 = &v
		}
	}

	return nil
}
