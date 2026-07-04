// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CLI-RSSI-MeasResourceId-r19 (line 6001).
// CLI-RSSI-MeasResourceId-r19 ::=      INTEGER(0..maxNrofCLI-RSSI-MeasResources-1-r19)

var cLIRSSIMeasResourceIdR19Constraints = per.Constrained(0, common.MaxNrofCLI_RSSI_MeasResources_1_r19)

type CLI_RSSI_MeasResourceId_r19 struct {
	Value int64
}

func (v *CLI_RSSI_MeasResourceId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cLIRSSIMeasResourceIdR19Constraints)
}

func (v *CLI_RSSI_MeasResourceId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cLIRSSIMeasResourceIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
