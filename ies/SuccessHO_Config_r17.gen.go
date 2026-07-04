// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SuccessHO-Config-r17 (line 26406).

var successHOConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdPercentageT304-r17", Optional: true},
		{Name: "thresholdPercentageT310-r17", Optional: true},
		{Name: "thresholdPercentageT312-r17", Optional: true},
		{Name: "sourceDAPS-FailureReporting-r17", Optional: true},
	},
}

const (
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_P40    = 0
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_P60    = 1
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_P80    = 2
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare5 = 3
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare4 = 4
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare3 = 5
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare2 = 6
	SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare1 = 7
)

var successHOConfigR17ThresholdPercentageT304R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessHO_Config_r17_ThresholdPercentageT304_r17_P40, SuccessHO_Config_r17_ThresholdPercentageT304_r17_P60, SuccessHO_Config_r17_ThresholdPercentageT304_r17_P80, SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare5, SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare4, SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare3, SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare2, SuccessHO_Config_r17_ThresholdPercentageT304_r17_Spare1},
}

const (
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_P40    = 0
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_P60    = 1
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_P80    = 2
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare5 = 3
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare4 = 4
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare3 = 5
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare2 = 6
	SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare1 = 7
)

var successHOConfigR17ThresholdPercentageT310R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessHO_Config_r17_ThresholdPercentageT310_r17_P40, SuccessHO_Config_r17_ThresholdPercentageT310_r17_P60, SuccessHO_Config_r17_ThresholdPercentageT310_r17_P80, SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare5, SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare4, SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare3, SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare2, SuccessHO_Config_r17_ThresholdPercentageT310_r17_Spare1},
}

const (
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_P20    = 0
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_P40    = 1
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_P60    = 2
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_P80    = 3
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare4 = 4
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare3 = 5
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare2 = 6
	SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare1 = 7
)

var successHOConfigR17ThresholdPercentageT312R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessHO_Config_r17_ThresholdPercentageT312_r17_P20, SuccessHO_Config_r17_ThresholdPercentageT312_r17_P40, SuccessHO_Config_r17_ThresholdPercentageT312_r17_P60, SuccessHO_Config_r17_ThresholdPercentageT312_r17_P80, SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare4, SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare3, SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare2, SuccessHO_Config_r17_ThresholdPercentageT312_r17_Spare1},
}

const (
	SuccessHO_Config_r17_SourceDAPS_FailureReporting_r17_True = 0
)

var successHOConfigR17SourceDAPSFailureReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuccessHO_Config_r17_SourceDAPS_FailureReporting_r17_True},
}

type SuccessHO_Config_r17 struct {
	ThresholdPercentageT304_r17     *int64
	ThresholdPercentageT310_r17     *int64
	ThresholdPercentageT312_r17     *int64
	SourceDAPS_FailureReporting_r17 *int64
}

func (ie *SuccessHO_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(successHOConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ThresholdPercentageT304_r17 != nil, ie.ThresholdPercentageT310_r17 != nil, ie.ThresholdPercentageT312_r17 != nil, ie.SourceDAPS_FailureReporting_r17 != nil}); err != nil {
		return err
	}

	// 3. thresholdPercentageT304-r17: enumerated
	{
		if ie.ThresholdPercentageT304_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ThresholdPercentageT304_r17, successHOConfigR17ThresholdPercentageT304R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. thresholdPercentageT310-r17: enumerated
	{
		if ie.ThresholdPercentageT310_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ThresholdPercentageT310_r17, successHOConfigR17ThresholdPercentageT310R17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. thresholdPercentageT312-r17: enumerated
	{
		if ie.ThresholdPercentageT312_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ThresholdPercentageT312_r17, successHOConfigR17ThresholdPercentageT312R17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sourceDAPS-FailureReporting-r17: enumerated
	{
		if ie.SourceDAPS_FailureReporting_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SourceDAPS_FailureReporting_r17, successHOConfigR17SourceDAPSFailureReportingR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SuccessHO_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(successHOConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. thresholdPercentageT304-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(successHOConfigR17ThresholdPercentageT304R17Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdPercentageT304_r17 = &idx
		}
	}

	// 4. thresholdPercentageT310-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(successHOConfigR17ThresholdPercentageT310R17Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdPercentageT310_r17 = &idx
		}
	}

	// 5. thresholdPercentageT312-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(successHOConfigR17ThresholdPercentageT312R17Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdPercentageT312_r17 = &idx
		}
	}

	// 6. sourceDAPS-FailureReporting-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(successHOConfigR17SourceDAPSFailureReportingR17Constraints)
			if err != nil {
				return err
			}
			ie.SourceDAPS_FailureReporting_r17 = &idx
		}
	}

	return nil
}
