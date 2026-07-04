// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ApplicabilitySetConfigCSI-r19 (line 26553).

var applicabilitySetConfigCSIR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "applicabilitySetConfigId-r19"},
		{Name: "resourcesForChannelMeasurement-r19", Optional: true},
		{Name: "resourcesForChannelPrediction-r19", Optional: true},
		{Name: "associatedIdForChannelMeasurement-r19", Optional: true},
		{Name: "associatedIdForChannelPrediction-r19", Optional: true},
		{Name: "reportQuantity-r19", Optional: true},
		{Name: "reportConfigType-r19", Optional: true},
		{Name: "nrofReportedPredictedRS-r19", Optional: true},
		{Name: "nrofTimeInstance-r19", Optional: true},
		{Name: "timeGap-r19", Optional: true},
	},
}

var applicabilitySetConfigCSI_r19ReportQuantityR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "p-CRI-r19"},
		{Name: "p-SSB-Index-r19"},
		{Name: "p-CRI-RSRP-r19"},
		{Name: "p-SSB-Index-RSRP-r19"},
	},
}

const (
	ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_CRI_r19            = 0
	ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_SSB_Index_r19      = 1
	ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_CRI_RSRP_r19       = 2
	ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_SSB_Index_RSRP_r19 = 3
)

var applicabilitySetConfigCSI_r19ReportConfigTypeR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodic"},
		{Name: "semiPersistentOnPUCCH"},
		{Name: "semiPersistentOnPUSCH"},
		{Name: "aperiodic"},
	},
}

const (
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_Periodic              = 0
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUCCH = 1
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH = 2
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_Aperiodic             = 3
)

const (
	ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N1 = 0
	ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N2 = 1
	ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N3 = 2
	ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N4 = 3
)

var applicabilitySetConfigCSIR19NrofReportedPredictedRSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N1, ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N2, ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N3, ApplicabilitySetConfigCSI_r19_NrofReportedPredictedRS_r19_N4},
}

const (
	ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N1 = 0
	ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N2 = 1
	ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N4 = 2
	ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N8 = 3
)

var applicabilitySetConfigCSIR19NrofTimeInstanceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N1, ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N2, ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N4, ApplicabilitySetConfigCSI_r19_NrofTimeInstance_r19_N8},
}

const (
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms10   = 0
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms20   = 1
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms40   = 2
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms80   = 3
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms160  = 4
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Spare3 = 5
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Spare2 = 6
	ApplicabilitySetConfigCSI_r19_TimeGap_r19_Spare1 = 7
)

var applicabilitySetConfigCSIR19TimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms10, ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms20, ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms40, ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms80, ApplicabilitySetConfigCSI_r19_TimeGap_r19_Ms160, ApplicabilitySetConfigCSI_r19_TimeGap_r19_Spare3, ApplicabilitySetConfigCSI_r19_TimeGap_r19_Spare2, ApplicabilitySetConfigCSI_r19_TimeGap_r19_Spare1},
}

var applicabilitySetConfigCSIR19ReportConfigTypeR19PeriodicPucchCSIResourceListR19Constraints = per.SizeRange(1, common.MaxNrofBWPs)

var applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUCCHPucchCSIResourceListR19Constraints = per.SizeRange(1, common.MaxNrofBWPs)

const (
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl5   = 0
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl10  = 1
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl20  = 2
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl40  = 3
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl80  = 4
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl160 = 5
	ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl320 = 6
)

var applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUSCHReportSlotConfigR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl5, ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl10, ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl20, ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl40, ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl80, ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl160, ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH_ReportSlotConfig_r19_Sl320},
}

var applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUSCHReportSlotOffsetListR19Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations)

var applicabilitySetConfigCSIR19ReportConfigTypeR19AperiodicReportSlotOffsetListR19Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations)

