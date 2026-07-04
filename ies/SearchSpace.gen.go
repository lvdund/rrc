// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SearchSpace (line 14370).

var searchSpaceConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "searchSpaceId"},
		{Name: "controlResourceSetId", Optional: true},
		{Name: "monitoringSlotPeriodicityAndOffset", Optional: true},
		{Name: "duration", Optional: true},
		{Name: "monitoringSymbolsWithinSlot", Optional: true},
		{Name: "nrofCandidates", Optional: true},
		{Name: "searchSpaceType", Optional: true},
	},
}

var searchSpaceMonitoringSlotPeriodicityAndOffsetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl1"},
		{Name: "sl2"},
		{Name: "sl4"},
		{Name: "sl5"},
		{Name: "sl8"},
		{Name: "sl10"},
		{Name: "sl16"},
		{Name: "sl20"},
		{Name: "sl40"},
		{Name: "sl80"},
		{Name: "sl160"},
		{Name: "sl320"},
		{Name: "sl640"},
		{Name: "sl1280"},
		{Name: "sl2560"},
	},
}

const (
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl1    = 0
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl2    = 1
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl4    = 2
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl5    = 3
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl8    = 4
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl10   = 5
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl16   = 6
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl20   = 7
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl40   = 8
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl80   = 9
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl160  = 10
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl320  = 11
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl640  = 12
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl1280 = 13
	SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl2560 = 14
)

var searchSpaceDurationConstraints = per.Constrained(2, 2559)

var searchSpaceMonitoringSymbolsWithinSlotConstraints = per.FixedSize(14)

var searchSpaceSearchSpaceTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "common"},
		{Name: "ue-Specific"},
	},
}

const (
	SearchSpace_SearchSpaceType_Common      = 0
	SearchSpace_SearchSpaceType_Ue_Specific = 1
)

const (
	SearchSpace_NrofCandidates_AggregationLevel1_N0 = 0
	SearchSpace_NrofCandidates_AggregationLevel1_N1 = 1
	SearchSpace_NrofCandidates_AggregationLevel1_N2 = 2
	SearchSpace_NrofCandidates_AggregationLevel1_N3 = 3
	SearchSpace_NrofCandidates_AggregationLevel1_N4 = 4
	SearchSpace_NrofCandidates_AggregationLevel1_N5 = 5
	SearchSpace_NrofCandidates_AggregationLevel1_N6 = 6
	SearchSpace_NrofCandidates_AggregationLevel1_N8 = 7
)

var searchSpaceNrofCandidatesAggregationLevel1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_NrofCandidates_AggregationLevel1_N0, SearchSpace_NrofCandidates_AggregationLevel1_N1, SearchSpace_NrofCandidates_AggregationLevel1_N2, SearchSpace_NrofCandidates_AggregationLevel1_N3, SearchSpace_NrofCandidates_AggregationLevel1_N4, SearchSpace_NrofCandidates_AggregationLevel1_N5, SearchSpace_NrofCandidates_AggregationLevel1_N6, SearchSpace_NrofCandidates_AggregationLevel1_N8},
}

const (
	SearchSpace_NrofCandidates_AggregationLevel2_N0 = 0
	SearchSpace_NrofCandidates_AggregationLevel2_N1 = 1
	SearchSpace_NrofCandidates_AggregationLevel2_N2 = 2
	SearchSpace_NrofCandidates_AggregationLevel2_N3 = 3
	SearchSpace_NrofCandidates_AggregationLevel2_N4 = 4
	SearchSpace_NrofCandidates_AggregationLevel2_N5 = 5
	SearchSpace_NrofCandidates_AggregationLevel2_N6 = 6
	SearchSpace_NrofCandidates_AggregationLevel2_N8 = 7
)

var searchSpaceNrofCandidatesAggregationLevel2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_NrofCandidates_AggregationLevel2_N0, SearchSpace_NrofCandidates_AggregationLevel2_N1, SearchSpace_NrofCandidates_AggregationLevel2_N2, SearchSpace_NrofCandidates_AggregationLevel2_N3, SearchSpace_NrofCandidates_AggregationLevel2_N4, SearchSpace_NrofCandidates_AggregationLevel2_N5, SearchSpace_NrofCandidates_AggregationLevel2_N6, SearchSpace_NrofCandidates_AggregationLevel2_N8},
}

