// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DataInactivityTimer (line 9055).
// DataInactivityTimer ::=         ENUMERATED {s1, s2, s3, s5, s7, s10, s15, s20, s40, s50, s60, s80, s100, s120, s150, s180}

const (
	DataInactivityTimer_S1   = 0
	DataInactivityTimer_S2   = 1
	DataInactivityTimer_S3   = 2
	DataInactivityTimer_S5   = 3
	DataInactivityTimer_S7   = 4
	DataInactivityTimer_S10  = 5
	DataInactivityTimer_S15  = 6
	DataInactivityTimer_S20  = 7
	DataInactivityTimer_S40  = 8
	DataInactivityTimer_S50  = 9
	DataInactivityTimer_S60  = 10
	DataInactivityTimer_S80  = 11
	DataInactivityTimer_S100 = 12
	DataInactivityTimer_S120 = 13
	DataInactivityTimer_S150 = 14
	DataInactivityTimer_S180 = 15
)

var dataInactivityTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DataInactivityTimer_S1, DataInactivityTimer_S2, DataInactivityTimer_S3, DataInactivityTimer_S5, DataInactivityTimer_S7, DataInactivityTimer_S10, DataInactivityTimer_S15, DataInactivityTimer_S20, DataInactivityTimer_S40, DataInactivityTimer_S50, DataInactivityTimer_S60, DataInactivityTimer_S80, DataInactivityTimer_S100, DataInactivityTimer_S120, DataInactivityTimer_S150, DataInactivityTimer_S180},
}

type DataInactivityTimer struct {
	Value int64
}

func (v *DataInactivityTimer) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, dataInactivityTimerConstraints)
}

func (v *DataInactivityTimer) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(dataInactivityTimerConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
