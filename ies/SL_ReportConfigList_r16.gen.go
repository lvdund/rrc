// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ReportConfigList-r16 (line 27872).
// SL-ReportConfigList-r16 ::=           SEQUENCE (SIZE (1..maxNrofSL-ReportConfigId-r16)) OF SL-ReportConfigInfo-r16

var sLReportConfigListR16SizeConstraints = per.SizeRange(1, common.MaxNrofSL_ReportConfigId_r16)

type SL_ReportConfigList_r16 []SL_ReportConfigInfo_r16

func (ie *SL_ReportConfigList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLReportConfigListR16SizeConstraints)
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

func (ie *SL_ReportConfigList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLReportConfigListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_ReportConfigList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
