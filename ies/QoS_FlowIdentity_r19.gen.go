// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: QoS-FlowIdentity-r19 (line 9087).

var qoSFlowIdentityR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdu-SessionID-r19"},
		{Name: "qfi-r19"},
	},
}

type QoS_FlowIdentity_r19 struct {
	Pdu_SessionID_r19 PDU_SessionID
	Qfi_r19           QFI
}

func (ie *QoS_FlowIdentity_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowIdentityR19Constraints)
	_ = seq

	// 1. pdu-SessionID-r19: ref
	{
		if err := ie.Pdu_SessionID_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. qfi-r19: ref
	{
		if err := ie.Qfi_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *QoS_FlowIdentity_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowIdentityR19Constraints)
	_ = seq

	// 1. pdu-SessionID-r19: ref
	{
		if err := ie.Pdu_SessionID_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. qfi-r19: ref
	{
		if err := ie.Qfi_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
