// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-CSI-ReportConfigId-r18 (line 8909).
// LTM-CSI-ReportConfigId-r18 ::=            INTEGER (0..maxNrofLTM-CSI-ReportConfigurations-1-r18)

var lTMCSIReportConfigIdR18Constraints = per.Constrained(0, common.MaxNrofLTM_CSI_ReportConfigurations_1_r18)

type LTM_CSI_ReportConfigId_r18 struct {
	Value int64
}

func (v *LTM_CSI_ReportConfigId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, lTMCSIReportConfigIdR18Constraints)
}

func (v *LTM_CSI_ReportConfigId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(lTMCSIReportConfigIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
