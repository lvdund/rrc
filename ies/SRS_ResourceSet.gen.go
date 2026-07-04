// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-ResourceSet (line 15338).

var sRSResourceSetConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-ResourceSetId"},
		{Name: "srs-ResourceIdList", Optional: true},
		{Name: "resourceType"},
		{Name: "usage"},
		{Name: "alpha", Optional: true},
		{Name: "p0", Optional: true},
		{Name: "pathlossReferenceRS", Optional: true},
		{Name: "srs-PowerControlAdjustmentStates", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var sRSResourceSetSrsResourceIdListConstraints = per.SizeRange(1, common.MaxNrofSRS_ResourcesPerSet)

var sRS_ResourceSetResourceTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "aperiodic"},
		{Name: "semi-persistent"},
		{Name: "periodic"},
	},
}

const (
	SRS_ResourceSet_ResourceType_Aperiodic       = 0
	SRS_ResourceSet_ResourceType_Semi_Persistent = 1
	SRS_ResourceSet_ResourceType_Periodic        = 2
)

const (
	SRS_ResourceSet_Usage_BeamManagement   = 0
	SRS_ResourceSet_Usage_Codebook         = 1
	SRS_ResourceSet_Usage_NonCodebook      = 2
	SRS_ResourceSet_Usage_AntennaSwitching = 3
)

var sRSResourceSetUsageConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Usage_BeamManagement, SRS_ResourceSet_Usage_Codebook, SRS_ResourceSet_Usage_NonCodebook, SRS_ResourceSet_Usage_AntennaSwitching},
}

var sRSResourceSetP0Constraints = per.Constrained(-202, 24)

const (
	SRS_ResourceSet_Srs_PowerControlAdjustmentStates_SameAsFci2         = 0
	SRS_ResourceSet_Srs_PowerControlAdjustmentStates_SeparateClosedLoop = 1
)

var sRSResourceSetSrsPowerControlAdjustmentStatesConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Srs_PowerControlAdjustmentStates_SameAsFci2, SRS_ResourceSet_Srs_PowerControlAdjustmentStates_SeparateClosedLoop},
}

var sRSResourceSetExtPathlossReferenceRSListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_ResourceSet_Ext_PathlossReferenceRSList_r16_Release = 0
	SRS_ResourceSet_Ext_PathlossReferenceRSList_r16_Setup   = 1
)

const (
	SRS_ResourceSet_Ext_UsagePDC_r17_True = 0
)

var sRSResourceSetExtUsagePDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Ext_UsagePDC_r17_True},
}

var sRSResourceSetExtAvailableSlotOffsetListR17Constraints = per.SizeRange(1, 4)

const (
	SRS_ResourceSet_Ext_FollowUnifiedTCI_StateSRS_r17_Enabled = 0
)

var sRSResourceSetExtFollowUnifiedTCIStateSRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Ext_FollowUnifiedTCI_StateSRS_r17_Enabled},
}

const (
	SRS_ResourceSet_Ext_ApplyIndicatedTCI_State_r18_First  = 0
	SRS_ResourceSet_Ext_ApplyIndicatedTCI_State_r18_Second = 1
)

var sRSResourceSetExtApplyIndicatedTCIStateR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Ext_ApplyIndicatedTCI_State_r18_First, SRS_ResourceSet_Ext_ApplyIndicatedTCI_State_r18_Second},
}

const (
	SRS_ResourceSet_Ext_SymbolType_r19_Sbfd     = 0
	SRS_ResourceSet_Ext_SymbolType_r19_Non_Sbfd = 1
)

var sRSResourceSetExtSymbolTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Ext_SymbolType_r19_Sbfd, SRS_ResourceSet_Ext_SymbolType_r19_Non_Sbfd},
}

const (
	SRS_ResourceSet_Ext_Srs_PortGrouping_r19_Enabled = 0
)

var sRSResourceSetExtSrsPortGroupingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Ext_Srs_PortGrouping_r19_Enabled},
}

const (
	SRS_ResourceSet_Ext_FourPortSRS_3Tx_r19_Enabled = 0
)

