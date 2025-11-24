package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type1 struct {
	SinglePanel CodebookParameters_type1_singlePanel `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	MultiPanel  *CodebookParameters_type1_multiPanel `lb:1,ub:maxNrofCSI_RS_Resources,optional`
}

func (ie *CodebookParameters_type1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MultiPanel != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SinglePanel.Encode(w); err != nil {
		return utils.WrapError("Encode SinglePanel", err)
	}
	if ie.MultiPanel != nil {
		if err = ie.MultiPanel.Encode(w); err != nil {
			return utils.WrapError("Encode MultiPanel", err)
		}
	}
	return nil
}

func (ie *CodebookParameters_type1) Decode(r *uper.UperReader) error {
	var err error
	var MultiPanelPresent bool
	if MultiPanelPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SinglePanel.Decode(r); err != nil {
		return utils.WrapError("Decode SinglePanel", err)
	}
	if MultiPanelPresent {
		ie.MultiPanel = new(CodebookParameters_type1_multiPanel)
		if err = ie.MultiPanel.Decode(r); err != nil {
			return utils.WrapError("Decode MultiPanel", err)
		}
	}
	return nil
}
