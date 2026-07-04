// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDCP-r18 (line 7297).

var tDCPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "delayDSetofLengthY-r18"},
		{Name: "phaseReporting-r18", Optional: true},
	},
}

var tDCPR18DelayDSetofLengthYR18Constraints = per.SizeRange(1, common.MaxNrofdelayD_r18)

const (
	TDCP_r18_PhaseReporting_r18_Enable = 0
)

var tDCPR18PhaseReportingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TDCP_r18_PhaseReporting_r18_Enable},
}

type TDCP_r18 struct {
	DelayDSetofLengthY_r18 []DelayD
	PhaseReporting_r18     *int64
}

func (ie *TDCP_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDCPR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PhaseReporting_r18 != nil}); err != nil {
		return err
	}

	// 2. delayDSetofLengthY-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(tDCPR18DelayDSetofLengthYR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.DelayDSetofLengthY_r18))); err != nil {
			return err
		}
		for i := range ie.DelayDSetofLengthY_r18 {
			if err := ie.DelayDSetofLengthY_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. phaseReporting-r18: enumerated
	{
		if ie.PhaseReporting_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PhaseReporting_r18, tDCPR18PhaseReportingR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TDCP_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDCPR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. delayDSetofLengthY-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(tDCPR18DelayDSetofLengthYR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.DelayDSetofLengthY_r18 = make([]DelayD, n)
		for i := int64(0); i < n; i++ {
			if err := ie.DelayDSetofLengthY_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. phaseReporting-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(tDCPR18PhaseReportingR18Constraints)
			if err != nil {
				return err
			}
			ie.PhaseReporting_r18 = &idx
		}
	}

	return nil
}
