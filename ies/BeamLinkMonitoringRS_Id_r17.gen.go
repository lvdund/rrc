// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BeamLinkMonitoringRS-Id-r17 (line 13243).
// BeamLinkMonitoringRS-Id-r17 ::=     INTEGER (0..maxNrofFailureDetectionResources-1-r17)

var beamLinkMonitoringRSIdR17Constraints = per.Constrained(0, common.MaxNrofFailureDetectionResources_1_r17)

type BeamLinkMonitoringRS_Id_r17 struct {
	Value int64
}

func (v *BeamLinkMonitoringRS_Id_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, beamLinkMonitoringRSIdR17Constraints)
}

func (v *BeamLinkMonitoringRS_Id_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(beamLinkMonitoringRSIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
