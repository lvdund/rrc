// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DataCollectionCandidateConfigId-r19 (line 7633).
// DataCollectionCandidateConfigId-r19 ::=            INTEGER (0..maxNrofDataCollectionCandidateConfigs-1-r19)

var dataCollectionCandidateConfigIdR19Constraints = per.Constrained(0, common.MaxNrofDataCollectionCandidateConfigs_1_r19)

type DataCollectionCandidateConfigId_r19 struct {
	Value int64
}

func (v *DataCollectionCandidateConfigId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, dataCollectionCandidateConfigIdR19Constraints)
}

func (v *DataCollectionCandidateConfigId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(dataCollectionCandidateConfigIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