const (
	SearchSpace_NrofCandidates_AggregationLevel4_N0 = 0
	SearchSpace_NrofCandidates_AggregationLevel4_N1 = 1
	SearchSpace_NrofCandidates_AggregationLevel4_N2 = 2
	SearchSpace_NrofCandidates_AggregationLevel4_N3 = 3
	SearchSpace_NrofCandidates_AggregationLevel4_N4 = 4
	SearchSpace_NrofCandidates_AggregationLevel4_N5 = 5
	SearchSpace_NrofCandidates_AggregationLevel4_N6 = 6
	SearchSpace_NrofCandidates_AggregationLevel4_N8 = 7
)

var searchSpaceNrofCandidatesAggregationLevel4Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_NrofCandidates_AggregationLevel4_N0, SearchSpace_NrofCandidates_AggregationLevel4_N1, SearchSpace_NrofCandidates_AggregationLevel4_N2, SearchSpace_NrofCandidates_AggregationLevel4_N3, SearchSpace_NrofCandidates_AggregationLevel4_N4, SearchSpace_NrofCandidates_AggregationLevel4_N5, SearchSpace_NrofCandidates_AggregationLevel4_N6, SearchSpace_NrofCandidates_AggregationLevel4_N8},
}

const (
	SearchSpace_NrofCandidates_AggregationLevel8_N0 = 0
	SearchSpace_NrofCandidates_AggregationLevel8_N1 = 1
	SearchSpace_NrofCandidates_AggregationLevel8_N2 = 2
	SearchSpace_NrofCandidates_AggregationLevel8_N3 = 3
	SearchSpace_NrofCandidates_AggregationLevel8_N4 = 4
	SearchSpace_NrofCandidates_AggregationLevel8_N5 = 5
	SearchSpace_NrofCandidates_AggregationLevel8_N6 = 6
	SearchSpace_NrofCandidates_AggregationLevel8_N8 = 7
)

var searchSpaceNrofCandidatesAggregationLevel8Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_NrofCandidates_AggregationLevel8_N0, SearchSpace_NrofCandidates_AggregationLevel8_N1, SearchSpace_NrofCandidates_AggregationLevel8_N2, SearchSpace_NrofCandidates_AggregationLevel8_N3, SearchSpace_NrofCandidates_AggregationLevel8_N4, SearchSpace_NrofCandidates_AggregationLevel8_N5, SearchSpace_NrofCandidates_AggregationLevel8_N6, SearchSpace_NrofCandidates_AggregationLevel8_N8},
}

const (
	SearchSpace_NrofCandidates_AggregationLevel16_N0 = 0
	SearchSpace_NrofCandidates_AggregationLevel16_N1 = 1
	SearchSpace_NrofCandidates_AggregationLevel16_N2 = 2
	SearchSpace_NrofCandidates_AggregationLevel16_N3 = 3
	SearchSpace_NrofCandidates_AggregationLevel16_N4 = 4
	SearchSpace_NrofCandidates_AggregationLevel16_N5 = 5
	SearchSpace_NrofCandidates_AggregationLevel16_N6 = 6
	SearchSpace_NrofCandidates_AggregationLevel16_N8 = 7
)

var searchSpaceNrofCandidatesAggregationLevel16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_NrofCandidates_AggregationLevel16_N0, SearchSpace_NrofCandidates_AggregationLevel16_N1, SearchSpace_NrofCandidates_AggregationLevel16_N2, SearchSpace_NrofCandidates_AggregationLevel16_N3, SearchSpace_NrofCandidates_AggregationLevel16_N4, SearchSpace_NrofCandidates_AggregationLevel16_N5, SearchSpace_NrofCandidates_AggregationLevel16_N6, SearchSpace_NrofCandidates_AggregationLevel16_N8},
}

var searchSpaceSearchSpaceTypeCommonConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dci-Format0-0-AndFormat1-0", Optional: true},
		{Name: "dci-Format2-0", Optional: true},
		{Name: "dci-Format2-1", Optional: true},
		{Name: "dci-Format2-2", Optional: true},
		{Name: "dci-Format2-3", Optional: true},
	},
}

var searchSpaceSearchSpaceTypeCommonDciFormat00AndFormat10Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceSearchSpaceTypeCommonDciFormat20Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofCandidates-SFI"},
	},
}

var searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "aggregationLevel1", Optional: true},
		{Name: "aggregationLevel2", Optional: true},
		{Name: "aggregationLevel4", Optional: true},
		{Name: "aggregationLevel8", Optional: true},
		{Name: "aggregationLevel16", Optional: true},
	},
}

const (
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel1_N1 = 0
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel1_N2 = 1
)

var searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel1_N1, SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel1_N2},
}

const (
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel2_N1 = 0
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel2_N2 = 1
)

var searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel2_N1, SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel2_N2},
}

const (
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel4_N1 = 0
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel4_N2 = 1
)

var searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel4Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel4_N1, SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel4_N2},
}

const (
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel8_N1 = 0
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel8_N2 = 1
)

var searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel8Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel8_N1, SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel8_N2},
}

const (
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel16_N1 = 0
	SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel16_N2 = 1
)

var searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel16_N1, SearchSpace_SearchSpaceType_Common_Dci_Format2_0_NrofCandidates_SFI_AggregationLevel16_N2},
}

var searchSpaceSearchSpaceTypeCommonDciFormat21Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceSearchSpaceTypeCommonDciFormat22Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceSearchSpaceTypeCommonDciFormat23Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy1", Optional: true},
		{Name: "dummy2"},
	},
}

const (
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl1  = 0
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl2  = 1
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl4  = 2
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl5  = 3
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl8  = 4
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl10 = 5
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl16 = 6
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl20 = 7
)

var searchSpaceSearchSpaceTypeCommonDciFormat23Dummy1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl1, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl2, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl4, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl5, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl8, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl10, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl16, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy1_Sl20},
}

const (
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy2_N1 = 0
	SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy2_N2 = 1
)

var searchSpaceSearchSpaceTypeCommonDciFormat23Dummy2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy2_N1, SearchSpace_SearchSpaceType_Common_Dci_Format2_3_Dummy2_N2},
}

var searchSpaceSearchSpaceTypeUeSpecificConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dci-Formats"},
	},
}

const (
	SearchSpace_SearchSpaceType_Ue_Specific_Dci_Formats_Formats0_0_And_1_0 = 0
	SearchSpace_SearchSpaceType_Ue_Specific_Dci_Formats_Formats0_1_And_1_1 = 1
)

var searchSpaceSearchSpaceTypeUeSpecificDciFormatsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpace_SearchSpaceType_Ue_Specific_Dci_Formats_Formats0_0_And_1_0, SearchSpace_SearchSpaceType_Ue_Specific_Dci_Formats_Formats0_1_And_1_1},
}

type SearchSpace struct {
	SearchSpaceId                      SearchSpaceId
	ControlResourceSetId               *ControlResourceSetId
	MonitoringSlotPeriodicityAndOffset *struct {
		Choice int
		Sl1    *struct{}
		Sl2    *int64
		Sl4    *int64
		Sl5    *int64
		Sl8    *int64
		Sl10   *int64
		Sl16   *int64
		Sl20   *int64
		Sl40   *int64
		Sl80   *int64
		Sl160  *int64
		Sl320  *int64
		Sl640  *int64
		Sl1280 *int64
		Sl2560 *int64
	}
	Duration                    *int64
	MonitoringSymbolsWithinSlot *per.BitString
	NrofCandidates              *struct {
		AggregationLevel1  int64
		AggregationLevel2  int64
		AggregationLevel4  int64
		AggregationLevel8  int64
		AggregationLevel16 int64
	}
	SearchSpaceType *struct {
		Choice int
		Common *struct {
			Dci_Format0_0_AndFormat1_0 *struct{}
			Dci_Format2_0              *struct {
				NrofCandidates_SFI struct {
					AggregationLevel1  *int64
					AggregationLevel2  *int64
					AggregationLevel4  *int64
					AggregationLevel8  *int64
					AggregationLevel16 *int64
				}
			}
			Dci_Format2_1 *struct{}
			Dci_Format2_2 *struct{}
			Dci_Format2_3 *struct {
				Dummy1 *int64
				Dummy2 int64
			}
		}
		Ue_Specific *struct{ Dci_Formats int64 }
	}
}

