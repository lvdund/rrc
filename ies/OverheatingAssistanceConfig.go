package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistanceConfig struct {
	OverheatingIndicationProhibitTimer OverheatingAssistanceConfig_overheatingIndicationProhibitTimer `madatory`
}

func (ie *OverheatingAssistanceConfig) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.OverheatingIndicationProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode OverheatingIndicationProhibitTimer", err)
	}
	return nil
}

func (ie *OverheatingAssistanceConfig) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.OverheatingIndicationProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode OverheatingIndicationProhibitTimer", err)
	}
	return nil
}
