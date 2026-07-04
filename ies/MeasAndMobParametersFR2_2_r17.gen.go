// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersFR2-2-r17 (line 21360).

var measAndMobParametersFR22R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "handoverInterF-r17", Optional: true},
		{Name: "handoverLTE-EPC-r17", Optional: true},
		{Name: "handoverLTE-5GC-r17", Optional: true},
		{Name: "idleInactiveNR-MeasReport-r17", Optional: true},
	},
}

const (
	MeasAndMobParametersFR2_2_r17_HandoverInterF_r17_Supported = 0
)

var measAndMobParametersFR22R17HandoverInterFR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFR2_2_r17_HandoverInterF_r17_Supported},
}

const (
	MeasAndMobParametersFR2_2_r17_HandoverLTE_EPC_r17_Supported = 0
)

var measAndMobParametersFR22R17HandoverLTEEPCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFR2_2_r17_HandoverLTE_EPC_r17_Supported},
}

const (
	MeasAndMobParametersFR2_2_r17_HandoverLTE_5GC_r17_Supported = 0
)

var measAndMobParametersFR22R17HandoverLTE5GCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFR2_2_r17_HandoverLTE_5GC_r17_Supported},
}

const (
	MeasAndMobParametersFR2_2_r17_IdleInactiveNR_MeasReport_r17_Supported = 0
)

var measAndMobParametersFR22R17IdleInactiveNRMeasReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFR2_2_r17_IdleInactiveNR_MeasReport_r17_Supported},
}

type MeasAndMobParametersFR2_2_r17 struct {
	HandoverInterF_r17            *int64
	HandoverLTE_EPC_r17           *int64
	HandoverLTE_5GC_r17           *int64
	IdleInactiveNR_MeasReport_r17 *int64
}

func (ie *MeasAndMobParametersFR2_2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.HandoverInterF_r17 != nil, ie.HandoverLTE_EPC_r17 != nil, ie.HandoverLTE_5GC_r17 != nil, ie.IdleInactiveNR_MeasReport_r17 != nil}); err != nil {
		return err
	}

	// 3. handoverInterF-r17: enumerated
	{
		if ie.HandoverInterF_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HandoverInterF_r17, measAndMobParametersFR22R17HandoverInterFR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. handoverLTE-EPC-r17: enumerated
	{
		if ie.HandoverLTE_EPC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HandoverLTE_EPC_r17, measAndMobParametersFR22R17HandoverLTEEPCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. handoverLTE-5GC-r17: enumerated
	{
		if ie.HandoverLTE_5GC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HandoverLTE_5GC_r17, measAndMobParametersFR22R17HandoverLTE5GCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. idleInactiveNR-MeasReport-r17: enumerated
	{
		if ie.IdleInactiveNR_MeasReport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.IdleInactiveNR_MeasReport_r17, measAndMobParametersFR22R17IdleInactiveNRMeasReportR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersFR2_2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersFR22R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. handoverInterF-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFR22R17HandoverInterFR17Constraints)
			if err != nil {
				return err
			}
			ie.HandoverInterF_r17 = &idx
		}
	}

	// 4. handoverLTE-EPC-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFR22R17HandoverLTEEPCR17Constraints)
			if err != nil {
				return err
			}
			ie.HandoverLTE_EPC_r17 = &idx
		}
	}

	// 5. handoverLTE-5GC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFR22R17HandoverLTE5GCR17Constraints)
			if err != nil {
				return err
			}
			ie.HandoverLTE_5GC_r17 = &idx
		}
	}

	// 6. idleInactiveNR-MeasReport-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFR22R17IdleInactiveNRMeasReportR17Constraints)
			if err != nil {
				return err
			}
			ie.IdleInactiveNR_MeasReport_r17 = &idx
		}
	}

	return nil
}
