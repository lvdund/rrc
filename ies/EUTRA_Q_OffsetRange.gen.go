// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-Q-OffsetRange (line 26207).

const (
	EUTRA_Q_OffsetRange_DB_24 = 0
	EUTRA_Q_OffsetRange_DB_22 = 1
	EUTRA_Q_OffsetRange_DB_20 = 2
	EUTRA_Q_OffsetRange_DB_18 = 3
	EUTRA_Q_OffsetRange_DB_16 = 4
	EUTRA_Q_OffsetRange_DB_14 = 5
	EUTRA_Q_OffsetRange_DB_12 = 6
	EUTRA_Q_OffsetRange_DB_10 = 7
	EUTRA_Q_OffsetRange_DB_8  = 8
	EUTRA_Q_OffsetRange_DB_6  = 9
	EUTRA_Q_OffsetRange_DB_5  = 10
	EUTRA_Q_OffsetRange_DB_4  = 11
	EUTRA_Q_OffsetRange_DB_3  = 12
	EUTRA_Q_OffsetRange_DB_2  = 13
	EUTRA_Q_OffsetRange_DB_1  = 14
	EUTRA_Q_OffsetRange_DB0   = 15
	EUTRA_Q_OffsetRange_DB1   = 16
	EUTRA_Q_OffsetRange_DB2   = 17
	EUTRA_Q_OffsetRange_DB3   = 18
	EUTRA_Q_OffsetRange_DB4   = 19
	EUTRA_Q_OffsetRange_DB5   = 20
	EUTRA_Q_OffsetRange_DB6   = 21
	EUTRA_Q_OffsetRange_DB8   = 22
	EUTRA_Q_OffsetRange_DB10  = 23
	EUTRA_Q_OffsetRange_DB12  = 24
	EUTRA_Q_OffsetRange_DB14  = 25
	EUTRA_Q_OffsetRange_DB16  = 26
	EUTRA_Q_OffsetRange_DB18  = 27
	EUTRA_Q_OffsetRange_DB20  = 28
	EUTRA_Q_OffsetRange_DB22  = 29
	EUTRA_Q_OffsetRange_DB24  = 30
)

var eUTRAQOffsetRangeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_Q_OffsetRange_DB_24, EUTRA_Q_OffsetRange_DB_22, EUTRA_Q_OffsetRange_DB_20, EUTRA_Q_OffsetRange_DB_18, EUTRA_Q_OffsetRange_DB_16, EUTRA_Q_OffsetRange_DB_14, EUTRA_Q_OffsetRange_DB_12, EUTRA_Q_OffsetRange_DB_10, EUTRA_Q_OffsetRange_DB_8, EUTRA_Q_OffsetRange_DB_6, EUTRA_Q_OffsetRange_DB_5, EUTRA_Q_OffsetRange_DB_4, EUTRA_Q_OffsetRange_DB_3, EUTRA_Q_OffsetRange_DB_2, EUTRA_Q_OffsetRange_DB_1, EUTRA_Q_OffsetRange_DB0, EUTRA_Q_OffsetRange_DB1, EUTRA_Q_OffsetRange_DB2, EUTRA_Q_OffsetRange_DB3, EUTRA_Q_OffsetRange_DB4, EUTRA_Q_OffsetRange_DB5, EUTRA_Q_OffsetRange_DB6, EUTRA_Q_OffsetRange_DB8, EUTRA_Q_OffsetRange_DB10, EUTRA_Q_OffsetRange_DB12, EUTRA_Q_OffsetRange_DB14, EUTRA_Q_OffsetRange_DB16, EUTRA_Q_OffsetRange_DB18, EUTRA_Q_OffsetRange_DB20, EUTRA_Q_OffsetRange_DB22, EUTRA_Q_OffsetRange_DB24},
}

type EUTRA_Q_OffsetRange struct {
	Value int64
}

func (v *EUTRA_Q_OffsetRange) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, eUTRAQOffsetRangeConstraints)
}

func (v *EUTRA_Q_OffsetRange) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(eUTRAQOffsetRangeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
