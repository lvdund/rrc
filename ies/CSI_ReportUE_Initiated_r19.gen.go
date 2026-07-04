// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ReportUE-Initiated-r19 (line 7359).

var cSIReportUEInitiatedR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventTypeUE-Initiated-r19"},
		{Name: "reportTransmissionMode-r19"},
		{Name: "nrofReportedRS-UE-Initiated-r19"},
		{Name: "tci-ServCellIndex-r19", Optional: true},
		{Name: "currentBeamReport-r19", Optional: true},
		{Name: "eventCountWindow-r19", Optional: true},
		{Name: "ue-InitiatedResourceConfig-r19"},
	},
}

var cSI_ReportUE_Initiated_r19EventTypeUEInitiatedR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "event1-r19"},
		{Name: "event2-r19"},
		{Name: "event7-r19"},
	},
}

const (
	CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event1_r19 = 0
	CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event2_r19 = 1
	CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event7_r19 = 2
)

var cSI_ReportUE_Initiated_r19ReportTransmissionModeR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "modeA-r19"},
		{Name: "modeB-r19"},
	},
}

const (
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeA_r19 = 0
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19 = 1
)

const (
	CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N1 = 0
	CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N2 = 1
	CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N3 = 2
	CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N4 = 3
)

var cSIReportUEInitiatedR19NrofReportedRSUEInitiatedR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N1, CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N2, CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N3, CSI_ReportUE_Initiated_r19_NrofReportedRS_UE_Initiated_r19_N4},
}

const (
	CSI_ReportUE_Initiated_r19_CurrentBeamReport_r19_Enabled = 0
)

var cSIReportUEInitiatedR19CurrentBeamReportR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportUE_Initiated_r19_CurrentBeamReport_r19_Enabled},
}

var cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event1R19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "event1Threshold-r19"},
	},
}

var cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventThreshold-r19"},
		{Name: "conditionFulfillmentIndicator-r19", Optional: true},
	},
}

const (
	CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event2_r19_ConditionFulfillmentIndicator_r19_Enabled = 0
)

var cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19ConditionFulfillmentIndicatorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event2_r19_ConditionFulfillmentIndicator_r19_Enabled},
}

var cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventThreshold-r19"},
		{Name: "valueOfQ-r19"},
		{Name: "conditionFulfillmentIndicator-r19", Optional: true},
	},
}

const (
	CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event7_r19_ConditionFulfillmentIndicator_r19_Enabled = 0
)

var cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19ConditionFulfillmentIndicatorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event7_r19_ConditionFulfillmentIndicator_r19_Enabled},
}

var cSIReportUEInitiatedR19ReportTransmissionModeR19ModeAR19ReportSlotOffsetListR19Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportUEInitiatedR19ReportTransmissionModeR19ModeBR19PuschResourceListOfModeBR19Constraints = per.SizeRange(1, common.MaxNrofBWPs)

const (
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym0   = 0
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym1   = 1
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym2   = 2
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym4   = 3
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym8   = 4
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym16  = 5
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym32  = 6
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym64  = 7
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym128 = 8
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym256 = 9
	CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym512 = 10
)

var cSIReportUEInitiatedR19ReportTransmissionModeR19ModeBR19MinimumPucchPuschOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym0, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym1, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym2, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym4, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym8, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym16, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym32, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym64, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym128, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym256, CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19_MinimumPucch_PuschOffset_r19_Sym512},
}

const (
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms4    = 0
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms5    = 1
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms8    = 2
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms10   = 3
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms16   = 4
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms20   = 5
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms40   = 6
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms80   = 7
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms160  = 8
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms320  = 9
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms640  = 10
	CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms1280 = 11
)

var cSIReportUEInitiatedR19EventCountWindowR19EventDetectionTimeWindowR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms4, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms5, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms8, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms10, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms16, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms20, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms40, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms80, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms160, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms320, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms640, CSI_ReportUE_Initiated_r19_EventCountWindow_r19_EventDetectionTimeWindow_r19_Ms1280},
}

var cSIReportUEInitiatedR19UeInitiatedResourceConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-Cell-r19"},
		{Name: "ue-InitiatedResourceList-r19"},
	},
}

const (
	CSI_ReportUE_Initiated_r19_Ue_InitiatedResourceConfig_r19_Pucch_Cell_r19_SpCell      = 0
	CSI_ReportUE_Initiated_r19_Ue_InitiatedResourceConfig_r19_Pucch_Cell_r19_Pucch_Scell = 1
)

