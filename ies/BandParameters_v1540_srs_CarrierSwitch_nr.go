package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540_srs_CarrierSwitch_nr struct {
	Srs_SwitchingTimesListNR []SRS_SwitchingTimeNR `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_nr) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Srs_SwitchingTimesListNR := utils.NewSequence[*SRS_SwitchingTimeNR]([]*SRS_SwitchingTimeNR{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Srs_SwitchingTimesListNR {
		tmp_Srs_SwitchingTimesListNR.Value = append(tmp_Srs_SwitchingTimesListNR.Value, &i)
	}
	if err = tmp_Srs_SwitchingTimesListNR.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_SwitchingTimesListNR", err)
	}
	return nil
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_nr) Decode(r *uper.UperReader) error {
	var err error
	tmp_Srs_SwitchingTimesListNR := utils.NewSequence[*SRS_SwitchingTimeNR]([]*SRS_SwitchingTimeNR{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_Srs_SwitchingTimesListNR := func() *SRS_SwitchingTimeNR {
		return new(SRS_SwitchingTimeNR)
	}
	if err = tmp_Srs_SwitchingTimesListNR.Decode(r, fn_Srs_SwitchingTimesListNR); err != nil {
		return utils.WrapError("Decode Srs_SwitchingTimesListNR", err)
	}
	ie.Srs_SwitchingTimesListNR = []SRS_SwitchingTimeNR{}
	for _, i := range tmp_Srs_SwitchingTimesListNR.Value {
		ie.Srs_SwitchingTimesListNR = append(ie.Srs_SwitchingTimesListNR, *i)
	}
	return nil
}
