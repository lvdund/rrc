package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityInformationSidelink_r16_IEs struct {
	AccessStratumReleaseSidelink_r16           AccessStratumReleaseSidelink_r16           `madatory`
	Pdcp_ParametersSidelink_r16                *PDCP_ParametersSidelink_r16               `optional`
	Rlc_ParametersSidelink_r16                 *RLC_ParametersSidelink_r16                `optional`
	SupportedBandCombinationListSidelinkNR_r16 *BandCombinationListSidelinkNR_r16         `optional`
	SupportedBandListSidelink_r16              []BandSidelinkPC5_r16                      `lb:1,ub:maxBands,optional`
	AppliedFreqBandListFilter_r16              *FreqBandList                              `optional`
	LateNonCriticalExtension                   *[]byte                                    `optional`
	NonCriticalExtension                       *UECapabilityInformationSidelink_v1700_IEs `optional`
}

func (ie *UECapabilityInformationSidelink_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcp_ParametersSidelink_r16 != nil, ie.Rlc_ParametersSidelink_r16 != nil, ie.SupportedBandCombinationListSidelinkNR_r16 != nil, len(ie.SupportedBandListSidelink_r16) > 0, ie.AppliedFreqBandListFilter_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.AccessStratumReleaseSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AccessStratumReleaseSidelink_r16", err)
	}
	if ie.Pdcp_ParametersSidelink_r16 != nil {
		if err = ie.Pdcp_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_ParametersSidelink_r16", err)
		}
	}
	if ie.Rlc_ParametersSidelink_r16 != nil {
		if err = ie.Rlc_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_ParametersSidelink_r16", err)
		}
	}
	if ie.SupportedBandCombinationListSidelinkNR_r16 != nil {
		if err = ie.SupportedBandCombinationListSidelinkNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationListSidelinkNR_r16", err)
		}
	}
	if len(ie.SupportedBandListSidelink_r16) > 0 {
		tmp_SupportedBandListSidelink_r16 := utils.NewSequence[*BandSidelinkPC5_r16]([]*BandSidelinkPC5_r16{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.SupportedBandListSidelink_r16 {
			tmp_SupportedBandListSidelink_r16.Value = append(tmp_SupportedBandListSidelink_r16.Value, &i)
		}
		if err = tmp_SupportedBandListSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandListSidelink_r16", err)
		}
	}
	if ie.AppliedFreqBandListFilter_r16 != nil {
		if err = ie.AppliedFreqBandListFilter_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AppliedFreqBandListFilter_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityInformationSidelink_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Pdcp_ParametersSidelink_r16Present bool
	if Pdcp_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlc_ParametersSidelink_r16Present bool
	if Rlc_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandCombinationListSidelinkNR_r16Present bool
	if SupportedBandCombinationListSidelinkNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandListSidelink_r16Present bool
	if SupportedBandListSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AppliedFreqBandListFilter_r16Present bool
	if AppliedFreqBandListFilter_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.AccessStratumReleaseSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AccessStratumReleaseSidelink_r16", err)
	}
	if Pdcp_ParametersSidelink_r16Present {
		ie.Pdcp_ParametersSidelink_r16 = new(PDCP_ParametersSidelink_r16)
		if err = ie.Pdcp_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_ParametersSidelink_r16", err)
		}
	}
	if Rlc_ParametersSidelink_r16Present {
		ie.Rlc_ParametersSidelink_r16 = new(RLC_ParametersSidelink_r16)
		if err = ie.Rlc_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rlc_ParametersSidelink_r16", err)
		}
	}
	if SupportedBandCombinationListSidelinkNR_r16Present {
		ie.SupportedBandCombinationListSidelinkNR_r16 = new(BandCombinationListSidelinkNR_r16)
		if err = ie.SupportedBandCombinationListSidelinkNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationListSidelinkNR_r16", err)
		}
	}
	if SupportedBandListSidelink_r16Present {
		tmp_SupportedBandListSidelink_r16 := utils.NewSequence[*BandSidelinkPC5_r16]([]*BandSidelinkPC5_r16{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_SupportedBandListSidelink_r16 := func() *BandSidelinkPC5_r16 {
			return new(BandSidelinkPC5_r16)
		}
		if err = tmp_SupportedBandListSidelink_r16.Decode(r, fn_SupportedBandListSidelink_r16); err != nil {
			return utils.WrapError("Decode SupportedBandListSidelink_r16", err)
		}
		ie.SupportedBandListSidelink_r16 = []BandSidelinkPC5_r16{}
		for _, i := range tmp_SupportedBandListSidelink_r16.Value {
			ie.SupportedBandListSidelink_r16 = append(ie.SupportedBandListSidelink_r16, *i)
		}
	}
	if AppliedFreqBandListFilter_r16Present {
		ie.AppliedFreqBandListFilter_r16 = new(FreqBandList)
		if err = ie.AppliedFreqBandListFilter_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AppliedFreqBandListFilter_r16", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UECapabilityInformationSidelink_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
