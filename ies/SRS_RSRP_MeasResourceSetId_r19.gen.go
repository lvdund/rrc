// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-RSRP-MeasResourceSetId-r19 (line 15821).
// SRS-RSRP-MeasResourceSetId-r19 ::=      INTEGER(0..maxNrofSRS-RSRP-MeasResourceSets-1-r19)

var sRSRSRPMeasResourceSetIdR19Constraints = per.Constrained(0, common.MaxNrofSRS_RSRP_MeasResourceSets_1_r19)

type SRS_RSRP_MeasResourceSetId_r19 struct {
	Value int64
}

func (v *SRS_RSRP_MeasResourceSetId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSRSRPMeasResourceSetIdR19Constraints)
}

func (v *SRS_RSRP_MeasResourceSetId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSRSRPMeasResourceSetIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
