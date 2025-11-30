package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_MBSFN_SubframeConfigList struct {
	Value []EUTRA_MBSFN_SubframeConfig `lb:1,ub:maxMBSFN_Allocations,madatory`
}

func (ie *EUTRA_MBSFN_SubframeConfigList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_MBSFN_SubframeConfig]([]*EUTRA_MBSFN_SubframeConfig{}, aper.Constraint{Lb: 1, Ub: maxMBSFN_Allocations}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode EUTRA_MBSFN_SubframeConfigList", err)
	}
	return nil
}

func (ie *EUTRA_MBSFN_SubframeConfigList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_MBSFN_SubframeConfig]([]*EUTRA_MBSFN_SubframeConfig{}, aper.Constraint{Lb: 1, Ub: maxMBSFN_Allocations}, false)
	fn := func() *EUTRA_MBSFN_SubframeConfig {
		return new(EUTRA_MBSFN_SubframeConfig)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode EUTRA_MBSFN_SubframeConfigList", err)
	}
	ie.Value = []EUTRA_MBSFN_SubframeConfig{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
