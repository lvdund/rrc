package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PagingUE_Identity_Choice_nothing uint64 = iota
	PagingUE_Identity_Choice_Ng_5G_S_TMSI
	PagingUE_Identity_Choice_FullI_RNTI
)

type PagingUE_Identity struct {
	Choice       uint64
	Ng_5G_S_TMSI *NG_5G_S_TMSI
	FullI_RNTI   *I_RNTI_Value
}

func (ie *PagingUE_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PagingUE_Identity_Choice_Ng_5G_S_TMSI:
		if err = ie.Ng_5G_S_TMSI.Encode(w); err != nil {
			err = utils.WrapError("Encode Ng_5G_S_TMSI", err)
		}
	case PagingUE_Identity_Choice_FullI_RNTI:
		if err = ie.FullI_RNTI.Encode(w); err != nil {
			err = utils.WrapError("Encode FullI_RNTI", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PagingUE_Identity) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PagingUE_Identity_Choice_Ng_5G_S_TMSI:
		ie.Ng_5G_S_TMSI = new(NG_5G_S_TMSI)
		if err = ie.Ng_5G_S_TMSI.Decode(r); err != nil {
			return utils.WrapError("Decode Ng_5G_S_TMSI", err)
		}
	case PagingUE_Identity_Choice_FullI_RNTI:
		ie.FullI_RNTI = new(I_RNTI_Value)
		if err = ie.FullI_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode FullI_RNTI", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
