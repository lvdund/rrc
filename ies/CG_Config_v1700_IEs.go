package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1700_IEs struct {
	CandidateCellInfoListCPC_r17 *CandidateCellInfoListCPC_r17          `optional`
	TwoPHRModeSCG_r17            *CG_Config_v1700_IEs_twoPHRModeSCG_r17 `optional`
	NonCriticalExtension         *CG_Config_v1730_IEs                   `optional`
}

func (ie *CG_Config_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CandidateCellInfoListCPC_r17 != nil, ie.TwoPHRModeSCG_r17 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CandidateCellInfoListCPC_r17 != nil {
		if err = ie.CandidateCellInfoListCPC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateCellInfoListCPC_r17", err)
		}
	}
	if ie.TwoPHRModeSCG_r17 != nil {
		if err = ie.TwoPHRModeSCG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPHRModeSCG_r17", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var CandidateCellInfoListCPC_r17Present bool
	if CandidateCellInfoListCPC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPHRModeSCG_r17Present bool
	if TwoPHRModeSCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CandidateCellInfoListCPC_r17Present {
		ie.CandidateCellInfoListCPC_r17 = new(CandidateCellInfoListCPC_r17)
		if err = ie.CandidateCellInfoListCPC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CandidateCellInfoListCPC_r17", err)
		}
	}
	if TwoPHRModeSCG_r17Present {
		ie.TwoPHRModeSCG_r17 = new(CG_Config_v1700_IEs_twoPHRModeSCG_r17)
		if err = ie.TwoPHRModeSCG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPHRModeSCG_r17", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1730_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
