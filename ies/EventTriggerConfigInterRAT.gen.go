// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EventTriggerConfigInterRAT (line 13399).

var eventTriggerConfigInterRATConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventId"},
		{Name: "rsType"},
		{Name: "reportInterval"},
		{Name: "reportAmount"},
		{Name: "reportQuantity"},
		{Name: "maxReportCells"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var eventTriggerConfigInterRATEventIdConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eventB1"},
		{Name: "eventB2"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "eventB1-UTRA-FDD-r16"},
		{Name: "eventB2-UTRA-FDD-r16"},
		{Name: "eventY1-Relay-r17"},
		{Name: "eventY2-Relay-r17"},
		{Name: "eventZ1-Relay-r18"},
	},
}

const (
	EventTriggerConfigInterRAT_EventId_EventB1 = 0
	EventTriggerConfigInterRAT_EventId_EventB2 = 1
)

const (
	EventTriggerConfigInterRAT_ReportAmount_r1       = 0
	EventTriggerConfigInterRAT_ReportAmount_r2       = 1
	EventTriggerConfigInterRAT_ReportAmount_r4       = 2
	EventTriggerConfigInterRAT_ReportAmount_r8       = 3
	EventTriggerConfigInterRAT_ReportAmount_r16      = 4
	EventTriggerConfigInterRAT_ReportAmount_r32      = 5
	EventTriggerConfigInterRAT_ReportAmount_r64      = 6
	EventTriggerConfigInterRAT_ReportAmount_Infinity = 7
)

var eventTriggerConfigInterRATReportAmountConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfigInterRAT_ReportAmount_r1, EventTriggerConfigInterRAT_ReportAmount_r2, EventTriggerConfigInterRAT_ReportAmount_r4, EventTriggerConfigInterRAT_ReportAmount_r8, EventTriggerConfigInterRAT_ReportAmount_r16, EventTriggerConfigInterRAT_ReportAmount_r32, EventTriggerConfigInterRAT_ReportAmount_r64, EventTriggerConfigInterRAT_ReportAmount_Infinity},
}

var eventTriggerConfigInterRATMaxReportCellsConstraints = per.Constrained(1, common.MaxCellReport)

const (
	EventTriggerConfigInterRAT_Ext_IncludeCommonLocationInfo_r16_True = 0
)

var eventTriggerConfigInterRATExtIncludeCommonLocationInfoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfigInterRAT_Ext_IncludeCommonLocationInfo_r16_True},
}

var eventTriggerConfigInterRATExtIncludeBTMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	EventTriggerConfigInterRAT_Ext_IncludeBT_Meas_r16_Release = 0
	EventTriggerConfigInterRAT_Ext_IncludeBT_Meas_r16_Setup   = 1
)

var eventTriggerConfigInterRATExtIncludeWLANMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	EventTriggerConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Release = 0
	EventTriggerConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Setup   = 1
)

var eventTriggerConfigInterRATExtIncludeSensorMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	EventTriggerConfigInterRAT_Ext_IncludeSensor_Meas_r16_Release = 0
	EventTriggerConfigInterRAT_Ext_IncludeSensor_Meas_r16_Setup   = 1
)

var eventTriggerConfigInterRATExtCellIndividualOffsetListR18Constraints = per.SizeRange(1, common.MaxCellMeasEUTRA)

var eventTriggerConfigInterRATEventIdEventB1Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "b1-ThresholdEUTRA"},
		{Name: "reportOnLeave"},
		{Name: "hysteresis"},
		{Name: "timeToTrigger"},
	},
}

var eventTriggerConfigInterRATEventIdEventB2Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "b2-Threshold1"},
		{Name: "b2-Threshold2EUTRA"},
		{Name: "reportOnLeave"},
		{Name: "hysteresis"},
		{Name: "timeToTrigger"},
	},
}

