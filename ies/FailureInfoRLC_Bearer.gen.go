// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureInfoRLC-Bearer (line 410).

var failureInfoRLCBearerConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellGroupId"},
		{Name: "logicalChannelIdentity"},
		{Name: "failureType"},
	},
}

const (
	FailureInfoRLC_Bearer_FailureType_Rlc_Failure = 0
	FailureInfoRLC_Bearer_FailureType_Spare3      = 1
	FailureInfoRLC_Bearer_FailureType_Spare2      = 2
	FailureInfoRLC_Bearer_FailureType_Spare1      = 3
)

var failureInfoRLCBearerFailureTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FailureInfoRLC_Bearer_FailureType_Rlc_Failure, FailureInfoRLC_Bearer_FailureType_Spare3, FailureInfoRLC_Bearer_FailureType_Spare2, FailureInfoRLC_Bearer_FailureType_Spare1},
}

type FailureInfoRLC_Bearer struct {
	CellGroupId            CellGroupId
	LogicalChannelIdentity LogicalChannelIdentity
	FailureType            int64
}

func (ie *FailureInfoRLC_Bearer) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureInfoRLCBearerConstraints)
	_ = seq

	// 1. cellGroupId: ref
	{
		if err := ie.CellGroupId.Encode(e); err != nil {
			return err
		}
	}

	// 2. logicalChannelIdentity: ref
	{
		if err := ie.LogicalChannelIdentity.Encode(e); err != nil {
			return err
		}
	}

	// 3. failureType: enumerated
	{
		if err := e.EncodeEnumerated(ie.FailureType, failureInfoRLCBearerFailureTypeConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FailureInfoRLC_Bearer) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureInfoRLCBearerConstraints)
	_ = seq

	// 1. cellGroupId: ref
	{
		if err := ie.CellGroupId.Decode(d); err != nil {
			return err
		}
	}

	// 2. logicalChannelIdentity: ref
	{
		if err := ie.LogicalChannelIdentity.Decode(d); err != nil {
			return err
		}
	}

	// 3. failureType: enumerated
	{
		v2, err := d.DecodeEnumerated(failureInfoRLCBearerFailureTypeConstraints)
		if err != nil {
			return err
		}
		ie.FailureType = v2
	}

	return nil
}
