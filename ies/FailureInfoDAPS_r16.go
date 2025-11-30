package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FailureInfoDAPS_r16 struct {
	FailureType_r16 FailureInfoDAPS_r16_failureType_r16 `madatory`
}

func (ie *FailureInfoDAPS_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.FailureType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FailureType_r16", err)
	}
	return nil
}

func (ie *FailureInfoDAPS_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.FailureType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FailureType_r16", err)
	}
	return nil
}
