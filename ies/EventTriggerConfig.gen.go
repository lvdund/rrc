// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EventTriggerConfig (line 13638).

var eventTriggerConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventId"},
		{Name: "rsType"},
		{Name: "reportInterval"},
		{Name: "reportAmount"},
		{Name: "reportQuantityCell"},
		{Name: "maxReportCells"},
		{Name: "reportQuantityRS-Indexes", Optional: true},
		{Name: "maxNrofRS-IndexesToReport", Optional: true},
		{Name: "includeBeamMeasurements"},
		{Name: "reportAddNeighMeas", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var eventTriggerConfigEventIdConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eventA1"},
		{Name: "eventA2"},
		{Name: "eventA3"},
		{Name: "eventA4"},
		{Name: "eventA5"},
		{Name: "eventA6"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "eventX1-r17"},
		{Name: "eventX2-r17"},
		{Name: "eventD1-r17"},
		{Name: "eventH1-r18"},
		{Name: "eventH2-r18"},
		{Name: "eventA3H1-r18"},
		{Name: "eventA3H2-r18"},
		{Name: "eventA4H1-r18"},
		{Name: "eventA4H2-r18"},
		{Name: "eventA5H1-r18"},
		{Name: "eventA5H2-r18"},
		{Name: "eventD2-r18"},
	},
}

const (
	EventTriggerConfig_EventId_EventA1 = 0
	EventTriggerConfig_EventId_EventA2 = 1
	EventTriggerConfig_EventId_EventA3 = 2
	EventTriggerConfig_EventId_EventA4 = 3
	EventTriggerConfig_EventId_EventA5 = 4
	EventTriggerConfig_EventId_EventA6 = 5
)

const (
	EventTriggerConfig_ReportAmount_r1       = 0
	EventTriggerConfig_ReportAmount_r2       = 1
	EventTriggerConfig_ReportAmount_r4       = 2
	EventTriggerConfig_ReportAmount_r8       = 3
	EventTriggerConfig_ReportAmount_r16      = 4
	EventTriggerConfig_ReportAmount_r32      = 5
	EventTriggerConfig_ReportAmount_r64      = 6
	EventTriggerConfig_ReportAmount_Infinity = 7
)

var eventTriggerConfigReportAmountConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfig_ReportAmount_r1, EventTriggerConfig_ReportAmount_r2, EventTriggerConfig_ReportAmount_r4, EventTriggerConfig_ReportAmount_r8, EventTriggerConfig_ReportAmount_r16, EventTriggerConfig_ReportAmount_r32, EventTriggerConfig_ReportAmount_r64, EventTriggerConfig_ReportAmount_Infinity},
}

var eventTriggerConfigMaxReportCellsConstraints = per.Constrained(1, common.MaxCellReport)

var eventTriggerConfigMaxNrofRSIndexesToReportConstraints = per.Constrained(1, common.MaxNrofIndexesToReport)

const (
	EventTriggerConfig_ReportAddNeighMeas_Setup = 0
)

var eventTriggerConfigReportAddNeighMeasConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfig_ReportAddNeighMeas_Setup},
}

const (
	EventTriggerConfig_Ext_IncludeCommonLocationInfo_r16_True = 0
)

var eventTriggerConfigExtIncludeCommonLocationInfoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfig_Ext_IncludeCommonLocationInfo_r16_True},
}

var eventTriggerConfigExtIncludeBTMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	EventTriggerConfig_Ext_IncludeBT_Meas_r16_Release = 0
	EventTriggerConfig_Ext_IncludeBT_Meas_r16_Setup   = 1
)

var eventTriggerConfigExtIncludeWLANMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	EventTriggerConfig_Ext_IncludeWLAN_Meas_r16_Release = 0
	EventTriggerConfig_Ext_IncludeWLAN_Meas_r16_Setup   = 1
)

var eventTriggerConfigExtIncludeSensorMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	EventTriggerConfig_Ext_IncludeSensor_Meas_r16_Release = 0
	EventTriggerConfig_Ext_IncludeSensor_Meas_r16_Setup   = 1
)

const (
	EventTriggerConfig_Ext_CoarseLocationRequest_r17_True = 0
)

var eventTriggerConfigExtCoarseLocationRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfig_Ext_CoarseLocationRequest_r17_True},
}

