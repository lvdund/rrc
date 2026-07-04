// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ReportSubConfigTriggerList-r18 (line 7458).
// CSI-ReportSubConfigTriggerList-r18 ::= SEQUENCE (SIZE(1..maxNrofCSI-ReportSubconfigPerCSI-ReportConfig-r18)) OF CSI-ReportSubConfigId-r18

var cSIReportSubConfigTriggerListR18SizeConstraints = per.SizeRange(1, common.MaxNrofCSI_ReportSubconfigPerCSI_ReportConfig_r18)

type CSI_ReportSubConfigTriggerList_r18 []CSI_ReportSubConfigId_r18

func (ie *CSI_ReportSubConfigTriggerList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cSIReportSubConfigTriggerListR18SizeConstraints)
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

func (ie *CSI_ReportSubConfigTriggerList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cSIReportSubConfigTriggerListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CSI_ReportSubConfigTriggerList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
