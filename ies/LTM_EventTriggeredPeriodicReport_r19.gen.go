// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-EventTriggeredPeriodicReport-r19 (line 8869).

var lTMEventTriggeredPeriodicReportR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportInterval-r19"},
		{Name: "reportAmount-r19"},
	},
}

const (
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms20    = 0
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms60    = 1
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms120   = 2
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms240   = 3
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms480   = 4
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms640   = 5
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms1024  = 6
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms2048  = 7
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms5120  = 8
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms10240 = 9
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms20480 = 10
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms40960 = 11
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min1    = 12
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min6    = 13
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min12   = 14
	LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min30   = 15
)

var lTMEventTriggeredPeriodicReportR19ReportIntervalR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms20, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms60, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms120, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms240, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms480, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms640, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms1024, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms2048, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms5120, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms10240, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms20480, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Ms40960, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min1, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min6, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min12, LTM_EventTriggeredPeriodicReport_r19_ReportInterval_r19_Min30},
}

const (
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r2       = 0
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r4       = 1
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r8       = 2
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r16      = 3
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r32      = 4
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r64      = 5
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_Infinity = 6
	LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_Spare1   = 7
)

var lTMEventTriggeredPeriodicReportR19ReportAmountR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r2, LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r4, LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r8, LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r16, LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r32, LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_r64, LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_Infinity, LTM_EventTriggeredPeriodicReport_r19_ReportAmount_r19_Spare1},
}

type LTM_EventTriggeredPeriodicReport_r19 struct {
	ReportInterval_r19 int64
	ReportAmount_r19   int64
}

func (ie *LTM_EventTriggeredPeriodicReport_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMEventTriggeredPeriodicReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. reportInterval-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportInterval_r19, lTMEventTriggeredPeriodicReportR19ReportIntervalR19Constraints); err != nil {
			return err
		}
	}

	// 3. reportAmount-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount_r19, lTMEventTriggeredPeriodicReportR19ReportAmountR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LTM_EventTriggeredPeriodicReport_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMEventTriggeredPeriodicReportR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. reportInterval-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(lTMEventTriggeredPeriodicReportR19ReportIntervalR19Constraints)
		if err != nil {
			return err
		}
		ie.ReportInterval_r19 = v0
	}

	// 3. reportAmount-r19: enumerated
	{
		v1, err := d.DecodeEnumerated(lTMEventTriggeredPeriodicReportR19ReportAmountR19Constraints)
		if err != nil {
			return err
		}
		ie.ReportAmount_r19 = v1
	}

	return nil
}
