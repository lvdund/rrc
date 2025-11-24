package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RepetitionSchemeConfig_v1630 struct {
	SlotBased_v1630 *SlotBased_v1630 `madatory,setuprelease`
}

func (ie *RepetitionSchemeConfig_v1630) Encode(w *uper.UperWriter) error {
	var err error
	tmp_SlotBased_v1630 := utils.SetupRelease[*SlotBased_v1630]{
		Setup: ie.SlotBased_v1630,
	}
	if err = tmp_SlotBased_v1630.Encode(w); err != nil {
		return utils.WrapError("Encode SlotBased_v1630", err)
	}
	return nil
}

func (ie *RepetitionSchemeConfig_v1630) Decode(r *uper.UperReader) error {
	var err error
	tmp_SlotBased_v1630 := utils.SetupRelease[*SlotBased_v1630]{}
	if err = tmp_SlotBased_v1630.Decode(r); err != nil {
		return utils.WrapError("Decode SlotBased_v1630", err)
	}
	ie.SlotBased_v1630 = tmp_SlotBased_v1630.Setup
	return nil
}
