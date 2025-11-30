package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RelaysTriggeredList_r17 struct {
	Value []SL_SourceIdentity_r17 `lb:1,ub:maxNrofRelayMeas_r17,madatory`
}

func (ie *RelaysTriggeredList_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_SourceIdentity_r17]([]*SL_SourceIdentity_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofRelayMeas_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode RelaysTriggeredList_r17", err)
	}
	return nil
}

func (ie *RelaysTriggeredList_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_SourceIdentity_r17]([]*SL_SourceIdentity_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofRelayMeas_r17}, false)
	fn := func() *SL_SourceIdentity_r17 {
		return new(SL_SourceIdentity_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode RelaysTriggeredList_r17", err)
	}
	ie.Value = []SL_SourceIdentity_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
