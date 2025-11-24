package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LBT_FailureRecoveryConfig_r16 struct {
	Lbt_FailureInstanceMaxCount_r16 LBT_FailureRecoveryConfig_r16_lbt_FailureInstanceMaxCount_r16 `madatory`
	Lbt_FailureDetectionTimer_r16   LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16   `madatory`
}

func (ie *LBT_FailureRecoveryConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Lbt_FailureInstanceMaxCount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Lbt_FailureInstanceMaxCount_r16", err)
	}
	if err = ie.Lbt_FailureDetectionTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Lbt_FailureDetectionTimer_r16", err)
	}
	return nil
}

func (ie *LBT_FailureRecoveryConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Lbt_FailureInstanceMaxCount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Lbt_FailureInstanceMaxCount_r16", err)
	}
	if err = ie.Lbt_FailureDetectionTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Lbt_FailureDetectionTimer_r16", err)
	}
	return nil
}
