// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TCI-UL-State-r17 (line 16035).

var tCIULStateR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tci-UL-StateId-r17"},
		{Name: "servingCellId-r17", Optional: true},
		{Name: "bwp-Id-r17", Optional: true},
		{Name: "referenceSignal-r17"},
		{Name: "additionalPCI-r17", Optional: true},
		{Name: "ul-powerControl-r17", Optional: true},
		{Name: "pathlossReferenceRS-Id-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var tCI_UL_State_r17ReferenceSignalR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index-r17"},
		{Name: "csi-RS-Index-r17"},
		{Name: "srs-r17"},
	},
}

const (
	TCI_UL_State_r17_ReferenceSignal_r17_Ssb_Index_r17    = 0
	TCI_UL_State_r17_ReferenceSignal_r17_Csi_RS_Index_r17 = 1
	TCI_UL_State_r17_ReferenceSignal_r17_Srs_r17          = 2
)

const (
	TCI_UL_State_r17_Ext_Tag_Id_Ptr_r18_N0 = 0
	TCI_UL_State_r17_Ext_Tag_Id_Ptr_r18_N1 = 1
)

var tCIULStateR17ExtTagIdPtrR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TCI_UL_State_r17_Ext_Tag_Id_Ptr_r18_N0, TCI_UL_State_r17_Ext_Tag_Id_Ptr_r18_N1},
}

const (
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB_12 = 0
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB_8  = 1
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB_4  = 2
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB0   = 3
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB4   = 4
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB8   = 5
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB12  = 6
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB16  = 7
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB20  = 8
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB24  = 9
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB28  = 10
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB32  = 11
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB36  = 12
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB40  = 13
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB44  = 14
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB48  = 15
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB52  = 16
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB56  = 17
	TCI_UL_State_r17_Ext_PathlossOffset_r19_DB60  = 18
)

var tCIULStateR17ExtPathlossOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TCI_UL_State_r17_Ext_PathlossOffset_r19_DB_12, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB_8, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB_4, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB0, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB4, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB8, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB12, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB16, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB20, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB24, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB28, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB32, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB36, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB40, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB44, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB48, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB52, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB56, TCI_UL_State_r17_Ext_PathlossOffset_r19_DB60},
}

type TCI_UL_State_r17 struct {
	Tci_UL_StateId_r17  TCI_UL_StateId_r17
	ServingCellId_r17   *ServCellIndex
	Bwp_Id_r17          *BWP_Id
	ReferenceSignal_r17 struct {
		Choice           int
		Ssb_Index_r17    *SSB_Index
		Csi_RS_Index_r17 *NZP_CSI_RS_ResourceId
		Srs_r17          *SRS_ResourceId
	}
	AdditionalPCI_r17          *AdditionalPCIIndex_r17
	Ul_PowerControl_r17        *Uplink_PowerControlId_r17
	PathlossReferenceRS_Id_r17 *PathlossReferenceRS_Id_r17
	Tag_Id_Ptr_r18             *int64
	PathlossOffset_r19         *int64
}

func (ie *TCI_UL_State_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tCIULStateR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Tag_Id_Ptr_r18 != nil
	hasExtGroup1 := ie.PathlossOffset_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServingCellId_r17 != nil, ie.Bwp_Id_r17 != nil, ie.AdditionalPCI_r17 != nil, ie.Ul_PowerControl_r17 != nil, ie.PathlossReferenceRS_Id_r17 != nil}); err != nil {
		return err
	}

	// 3. tci-UL-StateId-r17: ref
	{
		if err := ie.Tci_UL_StateId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. servingCellId-r17: ref
	{
		if ie.ServingCellId_r17 != nil {
			if err := ie.ServingCellId_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. bwp-Id-r17: ref
	{
		if ie.Bwp_Id_r17 != nil {
			if err := ie.Bwp_Id_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. referenceSignal-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(tCI_UL_State_r17ReferenceSignalR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal_r17.Choice {
		case TCI_UL_State_r17_ReferenceSignal_r17_Ssb_Index_r17:
			if err := ie.ReferenceSignal_r17.Ssb_Index_r17.Encode(e); err != nil {
				return err
			}
		case TCI_UL_State_r17_ReferenceSignal_r17_Csi_RS_Index_r17:
			if err := ie.ReferenceSignal_r17.Csi_RS_Index_r17.Encode(e); err != nil {
				return err
			}
		case TCI_UL_State_r17_ReferenceSignal_r17_Srs_r17:
			if err := ie.ReferenceSignal_r17.Srs_r17.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 7. additionalPCI-r17: ref
	{
		if ie.AdditionalPCI_r17 != nil {
			if err := ie.AdditionalPCI_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. ul-powerControl-r17: ref
	{
		if ie.Ul_PowerControl_r17 != nil {
			if err := ie.Ul_PowerControl_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. pathlossReferenceRS-Id-r17: ref
	{
		if ie.PathlossReferenceRS_Id_r17 != nil {
			if err := ie.PathlossReferenceRS_Id_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "tag-Id-ptr-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Tag_Id_Ptr_r18 != nil}); err != nil {
				return err
			}

			if ie.Tag_Id_Ptr_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Tag_Id_Ptr_r18, tCIULStateR17ExtTagIdPtrR18Constraints); err != nil {
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
					{Name: "pathlossOffset-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PathlossOffset_r19 != nil}); err != nil {
				return err
			}

			if ie.PathlossOffset_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffset_r19, tCIULStateR17ExtPathlossOffsetR19Constraints); err != nil {
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

func (ie *TCI_UL_State_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tCIULStateR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tci-UL-StateId-r17: ref
	{
		if err := ie.Tci_UL_StateId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. servingCellId-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ServingCellId_r17 = new(ServCellIndex)
			if err := ie.ServingCellId_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. bwp-Id-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Bwp_Id_r17 = new(BWP_Id)
			if err := ie.Bwp_Id_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. referenceSignal-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(tCI_UL_State_r17ReferenceSignalR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TCI_UL_State_r17_ReferenceSignal_r17_Ssb_Index_r17:
			ie.ReferenceSignal_r17.Ssb_Index_r17 = new(SSB_Index)
			if err := ie.ReferenceSignal_r17.Ssb_Index_r17.Decode(d); err != nil {
				return err
			}
		case TCI_UL_State_r17_ReferenceSignal_r17_Csi_RS_Index_r17:
			ie.ReferenceSignal_r17.Csi_RS_Index_r17 = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal_r17.Csi_RS_Index_r17.Decode(d); err != nil {
				return err
			}
		case TCI_UL_State_r17_ReferenceSignal_r17_Srs_r17:
			ie.ReferenceSignal_r17.Srs_r17 = new(SRS_ResourceId)
			if err := ie.ReferenceSignal_r17.Srs_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. additionalPCI-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
			if err := ie.AdditionalPCI_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. ul-powerControl-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Ul_PowerControl_r17 = new(Uplink_PowerControlId_r17)
			if err := ie.Ul_PowerControl_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. pathlossReferenceRS-Id-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.PathlossReferenceRS_Id_r17 = new(PathlossReferenceRS_Id_r17)
			if err := ie.PathlossReferenceRS_Id_r17.Decode(d); err != nil {
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
				{Name: "tag-Id-ptr-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(tCIULStateR17ExtTagIdPtrR18Constraints)
			if err != nil {
				return err
			}
			ie.Tag_Id_Ptr_r18 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pathlossOffset-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(tCIULStateR17ExtPathlossOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffset_r19 = &v
		}
	}

	return nil
}