type EventTriggerConfigInterRAT struct {
	EventId struct {
		Choice  int
		EventB1 *struct {
			B1_ThresholdEUTRA MeasTriggerQuantityEUTRA
			ReportOnLeave     bool
			Hysteresis        Hysteresis
			TimeToTrigger     TimeToTrigger
		}
		EventB2 *struct {
			B2_Threshold1      MeasTriggerQuantity
			B2_Threshold2EUTRA MeasTriggerQuantityEUTRA
			ReportOnLeave      bool
			Hysteresis         Hysteresis
			TimeToTrigger      TimeToTrigger
		}
	}
	RsType                        NR_RS_Type
	ReportInterval                ReportInterval
	ReportAmount                  int64
	ReportQuantity                MeasReportQuantity
	MaxReportCells                int64
	ReportQuantityUTRA_FDD_r16    *MeasReportQuantityUTRA_FDD_r16
	IncludeCommonLocationInfo_r16 *int64
	IncludeBT_Meas_r16            *struct {
		Choice  int
		Release *struct{}
		Setup   *BT_NameList_r16
	}
	IncludeWLAN_Meas_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *WLAN_NameList_r16
	}
	IncludeSensor_Meas_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *Sensor_NameList_r16
	}
	ReportQuantityRelay_r17      *SL_MeasReportQuantity_r16
	CellIndividualOffsetList_r18 []CellIndividualOffsetList_EUTRA_r18
}

