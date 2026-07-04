// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ClockQualityMetrics-r18 (line 6027).

var clockQualityMetricsR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "synchronisationState-r18", Optional: true},
		{Name: "tracebilityToUTC-r18", Optional: true},
		{Name: "tracebilityToGNSS-r18", Optional: true},
		{Name: "clockFrequencyStability-r18", Optional: true},
		{Name: "clockAccuracy-r18", Optional: true},
		{Name: "parentTimeSource-r18", Optional: true},
	},
}

const (
	ClockQualityMetrics_r18_SynchronisationState_r18_Locked   = 0
	ClockQualityMetrics_r18_SynchronisationState_r18_Holdover = 1
	ClockQualityMetrics_r18_SynchronisationState_r18_Freerun  = 2
	ClockQualityMetrics_r18_SynchronisationState_r18_Spare1   = 3
)

var clockQualityMetricsR18SynchronisationStateR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ClockQualityMetrics_r18_SynchronisationState_r18_Locked, ClockQualityMetrics_r18_SynchronisationState_r18_Holdover, ClockQualityMetrics_r18_SynchronisationState_r18_Freerun, ClockQualityMetrics_r18_SynchronisationState_r18_Spare1},
}

var clockQualityMetricsR18ClockFrequencyStabilityR18Constraints = per.FixedSize(16)

var clockQualityMetrics_r18ClockAccuracyR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "value"},
		{Name: "index"},
	},
}

const (
	ClockQualityMetrics_r18_ClockAccuracy_r18_Value = 0
	ClockQualityMetrics_r18_ClockAccuracy_r18_Index = 1
)

const (
	ClockQualityMetrics_r18_ParentTimeSource_r18_SyncE           = 0
	ClockQualityMetrics_r18_ParentTimeSource_r18_PTP             = 1
	ClockQualityMetrics_r18_ParentTimeSource_r18_GNSS            = 2
	ClockQualityMetrics_r18_ParentTimeSource_r18_AtomicClock     = 3
	ClockQualityMetrics_r18_ParentTimeSource_r18_TerrestialRadio = 4
	ClockQualityMetrics_r18_ParentTimeSource_r18_SerialTimeCode  = 5
	ClockQualityMetrics_r18_ParentTimeSource_r18_NTP             = 6
	ClockQualityMetrics_r18_ParentTimeSource_r18_Handset         = 7
	ClockQualityMetrics_r18_ParentTimeSource_r18_Other           = 8
	ClockQualityMetrics_r18_ParentTimeSource_r18_Spare7          = 9
	ClockQualityMetrics_r18_ParentTimeSource_r18_Spare6          = 10
	ClockQualityMetrics_r18_ParentTimeSource_r18_Spare5          = 11
	ClockQualityMetrics_r18_ParentTimeSource_r18_Spare4          = 12
	ClockQualityMetrics_r18_ParentTimeSource_r18_Spare3          = 13
	ClockQualityMetrics_r18_ParentTimeSource_r18_Spare2          = 14
	ClockQualityMetrics_r18_ParentTimeSource_r18_Spare1          = 15
)

var clockQualityMetricsR18ParentTimeSourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ClockQualityMetrics_r18_ParentTimeSource_r18_SyncE, ClockQualityMetrics_r18_ParentTimeSource_r18_PTP, ClockQualityMetrics_r18_ParentTimeSource_r18_GNSS, ClockQualityMetrics_r18_ParentTimeSource_r18_AtomicClock, ClockQualityMetrics_r18_ParentTimeSource_r18_TerrestialRadio, ClockQualityMetrics_r18_ParentTimeSource_r18_SerialTimeCode, ClockQualityMetrics_r18_ParentTimeSource_r18_NTP, ClockQualityMetrics_r18_ParentTimeSource_r18_Handset, ClockQualityMetrics_r18_ParentTimeSource_r18_Other, ClockQualityMetrics_r18_ParentTimeSource_r18_Spare7, ClockQualityMetrics_r18_ParentTimeSource_r18_Spare6, ClockQualityMetrics_r18_ParentTimeSource_r18_Spare5, ClockQualityMetrics_r18_ParentTimeSource_r18_Spare4, ClockQualityMetrics_r18_ParentTimeSource_r18_Spare3, ClockQualityMetrics_r18_ParentTimeSource_r18_Spare2, ClockQualityMetrics_r18_ParentTimeSource_r18_Spare1},
}