var sRSResourceSetExtFourPortSRS3TxR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_ResourceSet_Ext_FourPortSRS_3Tx_r19_Enabled},
}

var sRSResourceSetResourceTypeAperiodicConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aperiodicSRS-ResourceTrigger"},
		{Name: "csi-RS", Optional: true},
		{Name: "slotOffset", Optional: true},
	},
}

var sRSResourceSetResourceTypeSemiPersistentConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "associatedCSI-RS", Optional: true},
	},
}

var sRSResourceSetResourceTypePeriodicConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "associatedCSI-RS", Optional: true},
	},
}

type SRS_ResourceSet struct {
	Srs_ResourceSetId  SRS_ResourceSetId
	Srs_ResourceIdList []SRS_ResourceId
	ResourceType       struct {
		Choice    int
		Aperiodic *struct {
			AperiodicSRS_ResourceTrigger int64
			Csi_RS                       *NZP_CSI_RS_ResourceId
			SlotOffset                   *int64
		}
		Semi_Persistent *struct{ AssociatedCSI_RS *NZP_CSI_RS_ResourceId }
		Periodic        *struct{ AssociatedCSI_RS *NZP_CSI_RS_ResourceId }
	}
	Usage                            int64
	Alpha                            *Alpha
	P0                               *int64
	PathlossReferenceRS              *PathlossReferenceRS_Config
	Srs_PowerControlAdjustmentStates *int64
	PathlossReferenceRSList_r16      *struct {
		Choice  int
		Release *struct{}
		Setup   *PathlossReferenceRSList_r16
	}
	UsagePDC_r17                  *int64
	AvailableSlotOffsetList_r17   []AvailableSlotOffset_r17
	FollowUnifiedTCI_StateSRS_r17 *int64
	ApplyIndicatedTCI_State_r18   *int64
	SymbolType_r19                *int64
	AssociatedCSI_RS_Set_r19      *NZP_CSI_RS_ResourceSetId
	Srs_PortGrouping_r19          *int64
	FourPortSRS_3Tx_r19           *int64
}

