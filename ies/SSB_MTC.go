package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC struct {
	PeriodicityAndOffset SSB_MTC_periodicityAndOffset `lb:0,ub:4,madatory`
	Duration             SSB_MTC_duration             `madatory`
}

func (ie *SSB_MTC) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.PeriodicityAndOffset.Encode(w); err != nil {
		return utils.WrapError("Encode PeriodicityAndOffset", err)
	}
	if err = ie.Duration.Encode(w); err != nil {
		return utils.WrapError("Encode Duration", err)
	}
	return nil
}

func (ie *SSB_MTC) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.PeriodicityAndOffset.Decode(r); err != nil {
		return utils.WrapError("Decode PeriodicityAndOffset", err)
	}
	if err = ie.Duration.Decode(r); err != nil {
		return utils.WrapError("Decode Duration", err)
	}
	return nil
}