func (ie *SearchSpace) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ControlResourceSetId != nil, ie.MonitoringSlotPeriodicityAndOffset != nil, ie.Duration != nil, ie.MonitoringSymbolsWithinSlot != nil, ie.NrofCandidates != nil, ie.SearchSpaceType != nil}); err != nil {
		return err
	}

	// 2. searchSpaceId: ref
	{
		if err := ie.SearchSpaceId.Encode(e); err != nil {
			return err
		}
	}

	// 3. controlResourceSetId: ref
	{
		if ie.ControlResourceSetId != nil {
			if err := ie.ControlResourceSetId.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. monitoringSlotPeriodicityAndOffset: choice
	{
		if ie.MonitoringSlotPeriodicityAndOffset != nil {
			choiceEnc := e.NewChoiceEncoder(searchSpaceMonitoringSlotPeriodicityAndOffsetConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MonitoringSlotPeriodicityAndOffset).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MonitoringSlotPeriodicityAndOffset).Choice {
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl1:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl2:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl2), per.Constrained(0, 1)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl4:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl4), per.Constrained(0, 3)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl5:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl5), per.Constrained(0, 4)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl8:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl8), per.Constrained(0, 7)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl10:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl10), per.Constrained(0, 9)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl16:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl16), per.Constrained(0, 15)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl20:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl20), per.Constrained(0, 19)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl40:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl40), per.Constrained(0, 39)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl80:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl80), per.Constrained(0, 79)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl160:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl160), per.Constrained(0, 159)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl320:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl320), per.Constrained(0, 319)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl640:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl640), per.Constrained(0, 639)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl1280:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl1280), per.Constrained(0, 1279)); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl2560:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset).Sl2560), per.Constrained(0, 2559)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MonitoringSlotPeriodicityAndOffset).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. duration: integer
	{
		if ie.Duration != nil {
			if err := e.EncodeInteger(*ie.Duration, searchSpaceDurationConstraints); err != nil {
				return err
			}
		}
	}

	// 6. monitoringSymbolsWithinSlot: bit-string
	{
		if ie.MonitoringSymbolsWithinSlot != nil {
			if err := e.EncodeBitString(*ie.MonitoringSymbolsWithinSlot, searchSpaceMonitoringSymbolsWithinSlotConstraints); err != nil {
				return err
			}
		}
	}

	// 7. nrofCandidates: sequence
	{
		if ie.NrofCandidates != nil {
			c := ie.NrofCandidates
			if err := e.EncodeEnumerated(c.AggregationLevel1, searchSpaceNrofCandidatesAggregationLevel1Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.AggregationLevel2, searchSpaceNrofCandidatesAggregationLevel2Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.AggregationLevel4, searchSpaceNrofCandidatesAggregationLevel4Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.AggregationLevel8, searchSpaceNrofCandidatesAggregationLevel8Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.AggregationLevel16, searchSpaceNrofCandidatesAggregationLevel16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. searchSpaceType: choice
	{
		if ie.SearchSpaceType != nil {
			choiceEnc := e.NewChoiceEncoder(searchSpaceSearchSpaceTypeConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SearchSpaceType).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SearchSpaceType).Choice {
			case SearchSpace_SearchSpaceType_Common:
				c := (*(*ie.SearchSpaceType).Common)
				searchSpaceSearchSpaceTypeCommonSeq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeCommonConstraints)
				if err := searchSpaceSearchSpaceTypeCommonSeq.EncodePreamble([]bool{c.Dci_Format0_0_AndFormat1_0 != nil, c.Dci_Format2_0 != nil, c.Dci_Format2_1 != nil, c.Dci_Format2_2 != nil, c.Dci_Format2_3 != nil}); err != nil {
					return err
				}
				if c.Dci_Format0_0_AndFormat1_0 != nil {
					searchSpaceSearchSpaceTypeCommonDciFormat00AndFormat10Seq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeCommonDciFormat00AndFormat10Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat00AndFormat10Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
				if c.Dci_Format2_0 != nil {
					c := (*c.Dci_Format2_0)
					searchSpaceSearchSpaceTypeCommonDciFormat20Seq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeCommonDciFormat20Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat20Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_SFI
						searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIConstraints)
						if err := searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq.EncodePreamble([]bool{c.AggregationLevel1 != nil, c.AggregationLevel2 != nil, c.AggregationLevel4 != nil, c.AggregationLevel8 != nil, c.AggregationLevel16 != nil}); err != nil {
							return err
						}
						if c.AggregationLevel1 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel1), searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel1Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel2 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel2), searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel2Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel4 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel4), searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel4Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel8 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel8), searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel8Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel16), searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel16Constraints); err != nil {
								return err
							}
						}
					}
				}
				if c.Dci_Format2_1 != nil {
					searchSpaceSearchSpaceTypeCommonDciFormat21Seq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeCommonDciFormat21Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat21Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
				if c.Dci_Format2_2 != nil {
					searchSpaceSearchSpaceTypeCommonDciFormat22Seq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeCommonDciFormat22Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat22Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
				if c.Dci_Format2_3 != nil {
					c := (*c.Dci_Format2_3)
					searchSpaceSearchSpaceTypeCommonDciFormat23Seq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeCommonDciFormat23Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat23Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					if err := searchSpaceSearchSpaceTypeCommonDciFormat23Seq.EncodePreamble([]bool{c.Dummy1 != nil}); err != nil {
						return err
					}
					if c.Dummy1 != nil {
						if err := e.EncodeEnumerated((*c.Dummy1), searchSpaceSearchSpaceTypeCommonDciFormat23Dummy1Constraints); err != nil {
							return err
						}
					}
					if err := e.EncodeEnumerated(c.Dummy2, searchSpaceSearchSpaceTypeCommonDciFormat23Dummy2Constraints); err != nil {
						return err
					}
				}
			case SearchSpace_SearchSpaceType_Ue_Specific:
				c := (*(*ie.SearchSpaceType).Ue_Specific)
				searchSpaceSearchSpaceTypeUeSpecificSeq := e.NewSequenceEncoder(searchSpaceSearchSpaceTypeUeSpecificConstraints)
				if err := searchSpaceSearchSpaceTypeUeSpecificSeq.EncodeExtensionBit(false); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.Dci_Formats, searchSpaceSearchSpaceTypeUeSpecificDciFormatsConstraints); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SearchSpaceType).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SearchSpace) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. searchSpaceId: ref
	{
		if err := ie.SearchSpaceId.Decode(d); err != nil {
			return err
		}
	}

	// 3. controlResourceSetId: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ControlResourceSetId = new(ControlResourceSetId)
			if err := ie.ControlResourceSetId.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. monitoringSlotPeriodicityAndOffset: choice
	{
		if seq.IsComponentPresent(2) {
			ie.MonitoringSlotPeriodicityAndOffset = &struct {
				Choice int
				Sl1    *struct{}
				Sl2    *int64
				Sl4    *int64
				Sl5    *int64
				Sl8    *int64
				Sl10   *int64
				Sl16   *int64
				Sl20   *int64
				Sl40   *int64
				Sl80   *int64
				Sl160  *int64
				Sl320  *int64
				Sl640  *int64
				Sl1280 *int64
				Sl2560 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(searchSpaceMonitoringSlotPeriodicityAndOffsetConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MonitoringSlotPeriodicityAndOffset).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl1:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl1 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl2:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl2 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl2) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl4:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl4 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl4) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl5:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl5 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl5) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl8:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl8 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl8) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl10:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl10 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 9))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl10) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl16:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl16) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl20:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl20 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 19))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl20) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl40:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl40 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 39))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl40) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl80:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl80 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 79))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl80) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl160:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl160 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 159))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl160) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl320:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl320 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 319))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl320) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl640:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl640 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 639))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl640) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl1280:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl1280 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1279))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl1280) = v
			case SearchSpace_MonitoringSlotPeriodicityAndOffset_Sl2560:
				(*ie.MonitoringSlotPeriodicityAndOffset).Sl2560 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 2559))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset).Sl2560) = v
			}
		}
	}

	// 5. duration: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(searchSpaceDurationConstraints)
			if err != nil {
				return err
			}
			ie.Duration = &v
		}
	}

	// 6. monitoringSymbolsWithinSlot: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(searchSpaceMonitoringSymbolsWithinSlotConstraints)
			if err != nil {
				return err
			}
			ie.MonitoringSymbolsWithinSlot = &v
		}
	}

	// 7. nrofCandidates: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.NrofCandidates = &struct {
				AggregationLevel1  int64
				AggregationLevel2  int64
				AggregationLevel4  int64
				AggregationLevel8  int64
				AggregationLevel16 int64
			}{}
			c := ie.NrofCandidates
			{
				v, err := d.DecodeEnumerated(searchSpaceNrofCandidatesAggregationLevel1Constraints)
				if err != nil {
					return err
				}
				c.AggregationLevel1 = v
			}
			{
				v, err := d.DecodeEnumerated(searchSpaceNrofCandidatesAggregationLevel2Constraints)
				if err != nil {
					return err
				}
				c.AggregationLevel2 = v
			}
			{
				v, err := d.DecodeEnumerated(searchSpaceNrofCandidatesAggregationLevel4Constraints)
				if err != nil {
					return err
				}
				c.AggregationLevel4 = v
			}
			{
				v, err := d.DecodeEnumerated(searchSpaceNrofCandidatesAggregationLevel8Constraints)
				if err != nil {
					return err
				}
				c.AggregationLevel8 = v
			}
			{
				v, err := d.DecodeEnumerated(searchSpaceNrofCandidatesAggregationLevel16Constraints)
				if err != nil {
					return err
				}
				c.AggregationLevel16 = v
			}
		}
	}

	// 8. searchSpaceType: choice
	{
		if seq.IsComponentPresent(6) {
			ie.SearchSpaceType = &struct {
				Choice int
				Common *struct {
					Dci_Format0_0_AndFormat1_0 *struct{}
					Dci_Format2_0              *struct {
						NrofCandidates_SFI struct {
							AggregationLevel1  *int64
							AggregationLevel2  *int64
							AggregationLevel4  *int64
							AggregationLevel8  *int64
							AggregationLevel16 *int64
						}
					}
					Dci_Format2_1 *struct{}
					Dci_Format2_2 *struct{}
					Dci_Format2_3 *struct {
						Dummy1 *int64
						Dummy2 int64
					}
				}
				Ue_Specific *struct{ Dci_Formats int64 }
			}{}
			choiceDec := d.NewChoiceDecoder(searchSpaceSearchSpaceTypeConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SearchSpaceType).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SearchSpace_SearchSpaceType_Common:
				(*ie.SearchSpaceType).Common = &struct {
					Dci_Format0_0_AndFormat1_0 *struct{}
					Dci_Format2_0              *struct {
						NrofCandidates_SFI struct {
							AggregationLevel1  *int64
							AggregationLevel2  *int64
							AggregationLevel4  *int64
							AggregationLevel8  *int64
							AggregationLevel16 *int64
						}
					}
					Dci_Format2_1 *struct{}
					Dci_Format2_2 *struct{}
					Dci_Format2_3 *struct {
						Dummy1 *int64
						Dummy2 int64
					}
				}{}
				c := (*(*ie.SearchSpaceType).Common)
				searchSpaceSearchSpaceTypeCommonSeq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeCommonConstraints)
				if err := searchSpaceSearchSpaceTypeCommonSeq.DecodePreamble(); err != nil {
					return err
				}
				if searchSpaceSearchSpaceTypeCommonSeq.IsComponentPresent(0) {
					c.Dci_Format0_0_AndFormat1_0 = &struct{}{}
					searchSpaceSearchSpaceTypeCommonDciFormat00AndFormat10Seq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeCommonDciFormat00AndFormat10Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat00AndFormat10Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
				if searchSpaceSearchSpaceTypeCommonSeq.IsComponentPresent(1) {
					c.Dci_Format2_0 = &struct {
						NrofCandidates_SFI struct {
							AggregationLevel1  *int64
							AggregationLevel2  *int64
							AggregationLevel4  *int64
							AggregationLevel8  *int64
							AggregationLevel16 *int64
						}
					}{}
					c := (*c.Dci_Format2_0)
					searchSpaceSearchSpaceTypeCommonDciFormat20Seq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeCommonDciFormat20Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat20Seq.DecodeExtensionBit(); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_SFI
						searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIConstraints)
						if err := searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq.DecodePreamble(); err != nil {
							return err
						}
						if searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq.IsComponentPresent(0) {
							c.AggregationLevel1 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel1Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel1) = v
						}
						if searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq.IsComponentPresent(1) {
							c.AggregationLevel2 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel2Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel2) = v
						}
						if searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq.IsComponentPresent(2) {
							c.AggregationLevel4 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel4Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel4) = v
						}
						if searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq.IsComponentPresent(3) {
							c.AggregationLevel8 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel8Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel8) = v
						}
						if searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFISeq.IsComponentPresent(4) {
							c.AggregationLevel16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeCommonDciFormat20NrofCandidatesSFIAggregationLevel16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel16) = v
						}
					}
				}
				if searchSpaceSearchSpaceTypeCommonSeq.IsComponentPresent(2) {
					c.Dci_Format2_1 = &struct{}{}
					searchSpaceSearchSpaceTypeCommonDciFormat21Seq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeCommonDciFormat21Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat21Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
				if searchSpaceSearchSpaceTypeCommonSeq.IsComponentPresent(3) {
					c.Dci_Format2_2 = &struct{}{}
					searchSpaceSearchSpaceTypeCommonDciFormat22Seq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeCommonDciFormat22Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat22Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
				if searchSpaceSearchSpaceTypeCommonSeq.IsComponentPresent(4) {
					c.Dci_Format2_3 = &struct {
						Dummy1 *int64
						Dummy2 int64
					}{}
					c := (*c.Dci_Format2_3)
					searchSpaceSearchSpaceTypeCommonDciFormat23Seq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeCommonDciFormat23Constraints)
					if err := searchSpaceSearchSpaceTypeCommonDciFormat23Seq.DecodeExtensionBit(); err != nil {
						return err
					}
					if err := searchSpaceSearchSpaceTypeCommonDciFormat23Seq.DecodePreamble(); err != nil {
						return err
					}
					if searchSpaceSearchSpaceTypeCommonDciFormat23Seq.IsComponentPresent(0) {
						c.Dummy1 = new(int64)
						v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeCommonDciFormat23Dummy1Constraints)
						if err != nil {
							return err
						}
						(*c.Dummy1) = v
					}
					{
						v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeCommonDciFormat23Dummy2Constraints)
						if err != nil {
							return err
						}
						c.Dummy2 = v
					}
				}
			case SearchSpace_SearchSpaceType_Ue_Specific:
				(*ie.SearchSpaceType).Ue_Specific = &struct{ Dci_Formats int64 }{}
				c := (*(*ie.SearchSpaceType).Ue_Specific)
				searchSpaceSearchSpaceTypeUeSpecificSeq := d.NewSequenceDecoder(searchSpaceSearchSpaceTypeUeSpecificConstraints)
				if err := searchSpaceSearchSpaceTypeUeSpecificSeq.DecodeExtensionBit(); err != nil {
					return err
				}
				{
					v, err := d.DecodeEnumerated(searchSpaceSearchSpaceTypeUeSpecificDciFormatsConstraints)
					if err != nil {
						return err
					}
					c.Dci_Formats = v
				}
			}
		}
	}

	return nil
}
