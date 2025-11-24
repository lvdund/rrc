package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreqComb_r16 struct {
	AffectedCarrierFreqComb_r16 []ARFCN_ValueNR      `lb:2,ub:maxNrofServingCells,optional`
	VictimSystemType_r16        VictimSystemType_r16 `madatory`
}

func (ie *AffectedCarrierFreqComb_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.AffectedCarrierFreqComb_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.AffectedCarrierFreqComb_r16) > 0 {
		tmp_AffectedCarrierFreqComb_r16 := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 2, Ub: maxNrofServingCells}, false)
		for _, i := range ie.AffectedCarrierFreqComb_r16 {
			tmp_AffectedCarrierFreqComb_r16.Value = append(tmp_AffectedCarrierFreqComb_r16.Value, &i)
		}
		if err = tmp_AffectedCarrierFreqComb_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AffectedCarrierFreqComb_r16", err)
		}
	}
	if err = ie.VictimSystemType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode VictimSystemType_r16", err)
	}
	return nil
}

func (ie *AffectedCarrierFreqComb_r16) Decode(r *uper.UperReader) error {
	var err error
	var AffectedCarrierFreqComb_r16Present bool
	if AffectedCarrierFreqComb_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AffectedCarrierFreqComb_r16Present {
		tmp_AffectedCarrierFreqComb_r16 := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 2, Ub: maxNrofServingCells}, false)
		fn_AffectedCarrierFreqComb_r16 := func() *ARFCN_ValueNR {
			return new(ARFCN_ValueNR)
		}
		if err = tmp_AffectedCarrierFreqComb_r16.Decode(r, fn_AffectedCarrierFreqComb_r16); err != nil {
			return utils.WrapError("Decode AffectedCarrierFreqComb_r16", err)
		}
		ie.AffectedCarrierFreqComb_r16 = []ARFCN_ValueNR{}
		for _, i := range tmp_AffectedCarrierFreqComb_r16.Value {
			ie.AffectedCarrierFreqComb_r16 = append(ie.AffectedCarrierFreqComb_r16, *i)
		}
	}
	if err = ie.VictimSystemType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode VictimSystemType_r16", err)
	}
	return nil
}
