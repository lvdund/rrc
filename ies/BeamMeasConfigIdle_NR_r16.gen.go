// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BeamMeasConfigIdle-NR-r16 (line 9315).

var beamMeasConfigIdleNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportQuantityRS-Indexes-r16"},
		{Name: "maxNrofRS-IndexesToReport-r16"},
		{Name: "includeBeamMeasurements-r16"},
	},
}

const (
	BeamMeasConfigIdle_NR_r16_ReportQuantityRS_Indexes_r16_Rsrp = 0
	BeamMeasConfigIdle_NR_r16_ReportQuantityRS_Indexes_r16_Rsrq = 1
	BeamMeasConfigIdle_NR_r16_ReportQuantityRS_Indexes_r16_Both = 2
)

var beamMeasConfigIdleNRR16ReportQuantityRSIndexesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamMeasConfigIdle_NR_r16_ReportQuantityRS_Indexes_r16_Rsrp, BeamMeasConfigIdle_NR_r16_ReportQuantityRS_Indexes_r16_Rsrq, BeamMeasConfigIdle_NR_r16_ReportQuantityRS_Indexes_r16_Both},
}

var beamMeasConfigIdleNRR16MaxNrofRSIndexesToReportR16Constraints = per.Constrained(1, common.MaxNrofIndexesToReport)

type BeamMeasConfigIdle_NR_r16 struct {
	ReportQuantityRS_Indexes_r16  int64
	MaxNrofRS_IndexesToReport_r16 int64
	IncludeBeamMeasurements_r16   bool
}

func (ie *BeamMeasConfigIdle_NR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamMeasConfigIdleNRR16Constraints)
	_ = seq

	// 1. reportQuantityRS-Indexes-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportQuantityRS_Indexes_r16, beamMeasConfigIdleNRR16ReportQuantityRSIndexesR16Constraints); err != nil {
			return err
		}
	}

	// 2. maxNrofRS-IndexesToReport-r16: integer
	{
		if err := e.EncodeInteger(ie.MaxNrofRS_IndexesToReport_r16, beamMeasConfigIdleNRR16MaxNrofRSIndexesToReportR16Constraints); err != nil {
			return err
		}
	}

	// 3. includeBeamMeasurements-r16: boolean
	{
		if err := e.EncodeBoolean(ie.IncludeBeamMeasurements_r16); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BeamMeasConfigIdle_NR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamMeasConfigIdleNRR16Constraints)
	_ = seq

	// 1. reportQuantityRS-Indexes-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(beamMeasConfigIdleNRR16ReportQuantityRSIndexesR16Constraints)
		if err != nil {
			return err
		}
		ie.ReportQuantityRS_Indexes_r16 = v0
	}

	// 2. maxNrofRS-IndexesToReport-r16: integer
	{
		v1, err := d.DecodeInteger(beamMeasConfigIdleNRR16MaxNrofRSIndexesToReportR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNrofRS_IndexesToReport_r16 = v1
	}

	// 3. includeBeamMeasurements-r16: boolean
	{
		v2, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.IncludeBeamMeasurements_r16 = v2
	}

	return nil
}
