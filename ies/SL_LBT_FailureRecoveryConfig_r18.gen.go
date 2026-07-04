// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-LBT-FailureRecoveryConfig-r18 (line 27398).

var sLLBTFailureRecoveryConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-LBT-FailureInstanceMaxCount-r18", Optional: true},
		{Name: "sl-LBT-FailureDetectionTimer-r18", Optional: true},
		{Name: "sl-LBT-RecoveryTimer-r18", Optional: true},
	},
}

const (
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N4     = 0
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N8     = 1
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N16    = 2
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N32    = 3
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N64    = 4
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N128   = 5
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_Spare2 = 6
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_Spare1 = 7
)

var sLLBTFailureRecoveryConfigR18SlLBTFailureInstanceMaxCountR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N4, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N8, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N16, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N32, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N64, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_N128, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_Spare2, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureInstanceMaxCount_r18_Spare1},
}

const (
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms10   = 0
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms20   = 1
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms40   = 2
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms80   = 3
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms160  = 4
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms320  = 5
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Spare2 = 6
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Spare1 = 7
)

var sLLBTFailureRecoveryConfigR18SlLBTFailureDetectionTimerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms10, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms20, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms40, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms80, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms160, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Ms320, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Spare2, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_FailureDetectionTimer_r18_Spare1},
}

const (
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms10   = 0
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms20   = 1
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms40   = 2
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms80   = 3
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms160  = 4
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms320  = 5
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Spare2 = 6
	SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Spare1 = 7
)

var sLLBTFailureRecoveryConfigR18SlLBTRecoveryTimerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms10, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms20, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms40, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms80, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms160, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Ms320, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Spare2, SL_LBT_FailureRecoveryConfig_r18_Sl_LBT_RecoveryTimer_r18_Spare1},
}

type SL_LBT_FailureRecoveryConfig_r18 struct {
	Sl_LBT_FailureInstanceMaxCount_r18 *int64
	Sl_LBT_FailureDetectionTimer_r18   *int64
	Sl_LBT_RecoveryTimer_r18           *int64
}

func (ie *SL_LBT_FailureRecoveryConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLLBTFailureRecoveryConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_LBT_FailureInstanceMaxCount_r18 != nil, ie.Sl_LBT_FailureDetectionTimer_r18 != nil, ie.Sl_LBT_RecoveryTimer_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-LBT-FailureInstanceMaxCount-r18: enumerated
	{
		if ie.Sl_LBT_FailureInstanceMaxCount_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_LBT_FailureInstanceMaxCount_r18, sLLBTFailureRecoveryConfigR18SlLBTFailureInstanceMaxCountR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-LBT-FailureDetectionTimer-r18: enumerated
	{
		if ie.Sl_LBT_FailureDetectionTimer_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_LBT_FailureDetectionTimer_r18, sLLBTFailureRecoveryConfigR18SlLBTFailureDetectionTimerR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-LBT-RecoveryTimer-r18: enumerated
	{
		if ie.Sl_LBT_RecoveryTimer_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_LBT_RecoveryTimer_r18, sLLBTFailureRecoveryConfigR18SlLBTRecoveryTimerR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_LBT_FailureRecoveryConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLLBTFailureRecoveryConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-LBT-FailureInstanceMaxCount-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLLBTFailureRecoveryConfigR18SlLBTFailureInstanceMaxCountR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LBT_FailureInstanceMaxCount_r18 = &idx
		}
	}

	// 4. sl-LBT-FailureDetectionTimer-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLLBTFailureRecoveryConfigR18SlLBTFailureDetectionTimerR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LBT_FailureDetectionTimer_r18 = &idx
		}
	}

	// 5. sl-LBT-RecoveryTimer-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLLBTFailureRecoveryConfigR18SlLBTRecoveryTimerR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LBT_RecoveryTimer_r18 = &idx
		}
	}

	return nil
}