type ApplicabilitySetConfigCSI_r19 struct {
	ApplicabilitySetConfigId_r19          ApplicabilitySetConfigId_r19
	ResourcesForChannelMeasurement_r19    *CSI_ResourceConfigId
	ResourcesForChannelPrediction_r19     *CSI_ResourceConfigId
	AssociatedIdForChannelMeasurement_r19 *AssociatedId_r19
	AssociatedIdForChannelPrediction_r19  *AssociatedId_r19
	ReportQuantity_r19                    *struct {
		Choice               int
		P_CRI_r19            *struct{}
		P_SSB_Index_r19      *struct{}
		P_CRI_RSRP_r19       *struct{}
		P_SSB_Index_RSRP_r19 *struct{}
	}
	ReportConfigType_r19 *struct {
		Choice   int
		Periodic *struct {
			ReportSlotConfig_r19       CSI_ReportPeriodicityAndOffset
			Pucch_CSI_ResourceList_r19 []PUCCH_CSI_Resource
		}
		SemiPersistentOnPUCCH *struct {
			ReportSlotConfig_r19       CSI_ReportPeriodicityAndOffset
			Pucch_CSI_ResourceList_r19 []PUCCH_CSI_Resource
		}
		SemiPersistentOnPUSCH *struct {
			ReportSlotConfig_r19     int64
			ReportSlotOffsetList_r19 []int64
			P0alpha_r19              P0_PUSCH_AlphaSetId
		}
		Aperiodic *struct{ ReportSlotOffsetList_r19 []int64 }
	}
	NrofReportedPredictedRS_r19 *int64
	NrofTimeInstance_r19        *int64
	TimeGap_r19                 *int64
}

