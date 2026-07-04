// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ReportedFeatureCombination-r18 (line 3144).

var reportedFeatureCombinationR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "redCap-r18", Optional: true},
		{Name: "smallData-r18", Optional: true},
		{Name: "nsag-r18", Optional: true},
		{Name: "msg3-Repetitions-r18", Optional: true},
		{Name: "msg1-Repetitions-r18", Optional: true},
		{Name: "eRedCap-r18", Optional: true},
		{Name: "triggered-S-NSSAI-List-r18", Optional: true},
	},
}

const (
	ReportedFeatureCombination_r18_RedCap_r18_True = 0
)

var reportedFeatureCombinationR18RedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportedFeatureCombination_r18_RedCap_r18_True},
}

const (
	ReportedFeatureCombination_r18_SmallData_r18_True = 0
)

var reportedFeatureCombinationR18SmallDataR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportedFeatureCombination_r18_SmallData_r18_True},
}

const (
	ReportedFeatureCombination_r18_Msg3_Repetitions_r18_True = 0
)

var reportedFeatureCombinationR18Msg3RepetitionsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportedFeatureCombination_r18_Msg3_Repetitions_r18_True},
}

const (
	ReportedFeatureCombination_r18_Msg1_Repetitions_r18_True = 0
)

var reportedFeatureCombinationR18Msg1RepetitionsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportedFeatureCombination_r18_Msg1_Repetitions_r18_True},
}

const (
	ReportedFeatureCombination_r18_ERedCap_r18_True = 0
)

var reportedFeatureCombinationR18ERedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportedFeatureCombination_r18_ERedCap_r18_True},
}

var reportedFeatureCombinationR18TriggeredSNSSAIListR18Constraints = per.SizeRange(1, common.MaxNrofS_NSSAI)

type ReportedFeatureCombination_r18 struct {
	RedCap_r18                 *int64
	SmallData_r18              *int64
	Nsag_r18                   *NSAG_List_r17
	Msg3_Repetitions_r18       *int64
	Msg1_Repetitions_r18       *int64
	ERedCap_r18                *int64
	Triggered_S_NSSAI_List_r18 []S_NSSAI
}

func (ie *ReportedFeatureCombination_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportedFeatureCombinationR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RedCap_r18 != nil, ie.SmallData_r18 != nil, ie.Nsag_r18 != nil, ie.Msg3_Repetitions_r18 != nil, ie.Msg1_Repetitions_r18 != nil, ie.ERedCap_r18 != nil, ie.Triggered_S_NSSAI_List_r18 != nil}); err != nil {
		return err
	}

	// 2. redCap-r18: enumerated
	{
		if ie.RedCap_r18 != nil {
			if err := e.EncodeEnumerated(*ie.RedCap_r18, reportedFeatureCombinationR18RedCapR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. smallData-r18: enumerated
	{
		if ie.SmallData_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SmallData_r18, reportedFeatureCombinationR18SmallDataR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. nsag-r18: ref
	{
		if ie.Nsag_r18 != nil {
			if err := ie.Nsag_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. msg3-Repetitions-r18: enumerated
	{
		if ie.Msg3_Repetitions_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Msg3_Repetitions_r18, reportedFeatureCombinationR18Msg3RepetitionsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. msg1-Repetitions-r18: enumerated
	{
		if ie.Msg1_Repetitions_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Msg1_Repetitions_r18, reportedFeatureCombinationR18Msg1RepetitionsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. eRedCap-r18: enumerated
	{
		if ie.ERedCap_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ERedCap_r18, reportedFeatureCombinationR18ERedCapR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. triggered-S-NSSAI-List-r18: sequence-of
	{
		if ie.Triggered_S_NSSAI_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(reportedFeatureCombinationR18TriggeredSNSSAIListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Triggered_S_NSSAI_List_r18))); err != nil {
				return err
			}
			for i := range ie.Triggered_S_NSSAI_List_r18 {
				if err := ie.Triggered_S_NSSAI_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *ReportedFeatureCombination_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportedFeatureCombinationR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. redCap-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(reportedFeatureCombinationR18RedCapR18Constraints)
			if err != nil {
				return err
			}
			ie.RedCap_r18 = &idx
		}
	}

	// 3. smallData-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(reportedFeatureCombinationR18SmallDataR18Constraints)
			if err != nil {
				return err
			}
			ie.SmallData_r18 = &idx
		}
	}

	// 4. nsag-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Nsag_r18 = new(NSAG_List_r17)
			if err := ie.Nsag_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. msg3-Repetitions-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(reportedFeatureCombinationR18Msg3RepetitionsR18Constraints)
			if err != nil {
				return err
			}
			ie.Msg3_Repetitions_r18 = &idx
		}
	}

	// 6. msg1-Repetitions-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(reportedFeatureCombinationR18Msg1RepetitionsR18Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_Repetitions_r18 = &idx
		}
	}

	// 7. eRedCap-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(reportedFeatureCombinationR18ERedCapR18Constraints)
			if err != nil {
				return err
			}
			ie.ERedCap_r18 = &idx
		}
	}

	// 8. triggered-S-NSSAI-List-r18: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(reportedFeatureCombinationR18TriggeredSNSSAIListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Triggered_S_NSSAI_List_r18 = make([]S_NSSAI, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Triggered_S_NSSAI_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
