// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ReportConfigId (line 13370).
// ReportConfigId ::=                          INTEGER (1..maxReportConfigId)

var reportConfigIdConstraints = per.Constrained(1, common.MaxReportConfigId)

type ReportConfigId struct {
	Value int64
}

func (v *ReportConfigId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, reportConfigIdConstraints)
}

func (v *ReportConfigId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(reportConfigIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
