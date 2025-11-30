package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RSSI_ResourceListConfigCLI_r16 struct {
	Value []RSSI_ResourceConfigCLI_r16 `lb:1,ub:maxNrofCLI_RSSI_Resources_r16,madatory`
}

func (ie *RSSI_ResourceListConfigCLI_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*RSSI_ResourceConfigCLI_r16]([]*RSSI_ResourceConfigCLI_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofCLI_RSSI_Resources_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode RSSI_ResourceListConfigCLI_r16", err)
	}
	return nil
}

func (ie *RSSI_ResourceListConfigCLI_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*RSSI_ResourceConfigCLI_r16]([]*RSSI_ResourceConfigCLI_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofCLI_RSSI_Resources_r16}, false)
	fn := func() *RSSI_ResourceConfigCLI_r16 {
		return new(RSSI_ResourceConfigCLI_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode RSSI_ResourceListConfigCLI_r16", err)
	}
	ie.Value = []RSSI_ResourceConfigCLI_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
