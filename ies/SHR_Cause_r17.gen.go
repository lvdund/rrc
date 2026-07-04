// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SHR-Cause-r17 (line 3470).

var sHRCauseR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "t304-cause-r17", Optional: true},
		{Name: "t310-cause-r17", Optional: true},
		{Name: "t312-cause-r17", Optional: true},
		{Name: "sourceDAPS-Failure-r17", Optional: true},
	},
}

const (
	SHR_Cause_r17_T304_Cause_r17_True = 0
)

var sHRCauseR17T304CauseR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SHR_Cause_r17_T304_Cause_r17_True},
}

const (
	SHR_Cause_r17_T310_Cause_r17_True = 0
)

var sHRCauseR17T310CauseR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SHR_Cause_r17_T310_Cause_r17_True},
}

const (
	SHR_Cause_r17_T312_Cause_r17_True = 0
)

var sHRCauseR17T312CauseR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SHR_Cause_r17_T312_Cause_r17_True},
}

const (
	SHR_Cause_r17_SourceDAPS_Failure_r17_True = 0
)

var sHRCauseR17SourceDAPSFailureR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SHR_Cause_r17_SourceDAPS_Failure_r17_True},
}

type SHR_Cause_r17 struct {
	T304_Cause_r17         *int64
	T310_Cause_r17         *int64
	T312_Cause_r17         *int64
	SourceDAPS_Failure_r17 *int64
}

func (ie *SHR_Cause_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sHRCauseR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T304_Cause_r17 != nil, ie.T310_Cause_r17 != nil, ie.T312_Cause_r17 != nil, ie.SourceDAPS_Failure_r17 != nil}); err != nil {
		return err
	}

	// 3. t304-cause-r17: enumerated
	{
		if ie.T304_Cause_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T304_Cause_r17, sHRCauseR17T304CauseR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. t310-cause-r17: enumerated
	{
		if ie.T310_Cause_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T310_Cause_r17, sHRCauseR17T310CauseR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. t312-cause-r17: enumerated
	{
		if ie.T312_Cause_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T312_Cause_r17, sHRCauseR17T312CauseR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sourceDAPS-Failure-r17: enumerated
	{
		if ie.SourceDAPS_Failure_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SourceDAPS_Failure_r17, sHRCauseR17SourceDAPSFailureR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SHR_Cause_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sHRCauseR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. t304-cause-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sHRCauseR17T304CauseR17Constraints)
			if err != nil {
				return err
			}
			ie.T304_Cause_r17 = &idx
		}
	}

	// 4. t310-cause-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sHRCauseR17T310CauseR17Constraints)
			if err != nil {
				return err
			}
			ie.T310_Cause_r17 = &idx
		}
	}

	// 5. t312-cause-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sHRCauseR17T312CauseR17Constraints)
			if err != nil {
				return err
			}
			ie.T312_Cause_r17 = &idx
		}
	}

	// 6. sourceDAPS-Failure-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sHRCauseR17SourceDAPSFailureR17Constraints)
			if err != nil {
				return err
			}
			ie.SourceDAPS_Failure_r17 = &idx
		}
	}

	return nil
}
