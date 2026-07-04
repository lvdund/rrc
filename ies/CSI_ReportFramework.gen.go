// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-ReportFramework (line 22483).

var cSIReportFrameworkConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberPeriodicCSI-PerBWP-ForCSI-Report"},
		{Name: "maxNumberAperiodicCSI-PerBWP-ForCSI-Report"},
		{Name: "maxNumberSemiPersistentCSI-PerBWP-ForCSI-Report"},
		{Name: "maxNumberPeriodicCSI-PerBWP-ForBeamReport"},
		{Name: "maxNumberAperiodicCSI-PerBWP-ForBeamReport"},
		{Name: "maxNumberAperiodicCSI-triggeringStatePerCC"},
		{Name: "maxNumberSemiPersistentCSI-PerBWP-ForBeamReport"},
		{Name: "simultaneousCSI-ReportsPerCC"},
	},
}

var cSIReportFrameworkMaxNumberPeriodicCSIPerBWPForCSIReportConstraints = per.Constrained(1, 4)

var cSIReportFrameworkMaxNumberAperiodicCSIPerBWPForCSIReportConstraints = per.Constrained(1, 4)

var cSIReportFrameworkMaxNumberSemiPersistentCSIPerBWPForCSIReportConstraints = per.Constrained(0, 4)

var cSIReportFrameworkMaxNumberPeriodicCSIPerBWPForBeamReportConstraints = per.Constrained(1, 4)

var cSIReportFrameworkMaxNumberAperiodicCSIPerBWPForBeamReportConstraints = per.Constrained(1, 4)

const (
	CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N3   = 0
	CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N7   = 1
	CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N15  = 2
	CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N31  = 3
	CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N63  = 4
	CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N128 = 5
)

var cSIReportFrameworkMaxNumberAperiodicCSITriggeringStatePerCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N3, CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N7, CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N15, CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N31, CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N63, CSI_ReportFramework_MaxNumberAperiodicCSI_TriggeringStatePerCC_N128},
}

var cSIReportFrameworkMaxNumberSemiPersistentCSIPerBWPForBeamReportConstraints = per.Constrained(0, 4)

var cSIReportFrameworkSimultaneousCSIReportsPerCCConstraints = per.Constrained(1, 8)

type CSI_ReportFramework struct {
	MaxNumberPeriodicCSI_PerBWP_ForCSI_Report       int64
	MaxNumberAperiodicCSI_PerBWP_ForCSI_Report      int64
	MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report int64
	MaxNumberPeriodicCSI_PerBWP_ForBeamReport       int64
	MaxNumberAperiodicCSI_PerBWP_ForBeamReport      int64
	MaxNumberAperiodicCSI_TriggeringStatePerCC      int64
	MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport int64
	SimultaneousCSI_ReportsPerCC                    int64
}

func (ie *CSI_ReportFramework) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIReportFrameworkConstraints)
	_ = seq

	// 1. maxNumberPeriodicCSI-PerBWP-ForCSI-Report: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberPeriodicCSI_PerBWP_ForCSI_Report, cSIReportFrameworkMaxNumberPeriodicCSIPerBWPForCSIReportConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberAperiodicCSI-PerBWP-ForCSI-Report: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_Report, cSIReportFrameworkMaxNumberAperiodicCSIPerBWPForCSIReportConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberSemiPersistentCSI-PerBWP-ForCSI-Report: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report, cSIReportFrameworkMaxNumberSemiPersistentCSIPerBWPForCSIReportConstraints); err != nil {
			return err
		}
	}

	// 4. maxNumberPeriodicCSI-PerBWP-ForBeamReport: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberPeriodicCSI_PerBWP_ForBeamReport, cSIReportFrameworkMaxNumberPeriodicCSIPerBWPForBeamReportConstraints); err != nil {
			return err
		}
	}

	// 5. maxNumberAperiodicCSI-PerBWP-ForBeamReport: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberAperiodicCSI_PerBWP_ForBeamReport, cSIReportFrameworkMaxNumberAperiodicCSIPerBWPForBeamReportConstraints); err != nil {
			return err
		}
	}

	// 6. maxNumberAperiodicCSI-triggeringStatePerCC: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberAperiodicCSI_TriggeringStatePerCC, cSIReportFrameworkMaxNumberAperiodicCSITriggeringStatePerCCConstraints); err != nil {
			return err
		}
	}

	// 7. maxNumberSemiPersistentCSI-PerBWP-ForBeamReport: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport, cSIReportFrameworkMaxNumberSemiPersistentCSIPerBWPForBeamReportConstraints); err != nil {
			return err
		}
	}

	// 8. simultaneousCSI-ReportsPerCC: integer
	{
		if err := e.EncodeInteger(ie.SimultaneousCSI_ReportsPerCC, cSIReportFrameworkSimultaneousCSIReportsPerCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_ReportFramework) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIReportFrameworkConstraints)
	_ = seq

	// 1. maxNumberPeriodicCSI-PerBWP-ForCSI-Report: integer
	{
		v0, err := d.DecodeInteger(cSIReportFrameworkMaxNumberPeriodicCSIPerBWPForCSIReportConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicCSI_PerBWP_ForCSI_Report = v0
	}

	// 2. maxNumberAperiodicCSI-PerBWP-ForCSI-Report: integer
	{
		v1, err := d.DecodeInteger(cSIReportFrameworkMaxNumberAperiodicCSIPerBWPForCSIReportConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_Report = v1
	}

	// 3. maxNumberSemiPersistentCSI-PerBWP-ForCSI-Report: integer
	{
		v2, err := d.DecodeInteger(cSIReportFrameworkMaxNumberSemiPersistentCSIPerBWPForCSIReportConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report = v2
	}

	// 4. maxNumberPeriodicCSI-PerBWP-ForBeamReport: integer
	{
		v3, err := d.DecodeInteger(cSIReportFrameworkMaxNumberPeriodicCSIPerBWPForBeamReportConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicCSI_PerBWP_ForBeamReport = v3
	}

	// 5. maxNumberAperiodicCSI-PerBWP-ForBeamReport: integer
	{
		v4, err := d.DecodeInteger(cSIReportFrameworkMaxNumberAperiodicCSIPerBWPForBeamReportConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicCSI_PerBWP_ForBeamReport = v4
	}

	// 6. maxNumberAperiodicCSI-triggeringStatePerCC: enumerated
	{
		v5, err := d.DecodeEnumerated(cSIReportFrameworkMaxNumberAperiodicCSITriggeringStatePerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicCSI_TriggeringStatePerCC = v5
	}

	// 7. maxNumberSemiPersistentCSI-PerBWP-ForBeamReport: integer
	{
		v6, err := d.DecodeInteger(cSIReportFrameworkMaxNumberSemiPersistentCSIPerBWPForBeamReportConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport = v6
	}

	// 8. simultaneousCSI-ReportsPerCC: integer
	{
		v7, err := d.DecodeInteger(cSIReportFrameworkSimultaneousCSIReportsPerCCConstraints)
		if err != nil {
			return err
		}
		ie.SimultaneousCSI_ReportsPerCC = v7
	}

	return nil
}
