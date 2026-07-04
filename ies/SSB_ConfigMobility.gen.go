// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-ConfigMobility (line 9501).

var sSBConfigMobilityConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-ToMeasure", Optional: true},
		{Name: "deriveSSB-IndexFromCell"},
		{Name: "ss-RSSI-Measurement", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var sSB_ConfigMobilitySsbToMeasureConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SSB_ConfigMobility_Ssb_ToMeasure_Release = 0
	SSB_ConfigMobility_Ssb_ToMeasure_Setup   = 1
)

var sSBConfigMobilityExtSsbPositionQCLCellsR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SSB_ConfigMobility_Ext_Ssb_PositionQCL_Cells_r17_Release = 0
	SSB_ConfigMobility_Ext_Ssb_PositionQCL_Cells_r17_Setup   = 1
)

var sSBConfigMobilityExtSsbToMeasureAltitudeBasedListR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SSB_ConfigMobility_Ext_Ssb_ToMeasureAltitudeBasedList_r18_Release = 0
	SSB_ConfigMobility_Ext_Ssb_ToMeasureAltitudeBasedList_r18_Setup   = 1
)

type SSB_ConfigMobility struct {
	Ssb_ToMeasure *struct {
		Choice  int
		Release *struct{}
		Setup   *SSB_ToMeasure
	}
	DeriveSSB_IndexFromCell               bool
	Ss_RSSI_Measurement                   *SS_RSSI_Measurement
	Ssb_PositionQCL_Common_r16            *SSB_PositionQCL_Relation_r16
	Ssb_PositionQCL_CellsToAddModList_r16 *SSB_PositionQCL_CellsToAddModList_r16
	Ssb_PositionQCL_CellsToRemoveList_r16 *PCI_List
	DeriveSSB_IndexFromCellInter_r17      *ServCellIndex
	Ssb_PositionQCL_Common_r17            *SSB_PositionQCL_Relation_r17
	Ssb_PositionQCL_Cells_r17             *struct {
		Choice  int
		Release *struct{}
		Setup   *SSB_PositionQCL_CellList_r17
	}
	Cca_CellsToAddModList_r17          *PCI_List
	Cca_CellsToRemoveList_r17          *PCI_List
	Ssb_ToMeasureAltitudeBasedList_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SSB_ToMeasureAltitudeBasedList_r18
	}
}

