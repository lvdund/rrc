// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RadioLinkMonitoringRS-Id (line 13248).
// RadioLinkMonitoringRS-Id ::=            INTEGER (0..maxNrofFailureDetectionResources-1)

var radioLinkMonitoringRSIdConstraints = per.Constrained(0, common.MaxNrofFailureDetectionResources_1)

type RadioLinkMonitoringRS_Id struct {
	Value int64
}

func (v *RadioLinkMonitoringRS_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, radioLinkMonitoringRSIdConstraints)
}

func (v *RadioLinkMonitoringRS_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(radioLinkMonitoringRSIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
