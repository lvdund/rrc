package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1640_IEs struct {
	ServCellInfoListSCG_NR_r16    *ServCellInfoListSCG_NR_r16    `optional`
	ServCellInfoListSCG_EUTRA_r16 *ServCellInfoListSCG_EUTRA_r16 `optional`
	NonCriticalExtension          *CG_Config_v1700_IEs           `optional`
}

func (ie *CG_Config_v1640_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ServCellInfoListSCG_NR_r16 != nil, ie.ServCellInfoListSCG_EUTRA_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ServCellInfoListSCG_NR_r16 != nil {
		if err = ie.ServCellInfoListSCG_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellInfoListSCG_NR_r16", err)
		}
	}
	if ie.ServCellInfoListSCG_EUTRA_r16 != nil {
		if err = ie.ServCellInfoListSCG_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellInfoListSCG_EUTRA_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1640_IEs) Decode(r *aper.AperReader) error {
	var err error
	var ServCellInfoListSCG_NR_r16Present bool
	if ServCellInfoListSCG_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ServCellInfoListSCG_EUTRA_r16Present bool
	if ServCellInfoListSCG_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ServCellInfoListSCG_NR_r16Present {
		ie.ServCellInfoListSCG_NR_r16 = new(ServCellInfoListSCG_NR_r16)
		if err = ie.ServCellInfoListSCG_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellInfoListSCG_NR_r16", err)
		}
	}
	if ServCellInfoListSCG_EUTRA_r16Present {
		ie.ServCellInfoListSCG_EUTRA_r16 = new(ServCellInfoListSCG_EUTRA_r16)
		if err = ie.ServCellInfoListSCG_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellInfoListSCG_EUTRA_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