func (ie *SSB_ConfigMobility) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBConfigMobilityConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ssb_PositionQCL_Common_r16 != nil || ie.Ssb_PositionQCL_CellsToAddModList_r16 != nil || ie.Ssb_PositionQCL_CellsToRemoveList_r16 != nil
	hasExtGroup1 := ie.DeriveSSB_IndexFromCellInter_r17 != nil || ie.Ssb_PositionQCL_Common_r17 != nil || ie.Ssb_PositionQCL_Cells_r17 != nil
	hasExtGroup2 := ie.Cca_CellsToAddModList_r17 != nil || ie.Cca_CellsToRemoveList_r17 != nil
	hasExtGroup3 := ie.Ssb_ToMeasureAltitudeBasedList_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_ToMeasure != nil, ie.Ss_RSSI_Measurement != nil}); err != nil {
		return err
	}

	// 3. ssb-ToMeasure: choice
	{
		if ie.Ssb_ToMeasure != nil {
			choiceEnc := e.NewChoiceEncoder(sSB_ConfigMobilitySsbToMeasureConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_ToMeasure).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ssb_ToMeasure).Choice {
			case SSB_ConfigMobility_Ssb_ToMeasure_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SSB_ConfigMobility_Ssb_ToMeasure_Setup:
				if err := (*ie.Ssb_ToMeasure).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ssb_ToMeasure).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. deriveSSB-IndexFromCell: boolean
	{
		if err := e.EncodeBoolean(ie.DeriveSSB_IndexFromCell); err != nil {
			return err
		}
	}

	// 5. ss-RSSI-Measurement: ref
	{
		if ie.Ss_RSSI_Measurement != nil {
			if err := ie.Ss_RSSI_Measurement.Encode(e); err != nil {
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
					{Name: "ssb-PositionQCL-Common-r16", Optional: true},
					{Name: "ssb-PositionQCL-CellsToAddModList-r16", Optional: true},
					{Name: "ssb-PositionQCL-CellsToRemoveList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ssb_PositionQCL_Common_r16 != nil, ie.Ssb_PositionQCL_CellsToAddModList_r16 != nil, ie.Ssb_PositionQCL_CellsToRemoveList_r16 != nil}); err != nil {
				return err
			}

			if ie.Ssb_PositionQCL_Common_r16 != nil {
				if err := ie.Ssb_PositionQCL_Common_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ssb_PositionQCL_CellsToAddModList_r16 != nil {
				if err := ie.Ssb_PositionQCL_CellsToAddModList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ssb_PositionQCL_CellsToRemoveList_r16 != nil {
				if err := ie.Ssb_PositionQCL_CellsToRemoveList_r16.Encode(ex); err != nil {
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
					{Name: "deriveSSB-IndexFromCellInter-r17", Optional: true},
					{Name: "ssb-PositionQCL-Common-r17", Optional: true},
					{Name: "ssb-PositionQCL-Cells-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DeriveSSB_IndexFromCellInter_r17 != nil, ie.Ssb_PositionQCL_Common_r17 != nil, ie.Ssb_PositionQCL_Cells_r17 != nil}); err != nil {
				return err
			}

			if ie.DeriveSSB_IndexFromCellInter_r17 != nil {
				if err := ie.DeriveSSB_IndexFromCellInter_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ssb_PositionQCL_Common_r17 != nil {
				if err := ie.Ssb_PositionQCL_Common_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ssb_PositionQCL_Cells_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sSBConfigMobilityExtSsbPositionQCLCellsR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_PositionQCL_Cells_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ssb_PositionQCL_Cells_r17).Choice {
				case SSB_ConfigMobility_Ext_Ssb_PositionQCL_Cells_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SSB_ConfigMobility_Ext_Ssb_PositionQCL_Cells_r17_Setup:
					if err := (*ie.Ssb_PositionQCL_Cells_r17).Setup.Encode(ex); err != nil {
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
					{Name: "cca-CellsToAddModList-r17", Optional: true},
					{Name: "cca-CellsToRemoveList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cca_CellsToAddModList_r17 != nil, ie.Cca_CellsToRemoveList_r17 != nil}); err != nil {
				return err
			}

			if ie.Cca_CellsToAddModList_r17 != nil {
				if err := ie.Cca_CellsToAddModList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Cca_CellsToRemoveList_r17 != nil {
				if err := ie.Cca_CellsToRemoveList_r17.Encode(ex); err != nil {
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
					{Name: "ssb-ToMeasureAltitudeBasedList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ssb_ToMeasureAltitudeBasedList_r18 != nil}); err != nil {
				return err
			}

			if ie.Ssb_ToMeasureAltitudeBasedList_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sSBConfigMobilityExtSsbToMeasureAltitudeBasedListR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_ToMeasureAltitudeBasedList_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ssb_ToMeasureAltitudeBasedList_r18).Choice {
				case SSB_ConfigMobility_Ext_Ssb_ToMeasureAltitudeBasedList_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SSB_ConfigMobility_Ext_Ssb_ToMeasureAltitudeBasedList_r18_Setup:
					if err := (*ie.Ssb_ToMeasureAltitudeBasedList_r18).Setup.Encode(ex); err != nil {
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

func (ie *SSB_ConfigMobility) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBConfigMobilityConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ssb-ToMeasure: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Ssb_ToMeasure = &struct {
				Choice  int
				Release *struct{}
				Setup   *SSB_ToMeasure
			}{}
			choiceDec := d.NewChoiceDecoder(sSB_ConfigMobilitySsbToMeasureConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_ToMeasure).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SSB_ConfigMobility_Ssb_ToMeasure_Release:
				(*ie.Ssb_ToMeasure).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SSB_ConfigMobility_Ssb_ToMeasure_Setup:
				(*ie.Ssb_ToMeasure).Setup = new(SSB_ToMeasure)
				if err := (*ie.Ssb_ToMeasure).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. deriveSSB-IndexFromCell: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.DeriveSSB_IndexFromCell = v1
	}

	// 5. ss-RSSI-Measurement: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ss_RSSI_Measurement = new(SS_RSSI_Measurement)
			if err := ie.Ss_RSSI_Measurement.Decode(d); err != nil {
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
				{Name: "ssb-PositionQCL-Common-r16", Optional: true},
				{Name: "ssb-PositionQCL-CellsToAddModList-r16", Optional: true},
				{Name: "ssb-PositionQCL-CellsToRemoveList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
			if err := ie.Ssb_PositionQCL_Common_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ssb_PositionQCL_CellsToAddModList_r16 = new(SSB_PositionQCL_CellsToAddModList_r16)
			if err := ie.Ssb_PositionQCL_CellsToAddModList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ssb_PositionQCL_CellsToRemoveList_r16 = new(PCI_List)
			if err := ie.Ssb_PositionQCL_CellsToRemoveList_r16.Decode(dx); err != nil {
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
				{Name: "deriveSSB-IndexFromCellInter-r17", Optional: true},
				{Name: "ssb-PositionQCL-Common-r17", Optional: true},
				{Name: "ssb-PositionQCL-Cells-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.DeriveSSB_IndexFromCellInter_r17 = new(ServCellIndex)
			if err := ie.DeriveSSB_IndexFromCellInter_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
			if err := ie.Ssb_PositionQCL_Common_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ssb_PositionQCL_Cells_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SSB_PositionQCL_CellList_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(sSBConfigMobilityExtSsbPositionQCLCellsR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_PositionQCL_Cells_r17).Choice = int(index)
			switch index {
			case SSB_ConfigMobility_Ext_Ssb_PositionQCL_Cells_r17_Release:
				(*ie.Ssb_PositionQCL_Cells_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SSB_ConfigMobility_Ext_Ssb_PositionQCL_Cells_r17_Setup:
				(*ie.Ssb_PositionQCL_Cells_r17).Setup = new(SSB_PositionQCL_CellList_r17)
				if err := (*ie.Ssb_PositionQCL_Cells_r17).Setup.Decode(dx); err != nil {
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
				{Name: "cca-CellsToAddModList-r17", Optional: true},
				{Name: "cca-CellsToRemoveList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cca_CellsToAddModList_r17 = new(PCI_List)
			if err := ie.Cca_CellsToAddModList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Cca_CellsToRemoveList_r17 = new(PCI_List)
			if err := ie.Cca_CellsToRemoveList_r17.Decode(dx); err != nil {
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
				{Name: "ssb-ToMeasureAltitudeBasedList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ssb_ToMeasureAltitudeBasedList_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SSB_ToMeasureAltitudeBasedList_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sSBConfigMobilityExtSsbToMeasureAltitudeBasedListR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_ToMeasureAltitudeBasedList_r18).Choice = int(index)
			switch index {
			case SSB_ConfigMobility_Ext_Ssb_ToMeasureAltitudeBasedList_r18_Release:
				(*ie.Ssb_ToMeasureAltitudeBasedList_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SSB_ConfigMobility_Ext_Ssb_ToMeasureAltitudeBasedList_r18_Setup:
				(*ie.Ssb_ToMeasureAltitudeBasedList_r18).Setup = new(SSB_ToMeasureAltitudeBasedList_r18)
				if err := (*ie.Ssb_ToMeasureAltitudeBasedList_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
