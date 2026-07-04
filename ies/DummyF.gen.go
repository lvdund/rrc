// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyF (line 20488).

var dummyFConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberPeriodicCSI-ReportPerBWP"},
		{Name: "maxNumberAperiodicCSI-ReportPerBWP"},
		{Name: "maxNumberSemiPersistentCSI-ReportPerBWP"},
		{Name: "simultaneousCSI-ReportsAllCC"},
	},
}

var dummyFMaxNumberPeriodicCSIReportPerBWPConstraints = per.Constrained(1, 4)

var dummyFMaxNumberAperiodicCSIReportPerBWPConstraints = per.Constrained(1, 4)

var dummyFMaxNumberSemiPersistentCSIReportPerBWPConstraints = per.Constrained(0, 4)

var dummyFSimultaneousCSIReportsAllCCConstraints = per.Constrained(5, 32)

type DummyF struct {
	MaxNumberPeriodicCSI_ReportPerBWP       int64
	MaxNumberAperiodicCSI_ReportPerBWP      int64
	MaxNumberSemiPersistentCSI_ReportPerBWP int64
	SimultaneousCSI_ReportsAllCC            int64
}

func (ie *DummyF) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyFConstraints)
	_ = seq

	// 1. maxNumberPeriodicCSI-ReportPerBWP: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberPeriodicCSI_ReportPerBWP, dummyFMaxNumberPeriodicCSIReportPerBWPConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberAperiodicCSI-ReportPerBWP: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberAperiodicCSI_ReportPerBWP, dummyFMaxNumberAperiodicCSIReportPerBWPConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberSemiPersistentCSI-ReportPerBWP: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberSemiPersistentCSI_ReportPerBWP, dummyFMaxNumberSemiPersistentCSIReportPerBWPConstraints); err != nil {
			return err
		}
	}

	// 4. simultaneousCSI-ReportsAllCC: integer
	{
		if err := e.EncodeInteger(ie.SimultaneousCSI_ReportsAllCC, dummyFSimultaneousCSIReportsAllCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyF) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyFConstraints)
	_ = seq

	// 1. maxNumberPeriodicCSI-ReportPerBWP: integer
	{
		v0, err := d.DecodeInteger(dummyFMaxNumberPeriodicCSIReportPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPeriodicCSI_ReportPerBWP = v0
	}

	// 2. maxNumberAperiodicCSI-ReportPerBWP: integer
	{
		v1, err := d.DecodeInteger(dummyFMaxNumberAperiodicCSIReportPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicCSI_ReportPerBWP = v1
	}

	// 3. maxNumberSemiPersistentCSI-ReportPerBWP: integer
	{
		v2, err := d.DecodeInteger(dummyFMaxNumberSemiPersistentCSIReportPerBWPConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSemiPersistentCSI_ReportPerBWP = v2
	}

	// 4. simultaneousCSI-ReportsAllCC: integer
	{
		v3, err := d.DecodeInteger(dummyFSimultaneousCSIReportsAllCCConstraints)
		if err != nil {
			return err
		}
		ie.SimultaneousCSI_ReportsAllCC = v3
	}

	return nil
}