func (ie *SRS_ResourceSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSResourceSetConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.PathlossReferenceRSList_r16 != nil
	hasExtGroup1 := ie.UsagePDC_r17 != nil || ie.AvailableSlotOffsetList_r17 != nil || ie.FollowUnifiedTCI_StateSRS_r17 != nil
	hasExtGroup2 := ie.ApplyIndicatedTCI_State_r18 != nil
	hasExtGroup3 := ie.SymbolType_r19 != nil || ie.AssociatedCSI_RS_Set_r19 != nil || ie.Srs_PortGrouping_r19 != nil || ie.FourPortSRS_3Tx_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_ResourceIdList != nil, ie.Alpha != nil, ie.P0 != nil, ie.PathlossReferenceRS != nil, ie.Srs_PowerControlAdjustmentStates != nil}); err != nil {
		return err
	}

	// 3. srs-ResourceSetId: ref
	{
		if err := ie.Srs_ResourceSetId.Encode(e); err != nil {
			return err
		}
	}

	// 4. srs-ResourceIdList: sequence-of
	{
		if ie.Srs_ResourceIdList != nil {
			seqOf := e.NewSequenceOfEncoder(sRSResourceSetSrsResourceIdListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_ResourceIdList))); err != nil {
				return err
			}
			for i := range ie.Srs_ResourceIdList {
				if err := ie.Srs_ResourceIdList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. resourceType: choice
	{
		choiceEnc := e.NewChoiceEncoder(sRS_ResourceSetResourceTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ResourceType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ResourceType.Choice {
		case SRS_ResourceSet_ResourceType_Aperiodic:
			c := (*ie.ResourceType.Aperiodic)
			sRSResourceSetResourceTypeAperiodicSeq := e.NewSequenceEncoder(sRSResourceSetResourceTypeAperiodicConstraints)
			if err := sRSResourceSetResourceTypeAperiodicSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sRSResourceSetResourceTypeAperiodicSeq.EncodePreamble([]bool{c.Csi_RS != nil, c.SlotOffset != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.AperiodicSRS_ResourceTrigger, per.Constrained(1, common.MaxNrofSRS_TriggerStates_1)); err != nil {
				return err
			}
			if c.Csi_RS != nil {
				if err := c.Csi_RS.Encode(e); err != nil {
					return err
				}
			}
			if c.SlotOffset != nil {
				if err := e.EncodeInteger((*c.SlotOffset), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		case SRS_ResourceSet_ResourceType_Semi_Persistent:
			c := (*ie.ResourceType.Semi_Persistent)
			sRSResourceSetResourceTypeSemiPersistentSeq := e.NewSequenceEncoder(sRSResourceSetResourceTypeSemiPersistentConstraints)
			if err := sRSResourceSetResourceTypeSemiPersistentSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sRSResourceSetResourceTypeSemiPersistentSeq.EncodePreamble([]bool{c.AssociatedCSI_RS != nil}); err != nil {
				return err
			}
			if c.AssociatedCSI_RS != nil {
				if err := c.AssociatedCSI_RS.Encode(e); err != nil {
					return err
				}
			}
		case SRS_ResourceSet_ResourceType_Periodic:
			c := (*ie.ResourceType.Periodic)
			sRSResourceSetResourceTypePeriodicSeq := e.NewSequenceEncoder(sRSResourceSetResourceTypePeriodicConstraints)
			if err := sRSResourceSetResourceTypePeriodicSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sRSResourceSetResourceTypePeriodicSeq.EncodePreamble([]bool{c.AssociatedCSI_RS != nil}); err != nil {
				return err
			}
			if c.AssociatedCSI_RS != nil {
				if err := c.AssociatedCSI_RS.Encode(e); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ResourceType.Choice), Constraint: "index out of range"}
		}
	}

	// 6. usage: enumerated
	{
		if err := e.EncodeEnumerated(ie.Usage, sRSResourceSetUsageConstraints); err != nil {
			return err
		}
	}

	// 7. alpha: ref
	{
		if ie.Alpha != nil {
			if err := ie.Alpha.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. p0: integer
	{
		if ie.P0 != nil {
			if err := e.EncodeInteger(*ie.P0, sRSResourceSetP0Constraints); err != nil {
				return err
			}
		}
	}

	// 9. pathlossReferenceRS: ref
	{
		if ie.PathlossReferenceRS != nil {
			if err := ie.PathlossReferenceRS.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. srs-PowerControlAdjustmentStates: enumerated
	{
		if ie.Srs_PowerControlAdjustmentStates != nil {
			if err := e.EncodeEnumerated(*ie.Srs_PowerControlAdjustmentStates, sRSResourceSetSrsPowerControlAdjustmentStatesConstraints); err != nil {
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
					{Name: "pathlossReferenceRSList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PathlossReferenceRSList_r16 != nil}); err != nil {
				return err
			}

			if ie.PathlossReferenceRSList_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(sRSResourceSetExtPathlossReferenceRSListR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.PathlossReferenceRSList_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.PathlossReferenceRSList_r16).Choice {
				case SRS_ResourceSet_Ext_PathlossReferenceRSList_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SRS_ResourceSet_Ext_PathlossReferenceRSList_r16_Setup:
					if err := (*ie.PathlossReferenceRSList_r16).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "usagePDC-r17", Optional: true},
					{Name: "availableSlotOffsetList-r17", Optional: true},
					{Name: "followUnifiedTCI-StateSRS-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.UsagePDC_r17 != nil, ie.AvailableSlotOffsetList_r17 != nil, ie.FollowUnifiedTCI_StateSRS_r17 != nil}); err != nil {
				return err
			}

			if ie.UsagePDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UsagePDC_r17, sRSResourceSetExtUsagePDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.AvailableSlotOffsetList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSResourceSetExtAvailableSlotOffsetListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AvailableSlotOffsetList_r17))); err != nil {
					return err
				}
				for i := range ie.AvailableSlotOffsetList_r17 {
					if err := ie.AvailableSlotOffsetList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FollowUnifiedTCI_StateSRS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.FollowUnifiedTCI_StateSRS_r17, sRSResourceSetExtFollowUnifiedTCIStateSRSR17Constraints); err != nil {
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
					{Name: "applyIndicatedTCI-State-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ApplyIndicatedTCI_State_r18 != nil}); err != nil {
				return err
			}

			if ie.ApplyIndicatedTCI_State_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ApplyIndicatedTCI_State_r18, sRSResourceSetExtApplyIndicatedTCIStateR18Constraints); err != nil {
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
					{Name: "symbolType-r19", Optional: true},
					{Name: "associatedCSI-RS-Set-r19", Optional: true},
					{Name: "srs-PortGrouping-r19", Optional: true},
					{Name: "fourPortSRS-3Tx-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SymbolType_r19 != nil, ie.AssociatedCSI_RS_Set_r19 != nil, ie.Srs_PortGrouping_r19 != nil, ie.FourPortSRS_3Tx_r19 != nil}); err != nil {
				return err
			}

			if ie.SymbolType_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.SymbolType_r19, sRSResourceSetExtSymbolTypeR19Constraints); err != nil {
					return err
				}
			}

			if ie.AssociatedCSI_RS_Set_r19 != nil {
				if err := ie.AssociatedCSI_RS_Set_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Srs_PortGrouping_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_PortGrouping_r19, sRSResourceSetExtSrsPortGroupingR19Constraints); err != nil {
					return err
				}
			}

			if ie.FourPortSRS_3Tx_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.FourPortSRS_3Tx_r19, sRSResourceSetExtFourPortSRS3TxR19Constraints); err != nil {
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

func (ie *SRS_ResourceSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-ResourceSetId: ref
	{
		if err := ie.Srs_ResourceSetId.Decode(d); err != nil {
			return err
		}
	}

	// 4. srs-ResourceIdList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sRSResourceSetSrsResourceIdListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_ResourceIdList = make([]SRS_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_ResourceIdList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. resourceType: choice
	{
		choiceDec := d.NewChoiceDecoder(sRS_ResourceSetResourceTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ResourceType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SRS_ResourceSet_ResourceType_Aperiodic:
			ie.ResourceType.Aperiodic = &struct {
				AperiodicSRS_ResourceTrigger int64
				Csi_RS                       *NZP_CSI_RS_ResourceId
				SlotOffset                   *int64
			}{}
			c := (*ie.ResourceType.Aperiodic)
			sRSResourceSetResourceTypeAperiodicSeq := d.NewSequenceDecoder(sRSResourceSetResourceTypeAperiodicConstraints)
			if err := sRSResourceSetResourceTypeAperiodicSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sRSResourceSetResourceTypeAperiodicSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSRS_TriggerStates_1))
				if err != nil {
					return err
				}
				c.AperiodicSRS_ResourceTrigger = v
			}
			if sRSResourceSetResourceTypeAperiodicSeq.IsComponentPresent(1) {
				c.Csi_RS = new(NZP_CSI_RS_ResourceId)
				if err := (*c.Csi_RS).Decode(d); err != nil {
					return err
				}
			}
			if sRSResourceSetResourceTypeAperiodicSeq.IsComponentPresent(2) {
				c.SlotOffset = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.SlotOffset) = v
			}
		case SRS_ResourceSet_ResourceType_Semi_Persistent:
			ie.ResourceType.Semi_Persistent = &struct{ AssociatedCSI_RS *NZP_CSI_RS_ResourceId }{}
			c := (*ie.ResourceType.Semi_Persistent)
			sRSResourceSetResourceTypeSemiPersistentSeq := d.NewSequenceDecoder(sRSResourceSetResourceTypeSemiPersistentConstraints)
			if err := sRSResourceSetResourceTypeSemiPersistentSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sRSResourceSetResourceTypeSemiPersistentSeq.DecodePreamble(); err != nil {
				return err
			}
			if sRSResourceSetResourceTypeSemiPersistentSeq.IsComponentPresent(0) {
				c.AssociatedCSI_RS = new(NZP_CSI_RS_ResourceId)
				if err := (*c.AssociatedCSI_RS).Decode(d); err != nil {
					return err
				}
			}
		case SRS_ResourceSet_ResourceType_Periodic:
			ie.ResourceType.Periodic = &struct{ AssociatedCSI_RS *NZP_CSI_RS_ResourceId }{}
			c := (*ie.ResourceType.Periodic)
			sRSResourceSetResourceTypePeriodicSeq := d.NewSequenceDecoder(sRSResourceSetResourceTypePeriodicConstraints)
			if err := sRSResourceSetResourceTypePeriodicSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sRSResourceSetResourceTypePeriodicSeq.DecodePreamble(); err != nil {
				return err
			}
			if sRSResourceSetResourceTypePeriodicSeq.IsComponentPresent(0) {
				c.AssociatedCSI_RS = new(NZP_CSI_RS_ResourceId)
				if err := (*c.AssociatedCSI_RS).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. usage: enumerated
	{
		v3, err := d.DecodeEnumerated(sRSResourceSetUsageConstraints)
		if err != nil {
			return err
		}
		ie.Usage = v3
	}

	// 7. alpha: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Alpha = new(Alpha)
			if err := ie.Alpha.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. p0: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(sRSResourceSetP0Constraints)
			if err != nil {
				return err
			}
			ie.P0 = &v
		}
	}

	// 9. pathlossReferenceRS: ref
	{
		if seq.IsComponentPresent(6) {
			ie.PathlossReferenceRS = new(PathlossReferenceRS_Config)
			if err := ie.PathlossReferenceRS.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. srs-PowerControlAdjustmentStates: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sRSResourceSetSrsPowerControlAdjustmentStatesConstraints)
			if err != nil {
				return err
			}
			ie.Srs_PowerControlAdjustmentStates = &idx
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
				{Name: "pathlossReferenceRSList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PathlossReferenceRSList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PathlossReferenceRSList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(sRSResourceSetExtPathlossReferenceRSListR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PathlossReferenceRSList_r16).Choice = int(index)
			switch index {
			case SRS_ResourceSet_Ext_PathlossReferenceRSList_r16_Release:
				(*ie.PathlossReferenceRSList_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SRS_ResourceSet_Ext_PathlossReferenceRSList_r16_Setup:
				(*ie.PathlossReferenceRSList_r16).Setup = new(PathlossReferenceRSList_r16)
				if err := (*ie.PathlossReferenceRSList_r16).Setup.Decode(dx); err != nil {
					return err
				}
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
				{Name: "usagePDC-r17", Optional: true},
				{Name: "availableSlotOffsetList-r17", Optional: true},
				{Name: "followUnifiedTCI-StateSRS-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sRSResourceSetExtUsagePDCR17Constraints)
			if err != nil {
				return err
			}
			ie.UsagePDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(sRSResourceSetExtAvailableSlotOffsetListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AvailableSlotOffsetList_r17 = make([]AvailableSlotOffset_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AvailableSlotOffsetList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(sRSResourceSetExtFollowUnifiedTCIStateSRSR17Constraints)
			if err != nil {
				return err
			}
			ie.FollowUnifiedTCI_StateSRS_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "applyIndicatedTCI-State-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sRSResourceSetExtApplyIndicatedTCIStateR18Constraints)
			if err != nil {
				return err
			}
			ie.ApplyIndicatedTCI_State_r18 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "symbolType-r19", Optional: true},
				{Name: "associatedCSI-RS-Set-r19", Optional: true},
				{Name: "srs-PortGrouping-r19", Optional: true},
				{Name: "fourPortSRS-3Tx-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sRSResourceSetExtSymbolTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.SymbolType_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AssociatedCSI_RS_Set_r19 = new(NZP_CSI_RS_ResourceSetId)
			if err := ie.AssociatedCSI_RS_Set_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(sRSResourceSetExtSrsPortGroupingR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_PortGrouping_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(sRSResourceSetExtFourPortSRS3TxR19Constraints)
			if err != nil {
				return err
			}
			ie.FourPortSRS_3Tx_r19 = &v
		}
	}

	return nil
}
