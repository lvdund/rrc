// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-RSRP-MeasResourceId-r19 (line 15805).
// SRS-RSRP-MeasResourceId-r19 ::=      INTEGER(0..maxNrofSRS-RSRP-MeasResources-1-r19)

var sRSRSRPMeasResourceIdR19Constraints = per.Constrained(0, common.MaxNrofSRS_RSRP_MeasResources_1_r19)

type SRS_RSRP_MeasResourceId_r19 struct {
	Value int64
}

func (v *SRS_RSRP_MeasResourceId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSRSRPMeasResourceIdR19Constraints)
}

func (v *SRS_RSRP_MeasResourceId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSRSRPMeasResourceIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
