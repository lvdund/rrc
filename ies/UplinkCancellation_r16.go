package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkCancellation_r16 struct {
	Ci_RNTI_r16                        RNTI_Value                           `madatory`
	Dci_PayloadSizeForCI_r16           int64                                `lb:0,ub:maxCI_DCI_PayloadSize_r16,madatory`
	Ci_ConfigurationPerServingCell_r16 []CI_ConfigurationPerServingCell_r16 `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *UplinkCancellation_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ci_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ci_RNTI_r16", err)
	}
	if err = w.WriteInteger(ie.Dci_PayloadSizeForCI_r16, &uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("WriteInteger Dci_PayloadSizeForCI_r16", err)
	}
	tmp_Ci_ConfigurationPerServingCell_r16 := utils.NewSequence[*CI_ConfigurationPerServingCell_r16]([]*CI_ConfigurationPerServingCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Ci_ConfigurationPerServingCell_r16 {
		tmp_Ci_ConfigurationPerServingCell_r16.Value = append(tmp_Ci_ConfigurationPerServingCell_r16.Value, &i)
	}
	if err = tmp_Ci_ConfigurationPerServingCell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ci_ConfigurationPerServingCell_r16", err)
	}
	return nil
}

func (ie *UplinkCancellation_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ci_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ci_RNTI_r16", err)
	}
	var tmp_int_Dci_PayloadSizeForCI_r16 int64
	if tmp_int_Dci_PayloadSizeForCI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("ReadInteger Dci_PayloadSizeForCI_r16", err)
	}
	ie.Dci_PayloadSizeForCI_r16 = tmp_int_Dci_PayloadSizeForCI_r16
	tmp_Ci_ConfigurationPerServingCell_r16 := utils.NewSequence[*CI_ConfigurationPerServingCell_r16]([]*CI_ConfigurationPerServingCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_Ci_ConfigurationPerServingCell_r16 := func() *CI_ConfigurationPerServingCell_r16 {
		return new(CI_ConfigurationPerServingCell_r16)
	}
	if err = tmp_Ci_ConfigurationPerServingCell_r16.Decode(r, fn_Ci_ConfigurationPerServingCell_r16); err != nil {
		return utils.WrapError("Decode Ci_ConfigurationPerServingCell_r16", err)
	}
	ie.Ci_ConfigurationPerServingCell_r16 = []CI_ConfigurationPerServingCell_r16{}
	for _, i := range tmp_Ci_ConfigurationPerServingCell_r16.Value {
		ie.Ci_ConfigurationPerServingCell_r16 = append(ie.Ci_ConfigurationPerServingCell_r16, *i)
	}
	return nil
}
