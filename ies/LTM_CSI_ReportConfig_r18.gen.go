// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-CSI-ReportConfig-r18 (line 8790).

var lTMCSIReportConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-CSI-ReportConfigId-r18"},
		{Name: "ltm-ResourcesForChannelMeasurement-r18"},
		{Name: "ltm-ReportConfigType-r18"},
		{Name: "ltm-ReportContent-r18"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var lTM_CSI_ReportConfig_r18LtmReportConfigTypeR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodic-r18"},
		{Name: "semiPersistentOnPUCCH-r18"},
		{Name: "semiPersistentOnPUSCH-r18"},
		{Name: "aperiodic-r18"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "eventTriggered-r19"},
	},
}

const (
	LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_Periodic_r18              = 0
	LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_SemiPersistentOnPUCCH_r18 = 1
	LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_SemiPersistentOnPUSCH_r18 = 2
	LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_Aperiodic_r18             = 3
)

var lTMCSIReportConfigR18LtmReportConfigTypeR18PeriodicR18PucchCSIResourceListR18Constraints = per.SizeRange(1, common.MaxNrofBWPs)

var lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUCCHR18PucchCSIResourceListR18Constraints = per.SizeRange(1, common.MaxNrofBWPs)

var lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListR18Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListDCI02R18Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListDCI01R18Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListR18Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListDCI02R18Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListDCI01R18Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

type LTM_CSI_ReportConfig_r18 struct {
	Ltm_CSI_ReportConfigId_r18             LTM_CSI_ReportConfigId_r18
	Ltm_ResourcesForChannelMeasurement_r18 LTM_CSI_ResourceConfigId_r18
	Ltm_ReportConfigType_r18               struct {
		Choice       int
		Periodic_r18 *struct {
			ReportSlotConfig_r18       CSI_ReportPeriodicityAndOffset
			Pucch_CSI_ResourceList_r18 []PUCCH_CSI_Resource
		}
		SemiPersistentOnPUCCH_r18 *struct {
			ReportSlotConfig_r18       CSI_ReportPeriodicityAndOffset
			Pucch_CSI_ResourceList_r18 []PUCCH_CSI_Resource
		}
		SemiPersistentOnPUSCH_r18 *struct {
			ReportSlotConfig_r18            CSI_ReportPeriodicityAndOffset
			ReportSlotOffsetList_r18        []int64
			ReportSlotOffsetListDCI_0_2_r18 []int64
			ReportSlotOffsetListDCI_0_1_r18 []int64
			P0alpha_r18                     P0_PUSCH_AlphaSetId
		}
		Aperiodic_r18 *struct {
			ReportSlotOffsetList_r18        []int64
			ReportSlotOffsetListDCI_0_2_r18 []int64
			ReportSlotOffsetListDCI_0_1_r18 []int64
		}
	}
	Ltm_ReportContent_r18         LTM_ReportContent_r18
	Ltm_ReportContent_v1900       *LTM_ReportContent_v1900
	Ltm_EarlyCSI_ReportConfig_r19 *LTM_EarlyCSI_ReportConfig_r19
}

