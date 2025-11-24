package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Info2 struct {
	Drx_onDurationTimer DRX_Info2_drx_onDurationTimer `lb:1,ub:31,madatory`
}

func (ie *DRX_Info2) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Drx_onDurationTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_onDurationTimer", err)
	}
	return nil
}

func (ie *DRX_Info2) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Drx_onDurationTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_onDurationTimer", err)
	}
	return nil
}