var cSIReportUEInitiatedR19UeInitiatedResourceConfigR19PucchCellR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportUE_Initiated_r19_Ue_InitiatedResourceConfig_r19_Pucch_Cell_r19_SpCell, CSI_ReportUE_Initiated_r19_Ue_InitiatedResourceConfig_r19_Pucch_Cell_r19_Pucch_Scell},
}

var cSIReportUEInitiatedR19UeInitiatedResourceConfigR19UeInitiatedResourceListR19Constraints = per.SizeRange(1, common.MaxNrofBWPs)

type CSI_ReportUE_Initiated_r19 struct {
	EventTypeUE_Initiated_r19 struct {
		Choice     int
		Event1_r19 *struct{ Event1Threshold_r19 RSRP_Range }
		Event2_r19 *struct {
			EventThreshold_r19                int64
			ConditionFulfillmentIndicator_r19 *int64
		}
		Event7_r19 *struct {
			EventThreshold_r19                int64
			ValueOfQ_r19                      int64
			ConditionFulfillmentIndicator_r19 *int64
		}
	}
	ReportTransmissionMode_r19 struct {
		Choice    int
		ModeA_r19 *struct{ ReportSlotOffsetList_r19 []int64 }
		ModeB_r19 *struct {
			Pusch_ResourceListOfModeB_r19 []PUSCH_ResourceOfModeB_r19
			MinimumPucch_PuschOffset_r19  int64
		}
	}
	NrofReportedRS_UE_Initiated_r19 int64
	Tci_ServCellIndex_r19           *ServCellIndex
	CurrentBeamReport_r19           *int64
	EventCountWindow_r19            *struct {
		EventInstanceCount_r19       int64
		EventDetectionTimeWindow_r19 int64
	}
	Ue_InitiatedResourceConfig_r19 struct {
		Pucch_Cell_r19               int64
		Ue_InitiatedResourceList_r19 []UE_InitiatedResource_r19
	}
}

