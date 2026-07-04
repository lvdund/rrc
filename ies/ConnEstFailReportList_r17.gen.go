// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConnEstFailReportList-r17 (line 3036).
// ConnEstFailReportList-r17 ::= SEQUENCE (SIZE (1..maxCEFReport-r17)) OF ConnEstFailReport-r16

var connEstFailReportListR17SizeConstraints = per.SizeRange(1, common.MaxCEFReport_r17)

type ConnEstFailReportList_r17 []ConnEstFailReport_r16

func (ie *ConnEstFailReportList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(connEstFailReportListR17SizeConstraints)
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

func (ie *ConnEstFailReportList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(connEstFailReportListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ConnEstFailReportList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