type ClockQualityMetrics_r18 struct {
	SynchronisationState_r18    *int64
	TracebilityToUTC_r18        *bool
	TracebilityToGNSS_r18       *bool
	ClockFrequencyStability_r18 *per.BitString
	ClockAccuracy_r18           *struct {
		Choice int
		Value  *int64
		Index  *int64
	}
	ParentTimeSource_r18 *int64
}

func (ie *ClockQualityMetrics_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(clockQualityMetricsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SynchronisationState_r18 != nil, ie.TracebilityToUTC_r18 != nil, ie.TracebilityToGNSS_r18 != nil, ie.ClockFrequencyStability_r18 != nil, ie.ClockAccuracy_r18 != nil, ie.ParentTimeSource_r18 != nil}); err != nil {
		return err
	}

	// 3. synchronisationState-r18: enumerated
	{
		if ie.SynchronisationState_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SynchronisationState_r18, clockQualityMetricsR18SynchronisationStateR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. tracebilityToUTC-r18: boolean
	{
		if ie.TracebilityToUTC_r18 != nil {
			if err := e.EncodeBoolean(*ie.TracebilityToUTC_r18); err != nil {
				return err
			}
		}
	}

	// 5. tracebilityToGNSS-r18: boolean
	{
		if ie.TracebilityToGNSS_r18 != nil {
			if err := e.EncodeBoolean(*ie.TracebilityToGNSS_r18); err != nil {
				return err
			}
		}
	}

	// 6. clockFrequencyStability-r18: bit-string
	{
		if ie.ClockFrequencyStability_r18 != nil {
			if err := e.EncodeBitString(*ie.ClockFrequencyStability_r18, clockQualityMetricsR18ClockFrequencyStabilityR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. clockAccuracy-r18: choice
	{
		if ie.ClockAccuracy_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(clockQualityMetrics_r18ClockAccuracyR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ClockAccuracy_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ClockAccuracy_r18).Choice {
			case ClockQualityMetrics_r18_ClockAccuracy_r18_Value:
				if err := e.EncodeInteger((*(*ie.ClockAccuracy_r18).Value), per.Constrained(1, 40000000)); err != nil {
					return err
				}
			case ClockQualityMetrics_r18_ClockAccuracy_r18_Index:
				if err := e.EncodeInteger((*(*ie.ClockAccuracy_r18).Index), per.Constrained(32, 47)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ClockAccuracy_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. parentTimeSource-r18: enumerated
	{
		if ie.ParentTimeSource_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ParentTimeSource_r18, clockQualityMetricsR18ParentTimeSourceR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ClockQualityMetrics_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(clockQualityMetricsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. synchronisationState-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(clockQualityMetricsR18SynchronisationStateR18Constraints)
			if err != nil {
				return err
			}
			ie.SynchronisationState_r18 = &idx
		}
	}

	// 4. tracebilityToUTC-r18: boolean
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.TracebilityToUTC_r18 = &v
		}
	}

	// 5. tracebilityToGNSS-r18: boolean
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.TracebilityToGNSS_r18 = &v
		}
	}

	// 6. clockFrequencyStability-r18: bit-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeBitString(clockQualityMetricsR18ClockFrequencyStabilityR18Constraints)
			if err != nil {
				return err
			}
			ie.ClockFrequencyStability_r18 = &v
		}
	}

	// 7. clockAccuracy-r18: choice
	{
		if seq.IsComponentPresent(4) {
			ie.ClockAccuracy_r18 = &struct {
				Choice int
				Value  *int64
				Index  *int64
			}{}
			choiceDec := d.NewChoiceDecoder(clockQualityMetrics_r18ClockAccuracyR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ClockAccuracy_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ClockQualityMetrics_r18_ClockAccuracy_r18_Value:
				(*ie.ClockAccuracy_r18).Value = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 40000000))
				if err != nil {
					return err
				}
				(*(*ie.ClockAccuracy_r18).Value) = v
			case ClockQualityMetrics_r18_ClockAccuracy_r18_Index:
				(*ie.ClockAccuracy_r18).Index = new(int64)
				v, err := d.DecodeInteger(per.Constrained(32, 47))
				if err != nil {
					return err
				}
				(*(*ie.ClockAccuracy_r18).Index) = v
			}
		}
	}

	// 8. parentTimeSource-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(clockQualityMetricsR18ParentTimeSourceR18Constraints)
			if err != nil {
				return err
			}
			ie.ParentTimeSource_r18 = &idx
		}
	}

	return nil
}
