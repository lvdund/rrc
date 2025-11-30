package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16 struct {
	Ra_Prioritization_r16      RA_Prioritization `madatory`
	Ra_PrioritizationForAI_r16 aper.BitString    `lb:2,ub:2,madatory`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ra_Prioritization_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_Prioritization_r16", err)
	}
	if err = w.WriteBitString(ie.Ra_PrioritizationForAI_r16.Bytes, uint(ie.Ra_PrioritizationForAI_r16.NumBits), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString Ra_PrioritizationForAI_r16", err)
	}
	return nil
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ra_Prioritization_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_Prioritization_r16", err)
	}
	var tmp_bs_Ra_PrioritizationForAI_r16 []byte
	var n_Ra_PrioritizationForAI_r16 uint
	if tmp_bs_Ra_PrioritizationForAI_r16, n_Ra_PrioritizationForAI_r16, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString Ra_PrioritizationForAI_r16", err)
	}
	ie.Ra_PrioritizationForAI_r16 = aper.BitString{
		Bytes:   tmp_bs_Ra_PrioritizationForAI_r16,
		NumBits: uint64(n_Ra_PrioritizationForAI_r16),
	}
	return nil
}
