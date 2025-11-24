package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_ConfigSIB1 struct {
	ControlResourceSetZero ControlResourceSetZero `madatory`
	SearchSpaceZero        SearchSpaceZero        `madatory`
}

func (ie *PDCCH_ConfigSIB1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ControlResourceSetZero.Encode(w); err != nil {
		return utils.WrapError("Encode ControlResourceSetZero", err)
	}
	if err = ie.SearchSpaceZero.Encode(w); err != nil {
		return utils.WrapError("Encode SearchSpaceZero", err)
	}
	return nil
}

func (ie *PDCCH_ConfigSIB1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ControlResourceSetZero.Decode(r); err != nil {
		return utils.WrapError("Decode ControlResourceSetZero", err)
	}
	if err = ie.SearchSpaceZero.Decode(r); err != nil {
		return utils.WrapError("Decode SearchSpaceZero", err)
	}
	return nil
}
