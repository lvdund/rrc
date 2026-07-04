// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PCI-RangeIndex (line 10979).
// PCI-RangeIndex ::=                  INTEGER (1..maxNrofPCI-Ranges)

var pCIRangeIndexConstraints = per.Constrained(1, common.MaxNrofPCI_Ranges)

type PCI_RangeIndex struct {
	Value int64
}

func (v *PCI_RangeIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pCIRangeIndexConstraints)
}

func (v *PCI_RangeIndex) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pCIRangeIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
