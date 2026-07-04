// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LogMeasResultListWLAN-r16 (line 26261).
// LogMeasResultListWLAN-r16 ::=    SEQUENCE (SIZE (1..maxWLAN-Id-Report-r16)) OF LogMeasResultWLAN-r16

var logMeasResultListWLANR16SizeConstraints = per.SizeRange(1, common.MaxWLAN_Id_Report_r16)

type LogMeasResultListWLAN_r16 []LogMeasResultWLAN_r16

func (ie *LogMeasResultListWLAN_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(logMeasResultListWLANR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *LogMeasResultListWLAN_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(logMeasResultListWLANR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LogMeasResultListWLAN_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
