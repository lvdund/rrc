// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PollPDU (line 14103).

const (
	PollPDU_P4       = 0
	PollPDU_P8       = 1
	PollPDU_P16      = 2
	PollPDU_P32      = 3
	PollPDU_P64      = 4
	PollPDU_P128     = 5
	PollPDU_P256     = 6
	PollPDU_P512     = 7
	PollPDU_P1024    = 8
	PollPDU_P2048    = 9
	PollPDU_P4096    = 10
	PollPDU_P6144    = 11
	PollPDU_P8192    = 12
	PollPDU_P12288   = 13
	PollPDU_P16384   = 14
	PollPDU_P20480   = 15
	PollPDU_P24576   = 16
	PollPDU_P28672   = 17
	PollPDU_P32768   = 18
	PollPDU_P40960   = 19
	PollPDU_P49152   = 20
	PollPDU_P57344   = 21
	PollPDU_P65536   = 22
	PollPDU_Infinity = 23
	PollPDU_Spare8   = 24
	PollPDU_Spare7   = 25
	PollPDU_Spare6   = 26
	PollPDU_Spare5   = 27
	PollPDU_Spare4   = 28
	PollPDU_Spare3   = 29
	PollPDU_Spare2   = 30
	PollPDU_Spare1   = 31
)

var pollPDUConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PollPDU_P4, PollPDU_P8, PollPDU_P16, PollPDU_P32, PollPDU_P64, PollPDU_P128, PollPDU_P256, PollPDU_P512, PollPDU_P1024, PollPDU_P2048, PollPDU_P4096, PollPDU_P6144, PollPDU_P8192, PollPDU_P12288, PollPDU_P16384, PollPDU_P20480, PollPDU_P24576, PollPDU_P28672, PollPDU_P32768, PollPDU_P40960, PollPDU_P49152, PollPDU_P57344, PollPDU_P65536, PollPDU_Infinity, PollPDU_Spare8, PollPDU_Spare7, PollPDU_Spare6, PollPDU_Spare5, PollPDU_Spare4, PollPDU_Spare3, PollPDU_Spare2, PollPDU_Spare1},
}

type PollPDU struct {
	Value int64
}

func (v *PollPDU) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, pollPDUConstraints)
}

func (v *PollPDU) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(pollPDUConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
