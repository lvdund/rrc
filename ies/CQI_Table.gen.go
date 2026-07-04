// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CQI-Table (line 6811).
// CQI-Table ::=                   ENUMERATED {table1, table2, table3, table4-r17}

const (
	CQI_Table_Table1     = 0
	CQI_Table_Table2     = 1
	CQI_Table_Table3     = 2
	CQI_Table_Table4_r17 = 3
)

var cQITableConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CQI_Table_Table1, CQI_Table_Table2, CQI_Table_Table3, CQI_Table_Table4_r17},
}

type CQI_Table struct {
	Value int64
}

func (v *CQI_Table) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, cQITableConstraints)
}

func (v *CQI_Table) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(cQITableConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
