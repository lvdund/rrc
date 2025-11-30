package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1640_IEs struct {
	ServCellInfoListMCG_NR_r16    *ServCellInfoListMCG_NR_r16    `optional`
	ServCellInfoListMCG_EUTRA_r16 *ServCellInfoListMCG_EUTRA_r16 `optional`
	NonCriticalExtension          *CG_ConfigInfo_v1700_IEs       `optional`
}

func (ie *CG_ConfigInfo_v1640_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ServCellInfoListMCG_NR_r16 != nil, ie.ServCellInfoListMCG_EUTRA_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ServCellInfoListMCG_NR_r16 != nil {
		if err = ie.ServCellInfoListMCG_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellInfoListMCG_NR_r16", err)
		}
	}
	if ie.ServCellInfoListMCG_EUTRA_r16 != nil {
		if err = ie.ServCellInfoListMCG_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellInfoListMCG_EUTRA_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1640_IEs) Decode(r *aper.AperReader) error {
	var err error
	var ServCellInfoListMCG_NR_r16Present bool
	if ServCellInfoListMCG_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ServCellInfoListMCG_EUTRA_r16Present bool
	if ServCellInfoListMCG_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ServCellInfoListMCG_NR_r16Present {
		ie.ServCellInfoListMCG_NR_r16 = new(ServCellInfoListMCG_NR_r16)
		if err = ie.ServCellInfoListMCG_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellInfoListMCG_NR_r16", err)
		}
	}
	if ServCellInfoListMCG_EUTRA_r16Present {
		ie.ServCellInfoListMCG_EUTRA_r16 = new(ServCellInfoListMCG_EUTRA_r16)
		if err = ie.ServCellInfoListMCG_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellInfoListMCG_EUTRA_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
