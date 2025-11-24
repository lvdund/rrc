package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_IEs_deprioritisationReq struct {
	DeprioritisationType  RRCRelease_IEs_deprioritisationReq_deprioritisationType  `madatory`
	DeprioritisationTimer RRCRelease_IEs_deprioritisationReq_deprioritisationTimer `madatory`
}

func (ie *RRCRelease_IEs_deprioritisationReq) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.DeprioritisationType.Encode(w); err != nil {
		return utils.WrapError("Encode DeprioritisationType", err)
	}
	if err = ie.DeprioritisationTimer.Encode(w); err != nil {
		return utils.WrapError("Encode DeprioritisationTimer", err)
	}
	return nil
}

func (ie *RRCRelease_IEs_deprioritisationReq) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.DeprioritisationType.Decode(r); err != nil {
		return utils.WrapError("Decode DeprioritisationType", err)
	}
	if err = ie.DeprioritisationTimer.Decode(r); err != nil {
		return utils.WrapError("Decode DeprioritisationTimer", err)
	}
	return nil
}
