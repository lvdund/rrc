// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-LoggedMeasurementConfigId-r19 (line 7005).
// CSI-LoggedMeasurementConfigId-r19 ::=            INTEGER (0..maxNrofLoggedMeasurementConfigurations-1-r19)

var cSILoggedMeasurementConfigIdR19Constraints = per.Constrained(0, common.MaxNrofLoggedMeasurementConfigurations_1_r19)

type CSI_LoggedMeasurementConfigId_r19 struct {
	Value int64
}

func (v *CSI_LoggedMeasurementConfigId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cSILoggedMeasurementConfigIdR19Constraints)
}

func (v *CSI_LoggedMeasurementConfigId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cSILoggedMeasurementConfigIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
