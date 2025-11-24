package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SecurityConfigSMC struct {
	SecurityAlgorithmConfig SecurityAlgorithmConfig `madatory`
}

func (ie *SecurityConfigSMC) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.SecurityAlgorithmConfig.Encode(w); err != nil {
		return utils.WrapError("Encode SecurityAlgorithmConfig", err)
	}
	return nil
}

func (ie *SecurityConfigSMC) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.SecurityAlgorithmConfig.Decode(r); err != nil {
		return utils.WrapError("Decode SecurityAlgorithmConfig", err)
	}
	return nil
}
