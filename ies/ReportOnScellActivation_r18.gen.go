// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ReportOnScellActivation-r18 (line 13914).

var reportOnScellActivationR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsType-r18"},
		{Name: "reportQuantityRS-Indexes-r18"},
		{Name: "maxNrofRS-IndexesToReport-r18"},
		{Name: "includeBeamMeasurements-r18"},
	},
}

var reportOnScellActivationR18MaxNrofRSIndexesToReportR18Constraints = per.Constrained(1, common.MaxNrofIndexesToReport)

type ReportOnScellActivation_r18 struct {
	RsType_r18                    NR_RS_Type
	ReportQuantityRS_Indexes_r18  MeasReportQuantity
	MaxNrofRS_IndexesToReport_r18 int64
	IncludeBeamMeasurements_r18   bool
}

func (ie *ReportOnScellActivation_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportOnScellActivationR18Constraints)
	_ = seq

	// 1. rsType-r18: ref
	{
		if err := ie.RsType_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. reportQuantityRS-Indexes-r18: ref
	{
		if err := ie.ReportQuantityRS_Indexes_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. maxNrofRS-IndexesToReport-r18: integer
	{
		if err := e.EncodeInteger(ie.MaxNrofRS_IndexesToReport_r18, reportOnScellActivationR18MaxNrofRSIndexesToReportR18Constraints); err != nil {
			return err
		}
	}

	// 4. includeBeamMeasurements-r18: boolean
	{
		if err := e.EncodeBoolean(ie.IncludeBeamMeasurements_r18); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReportOnScellActivation_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportOnScellActivationR18Constraints)
	_ = seq

	// 1. rsType-r18: ref
	{
		if err := ie.RsType_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. reportQuantityRS-Indexes-r18: ref
	{
		if err := ie.ReportQuantityRS_Indexes_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. maxNrofRS-IndexesToReport-r18: integer
	{
		v2, err := d.DecodeInteger(reportOnScellActivationR18MaxNrofRSIndexesToReportR18Constraints)
		if err != nil {
			return err
		}
		ie.MaxNrofRS_IndexesToReport_r18 = v2
	}

	// 4. includeBeamMeasurements-r18: boolean
	{
		v3, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.IncludeBeamMeasurements_r18 = v3
	}

	return nil
}
