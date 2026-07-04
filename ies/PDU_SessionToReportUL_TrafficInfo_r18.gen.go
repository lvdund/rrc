// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDU-SessionToReportUL-TrafficInfo-r18 (line 26522).

var pDUSessionToReportULTrafficInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdu-SessionID-r18"},
		{Name: "qfi-ToReportUL-TrafficInfoList-r18"},
	},
}

var pDUSessionToReportULTrafficInfoR18QfiToReportULTrafficInfoListR18Constraints = per.SizeRange(1, common.MaxNrofQFIs)

type PDU_SessionToReportUL_TrafficInfo_r18 struct {
	Pdu_SessionID_r18                  PDU_SessionID
	Qfi_ToReportUL_TrafficInfoList_r18 []QFI
}

func (ie *PDU_SessionToReportUL_TrafficInfo_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionToReportULTrafficInfoR18Constraints)
	_ = seq

	// 1. pdu-SessionID-r18: ref
	{
		if err := ie.Pdu_SessionID_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. qfi-ToReportUL-TrafficInfoList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pDUSessionToReportULTrafficInfoR18QfiToReportULTrafficInfoListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Qfi_ToReportUL_TrafficInfoList_r18))); err != nil {
			return err
		}
		for i := range ie.Qfi_ToReportUL_TrafficInfoList_r18 {
			if err := ie.Qfi_ToReportUL_TrafficInfoList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDU_SessionToReportUL_TrafficInfo_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionToReportULTrafficInfoR18Constraints)
	_ = seq

	// 1. pdu-SessionID-r18: ref
	{
		if err := ie.Pdu_SessionID_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. qfi-ToReportUL-TrafficInfoList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pDUSessionToReportULTrafficInfoR18QfiToReportULTrafficInfoListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Qfi_ToReportUL_TrafficInfoList_r18 = make([]QFI, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Qfi_ToReportUL_TrafficInfoList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
