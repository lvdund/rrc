// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TraceReference-r16 (line 26659).

var traceReferenceR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity-r16"},
		{Name: "traceId-r16"},
	},
}

var traceReferenceR16TraceIdR16Constraints = per.FixedSize(3)

type TraceReference_r16 struct {
	Plmn_Identity_r16 PLMN_Identity
	TraceId_r16       []byte
}

func (ie *TraceReference_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(traceReferenceR16Constraints)
	_ = seq

	// 1. plmn-Identity-r16: ref
	{
		if err := ie.Plmn_Identity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. traceId-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.TraceId_r16, traceReferenceR16TraceIdR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TraceReference_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(traceReferenceR16Constraints)
	_ = seq

	// 1. plmn-Identity-r16: ref
	{
		if err := ie.Plmn_Identity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. traceId-r16: octet-string
	{
		v1, err := d.DecodeOctetString(traceReferenceR16TraceIdR16Constraints)
		if err != nil {
			return err
		}
		ie.TraceId_r16 = v1
	}

	return nil
}
