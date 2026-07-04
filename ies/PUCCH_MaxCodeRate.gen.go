// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-MaxCodeRate (line 12061).
// PUCCH-MaxCodeRate ::=                   ENUMERATED {zeroDot08, zeroDot15, zeroDot25, zeroDot35, zeroDot45, zeroDot60, zeroDot80}

const (
	PUCCH_MaxCodeRate_ZeroDot08 = 0
	PUCCH_MaxCodeRate_ZeroDot15 = 1
	PUCCH_MaxCodeRate_ZeroDot25 = 2
	PUCCH_MaxCodeRate_ZeroDot35 = 3
	PUCCH_MaxCodeRate_ZeroDot45 = 4
	PUCCH_MaxCodeRate_ZeroDot60 = 5
	PUCCH_MaxCodeRate_ZeroDot80 = 6
)

var pUCCHMaxCodeRateConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_MaxCodeRate_ZeroDot08, PUCCH_MaxCodeRate_ZeroDot15, PUCCH_MaxCodeRate_ZeroDot25, PUCCH_MaxCodeRate_ZeroDot35, PUCCH_MaxCodeRate_ZeroDot45, PUCCH_MaxCodeRate_ZeroDot60, PUCCH_MaxCodeRate_ZeroDot80},
}

type PUCCH_MaxCodeRate struct {
	Value int64
}

func (v *PUCCH_MaxCodeRate) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, pUCCHMaxCodeRateConstraints)
}

func (v *PUCCH_MaxCodeRate) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(pUCCHMaxCodeRateConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
