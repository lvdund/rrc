// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDU-SessionUL-TrafficInfo-r18 (line 2763).

var pDUSessionULTrafficInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdu-SessionID-r18"},
		{Name: "qos-FlowUL-TrafficInfoList-r18"},
	},
}

var pDUSessionULTrafficInfoR18QosFlowULTrafficInfoListR18Constraints = per.SizeRange(1, common.MaxNrofQFIs)

type PDU_SessionUL_TrafficInfo_r18 struct {
	Pdu_SessionID_r18              PDU_SessionID
	Qos_FlowUL_TrafficInfoList_r18 []QOS_FlowUL_TrafficInfo_r18
}

func (ie *PDU_SessionUL_TrafficInfo_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionULTrafficInfoR18Constraints)
	_ = seq

	// 1. pdu-SessionID-r18: ref
	{
		if err := ie.Pdu_SessionID_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. qos-FlowUL-TrafficInfoList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pDUSessionULTrafficInfoR18QosFlowULTrafficInfoListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Qos_FlowUL_TrafficInfoList_r18))); err != nil {
			return err
		}
		for i := range ie.Qos_FlowUL_TrafficInfoList_r18 {
			if err := ie.Qos_FlowUL_TrafficInfoList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDU_SessionUL_TrafficInfo_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionULTrafficInfoR18Constraints)
	_ = seq

	// 1. pdu-SessionID-r18: ref
	{
		if err := ie.Pdu_SessionID_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. qos-FlowUL-TrafficInfoList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pDUSessionULTrafficInfoR18QosFlowULTrafficInfoListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Qos_FlowUL_TrafficInfoList_r18 = make([]QOS_FlowUL_TrafficInfo_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Qos_FlowUL_TrafficInfoList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
