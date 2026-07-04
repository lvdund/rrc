// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ReportSubConfigId-r18 (line 7453).
// CSI-ReportSubConfigId-r18 ::=              INTEGER (0..maxNrofCSI-ReportSubconfigPerCSI-ReportConfig-1-r18)

var cSIReportSubConfigIdR18Constraints = per.Constrained(0, common.MaxNrofCSI_ReportSubconfigPerCSI_ReportConfig_1_r18)

type CSI_ReportSubConfigId_r18 struct {
	Value int64
}

func (v *CSI_ReportSubConfigId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSIReportSubConfigIdR18Constraints)
}

func (v *CSI_ReportSubConfigId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSIReportSubConfigIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
