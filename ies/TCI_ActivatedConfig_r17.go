package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_ActivatedConfig_r17 struct {
	Pdcch_TCI_r17 []TCI_StateId  `lb:1,ub:5,madatory`
	Pdsch_TCI_r17 uper.BitString `lb:1,ub:maxNrofTCI_States,madatory`
}

func (ie *TCI_ActivatedConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Pdcch_TCI_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: 5}, false)
	for _, i := range ie.Pdcch_TCI_r17 {
		tmp_Pdcch_TCI_r17.Value = append(tmp_Pdcch_TCI_r17.Value, &i)
	}
	if err = tmp_Pdcch_TCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcch_TCI_r17", err)
	}
	if err = w.WriteBitString(ie.Pdsch_TCI_r17.Bytes, uint(ie.Pdsch_TCI_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false); err != nil {
		return utils.WrapError("WriteBitString Pdsch_TCI_r17", err)
	}
	return nil
}

func (ie *TCI_ActivatedConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_Pdcch_TCI_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: 5}, false)
	fn_Pdcch_TCI_r17 := func() *TCI_StateId {
		return new(TCI_StateId)
	}
	if err = tmp_Pdcch_TCI_r17.Decode(r, fn_Pdcch_TCI_r17); err != nil {
		return utils.WrapError("Decode Pdcch_TCI_r17", err)
	}
	ie.Pdcch_TCI_r17 = []TCI_StateId{}
	for _, i := range tmp_Pdcch_TCI_r17.Value {
		ie.Pdcch_TCI_r17 = append(ie.Pdcch_TCI_r17, *i)
	}
	var tmp_bs_Pdsch_TCI_r17 []byte
	var n_Pdsch_TCI_r17 uint
	if tmp_bs_Pdsch_TCI_r17, n_Pdsch_TCI_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false); err != nil {
		return utils.WrapError("ReadBitString Pdsch_TCI_r17", err)
	}
	ie.Pdsch_TCI_r17 = uper.BitString{
		Bytes:   tmp_bs_Pdsch_TCI_r17,
		NumBits: uint64(n_Pdsch_TCI_r17),
	}
	return nil
}
