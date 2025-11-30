package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UERadioPagingInformation_IEs struct {
	SupportedBandListNRForPaging []FreqBandIndicatorNR               `lb:1,ub:maxBands,optional`
	NonCriticalExtension         *UERadioPagingInformation_v15e0_IEs `optional`
}

func (ie *UERadioPagingInformation_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.SupportedBandListNRForPaging) > 0, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.SupportedBandListNRForPaging) > 0 {
		tmp_SupportedBandListNRForPaging := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.SupportedBandListNRForPaging {
			tmp_SupportedBandListNRForPaging.Value = append(tmp_SupportedBandListNRForPaging.Value, &i)
		}
		if err = tmp_SupportedBandListNRForPaging.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandListNRForPaging", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UERadioPagingInformation_IEs) Decode(r *aper.AperReader) error {
	var err error
	var SupportedBandListNRForPagingPresent bool
	if SupportedBandListNRForPagingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedBandListNRForPagingPresent {
		tmp_SupportedBandListNRForPaging := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_SupportedBandListNRForPaging := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_SupportedBandListNRForPaging.Decode(r, fn_SupportedBandListNRForPaging); err != nil {
			return utils.WrapError("Decode SupportedBandListNRForPaging", err)
		}
		ie.SupportedBandListNRForPaging = []FreqBandIndicatorNR{}
		for _, i := range tmp_SupportedBandListNRForPaging.Value {
			ie.SupportedBandListNRForPaging = append(ie.SupportedBandListNRForPaging, *i)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UERadioPagingInformation_v15e0_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