func (ie *LTM_CSI_ReportConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMCSIReportConfigR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ltm_ReportContent_v1900 != nil || ie.Ltm_EarlyCSI_ReportConfig_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. ltm-CSI-ReportConfigId-r18: ref
	{
		if err := ie.Ltm_CSI_ReportConfigId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. ltm-ResourcesForChannelMeasurement-r18: ref
	{
		if err := ie.Ltm_ResourcesForChannelMeasurement_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. ltm-ReportConfigType-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(lTM_CSI_ReportConfig_r18LtmReportConfigTypeR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Ltm_ReportConfigType_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Ltm_ReportConfigType_r18.Choice {
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_Periodic_r18:
			c := (*ie.Ltm_ReportConfigType_r18.Periodic_r18)
			if err := c.ReportSlotConfig_r18.Encode(e); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18PeriodicR18PucchCSIResourceListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Pucch_CSI_ResourceList_r18))); err != nil {
					return err
				}
				for i := range c.Pucch_CSI_ResourceList_r18 {
					if err := c.Pucch_CSI_ResourceList_r18[i].Encode(e); err != nil {
						return err
					}
				}
			}
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_SemiPersistentOnPUCCH_r18:
			c := (*ie.Ltm_ReportConfigType_r18.SemiPersistentOnPUCCH_r18)
			if err := c.ReportSlotConfig_r18.Encode(e); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUCCHR18PucchCSIResourceListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Pucch_CSI_ResourceList_r18))); err != nil {
					return err
				}
				for i := range c.Pucch_CSI_ResourceList_r18 {
					if err := c.Pucch_CSI_ResourceList_r18[i].Encode(e); err != nil {
						return err
					}
				}
			}
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_SemiPersistentOnPUSCH_r18:
			c := (*ie.Ltm_ReportConfigType_r18.SemiPersistentOnPUSCH_r18)
			if err := c.ReportSlotConfig_r18.Encode(e); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList_r18))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetList_r18 {
					if err := e.EncodeInteger(c.ReportSlotOffsetList_r18[i], per.Constrained(0, 128)); err != nil {
						return err
					}
				}
			}
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListDCI02R18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_2_r18))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetListDCI_0_2_r18 {
					if err := e.EncodeInteger(c.ReportSlotOffsetListDCI_0_2_r18[i], per.Constrained(0, 128)); err != nil {
						return err
					}
				}
			}
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListDCI01R18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_1_r18))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetListDCI_0_1_r18 {
					if err := e.EncodeInteger(c.ReportSlotOffsetListDCI_0_1_r18[i], per.Constrained(0, 128)); err != nil {
						return err
					}
				}
			}
			if err := c.P0alpha_r18.Encode(e); err != nil {
				return err
			}
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_Aperiodic_r18:
			c := (*ie.Ltm_ReportConfigType_r18.Aperiodic_r18)
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList_r18))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetList_r18 {
					if err := e.EncodeInteger(c.ReportSlotOffsetList_r18[i], per.Constrained(0, 128)); err != nil {
						return err
					}
				}
			}
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListDCI02R18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_2_r18))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetListDCI_0_2_r18 {
					if err := e.EncodeInteger(c.ReportSlotOffsetListDCI_0_2_r18[i], per.Constrained(0, 128)); err != nil {
						return err
					}
				}
			}
			{
				seqOf := e.NewSequenceOfEncoder(lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListDCI01R18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_1_r18))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetListDCI_0_1_r18 {
					if err := e.EncodeInteger(c.ReportSlotOffsetListDCI_0_1_r18[i], per.Constrained(0, 128)); err != nil {
						return err
					}
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Ltm_ReportConfigType_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 5. ltm-ReportContent-r18: ref
	{
		if err := ie.Ltm_ReportContent_r18.Encode(e); err != nil {
			return err
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
					{Name: "ltm-ReportContent-v1900", Optional: true},
					{Name: "ltm-EarlyCSI-ReportConfig-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_ReportContent_v1900 != nil, ie.Ltm_EarlyCSI_ReportConfig_r19 != nil}); err != nil {
				return err
			}

			if ie.Ltm_ReportContent_v1900 != nil {
				if err := ie.Ltm_ReportContent_v1900.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_EarlyCSI_ReportConfig_r19 != nil {
				if err := ie.Ltm_EarlyCSI_ReportConfig_r19.Encode(ex); err != nil {
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

func (ie *LTM_CSI_ReportConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMCSIReportConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ltm-CSI-ReportConfigId-r18: ref
	{
		if err := ie.Ltm_CSI_ReportConfigId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. ltm-ResourcesForChannelMeasurement-r18: ref
	{
		if err := ie.Ltm_ResourcesForChannelMeasurement_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. ltm-ReportConfigType-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(lTM_CSI_ReportConfig_r18LtmReportConfigTypeR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Ltm_ReportConfigType_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_Periodic_r18:
			ie.Ltm_ReportConfigType_r18.Periodic_r18 = &struct {
				ReportSlotConfig_r18       CSI_ReportPeriodicityAndOffset
				Pucch_CSI_ResourceList_r18 []PUCCH_CSI_Resource
			}{}
			c := (*ie.Ltm_ReportConfigType_r18.Periodic_r18)
			{
				if err := c.ReportSlotConfig_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18PeriodicR18PucchCSIResourceListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Pucch_CSI_ResourceList_r18 = make([]PUCCH_CSI_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.Pucch_CSI_ResourceList_r18[i].Decode(d); err != nil {
						return err
					}
				}
			}
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_SemiPersistentOnPUCCH_r18:
			ie.Ltm_ReportConfigType_r18.SemiPersistentOnPUCCH_r18 = &struct {
				ReportSlotConfig_r18       CSI_ReportPeriodicityAndOffset
				Pucch_CSI_ResourceList_r18 []PUCCH_CSI_Resource
			}{}
			c := (*ie.Ltm_ReportConfigType_r18.SemiPersistentOnPUCCH_r18)
			{
				if err := c.ReportSlotConfig_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUCCHR18PucchCSIResourceListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Pucch_CSI_ResourceList_r18 = make([]PUCCH_CSI_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.Pucch_CSI_ResourceList_r18[i].Decode(d); err != nil {
						return err
					}
				}
			}
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_SemiPersistentOnPUSCH_r18:
			ie.Ltm_ReportConfigType_r18.SemiPersistentOnPUSCH_r18 = &struct {
				ReportSlotConfig_r18            CSI_ReportPeriodicityAndOffset
				ReportSlotOffsetList_r18        []int64
				ReportSlotOffsetListDCI_0_2_r18 []int64
				ReportSlotOffsetListDCI_0_1_r18 []int64
				P0alpha_r18                     P0_PUSCH_AlphaSetId
			}{}
			c := (*ie.Ltm_ReportConfigType_r18.SemiPersistentOnPUSCH_r18)
			{
				if err := c.ReportSlotConfig_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetList_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList_r18[i] = v
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListDCI02R18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_2_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_2_r18[i] = v
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18SemiPersistentOnPUSCHR18ReportSlotOffsetListDCI01R18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_1_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_1_r18[i] = v
				}
			}
			{
				if err := c.P0alpha_r18.Decode(d); err != nil {
					return err
				}
			}
		case LTM_CSI_ReportConfig_r18_Ltm_ReportConfigType_r18_Aperiodic_r18:
			ie.Ltm_ReportConfigType_r18.Aperiodic_r18 = &struct {
				ReportSlotOffsetList_r18        []int64
				ReportSlotOffsetListDCI_0_2_r18 []int64
				ReportSlotOffsetListDCI_0_1_r18 []int64
			}{}
			c := (*ie.Ltm_ReportConfigType_r18.Aperiodic_r18)
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetList_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList_r18[i] = v
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListDCI02R18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_2_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_2_r18[i] = v
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(lTMCSIReportConfigR18LtmReportConfigTypeR18AperiodicR18ReportSlotOffsetListDCI01R18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_1_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_1_r18[i] = v
				}
			}
		}
	}

	// 5. ltm-ReportContent-r18: ref
	{
		if err := ie.Ltm_ReportContent_r18.Decode(d); err != nil {
			return err
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
				{Name: "ltm-ReportContent-v1900", Optional: true},
				{Name: "ltm-EarlyCSI-ReportConfig-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_ReportContent_v1900 = new(LTM_ReportContent_v1900)
			if err := ie.Ltm_ReportContent_v1900.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ltm_EarlyCSI_ReportConfig_r19 = new(LTM_EarlyCSI_ReportConfig_r19)
			if err := ie.Ltm_EarlyCSI_ReportConfig_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
