// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ReportConfigId-r16 (line 27880).
// SL-ReportConfigId-r16 ::=             INTEGER (1..maxNrofSL-ReportConfigId-r16)

var sLReportConfigIdR16Constraints = per.Constrained(1, common.MaxNrofSL_ReportConfigId_r16)

type SL_ReportConfigId_r16 struct {
	Value int64
}

func (v *SL_ReportConfigId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLReportConfigIdR16Constraints)
}

func (v *SL_ReportConfigId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLReportConfigIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
