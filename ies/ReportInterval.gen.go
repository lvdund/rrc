// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReportInterval (line 13987).

const (
	ReportInterval_Ms120   = 0
	ReportInterval_Ms240   = 1
	ReportInterval_Ms480   = 2
	ReportInterval_Ms640   = 3
	ReportInterval_Ms1024  = 4
	ReportInterval_Ms2048  = 5
	ReportInterval_Ms5120  = 6
	ReportInterval_Ms10240 = 7
	ReportInterval_Ms20480 = 8
	ReportInterval_Ms40960 = 9
	ReportInterval_Min1    = 10
	ReportInterval_Min6    = 11
	ReportInterval_Min12   = 12
	ReportInterval_Min30   = 13
)

var reportIntervalConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportInterval_Ms120, ReportInterval_Ms240, ReportInterval_Ms480, ReportInterval_Ms640, ReportInterval_Ms1024, ReportInterval_Ms2048, ReportInterval_Ms5120, ReportInterval_Ms10240, ReportInterval_Ms20480, ReportInterval_Ms40960, ReportInterval_Min1, ReportInterval_Min6, ReportInterval_Min12, ReportInterval_Min30},
}

type ReportInterval struct {
	Value int64
}

func (v *ReportInterval) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, reportIntervalConstraints)
}

func (v *ReportInterval) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(reportIntervalConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
