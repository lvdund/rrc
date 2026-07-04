// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DSR-ReportingThreshold-r19 (line 9085).
// DSR-ReportingThreshold-r19 ::= INTEGER (1..64)

var dSRReportingThresholdR19Constraints = per.Constrained(1, 64)

type DSR_ReportingThreshold_r19 struct {
	Value int64
}

func (v *DSR_ReportingThreshold_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, dSRReportingThresholdR19Constraints)
}

func (v *DSR_ReportingThreshold_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(dSRReportingThresholdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
