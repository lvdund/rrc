// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ReportConfigId (line 7432).
// CSI-ReportConfigId ::=              INTEGER (0..maxNrofCSI-ReportConfigurations-1)

var cSIReportConfigIdConstraints = per.Constrained(0, common.MaxNrofCSI_ReportConfigurations_1)

type CSI_ReportConfigId struct {
	Value int64
}

func (v *CSI_ReportConfigId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSIReportConfigIdConstraints)
}

func (v *CSI_ReportConfigId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSIReportConfigIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