var eventTriggerConfigNumberOfTriggeringCellsR18Constraints = per.Constrained(2, common.MaxCellReport)

var eventTriggerConfigExtCellIndividualOffsetListR18Constraints = per.SizeRange(1, common.MaxNrofCellMeas)

const (
	EventTriggerConfig_Ext_ReportOnBestCellChange_r18_N1 = 0
	EventTriggerConfig_Ext_ReportOnBestCellChange_r18_N2 = 1
)

var eventTriggerConfigExtReportOnBestCellChangeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfig_Ext_ReportOnBestCellChange_r18_N1, EventTriggerConfig_Ext_ReportOnBestCellChange_r18_N2},
}

const (
	EventTriggerConfig_Ext_EnteringLeavingReport_r18_True = 0
)

var eventTriggerConfigExtEnteringLeavingReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfig_Ext_EnteringLeavingReport_r18_True},
}

type EventTriggerConfig struct {
	EventId struct {
		Choice  int
		EventA1 *struct {
			A1_Threshold  MeasTriggerQuantity
			ReportOnLeave bool
			Hysteresis    Hysteresis
			TimeToTrigger TimeToTrigger
		}
		EventA2 *struct {
			A2_Threshold  MeasTriggerQuantity
			ReportOnLeave bool
			Hysteresis    Hysteresis
			TimeToTrigger TimeToTrigger
		}
		EventA3 *struct {
			A3_Offset          MeasTriggerQuantityOffset
			ReportOnLeave      bool
			Hysteresis         Hysteresis
			TimeToTrigger      TimeToTrigger
			UseAllowedCellList bool
		}
		EventA4 *struct {
			A4_Threshold       MeasTriggerQuantity
			ReportOnLeave      bool
			Hysteresis         Hysteresis
			TimeToTrigger      TimeToTrigger
			UseAllowedCellList bool
		}
		EventA5 *struct {
			A5_Threshold1      MeasTriggerQuantity
			A5_Threshold2      MeasTriggerQuantity
			ReportOnLeave      bool
			Hysteresis         Hysteresis
			TimeToTrigger      TimeToTrigger
			UseAllowedCellList bool
		}
		EventA6 *struct {
			A6_Offset          MeasTriggerQuantityOffset
			ReportOnLeave      bool
			Hysteresis         Hysteresis
			TimeToTrigger      TimeToTrigger
			UseAllowedCellList bool
		}
	}
	RsType                        NR_RS_Type
	ReportInterval                ReportInterval
	ReportAmount                  int64
	ReportQuantityCell            MeasReportQuantity
	MaxReportCells                int64
	ReportQuantityRS_Indexes      *MeasReportQuantity
	MaxNrofRS_IndexesToReport     *int64
	IncludeBeamMeasurements       bool
	ReportAddNeighMeas            *int64
	MeasRSSI_ReportConfig_r16     *MeasRSSI_ReportConfig_r16
	UseT312_r16                   *bool
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
	CoarseLocationRequest_r17    *int64
	ReportQuantityRelay_r17      *SL_MeasReportQuantity_r16
	NumberOfTriggeringCells_r18  *int64
	CellIndividualOffsetList_r18 []CellIndividualOffsetList_r18
	EventX1_SD_Threshold1_r18    *SL_MeasTriggerQuantity_r16
	EventX2_SD_Threshold_r18     *SL_MeasTriggerQuantity_r16
	ReportOnBestCellChange_r18   *int64
	EnteringLeavingReport_r18    *int64
	EventD2PhysCellId_r19        *PhysCellId
}