func (ie *ApplicabilitySetConfigCSI_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(applicabilitySetConfigCSIR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResourcesForChannelMeasurement_r19 != nil, ie.ResourcesForChannelPrediction_r19 != nil, ie.AssociatedIdForChannelMeasurement_r19 != nil, ie.AssociatedIdForChannelPrediction_r19 != nil, ie.ReportQuantity_r19 != nil, ie.ReportConfigType_r19 != nil, ie.NrofReportedPredictedRS_r19 != nil, ie.NrofTimeInstance_r19 != nil, ie.TimeGap_r19 != nil}); err != nil {
		return err
	}

	// 3. applicabilitySetConfigId-r19: ref
	{
		if err := ie.ApplicabilitySetConfigId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. resourcesForChannelMeasurement-r19: ref
	{
		if ie.ResourcesForChannelMeasurement_r19 != nil {
			if err := ie.ResourcesForChannelMeasurement_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. resourcesForChannelPrediction-r19: ref
	{
		if ie.ResourcesForChannelPrediction_r19 != nil {
			if err := ie.ResourcesForChannelPrediction_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. associatedIdForChannelMeasurement-r19: ref
	{
		if ie.AssociatedIdForChannelMeasurement_r19 != nil {
			if err := ie.AssociatedIdForChannelMeasurement_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. associatedIdForChannelPrediction-r19: ref
	{
		if ie.AssociatedIdForChannelPrediction_r19 != nil {
			if err := ie.AssociatedIdForChannelPrediction_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. reportQuantity-r19: choice
	{
		if ie.ReportQuantity_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(applicabilitySetConfigCSI_r19ReportQuantityR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ReportQuantity_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ReportQuantity_r19).Choice {
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_CRI_r19:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_SSB_Index_r19:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_CRI_RSRP_r19:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_SSB_Index_RSRP_r19:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ReportQuantity_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. reportConfigType-r19: choice
	{
		if ie.ReportConfigType_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(applicabilitySetConfigCSI_r19ReportConfigTypeR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ReportConfigType_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ReportConfigType_r19).Choice {
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_Periodic:
				c := (*(*ie.ReportConfigType_r19).Periodic)
				if err := c.ReportSlotConfig_r19.Encode(e); err != nil {
					return err
				}
				{
					seqOf := e.NewSequenceOfEncoder(applicabilitySetConfigCSIR19ReportConfigTypeR19PeriodicPucchCSIResourceListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Pucch_CSI_ResourceList_r19))); err != nil {
						return err
					}
					for i := range c.Pucch_CSI_ResourceList_r19 {
						if err := c.Pucch_CSI_ResourceList_r19[i].Encode(e); err != nil {
							return err
						}
					}
				}
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUCCH:
				c := (*(*ie.ReportConfigType_r19).SemiPersistentOnPUCCH)
				if err := c.ReportSlotConfig_r19.Encode(e); err != nil {
					return err
				}
				{
					seqOf := e.NewSequenceOfEncoder(applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUCCHPucchCSIResourceListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Pucch_CSI_ResourceList_r19))); err != nil {
						return err
					}
					for i := range c.Pucch_CSI_ResourceList_r19 {
						if err := c.Pucch_CSI_ResourceList_r19[i].Encode(e); err != nil {
							return err
						}
					}
				}
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH:
				c := (*(*ie.ReportConfigType_r19).SemiPersistentOnPUSCH)
				if err := e.EncodeEnumerated(c.ReportSlotConfig_r19, applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUSCHReportSlotConfigR19Constraints); err != nil {
					return err
				}
				{
					seqOf := e.NewSequenceOfEncoder(applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUSCHReportSlotOffsetListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList_r19))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetList_r19 {
						if err := e.EncodeInteger(c.ReportSlotOffsetList_r19[i], per.Constrained(0, 32)); err != nil {
							return err
						}
					}
				}
				if err := c.P0alpha_r19.Encode(e); err != nil {
					return err
				}
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_Aperiodic:
				c := (*(*ie.ReportConfigType_r19).Aperiodic)
				{
					seqOf := e.NewSequenceOfEncoder(applicabilitySetConfigCSIR19ReportConfigTypeR19AperiodicReportSlotOffsetListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList_r19))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetList_r19 {
						if err := e.EncodeInteger(c.ReportSlotOffsetList_r19[i], per.Constrained(0, 32)); err != nil {
							return err
						}
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ReportConfigType_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. nrofReportedPredictedRS-r19: enumerated
	{
		if ie.NrofReportedPredictedRS_r19 != nil {
			if err := e.EncodeEnumerated(*ie.NrofReportedPredictedRS_r19, applicabilitySetConfigCSIR19NrofReportedPredictedRSR19Constraints); err != nil {
				return err
			}
		}
	}

	// 11. nrofTimeInstance-r19: enumerated
	{
		if ie.NrofTimeInstance_r19 != nil {
			if err := e.EncodeEnumerated(*ie.NrofTimeInstance_r19, applicabilitySetConfigCSIR19NrofTimeInstanceR19Constraints); err != nil {
				return err
			}
		}
	}

	// 12. timeGap-r19: enumerated
	{
		if ie.TimeGap_r19 != nil {
			if err := e.EncodeEnumerated(*ie.TimeGap_r19, applicabilitySetConfigCSIR19TimeGapR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ApplicabilitySetConfigCSI_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(applicabilitySetConfigCSIR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. applicabilitySetConfigId-r19: ref
	{
		if err := ie.ApplicabilitySetConfigId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. resourcesForChannelMeasurement-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ResourcesForChannelMeasurement_r19 = new(CSI_ResourceConfigId)
			if err := ie.ResourcesForChannelMeasurement_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. resourcesForChannelPrediction-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ResourcesForChannelPrediction_r19 = new(CSI_ResourceConfigId)
			if err := ie.ResourcesForChannelPrediction_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. associatedIdForChannelMeasurement-r19: ref
	{
		if seq.IsComponentPresent(3) {
			ie.AssociatedIdForChannelMeasurement_r19 = new(AssociatedId_r19)
			if err := ie.AssociatedIdForChannelMeasurement_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. associatedIdForChannelPrediction-r19: ref
	{
		if seq.IsComponentPresent(4) {
			ie.AssociatedIdForChannelPrediction_r19 = new(AssociatedId_r19)
			if err := ie.AssociatedIdForChannelPrediction_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. reportQuantity-r19: choice
	{
		if seq.IsComponentPresent(5) {
			ie.ReportQuantity_r19 = &struct {
				Choice               int
				P_CRI_r19            *struct{}
				P_SSB_Index_r19      *struct{}
				P_CRI_RSRP_r19       *struct{}
				P_SSB_Index_RSRP_r19 *struct{}
			}{}
			choiceDec := d.NewChoiceDecoder(applicabilitySetConfigCSI_r19ReportQuantityR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ReportQuantity_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_CRI_r19:
				(*ie.ReportQuantity_r19).P_CRI_r19 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_SSB_Index_r19:
				(*ie.ReportQuantity_r19).P_SSB_Index_r19 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_CRI_RSRP_r19:
				(*ie.ReportQuantity_r19).P_CRI_RSRP_r19 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ApplicabilitySetConfigCSI_r19_ReportQuantity_r19_P_SSB_Index_RSRP_r19:
				(*ie.ReportQuantity_r19).P_SSB_Index_RSRP_r19 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			}
		}
	}

	// 9. reportConfigType-r19: choice
	{
		if seq.IsComponentPresent(6) {
			ie.ReportConfigType_r19 = &struct {
				Choice   int
				Periodic *struct {
					ReportSlotConfig_r19       CSI_ReportPeriodicityAndOffset
					Pucch_CSI_ResourceList_r19 []PUCCH_CSI_Resource
				}
				SemiPersistentOnPUCCH *struct {
					ReportSlotConfig_r19       CSI_ReportPeriodicityAndOffset
					Pucch_CSI_ResourceList_r19 []PUCCH_CSI_Resource
				}
				SemiPersistentOnPUSCH *struct {
					ReportSlotConfig_r19     int64
					ReportSlotOffsetList_r19 []int64
					P0alpha_r19              P0_PUSCH_AlphaSetId
				}
				Aperiodic *struct{ ReportSlotOffsetList_r19 []int64 }
			}{}
			choiceDec := d.NewChoiceDecoder(applicabilitySetConfigCSI_r19ReportConfigTypeR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ReportConfigType_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_Periodic:
				(*ie.ReportConfigType_r19).Periodic = &struct {
					ReportSlotConfig_r19       CSI_ReportPeriodicityAndOffset
					Pucch_CSI_ResourceList_r19 []PUCCH_CSI_Resource
				}{}
				c := (*(*ie.ReportConfigType_r19).Periodic)
				{
					if err := c.ReportSlotConfig_r19.Decode(d); err != nil {
						return err
					}
				}
				{
					seqOf := d.NewSequenceOfDecoder(applicabilitySetConfigCSIR19ReportConfigTypeR19PeriodicPucchCSIResourceListR19Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Pucch_CSI_ResourceList_r19 = make([]PUCCH_CSI_Resource, n)
					for i := int64(0); i < n; i++ {
						if err := c.Pucch_CSI_ResourceList_r19[i].Decode(d); err != nil {
							return err
						}
					}
				}
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUCCH:
				(*ie.ReportConfigType_r19).SemiPersistentOnPUCCH = &struct {
					ReportSlotConfig_r19       CSI_ReportPeriodicityAndOffset
					Pucch_CSI_ResourceList_r19 []PUCCH_CSI_Resource
				}{}
				c := (*(*ie.ReportConfigType_r19).SemiPersistentOnPUCCH)
				{
					if err := c.ReportSlotConfig_r19.Decode(d); err != nil {
						return err
					}
				}
				{
					seqOf := d.NewSequenceOfDecoder(applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUCCHPucchCSIResourceListR19Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Pucch_CSI_ResourceList_r19 = make([]PUCCH_CSI_Resource, n)
					for i := int64(0); i < n; i++ {
						if err := c.Pucch_CSI_ResourceList_r19[i].Decode(d); err != nil {
							return err
						}
					}
				}
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_SemiPersistentOnPUSCH:
				(*ie.ReportConfigType_r19).SemiPersistentOnPUSCH = &struct {
					ReportSlotConfig_r19     int64
					ReportSlotOffsetList_r19 []int64
					P0alpha_r19              P0_PUSCH_AlphaSetId
				}{}
				c := (*(*ie.ReportConfigType_r19).SemiPersistentOnPUSCH)
				{
					v, err := d.DecodeEnumerated(applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUSCHReportSlotConfigR19Constraints)
					if err != nil {
						return err
					}
					c.ReportSlotConfig_r19 = v
				}
				{
					seqOf := d.NewSequenceOfDecoder(applicabilitySetConfigCSIR19ReportConfigTypeR19SemiPersistentOnPUSCHReportSlotOffsetListR19Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList_r19 = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := d.DecodeInteger(per.Constrained(0, 32))
						if err != nil {
							return err
						}
						c.ReportSlotOffsetList_r19[i] = v
					}
				}
				{
					if err := c.P0alpha_r19.Decode(d); err != nil {
						return err
					}
				}
			case ApplicabilitySetConfigCSI_r19_ReportConfigType_r19_Aperiodic:
				(*ie.ReportConfigType_r19).Aperiodic = &struct{ ReportSlotOffsetList_r19 []int64 }{}
				c := (*(*ie.ReportConfigType_r19).Aperiodic)
				{
					seqOf := d.NewSequenceOfDecoder(applicabilitySetConfigCSIR19ReportConfigTypeR19AperiodicReportSlotOffsetListR19Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList_r19 = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := d.DecodeInteger(per.Constrained(0, 32))
						if err != nil {
							return err
						}
						c.ReportSlotOffsetList_r19[i] = v
					}
				}
			}
		}
	}

	// 10. nrofReportedPredictedRS-r19: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(applicabilitySetConfigCSIR19NrofReportedPredictedRSR19Constraints)
			if err != nil {
				return err
			}
			ie.NrofReportedPredictedRS_r19 = &idx
		}
	}

	// 11. nrofTimeInstance-r19: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(applicabilitySetConfigCSIR19NrofTimeInstanceR19Constraints)
			if err != nil {
				return err
			}
			ie.NrofTimeInstance_r19 = &idx
		}
	}

	// 12. timeGap-r19: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(applicabilitySetConfigCSIR19TimeGapR19Constraints)
			if err != nil {
				return err
			}
			ie.TimeGap_r19 = &idx
		}
	}

	return nil
}
