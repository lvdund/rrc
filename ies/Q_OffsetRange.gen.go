// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Q-OffsetRange (line 12754).

const (
	Q_OffsetRange_DB_24 = 0
	Q_OffsetRange_DB_22 = 1
	Q_OffsetRange_DB_20 = 2
	Q_OffsetRange_DB_18 = 3
	Q_OffsetRange_DB_16 = 4
	Q_OffsetRange_DB_14 = 5
	Q_OffsetRange_DB_12 = 6
	Q_OffsetRange_DB_10 = 7
	Q_OffsetRange_DB_8  = 8
	Q_OffsetRange_DB_6  = 9
	Q_OffsetRange_DB_5  = 10
	Q_OffsetRange_DB_4  = 11
	Q_OffsetRange_DB_3  = 12
	Q_OffsetRange_DB_2  = 13
	Q_OffsetRange_DB_1  = 14
	Q_OffsetRange_DB0   = 15
	Q_OffsetRange_DB1   = 16
	Q_OffsetRange_DB2   = 17
	Q_OffsetRange_DB3   = 18
	Q_OffsetRange_DB4   = 19
	Q_OffsetRange_DB5   = 20
	Q_OffsetRange_DB6   = 21
	Q_OffsetRange_DB8   = 22
	Q_OffsetRange_DB10  = 23
	Q_OffsetRange_DB12  = 24
	Q_OffsetRange_DB14  = 25
	Q_OffsetRange_DB16  = 26
	Q_OffsetRange_DB18  = 27
	Q_OffsetRange_DB20  = 28
	Q_OffsetRange_DB22  = 29
	Q_OffsetRange_DB24  = 30
)

var qOffsetRangeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Q_OffsetRange_DB_24, Q_OffsetRange_DB_22, Q_OffsetRange_DB_20, Q_OffsetRange_DB_18, Q_OffsetRange_DB_16, Q_OffsetRange_DB_14, Q_OffsetRange_DB_12, Q_OffsetRange_DB_10, Q_OffsetRange_DB_8, Q_OffsetRange_DB_6, Q_OffsetRange_DB_5, Q_OffsetRange_DB_4, Q_OffsetRange_DB_3, Q_OffsetRange_DB_2, Q_OffsetRange_DB_1, Q_OffsetRange_DB0, Q_OffsetRange_DB1, Q_OffsetRange_DB2, Q_OffsetRange_DB3, Q_OffsetRange_DB4, Q_OffsetRange_DB5, Q_OffsetRange_DB6, Q_OffsetRange_DB8, Q_OffsetRange_DB10, Q_OffsetRange_DB12, Q_OffsetRange_DB14, Q_OffsetRange_DB16, Q_OffsetRange_DB18, Q_OffsetRange_DB20, Q_OffsetRange_DB22, Q_OffsetRange_DB24},
}

type Q_OffsetRange struct {
	Value int64
}

func (v *Q_OffsetRange) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, qOffsetRangeConstraints)
}

func (v *Q_OffsetRange) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(qOffsetRangeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
