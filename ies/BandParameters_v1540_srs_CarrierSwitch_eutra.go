package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540_srs_CarrierSwitch_eutra struct {
	Srs_SwitchingTimesListEUTRA []SRS_SwitchingTimeEUTRA `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_eutra) Encode(w *aper.AperWriter) error {
	var err error
	tmp_Srs_SwitchingTimesListEUTRA := utils.NewSequence[*SRS_SwitchingTimeEUTRA]([]*SRS_SwitchingTimeEUTRA{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Srs_SwitchingTimesListEUTRA {
		tmp_Srs_SwitchingTimesListEUTRA.Value = append(tmp_Srs_SwitchingTimesListEUTRA.Value, &i)
	}
	if err = tmp_Srs_SwitchingTimesListEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_SwitchingTimesListEUTRA", err)
	}
	return nil
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_eutra) Decode(r *aper.AperReader) error {
	var err error
	tmp_Srs_SwitchingTimesListEUTRA := utils.NewSequence[*SRS_SwitchingTimeEUTRA]([]*SRS_SwitchingTimeEUTRA{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_Srs_SwitchingTimesListEUTRA := func() *SRS_SwitchingTimeEUTRA {
		return new(SRS_SwitchingTimeEUTRA)
	}
	if err = tmp_Srs_SwitchingTimesListEUTRA.Decode(r, fn_Srs_SwitchingTimesListEUTRA); err != nil {
		return utils.WrapError("Decode Srs_SwitchingTimesListEUTRA", err)
	}
	ie.Srs_SwitchingTimesListEUTRA = []SRS_SwitchingTimeEUTRA{}
	for _, i := range tmp_Srs_SwitchingTimesListEUTRA.Value {
		ie.Srs_SwitchingTimesListEUTRA = append(ie.Srs_SwitchingTimesListEUTRA, *i)
	}
	return nil
}