func (ie *CSI_ReportUE_Initiated_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIReportUEInitiatedR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tci_ServCellIndex_r19 != nil, ie.CurrentBeamReport_r19 != nil, ie.EventCountWindow_r19 != nil}); err != nil {
		return err
	}

	// 3. eventTypeUE-Initiated-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_ReportUE_Initiated_r19EventTypeUEInitiatedR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.EventTypeUE_Initiated_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.EventTypeUE_Initiated_r19.Choice {
		case CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event1_r19:
			c := (*ie.EventTypeUE_Initiated_r19.Event1_r19)
			cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event1R19Seq := e.NewSequenceEncoder(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event1R19Constraints)
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event1R19Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.Event1Threshold_r19.Encode(e); err != nil {
				return err
			}
		case CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event2_r19:
			c := (*ie.EventTypeUE_Initiated_r19.Event2_r19)
			cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Seq := e.NewSequenceEncoder(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Constraints)
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Seq.EncodePreamble([]bool{c.ConditionFulfillmentIndicator_r19 != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.EventThreshold_r19, per.Constrained(0, 31)); err != nil {
				return err
			}
			if c.ConditionFulfillmentIndicator_r19 != nil {
				if err := e.EncodeEnumerated((*c.ConditionFulfillmentIndicator_r19), cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19ConditionFulfillmentIndicatorR19Constraints); err != nil {
					return err
				}
			}
		case CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event7_r19:
			c := (*ie.EventTypeUE_Initiated_r19.Event7_r19)
			cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Seq := e.NewSequenceEncoder(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Constraints)
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Seq.EncodePreamble([]bool{c.ConditionFulfillmentIndicator_r19 != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.EventThreshold_r19, per.Constrained(0, 31)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ValueOfQ_r19, per.Constrained(1, 8)); err != nil {
				return err
			}
			if c.ConditionFulfillmentIndicator_r19 != nil {
				if err := e.EncodeEnumerated((*c.ConditionFulfillmentIndicator_r19), cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19ConditionFulfillmentIndicatorR19Constraints); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.EventTypeUE_Initiated_r19.Choice), Constraint: "index out of range"}
		}
	}

	// 4. reportTransmissionMode-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_ReportUE_Initiated_r19ReportTransmissionModeR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportTransmissionMode_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportTransmissionMode_r19.Choice {
		case CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeA_r19:
			c := (*ie.ReportTransmissionMode_r19.ModeA_r19)
			{
				seqOf := e.NewSequenceOfEncoder(cSIReportUEInitiatedR19ReportTransmissionModeR19ModeAR19ReportSlotOffsetListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList_r19))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetList_r19 {
					if err := e.EncodeInteger(c.ReportSlotOffsetList_r19[i], per.Constrained(0, 128)); err != nil {
						return err
					}
				}
			}
		case CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19:
			c := (*ie.ReportTransmissionMode_r19.ModeB_r19)
			{
				seqOf := e.NewSequenceOfEncoder(cSIReportUEInitiatedR19ReportTransmissionModeR19ModeBR19PuschResourceListOfModeBR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Pusch_ResourceListOfModeB_r19))); err != nil {
					return err
				}
				for i := range c.Pusch_ResourceListOfModeB_r19 {
					if err := c.Pusch_ResourceListOfModeB_r19[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.MinimumPucch_PuschOffset_r19, cSIReportUEInitiatedR19ReportTransmissionModeR19ModeBR19MinimumPucchPuschOffsetR19Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportTransmissionMode_r19.Choice), Constraint: "index out of range"}
		}
	}

	// 5. nrofReportedRS-UE-Initiated-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofReportedRS_UE_Initiated_r19, cSIReportUEInitiatedR19NrofReportedRSUEInitiatedR19Constraints); err != nil {
			return err
		}
	}

	// 6. tci-ServCellIndex-r19: ref
	{
		if ie.Tci_ServCellIndex_r19 != nil {
			if err := ie.Tci_ServCellIndex_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. currentBeamReport-r19: enumerated
	{
		if ie.CurrentBeamReport_r19 != nil {
			if err := e.EncodeEnumerated(*ie.CurrentBeamReport_r19, cSIReportUEInitiatedR19CurrentBeamReportR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. eventCountWindow-r19: sequence
	{
		if ie.EventCountWindow_r19 != nil {
			c := ie.EventCountWindow_r19
			if err := e.EncodeInteger(c.EventInstanceCount_r19, per.Constrained(2, 16)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.EventDetectionTimeWindow_r19, cSIReportUEInitiatedR19EventCountWindowR19EventDetectionTimeWindowR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. ue-InitiatedResourceConfig-r19: sequence
	{
		{
			c := &ie.Ue_InitiatedResourceConfig_r19
			cSIReportUEInitiatedR19UeInitiatedResourceConfigR19Seq := e.NewSequenceEncoder(cSIReportUEInitiatedR19UeInitiatedResourceConfigR19Constraints)
			if err := cSIReportUEInitiatedR19UeInitiatedResourceConfigR19Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Pucch_Cell_r19, cSIReportUEInitiatedR19UeInitiatedResourceConfigR19PucchCellR19Constraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(cSIReportUEInitiatedR19UeInitiatedResourceConfigR19UeInitiatedResourceListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.Ue_InitiatedResourceList_r19))); err != nil {
					return err
				}
				for i := range c.Ue_InitiatedResourceList_r19 {
					if err := c.Ue_InitiatedResourceList_r19[i].Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *CSI_ReportUE_Initiated_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIReportUEInitiatedR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. eventTypeUE-Initiated-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_ReportUE_Initiated_r19EventTypeUEInitiatedR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.EventTypeUE_Initiated_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event1_r19:
			ie.EventTypeUE_Initiated_r19.Event1_r19 = &struct{ Event1Threshold_r19 RSRP_Range }{}
			c := (*ie.EventTypeUE_Initiated_r19.Event1_r19)
			cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event1R19Seq := d.NewSequenceDecoder(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event1R19Constraints)
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event1R19Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.Event1Threshold_r19.Decode(d); err != nil {
					return err
				}
			}
		case CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event2_r19:
			ie.EventTypeUE_Initiated_r19.Event2_r19 = &struct {
				EventThreshold_r19                int64
				ConditionFulfillmentIndicator_r19 *int64
			}{}
			c := (*ie.EventTypeUE_Initiated_r19.Event2_r19)
			cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Seq := d.NewSequenceDecoder(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Constraints)
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 31))
				if err != nil {
					return err
				}
				c.EventThreshold_r19 = v
			}
			if cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19Seq.IsComponentPresent(1) {
				c.ConditionFulfillmentIndicator_r19 = new(int64)
				v, err := d.DecodeEnumerated(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event2R19ConditionFulfillmentIndicatorR19Constraints)
				if err != nil {
					return err
				}
				(*c.ConditionFulfillmentIndicator_r19) = v
			}
		case CSI_ReportUE_Initiated_r19_EventTypeUE_Initiated_r19_Event7_r19:
			ie.EventTypeUE_Initiated_r19.Event7_r19 = &struct {
				EventThreshold_r19                int64
				ValueOfQ_r19                      int64
				ConditionFulfillmentIndicator_r19 *int64
			}{}
			c := (*ie.EventTypeUE_Initiated_r19.Event7_r19)
			cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Seq := d.NewSequenceDecoder(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Constraints)
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 31))
				if err != nil {
					return err
				}
				c.EventThreshold_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.ValueOfQ_r19 = v
			}
			if cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19Seq.IsComponentPresent(2) {
				c.ConditionFulfillmentIndicator_r19 = new(int64)
				v, err := d.DecodeEnumerated(cSIReportUEInitiatedR19EventTypeUEInitiatedR19Event7R19ConditionFulfillmentIndicatorR19Constraints)
				if err != nil {
					return err
				}
				(*c.ConditionFulfillmentIndicator_r19) = v
			}
		}
	}

	// 4. reportTransmissionMode-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_ReportUE_Initiated_r19ReportTransmissionModeR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportTransmissionMode_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeA_r19:
			ie.ReportTransmissionMode_r19.ModeA_r19 = &struct{ ReportSlotOffsetList_r19 []int64 }{}
			c := (*ie.ReportTransmissionMode_r19.ModeA_r19)
			{
				seqOf := d.NewSequenceOfDecoder(cSIReportUEInitiatedR19ReportTransmissionModeR19ModeAR19ReportSlotOffsetListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList_r19[i] = v
				}
			}
		case CSI_ReportUE_Initiated_r19_ReportTransmissionMode_r19_ModeB_r19:
			ie.ReportTransmissionMode_r19.ModeB_r19 = &struct {
				Pusch_ResourceListOfModeB_r19 []PUSCH_ResourceOfModeB_r19
				MinimumPucch_PuschOffset_r19  int64
			}{}
			c := (*ie.ReportTransmissionMode_r19.ModeB_r19)
			{
				seqOf := d.NewSequenceOfDecoder(cSIReportUEInitiatedR19ReportTransmissionModeR19ModeBR19PuschResourceListOfModeBR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Pusch_ResourceListOfModeB_r19 = make([]PUSCH_ResourceOfModeB_r19, n)
				for i := int64(0); i < n; i++ {
					if err := c.Pusch_ResourceListOfModeB_r19[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				v, err := d.DecodeEnumerated(cSIReportUEInitiatedR19ReportTransmissionModeR19ModeBR19MinimumPucchPuschOffsetR19Constraints)
				if err != nil {
					return err
				}
				c.MinimumPucch_PuschOffset_r19 = v
			}
		}
	}

	// 5. nrofReportedRS-UE-Initiated-r19: enumerated
	{
		v2, err := d.DecodeEnumerated(cSIReportUEInitiatedR19NrofReportedRSUEInitiatedR19Constraints)
		if err != nil {
			return err
		}
		ie.NrofReportedRS_UE_Initiated_r19 = v2
	}

	// 6. tci-ServCellIndex-r19: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Tci_ServCellIndex_r19 = new(ServCellIndex)
			if err := ie.Tci_ServCellIndex_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. currentBeamReport-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cSIReportUEInitiatedR19CurrentBeamReportR19Constraints)
			if err != nil {
				return err
			}
			ie.CurrentBeamReport_r19 = &idx
		}
	}

	// 8. eventCountWindow-r19: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.EventCountWindow_r19 = &struct {
				EventInstanceCount_r19       int64
				EventDetectionTimeWindow_r19 int64
			}{}
			c := ie.EventCountWindow_r19
			{
				v, err := d.DecodeInteger(per.Constrained(2, 16))
				if err != nil {
					return err
				}
				c.EventInstanceCount_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(cSIReportUEInitiatedR19EventCountWindowR19EventDetectionTimeWindowR19Constraints)
				if err != nil {
					return err
				}
				c.EventDetectionTimeWindow_r19 = v
			}
		}
	}

	// 9. ue-InitiatedResourceConfig-r19: sequence
	{
		{
			c := &ie.Ue_InitiatedResourceConfig_r19
			cSIReportUEInitiatedR19UeInitiatedResourceConfigR19Seq := d.NewSequenceDecoder(cSIReportUEInitiatedR19UeInitiatedResourceConfigR19Constraints)
			if err := cSIReportUEInitiatedR19UeInitiatedResourceConfigR19Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(cSIReportUEInitiatedR19UeInitiatedResourceConfigR19PucchCellR19Constraints)
				if err != nil {
					return err
				}
				c.Pucch_Cell_r19 = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(cSIReportUEInitiatedR19UeInitiatedResourceConfigR19UeInitiatedResourceListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Ue_InitiatedResourceList_r19 = make([]UE_InitiatedResource_r19, n)
				for i := int64(0); i < n; i++ {
					if err := c.Ue_InitiatedResourceList_r19[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