func (ie *EventTriggerConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eventTriggerConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MeasRSSI_ReportConfig_r16 != nil || ie.UseT312_r16 != nil || ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil
	hasExtGroup1 := ie.CoarseLocationRequest_r17 != nil || ie.ReportQuantityRelay_r17 != nil
	hasExtGroup2 := ie.NumberOfTriggeringCells_r18 != nil || ie.CellIndividualOffsetList_r18 != nil || ie.EventX1_SD_Threshold1_r18 != nil || ie.EventX2_SD_Threshold_r18 != nil || ie.ReportOnBestCellChange_r18 != nil || ie.EnteringLeavingReport_r18 != nil
	hasExtGroup3 := ie.EventD2PhysCellId_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReportQuantityRS_Indexes != nil, ie.MaxNrofRS_IndexesToReport != nil, ie.ReportAddNeighMeas != nil}); err != nil {
		return err
	}

	// 3. eventId: choice
	{
		choiceEnc := e.NewChoiceEncoder(eventTriggerConfigEventIdConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.EventId.Choice), false, nil); err != nil {
			return err
		}
		switch ie.EventId.Choice {
		case EventTriggerConfig_EventId_EventA1:
			c := (*ie.EventId.EventA1)
			if err := c.A1_Threshold.Encode(e); err != nil {
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
		case EventTriggerConfig_EventId_EventA2:
			c := (*ie.EventId.EventA2)
			if err := c.A2_Threshold.Encode(e); err != nil {
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
		case EventTriggerConfig_EventId_EventA3:
			c := (*ie.EventId.EventA3)
			if err := c.A3_Offset.Encode(e); err != nil {
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
			if err := e.EncodeBoolean(c.UseAllowedCellList); err != nil {
				return err
			}
		case EventTriggerConfig_EventId_EventA4:
			c := (*ie.EventId.EventA4)
			if err := c.A4_Threshold.Encode(e); err != nil {
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
			if err := e.EncodeBoolean(c.UseAllowedCellList); err != nil {
				return err
			}
		case EventTriggerConfig_EventId_EventA5:
			c := (*ie.EventId.EventA5)
			if err := c.A5_Threshold1.Encode(e); err != nil {
				return err
			}
			if err := c.A5_Threshold2.Encode(e); err != nil {
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
			if err := e.EncodeBoolean(c.UseAllowedCellList); err != nil {
				return err
			}
		case EventTriggerConfig_EventId_EventA6:
			c := (*ie.EventId.EventA6)
			if err := c.A6_Offset.Encode(e); err != nil {
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
			if err := e.EncodeBoolean(c.UseAllowedCellList); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.EventId.Choice), Constraint: "index out of range"}
		}
	}

	// 4. rsType: ref
	{
		if err := ie.RsType.Encode(e); err != nil {
			return err
		}
	}

	// 5. reportInterval: ref
	{
		if err := ie.ReportInterval.Encode(e); err != nil {
			return err
		}
	}

	// 6. reportAmount: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount, eventTriggerConfigReportAmountConstraints); err != nil {
			return err
		}
	}

	// 7. reportQuantityCell: ref
	{
		if err := ie.ReportQuantityCell.Encode(e); err != nil {
			return err
		}
	}

	// 8. maxReportCells: integer
	{
		if err := e.EncodeInteger(ie.MaxReportCells, eventTriggerConfigMaxReportCellsConstraints); err != nil {
			return err
		}
	}

	// 9. reportQuantityRS-Indexes: ref
	{
		if ie.ReportQuantityRS_Indexes != nil {
			if err := ie.ReportQuantityRS_Indexes.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. maxNrofRS-IndexesToReport: integer
	{
		if ie.MaxNrofRS_IndexesToReport != nil {
			if err := e.EncodeInteger(*ie.MaxNrofRS_IndexesToReport, eventTriggerConfigMaxNrofRSIndexesToReportConstraints); err != nil {
				return err
			}
		}
	}

	// 11. includeBeamMeasurements: boolean
	{
		if err := e.EncodeBoolean(ie.IncludeBeamMeasurements); err != nil {
			return err
		}
	}

	// 12. reportAddNeighMeas: enumerated
	{
		if ie.ReportAddNeighMeas != nil {
			if err := e.EncodeEnumerated(*ie.ReportAddNeighMeas, eventTriggerConfigReportAddNeighMeasConstraints); err != nil {
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
					{Name: "measRSSI-ReportConfig-r16", Optional: true},
					{Name: "useT312-r16", Optional: true},
					{Name: "includeCommonLocationInfo-r16", Optional: true},
					{Name: "includeBT-Meas-r16", Optional: true},
					{Name: "includeWLAN-Meas-r16", Optional: true},
					{Name: "includeSensor-Meas-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasRSSI_ReportConfig_r16 != nil, ie.UseT312_r16 != nil, ie.IncludeCommonLocationInfo_r16 != nil, ie.IncludeBT_Meas_r16 != nil, ie.IncludeWLAN_Meas_r16 != nil, ie.IncludeSensor_Meas_r16 != nil}); err != nil {
				return err
			}

			if ie.MeasRSSI_ReportConfig_r16 != nil {
				if err := ie.MeasRSSI_ReportConfig_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.UseT312_r16 != nil {
				if err := ex.EncodeBoolean(*ie.UseT312_r16); err != nil {
					return err
				}
			}

			if ie.IncludeCommonLocationInfo_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IncludeCommonLocationInfo_r16, eventTriggerConfigExtIncludeCommonLocationInfoR16Constraints); err != nil {
					return err
				}
			}

			if ie.IncludeBT_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(eventTriggerConfigExtIncludeBTMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeBT_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeBT_Meas_r16).Choice {
				case EventTriggerConfig_Ext_IncludeBT_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case EventTriggerConfig_Ext_IncludeBT_Meas_r16_Setup:
					if err := (*ie.IncludeBT_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeWLAN_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(eventTriggerConfigExtIncludeWLANMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeWLAN_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeWLAN_Meas_r16).Choice {
				case EventTriggerConfig_Ext_IncludeWLAN_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case EventTriggerConfig_Ext_IncludeWLAN_Meas_r16_Setup:
					if err := (*ie.IncludeWLAN_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeSensor_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(eventTriggerConfigExtIncludeSensorMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeSensor_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeSensor_Meas_r16).Choice {
				case EventTriggerConfig_Ext_IncludeSensor_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case EventTriggerConfig_Ext_IncludeSensor_Meas_r16_Setup:
					if err := (*ie.IncludeSensor_Meas_r16).Setup.Encode(ex); err != nil {
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
					{Name: "coarseLocationRequest-r17", Optional: true},
					{Name: "reportQuantityRelay-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CoarseLocationRequest_r17 != nil, ie.ReportQuantityRelay_r17 != nil}); err != nil {
				return err
			}

			if ie.CoarseLocationRequest_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.CoarseLocationRequest_r17, eventTriggerConfigExtCoarseLocationRequestR17Constraints); err != nil {
					return err
				}
			}

			if ie.ReportQuantityRelay_r17 != nil {
				if err := ie.ReportQuantityRelay_r17.Encode(ex); err != nil {
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
					{Name: "numberOfTriggeringCells-r18", Optional: true},
					{Name: "cellIndividualOffsetList-r18", Optional: true},
					{Name: "eventX1-SD-Threshold1-r18", Optional: true},
					{Name: "eventX2-SD-Threshold-r18", Optional: true},
					{Name: "reportOnBestCellChange-r18", Optional: true},
					{Name: "enteringLeavingReport-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NumberOfTriggeringCells_r18 != nil, ie.CellIndividualOffsetList_r18 != nil, ie.EventX1_SD_Threshold1_r18 != nil, ie.EventX2_SD_Threshold_r18 != nil, ie.ReportOnBestCellChange_r18 != nil, ie.EnteringLeavingReport_r18 != nil}); err != nil {
				return err
			}

			if ie.NumberOfTriggeringCells_r18 != nil {
				if err := ex.EncodeInteger(*ie.NumberOfTriggeringCells_r18, eventTriggerConfigNumberOfTriggeringCellsR18Constraints); err != nil {
					return err
				}
			}

			if ie.CellIndividualOffsetList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(eventTriggerConfigExtCellIndividualOffsetListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CellIndividualOffsetList_r18))); err != nil {
					return err
				}
				for i := range ie.CellIndividualOffsetList_r18 {
					if err := ie.CellIndividualOffsetList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.EventX1_SD_Threshold1_r18 != nil {
				if err := ie.EventX1_SD_Threshold1_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.EventX2_SD_Threshold_r18 != nil {
				if err := ie.EventX2_SD_Threshold_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ReportOnBestCellChange_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ReportOnBestCellChange_r18, eventTriggerConfigExtReportOnBestCellChangeR18Constraints); err != nil {
					return err
				}
			}

			if ie.EnteringLeavingReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnteringLeavingReport_r18, eventTriggerConfigExtEnteringLeavingReportR18Constraints); err != nil {
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
					{Name: "eventD2PhysCellId-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EventD2PhysCellId_r19 != nil}); err != nil {
				return err
			}

			if ie.EventD2PhysCellId_r19 != nil {
				if err := ie.EventD2PhysCellId_r19.Encode(ex); err != nil {
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

func (ie *EventTriggerConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eventTriggerConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. eventId: choice
	{
		choiceDec := d.NewChoiceDecoder(eventTriggerConfigEventIdConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.EventId.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case EventTriggerConfig_EventId_EventA1:
			ie.EventId.EventA1 = &struct {
				A1_Threshold  MeasTriggerQuantity
				ReportOnLeave bool
				Hysteresis    Hysteresis
				TimeToTrigger TimeToTrigger
			}{}
			c := (*ie.EventId.EventA1)
			{
				if err := c.A1_Threshold.Decode(d); err != nil {
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
		case EventTriggerConfig_EventId_EventA2:
			ie.EventId.EventA2 = &struct {
				A2_Threshold  MeasTriggerQuantity
				ReportOnLeave bool
				Hysteresis    Hysteresis
				TimeToTrigger TimeToTrigger
			}{}
			c := (*ie.EventId.EventA2)
			{
				if err := c.A2_Threshold.Decode(d); err != nil {
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
		case EventTriggerConfig_EventId_EventA3:
			ie.EventId.EventA3 = &struct {
				A3_Offset          MeasTriggerQuantityOffset
				ReportOnLeave      bool
				Hysteresis         Hysteresis
				TimeToTrigger      TimeToTrigger
				UseAllowedCellList bool
			}{}
			c := (*ie.EventId.EventA3)
			{
				if err := c.A3_Offset.Decode(d); err != nil {
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
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.UseAllowedCellList = v
			}
		case EventTriggerConfig_EventId_EventA4:
			ie.EventId.EventA4 = &struct {
				A4_Threshold       MeasTriggerQuantity
				ReportOnLeave      bool
				Hysteresis         Hysteresis
				TimeToTrigger      TimeToTrigger
				UseAllowedCellList bool
			}{}
			c := (*ie.EventId.EventA4)
			{
				if err := c.A4_Threshold.Decode(d); err != nil {
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
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.UseAllowedCellList = v
			}
		case EventTriggerConfig_EventId_EventA5:
			ie.EventId.EventA5 = &struct {
				A5_Threshold1      MeasTriggerQuantity
				A5_Threshold2      MeasTriggerQuantity
				ReportOnLeave      bool
				Hysteresis         Hysteresis
				TimeToTrigger      TimeToTrigger
				UseAllowedCellList bool
			}{}
			c := (*ie.EventId.EventA5)
			{
				if err := c.A5_Threshold1.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.A5_Threshold2.Decode(d); err != nil {
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
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.UseAllowedCellList = v
			}
		case EventTriggerConfig_EventId_EventA6:
			ie.EventId.EventA6 = &struct {
				A6_Offset          MeasTriggerQuantityOffset
				ReportOnLeave      bool
				Hysteresis         Hysteresis
				TimeToTrigger      TimeToTrigger
				UseAllowedCellList bool
			}{}
			c := (*ie.EventId.EventA6)
			{
				if err := c.A6_Offset.Decode(d); err != nil {
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
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.UseAllowedCellList = v
			}
		}
	}

	// 4. rsType: ref
	{
		if err := ie.RsType.Decode(d); err != nil {
			return err
		}
	}

	// 5. reportInterval: ref
	{
		if err := ie.ReportInterval.Decode(d); err != nil {
			return err
		}
	}

	// 6. reportAmount: enumerated
	{
		v3, err := d.DecodeEnumerated(eventTriggerConfigReportAmountConstraints)
		if err != nil {
			return err
		}
		ie.ReportAmount = v3
	}

	// 7. reportQuantityCell: ref
	{
		if err := ie.ReportQuantityCell.Decode(d); err != nil {
			return err
		}
	}

	// 8. maxReportCells: integer
	{
		v5, err := d.DecodeInteger(eventTriggerConfigMaxReportCellsConstraints)
		if err != nil {
			return err
		}
		ie.MaxReportCells = v5
	}

	// 9. reportQuantityRS-Indexes: ref
	{
		if seq.IsComponentPresent(6) {
			ie.ReportQuantityRS_Indexes = new(MeasReportQuantity)
			if err := ie.ReportQuantityRS_Indexes.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. maxNrofRS-IndexesToReport: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(eventTriggerConfigMaxNrofRSIndexesToReportConstraints)
			if err != nil {
				return err
			}
			ie.MaxNrofRS_IndexesToReport = &v
		}
	}

	// 11. includeBeamMeasurements: boolean
	{
		v8, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.IncludeBeamMeasurements = v8
	}

	// 12. reportAddNeighMeas: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(eventTriggerConfigReportAddNeighMeasConstraints)
			if err != nil {
				return err
			}
			ie.ReportAddNeighMeas = &idx
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
				{Name: "measRSSI-ReportConfig-r16", Optional: true},
				{Name: "useT312-r16", Optional: true},
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
			ie.MeasRSSI_ReportConfig_r16 = new(MeasRSSI_ReportConfig_r16)
			if err := ie.MeasRSSI_ReportConfig_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.UseT312_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(eventTriggerConfigExtIncludeCommonLocationInfoR16Constraints)
			if err != nil {
				return err
			}
			ie.IncludeCommonLocationInfo_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.IncludeBT_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BT_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(eventTriggerConfigExtIncludeBTMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeBT_Meas_r16).Choice = int(index)
			switch index {
			case EventTriggerConfig_Ext_IncludeBT_Meas_r16_Release:
				(*ie.IncludeBT_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case EventTriggerConfig_Ext_IncludeBT_Meas_r16_Setup:
				(*ie.IncludeBT_Meas_r16).Setup = new(BT_NameList_r16)
				if err := (*ie.IncludeBT_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.IncludeWLAN_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *WLAN_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(eventTriggerConfigExtIncludeWLANMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeWLAN_Meas_r16).Choice = int(index)
			switch index {
			case EventTriggerConfig_Ext_IncludeWLAN_Meas_r16_Release:
				(*ie.IncludeWLAN_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case EventTriggerConfig_Ext_IncludeWLAN_Meas_r16_Setup:
				(*ie.IncludeWLAN_Meas_r16).Setup = new(WLAN_NameList_r16)
				if err := (*ie.IncludeWLAN_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.IncludeSensor_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Sensor_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(eventTriggerConfigExtIncludeSensorMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeSensor_Meas_r16).Choice = int(index)
			switch index {
			case EventTriggerConfig_Ext_IncludeSensor_Meas_r16_Release:
				(*ie.IncludeSensor_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case EventTriggerConfig_Ext_IncludeSensor_Meas_r16_Setup:
				(*ie.IncludeSensor_Meas_r16).Setup = new(Sensor_NameList_r16)
				if err := (*ie.IncludeSensor_Meas_r16).Setup.Decode(dx); err != nil {
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
				{Name: "coarseLocationRequest-r17", Optional: true},
				{Name: "reportQuantityRelay-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(eventTriggerConfigExtCoarseLocationRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.CoarseLocationRequest_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ReportQuantityRelay_r17 = new(SL_MeasReportQuantity_r16)
			if err := ie.ReportQuantityRelay_r17.Decode(dx); err != nil {
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
				{Name: "numberOfTriggeringCells-r18", Optional: true},
				{Name: "cellIndividualOffsetList-r18", Optional: true},
				{Name: "eventX1-SD-Threshold1-r18", Optional: true},
				{Name: "eventX2-SD-Threshold-r18", Optional: true},
				{Name: "reportOnBestCellChange-r18", Optional: true},
				{Name: "enteringLeavingReport-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(eventTriggerConfigNumberOfTriggeringCellsR18Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfTriggeringCells_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(eventTriggerConfigExtCellIndividualOffsetListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CellIndividualOffsetList_r18 = make([]CellIndividualOffsetList_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CellIndividualOffsetList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.EventX1_SD_Threshold1_r18 = new(SL_MeasTriggerQuantity_r16)
			if err := ie.EventX1_SD_Threshold1_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.EventX2_SD_Threshold_r18 = new(SL_MeasTriggerQuantity_r16)
			if err := ie.EventX2_SD_Threshold_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(eventTriggerConfigExtReportOnBestCellChangeR18Constraints)
			if err != nil {
				return err
			}
			ie.ReportOnBestCellChange_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(eventTriggerConfigExtEnteringLeavingReportR18Constraints)
			if err != nil {
				return err
			}
			ie.EnteringLeavingReport_r18 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "eventD2PhysCellId-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.EventD2PhysCellId_r19 = new(PhysCellId)
			if err := ie.EventD2PhysCellId_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
