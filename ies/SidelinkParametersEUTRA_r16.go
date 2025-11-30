package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkParametersEUTRA_r16 struct {
	Sl_ParametersEUTRA1_r16            *[]byte                 `optional`
	Sl_ParametersEUTRA2_r16            *[]byte                 `optional`
	Sl_ParametersEUTRA3_r16            *[]byte                 `optional`
	SupportedBandListSidelinkEUTRA_r16 []BandSidelinkEUTRA_r16 `lb:1,ub:maxBandsEUTRA,optional`
}

func (ie *SidelinkParametersEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ParametersEUTRA1_r16 != nil, ie.Sl_ParametersEUTRA2_r16 != nil, ie.Sl_ParametersEUTRA3_r16 != nil, len(ie.SupportedBandListSidelinkEUTRA_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_ParametersEUTRA1_r16 != nil {
		if err = w.WriteOctetString(*ie.Sl_ParametersEUTRA1_r16, nil, false); err != nil {
			return utils.WrapError("Encode Sl_ParametersEUTRA1_r16", err)
		}
	}
	if ie.Sl_ParametersEUTRA2_r16 != nil {
		if err = w.WriteOctetString(*ie.Sl_ParametersEUTRA2_r16, nil, false); err != nil {
			return utils.WrapError("Encode Sl_ParametersEUTRA2_r16", err)
		}
	}
	if ie.Sl_ParametersEUTRA3_r16 != nil {
		if err = w.WriteOctetString(*ie.Sl_ParametersEUTRA3_r16, nil, false); err != nil {
			return utils.WrapError("Encode Sl_ParametersEUTRA3_r16", err)
		}
	}
	if len(ie.SupportedBandListSidelinkEUTRA_r16) > 0 {
		tmp_SupportedBandListSidelinkEUTRA_r16 := utils.NewSequence[*BandSidelinkEUTRA_r16]([]*BandSidelinkEUTRA_r16{}, aper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		for _, i := range ie.SupportedBandListSidelinkEUTRA_r16 {
			tmp_SupportedBandListSidelinkEUTRA_r16.Value = append(tmp_SupportedBandListSidelinkEUTRA_r16.Value, &i)
		}
		if err = tmp_SupportedBandListSidelinkEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandListSidelinkEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *SidelinkParametersEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_ParametersEUTRA1_r16Present bool
	if Sl_ParametersEUTRA1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ParametersEUTRA2_r16Present bool
	if Sl_ParametersEUTRA2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ParametersEUTRA3_r16Present bool
	if Sl_ParametersEUTRA3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandListSidelinkEUTRA_r16Present bool
	if SupportedBandListSidelinkEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_ParametersEUTRA1_r16Present {
		var tmp_os_Sl_ParametersEUTRA1_r16 []byte
		if tmp_os_Sl_ParametersEUTRA1_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Sl_ParametersEUTRA1_r16", err)
		}
		ie.Sl_ParametersEUTRA1_r16 = &tmp_os_Sl_ParametersEUTRA1_r16
	}
	if Sl_ParametersEUTRA2_r16Present {
		var tmp_os_Sl_ParametersEUTRA2_r16 []byte
		if tmp_os_Sl_ParametersEUTRA2_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Sl_ParametersEUTRA2_r16", err)
		}
		ie.Sl_ParametersEUTRA2_r16 = &tmp_os_Sl_ParametersEUTRA2_r16
	}
	if Sl_ParametersEUTRA3_r16Present {
		var tmp_os_Sl_ParametersEUTRA3_r16 []byte
		if tmp_os_Sl_ParametersEUTRA3_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Sl_ParametersEUTRA3_r16", err)
		}
		ie.Sl_ParametersEUTRA3_r16 = &tmp_os_Sl_ParametersEUTRA3_r16
	}
	if SupportedBandListSidelinkEUTRA_r16Present {
		tmp_SupportedBandListSidelinkEUTRA_r16 := utils.NewSequence[*BandSidelinkEUTRA_r16]([]*BandSidelinkEUTRA_r16{}, aper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		fn_SupportedBandListSidelinkEUTRA_r16 := func() *BandSidelinkEUTRA_r16 {
			return new(BandSidelinkEUTRA_r16)
		}
		if err = tmp_SupportedBandListSidelinkEUTRA_r16.Decode(r, fn_SupportedBandListSidelinkEUTRA_r16); err != nil {
			return utils.WrapError("Decode SupportedBandListSidelinkEUTRA_r16", err)
		}
		ie.SupportedBandListSidelinkEUTRA_r16 = []BandSidelinkEUTRA_r16{}
		for _, i := range tmp_SupportedBandListSidelinkEUTRA_r16.Value {
			ie.SupportedBandListSidelinkEUTRA_r16 = append(ie.SupportedBandListSidelinkEUTRA_r16, *i)
		}
	}
	return nil
}
