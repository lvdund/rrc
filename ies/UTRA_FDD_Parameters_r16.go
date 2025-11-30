package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UTRA_FDD_Parameters_r16 struct {
	SupportedBandListUTRA_FDD_r16 []SupportedBandUTRA_FDD_r16 `lb:1,ub:maxBandsUTRA_FDD_r16,madatory`
}

func (ie *UTRA_FDD_Parameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp_SupportedBandListUTRA_FDD_r16 := utils.NewSequence[*SupportedBandUTRA_FDD_r16]([]*SupportedBandUTRA_FDD_r16{}, aper.Constraint{Lb: 1, Ub: maxBandsUTRA_FDD_r16}, false)
	for _, i := range ie.SupportedBandListUTRA_FDD_r16 {
		tmp_SupportedBandListUTRA_FDD_r16.Value = append(tmp_SupportedBandListUTRA_FDD_r16.Value, &i)
	}
	if err = tmp_SupportedBandListUTRA_FDD_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedBandListUTRA_FDD_r16", err)
	}
	return nil
}

func (ie *UTRA_FDD_Parameters_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp_SupportedBandListUTRA_FDD_r16 := utils.NewSequence[*SupportedBandUTRA_FDD_r16]([]*SupportedBandUTRA_FDD_r16{}, aper.Constraint{Lb: 1, Ub: maxBandsUTRA_FDD_r16}, false)
	fn_SupportedBandListUTRA_FDD_r16 := func() *SupportedBandUTRA_FDD_r16 {
		return new(SupportedBandUTRA_FDD_r16)
	}
	if err = tmp_SupportedBandListUTRA_FDD_r16.Decode(r, fn_SupportedBandListUTRA_FDD_r16); err != nil {
		return utils.WrapError("Decode SupportedBandListUTRA_FDD_r16", err)
	}
	ie.SupportedBandListUTRA_FDD_r16 = []SupportedBandUTRA_FDD_r16{}
	for _, i := range tmp_SupportedBandListUTRA_FDD_r16.Value {
		ie.SupportedBandListUTRA_FDD_r16 = append(ie.SupportedBandListUTRA_FDD_r16, *i)
	}
	return nil
}
