// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RxDRX-ReportList-v1700 (line 2178).
// SL-RxDRX-ReportList-v1700 ::=          SEQUENCE (SIZE (1..maxNrofSL-Dest-r16)) OF SL-RxDRX-Report-v1700

var sLRxDRXReportListV1700SizeConstraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SL_RxDRX_ReportList_v1700 []SL_RxDRX_Report_v1700

func (ie *SL_RxDRX_ReportList_v1700) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLRxDRXReportListV1700SizeConstraints)
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

func (ie *SL_RxDRX_ReportList_v1700) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLRxDRXReportListV1700SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_RxDRX_ReportList_v1700, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
