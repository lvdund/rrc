// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GroupB-ConfiguredTwoStepRA-r16 (line 12909).

var groupBConfiguredTwoStepRAR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ra-MsgA-SizeGroupA-r16"},
		{Name: "messagePowerOffsetGroupB-r16"},
		{Name: "numberOfRA-PreamblesGroupA-r16"},
	},
}

const (
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B56    = 0
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B144   = 1
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B208   = 2
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B256   = 3
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B282   = 4
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B480   = 5
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B640   = 6
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B800   = 7
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B1000  = 8
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B72    = 9
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare6 = 10
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare5 = 11
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare4 = 12
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare3 = 13
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare2 = 14
	GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare1 = 15
)

var groupBConfiguredTwoStepRAR16RaMsgASizeGroupAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B56, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B144, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B208, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B256, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B282, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B480, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B640, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B800, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B1000, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_B72, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare6, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare5, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare4, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare3, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare2, GroupB_ConfiguredTwoStepRA_r16_Ra_MsgA_SizeGroupA_r16_Spare1},
}

const (
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_Minusinfinity = 0
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB0           = 1
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB5           = 2
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB8           = 3
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB10          = 4
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB12          = 5
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB15          = 6
	GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB18          = 7
)

var groupBConfiguredTwoStepRAR16MessagePowerOffsetGroupBR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_Minusinfinity, GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB0, GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB5, GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB8, GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB10, GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB12, GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB15, GroupB_ConfiguredTwoStepRA_r16_MessagePowerOffsetGroupB_r16_DB18},
}

var groupBConfiguredTwoStepRAR16NumberOfRAPreamblesGroupAR16Constraints = per.Constrained(1, 64)

type GroupB_ConfiguredTwoStepRA_r16 struct {
	Ra_MsgA_SizeGroupA_r16         int64
	MessagePowerOffsetGroupB_r16   int64
	NumberOfRA_PreamblesGroupA_r16 int64
}

func (ie *GroupB_ConfiguredTwoStepRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(groupBConfiguredTwoStepRAR16Constraints)
	_ = seq

	// 1. ra-MsgA-SizeGroupA-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ra_MsgA_SizeGroupA_r16, groupBConfiguredTwoStepRAR16RaMsgASizeGroupAR16Constraints); err != nil {
			return err
		}
	}

	// 2. messagePowerOffsetGroupB-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MessagePowerOffsetGroupB_r16, groupBConfiguredTwoStepRAR16MessagePowerOffsetGroupBR16Constraints); err != nil {
			return err
		}
	}

	// 3. numberOfRA-PreamblesGroupA-r16: integer
	{
		if err := e.EncodeInteger(ie.NumberOfRA_PreamblesGroupA_r16, groupBConfiguredTwoStepRAR16NumberOfRAPreamblesGroupAR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *GroupB_ConfiguredTwoStepRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(groupBConfiguredTwoStepRAR16Constraints)
	_ = seq

	// 1. ra-MsgA-SizeGroupA-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(groupBConfiguredTwoStepRAR16RaMsgASizeGroupAR16Constraints)
		if err != nil {
			return err
		}
		ie.Ra_MsgA_SizeGroupA_r16 = v0
	}

	// 2. messagePowerOffsetGroupB-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(groupBConfiguredTwoStepRAR16MessagePowerOffsetGroupBR16Constraints)
		if err != nil {
			return err
		}
		ie.MessagePowerOffsetGroupB_r16 = v1
	}

	// 3. numberOfRA-PreamblesGroupA-r16: integer
	{
		v2, err := d.DecodeInteger(groupBConfiguredTwoStepRAR16NumberOfRAPreamblesGroupAR16Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfRA_PreamblesGroupA_r16 = v2
	}

	return nil
}
