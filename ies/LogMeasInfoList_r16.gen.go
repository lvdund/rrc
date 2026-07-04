// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LogMeasInfoList-r16 (line 3001).
// LogMeasInfoList-r16 ::=              SEQUENCE (SIZE (1..maxLogMeasReport-r16)) OF LogMeasInfo-r16

var logMeasInfoListR16SizeConstraints = per.SizeRange(1, common.MaxLogMeasReport_r16)

type LogMeasInfoList_r16 []LogMeasInfo_r16

func (ie *LogMeasInfoList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(logMeasInfoListR16SizeConstraints)
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

func (ie *LogMeasInfoList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(logMeasInfoListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LogMeasInfoList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
