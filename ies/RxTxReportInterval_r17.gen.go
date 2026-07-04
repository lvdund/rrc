// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RxTxReportInterval-r17 (line 13905).
// RxTxReportInterval-r17 ::= ENUMERATED {ms80,ms120,ms160,ms240,ms320,ms480,ms640,ms1024,ms1280,ms2048,ms2560,ms5120,spare4,spare3,spare2,spare1}

const (
	RxTxReportInterval_r17_Ms80   = 0
	RxTxReportInterval_r17_Ms120  = 1
	RxTxReportInterval_r17_Ms160  = 2
	RxTxReportInterval_r17_Ms240  = 3
	RxTxReportInterval_r17_Ms320  = 4
	RxTxReportInterval_r17_Ms480  = 5
	RxTxReportInterval_r17_Ms640  = 6
	RxTxReportInterval_r17_Ms1024 = 7
	RxTxReportInterval_r17_Ms1280 = 8
	RxTxReportInterval_r17_Ms2048 = 9
	RxTxReportInterval_r17_Ms2560 = 10
	RxTxReportInterval_r17_Ms5120 = 11
	RxTxReportInterval_r17_Spare4 = 12
	RxTxReportInterval_r17_Spare3 = 13
	RxTxReportInterval_r17_Spare2 = 14
	RxTxReportInterval_r17_Spare1 = 15
)

var rxTxReportIntervalR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RxTxReportInterval_r17_Ms80, RxTxReportInterval_r17_Ms120, RxTxReportInterval_r17_Ms160, RxTxReportInterval_r17_Ms240, RxTxReportInterval_r17_Ms320, RxTxReportInterval_r17_Ms480, RxTxReportInterval_r17_Ms640, RxTxReportInterval_r17_Ms1024, RxTxReportInterval_r17_Ms1280, RxTxReportInterval_r17_Ms2048, RxTxReportInterval_r17_Ms2560, RxTxReportInterval_r17_Ms5120, RxTxReportInterval_r17_Spare4, RxTxReportInterval_r17_Spare3, RxTxReportInterval_r17_Spare2, RxTxReportInterval_r17_Spare1},
}

type RxTxReportInterval_r17 struct {
	Value int64
}

func (v *RxTxReportInterval_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, rxTxReportIntervalR17Constraints)
}

func (v *RxTxReportInterval_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(rxTxReportIntervalR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
