package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_resourceType_semi_persistent struct {
	PeriodicityAndOffset_sp SRS_PeriodicityAndOffset `madatory`
}

func (ie *SRS_Resource_resourceType_semi_persistent) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.PeriodicityAndOffset_sp.Encode(w); err != nil {
		return utils.WrapError("Encode PeriodicityAndOffset_sp", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceType_semi_persistent) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.PeriodicityAndOffset_sp.Decode(r); err != nil {
		return utils.WrapError("Decode PeriodicityAndOffset_sp", err)
	}
	return nil
}