func (ie *EventTriggerConfigInterRAT) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eventTriggerConfigInterRATConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReportQuantityUTRA_FDD_r16 != nil
	hasExtGroup1 := ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil
	hasExtGroup2 := ie.ReportQuantityRelay_r17 != nil
	hasExtGroup3 := ie.CellIndividualOffsetList_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. eventId: choice
	{
		choiceEnc := e.NewChoiceEncoder(eventTriggerConfigInterRATEventIdConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.EventId.Choice), false, nil); err != nil {
			return err
		}
		switch ie.EventId.Choice {
		case EventTriggerConfigInterRAT_EventId_EventB1:
			c := (*ie.EventId.EventB1)
			eventTriggerConfigInterRATEventIdEventB1Seq := e.NewSequenceEncoder(eventTriggerConfigInterRATEventIdEventB1Constraints)
			if err := eventTriggerConfigInterRATEventIdEventB1Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.B1_ThresholdEUTRA.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.ReportOnLeave); err != nil {
				return err
			}
			if err := c.Hysteresis.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger.Encode(e); err != nil {
				return err
			}
		case EventTriggerConfigInterRAT_EventId_EventB2:
			c := (*ie.EventId.EventB2)
			eventTriggerConfigInterRATEventIdEventB2Seq := e.NewSequenceEncoder(eventTriggerConfigInterRATEventIdEventB2Constraints)
			if err := eventTriggerConfigInterRATEventIdEventB2Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.B2_Threshold1.Encode(e); err != nil {
				return err
			}
			if err := c.B2_Threshold2EUTRA.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.ReportOnLeave); err != nil {
				return err
			}
			if err := c.Hysteresis.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.EventId.Choice), Constraint: "index out of range"}
		}
	}

	// 3. rsType: ref
	{
		if err := ie.RsType.Encode(e); err != nil {
			return err
		}
	}

	// 4. reportInterval: ref
	{
		if err := ie.ReportInterval.Encode(e); err != nil {
			return err
		}
	}

	// 5. reportAmount: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount, eventTriggerConfigInterRATReportAmountConstraints); err != nil {
			return err
		}
	}

	// 6. reportQuantity: ref
	{
		if err := ie.ReportQuantity.Encode(e); err != nil {
			return err
		}
	}

	// 7. maxReportCells: integer
	{
		if err := e.EncodeInteger(ie.MaxReportCells, eventTriggerConfigInterRATMaxReportCellsConstraints); err != nil {
			return err
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
					{Name: "reportQuantityUTRA-FDD-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportQuantityUTRA_FDD_r16 != nil}); err != nil {
				return err
			}

			if ie.ReportQuantityUTRA_FDD_r16 != nil {
				if err := ie.ReportQuantityUTRA_FDD_r16.Encode(ex); err != nil {
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
					{Name: "includeCommonLocationInfo-r16", Optional: true},
					{Name: "includeBT-Meas-r16", Optional: true},
					{Name: "includeWLAN-Meas-r16", Optional: true},
					{Name: "includeSensor-Meas-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IncludeCommonLocationInfo_r16 != nil, ie.IncludeBT_Meas_r16 != nil, ie.IncludeWLAN_Meas_r16 != nil, ie.IncludeSensor_Meas_r16 != nil}); err != nil {
				return err
			}

			if ie.IncludeCommonLocationInfo_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IncludeCommonLocationInfo_r16, eventTriggerConfigInterRATExtIncludeCommonLocationInfoR16Constraints); err != nil {
					return err
				}
			}

			if ie.IncludeBT_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(eventTriggerConfigInterRATExtIncludeBTMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeBT_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeBT_Meas_r16).Choice {
				case EventTriggerConfigInterRAT_Ext_IncludeBT_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case EventTriggerConfigInterRAT_Ext_IncludeBT_Meas_r16_Setup:
					if err := (*ie.IncludeBT_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeWLAN_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(eventTriggerConfigInterRATExtIncludeWLANMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeWLAN_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeWLAN_Meas_r16).Choice {
				case EventTriggerConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case EventTriggerConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Setup:
					if err := (*ie.IncludeWLAN_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeSensor_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(eventTriggerConfigInterRATExtIncludeSensorMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeSensor_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeSensor_Meas_r16).Choice {
				case EventTriggerConfigInterRAT_Ext_IncludeSensor_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case EventTriggerConfigInterRAT_Ext_IncludeSensor_Meas_r16_Setup:
					if err := (*ie.IncludeSensor_Meas_r16).Setup.Encode(ex); err != nil {
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
					{Name: "reportQuantityRelay-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportQuantityRelay_r17 != nil}); err != nil {
				return err
			}

			if ie.ReportQuantityRelay_r17 != nil {
				if err := ie.ReportQuantityRelay_r17.Encode(ex); err != nil {
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
					{Name: "cellIndividualOffsetList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CellIndividualOffsetList_r18 != nil}); err != nil {
				return err
			}

			if ie.CellIndividualOffsetList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(eventTriggerConfigInterRATExtCellIndividualOffsetListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CellIndividualOffsetList_r18))); err != nil {
					return err
				}
				for i := range ie.CellIndividualOffsetList_r18 {
					if err := ie.CellIndividualOffsetList_r18[i].Encode(ex); err != nil {
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

func (ie *EventTriggerConfigInterRAT) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eventTriggerConfigInterRATConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. eventId: choice
	{
		choiceDec := d.NewChoiceDecoder(eventTriggerConfigInterRATEventIdConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.EventId.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case EventTriggerConfigInterRAT_EventId_EventB1:
			ie.EventId.EventB1 = &struct {
				B1_ThresholdEUTRA MeasTriggerQuantityEUTRA
				ReportOnLeave     bool
				Hysteresis        Hysteresis
				TimeToTrigger     TimeToTrigger
			}{}
			c := (*ie.EventId.EventB1)
			eventTriggerConfigInterRATEventIdEventB1Seq := d.NewSequenceDecoder(eventTriggerConfigInterRATEventIdEventB1Constraints)
			if err := eventTriggerConfigInterRATEventIdEventB1Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.B1_ThresholdEUTRA.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.ReportOnLeave = v
			}
			{
				if err := c.Hysteresis.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.TimeToTrigger.Decode(d); err != nil {
					return err
				}
			}
		case EventTriggerConfigInterRAT_EventId_EventB2:
			ie.EventId.EventB2 = &struct {
				B2_Threshold1      MeasTriggerQuantity
				B2_Threshold2EUTRA MeasTriggerQuantityEUTRA
				ReportOnLeave      bool
				Hysteresis         Hysteresis
				TimeToTrigger      TimeToTrigger
			}{}
			c := (*ie.EventId.EventB2)
			eventTriggerConfigInterRATEventIdEventB2Seq := d.NewSequenceDecoder(eventTriggerConfigInterRATEventIdEventB2Constraints)
			if err := eventTriggerConfigInterRATEventIdEventB2Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.B2_Threshold1.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.B2_Threshold2EUTRA.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.ReportOnLeave = v
			}
			{
				if err := c.Hysteresis.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.TimeToTrigger.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. rsType: ref
	{
		if err := ie.RsType.Decode(d); err != nil {
			return err
		}
	}

	// 4. reportInterval: ref
	{
		if err := ie.ReportInterval.Decode(d); err != nil {
			return err
		}
	}

	// 5. reportAmount: enumerated
	{
		v3, err := d.DecodeEnumerated(eventTriggerConfigInterRATReportAmountConstraints)
		if err != nil {
			return err
		}
		ie.ReportAmount = v3
	}

	// 6. reportQuantity: ref
	{
		if err := ie.ReportQuantity.Decode(d); err != nil {
			return err
		}
	}

	// 7. maxReportCells: integer
	{
		v5, err := d.DecodeInteger(eventTriggerConfigInterRATMaxReportCellsConstraints)
		if err != nil {
			return err
		}
		ie.MaxReportCells = v5
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
				{Name: "reportQuantityUTRA-FDD-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ReportQuantityUTRA_FDD_r16 = new(MeasReportQuantityUTRA_FDD_r16)
			if err := ie.ReportQuantityUTRA_FDD_r16.Decode(dx); err != nil {
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
				{Name: "includeCommonLocationInfo-r16", Optional: true},
				{Name: "includeBT-Meas-r16", Optional: true},
				{Name: "includeWLAN-Meas-r16", Optional: true},
				{Name: "includeSensor-Meas-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(eventTriggerConfigInterRATExtIncludeCommonLocationInfoR16Constraints)
			if err != nil {
				return err
			}
			ie.IncludeCommonLocationInfo_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.IncludeBT_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BT_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(eventTriggerConfigInterRATExtIncludeBTMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeBT_Meas_r16).Choice = int(index)
			switch index {
			case EventTriggerConfigInterRAT_Ext_IncludeBT_Meas_r16_Release:
				(*ie.IncludeBT_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case EventTriggerConfigInterRAT_Ext_IncludeBT_Meas_r16_Setup:
				(*ie.IncludeBT_Meas_r16).Setup = new(BT_NameList_r16)
				if err := (*ie.IncludeBT_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.IncludeWLAN_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *WLAN_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(eventTriggerConfigInterRATExtIncludeWLANMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeWLAN_Meas_r16).Choice = int(index)
			switch index {
			case EventTriggerConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Release:
				(*ie.IncludeWLAN_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case EventTriggerConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Setup:
				(*ie.IncludeWLAN_Meas_r16).Setup = new(WLAN_NameList_r16)
				if err := (*ie.IncludeWLAN_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.IncludeSensor_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Sensor_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(eventTriggerConfigInterRATExtIncludeSensorMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeSensor_Meas_r16).Choice = int(index)
			switch index {
			case EventTriggerConfigInterRAT_Ext_IncludeSensor_Meas_r16_Release:
				(*ie.IncludeSensor_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case EventTriggerConfigInterRAT_Ext_IncludeSensor_Meas_r16_Setup:
				(*ie.IncludeSensor_Meas_r16).Setup = new(Sensor_NameList_r16)
				if err := (*ie.IncludeSensor_Meas_r16).Setup.Decode(dx); err != nil {
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
				{Name: "reportQuantityRelay-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ReportQuantityRelay_r17 = new(SL_MeasReportQuantity_r16)
			if err := ie.ReportQuantityRelay_r17.Decode(dx); err != nil {
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
				{Name: "cellIndividualOffsetList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(eventTriggerConfigInterRATExtCellIndividualOffsetListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CellIndividualOffsetList_r18 = make([]CellIndividualOffsetList_EUTRA_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CellIndividualOffsetList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
