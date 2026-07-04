// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LBT-FailureRecoveryConfig-r16 (line 8540).

var lBTFailureRecoveryConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lbt-FailureInstanceMaxCount-r16"},
		{Name: "lbt-FailureDetectionTimer-r16"},
	},
}

const (
	LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N4   = 0
	LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N8   = 1
	LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N16  = 2
	LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N32  = 3
	LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N64  = 4
	LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N128 = 5
)

var lBTFailureRecoveryConfigR16LbtFailureInstanceMaxCountR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N4, LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N8, LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N16, LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N32, LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N64, LBT_FailureRecoveryConfig_r16_Lbt_FailureInstanceMaxCount_r16_N128},
}

const (
	LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms10  = 0
	LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms20  = 1
	LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms40  = 2
	LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms80  = 3
	LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms160 = 4
	LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms320 = 5
)

var lBTFailureRecoveryConfigR16LbtFailureDetectionTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms10, LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms20, LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms40, LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms80, LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms160, LBT_FailureRecoveryConfig_r16_Lbt_FailureDetectionTimer_r16_Ms320},
}

type LBT_FailureRecoveryConfig_r16 struct {
	Lbt_FailureInstanceMaxCount_r16 int64
	Lbt_FailureDetectionTimer_r16   int64
}

func (ie *LBT_FailureRecoveryConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lBTFailureRecoveryConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. lbt-FailureInstanceMaxCount-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Lbt_FailureInstanceMaxCount_r16, lBTFailureRecoveryConfigR16LbtFailureInstanceMaxCountR16Constraints); err != nil {
			return err
		}
	}

	// 3. lbt-FailureDetectionTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Lbt_FailureDetectionTimer_r16, lBTFailureRecoveryConfigR16LbtFailureDetectionTimerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LBT_FailureRecoveryConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lBTFailureRecoveryConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. lbt-FailureInstanceMaxCount-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(lBTFailureRecoveryConfigR16LbtFailureInstanceMaxCountR16Constraints)
		if err != nil {
			return err
		}
		ie.Lbt_FailureInstanceMaxCount_r16 = v0
	}

	// 3. lbt-FailureDetectionTimer-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(lBTFailureRecoveryConfigR16LbtFailureDetectionTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.Lbt_FailureDetectionTimer_r16 = v1
	}

	return nil
}
