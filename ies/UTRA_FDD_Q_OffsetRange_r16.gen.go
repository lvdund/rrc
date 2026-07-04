// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UTRA-FDD-Q-OffsetRange-r16 (line 26687).

const (
	UTRA_FDD_Q_OffsetRange_r16_DB_24 = 0
	UTRA_FDD_Q_OffsetRange_r16_DB_22 = 1
	UTRA_FDD_Q_OffsetRange_r16_DB_20 = 2
	UTRA_FDD_Q_OffsetRange_r16_DB_18 = 3
	UTRA_FDD_Q_OffsetRange_r16_DB_16 = 4
	UTRA_FDD_Q_OffsetRange_r16_DB_14 = 5
	UTRA_FDD_Q_OffsetRange_r16_DB_12 = 6
	UTRA_FDD_Q_OffsetRange_r16_DB_10 = 7
	UTRA_FDD_Q_OffsetRange_r16_DB_8  = 8
	UTRA_FDD_Q_OffsetRange_r16_DB_6  = 9
	UTRA_FDD_Q_OffsetRange_r16_DB_5  = 10
	UTRA_FDD_Q_OffsetRange_r16_DB_4  = 11
	UTRA_FDD_Q_OffsetRange_r16_DB_3  = 12
	UTRA_FDD_Q_OffsetRange_r16_DB_2  = 13
	UTRA_FDD_Q_OffsetRange_r16_DB_1  = 14
	UTRA_FDD_Q_OffsetRange_r16_DB0   = 15
	UTRA_FDD_Q_OffsetRange_r16_DB1   = 16
	UTRA_FDD_Q_OffsetRange_r16_DB2   = 17
	UTRA_FDD_Q_OffsetRange_r16_DB3   = 18
	UTRA_FDD_Q_OffsetRange_r16_DB4   = 19
	UTRA_FDD_Q_OffsetRange_r16_DB5   = 20
	UTRA_FDD_Q_OffsetRange_r16_DB6   = 21
	UTRA_FDD_Q_OffsetRange_r16_DB8   = 22
	UTRA_FDD_Q_OffsetRange_r16_DB10  = 23
	UTRA_FDD_Q_OffsetRange_r16_DB12  = 24
	UTRA_FDD_Q_OffsetRange_r16_DB14  = 25
	UTRA_FDD_Q_OffsetRange_r16_DB16  = 26
	UTRA_FDD_Q_OffsetRange_r16_DB18  = 27
	UTRA_FDD_Q_OffsetRange_r16_DB20  = 28
	UTRA_FDD_Q_OffsetRange_r16_DB22  = 29
	UTRA_FDD_Q_OffsetRange_r16_DB24  = 30
)

var uTRAFDDQOffsetRangeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UTRA_FDD_Q_OffsetRange_r16_DB_24, UTRA_FDD_Q_OffsetRange_r16_DB_22, UTRA_FDD_Q_OffsetRange_r16_DB_20, UTRA_FDD_Q_OffsetRange_r16_DB_18, UTRA_FDD_Q_OffsetRange_r16_DB_16, UTRA_FDD_Q_OffsetRange_r16_DB_14, UTRA_FDD_Q_OffsetRange_r16_DB_12, UTRA_FDD_Q_OffsetRange_r16_DB_10, UTRA_FDD_Q_OffsetRange_r16_DB_8, UTRA_FDD_Q_OffsetRange_r16_DB_6, UTRA_FDD_Q_OffsetRange_r16_DB_5, UTRA_FDD_Q_OffsetRange_r16_DB_4, UTRA_FDD_Q_OffsetRange_r16_DB_3, UTRA_FDD_Q_OffsetRange_r16_DB_2, UTRA_FDD_Q_OffsetRange_r16_DB_1, UTRA_FDD_Q_OffsetRange_r16_DB0, UTRA_FDD_Q_OffsetRange_r16_DB1, UTRA_FDD_Q_OffsetRange_r16_DB2, UTRA_FDD_Q_OffsetRange_r16_DB3, UTRA_FDD_Q_OffsetRange_r16_DB4, UTRA_FDD_Q_OffsetRange_r16_DB5, UTRA_FDD_Q_OffsetRange_r16_DB6, UTRA_FDD_Q_OffsetRange_r16_DB8, UTRA_FDD_Q_OffsetRange_r16_DB10, UTRA_FDD_Q_OffsetRange_r16_DB12, UTRA_FDD_Q_OffsetRange_r16_DB14, UTRA_FDD_Q_OffsetRange_r16_DB16, UTRA_FDD_Q_OffsetRange_r16_DB18, UTRA_FDD_Q_OffsetRange_r16_DB20, UTRA_FDD_Q_OffsetRange_r16_DB22, UTRA_FDD_Q_OffsetRange_r16_DB24},
}

type UTRA_FDD_Q_OffsetRange_r16 struct {
	Value int64
}

func (v *UTRA_FDD_Q_OffsetRange_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, uTRAFDDQOffsetRangeR16Constraints)
}

func (v *UTRA_FDD_Q_OffsetRange_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(uTRAFDDQOffsetRangeR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
