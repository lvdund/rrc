package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_nothing uint64 = iota
	RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_Ng_5G_S_TMSI
	RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_Ng_5G_S_TMSI_Part2
)

type RRCSetupComplete_IEs_ng_5G_S_TMSI_Value struct {
	Choice             uint64
	Ng_5G_S_TMSI       *NG_5G_S_TMSI
	Ng_5G_S_TMSI_Part2 aper.BitString `lb:9,ub:9,madatory`
}

func (ie *RRCSetupComplete_IEs_ng_5G_S_TMSI_Value) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_Ng_5G_S_TMSI:
		if err = ie.Ng_5G_S_TMSI.Encode(w); err != nil {
			err = utils.WrapError("Encode Ng_5G_S_TMSI", err)
		}
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_Ng_5G_S_TMSI_Part2:
		if err = w.WriteBitString(ie.Ng_5G_S_TMSI_Part2.Bytes, uint(ie.Ng_5G_S_TMSI_Part2.NumBits), &aper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Ng_5G_S_TMSI_Part2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSetupComplete_IEs_ng_5G_S_TMSI_Value) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_Ng_5G_S_TMSI:
		ie.Ng_5G_S_TMSI = new(NG_5G_S_TMSI)
		if err = ie.Ng_5G_S_TMSI.Decode(r); err != nil {
			return utils.WrapError("Decode Ng_5G_S_TMSI", err)
		}
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_Ng_5G_S_TMSI_Part2:
		var tmp_bs_Ng_5G_S_TMSI_Part2 []byte
		var n_Ng_5G_S_TMSI_Part2 uint
		if tmp_bs_Ng_5G_S_TMSI_Part2, n_Ng_5G_S_TMSI_Part2, err = r.ReadBitString(&aper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Ng_5G_S_TMSI_Part2", err)
		}
		ie.Ng_5G_S_TMSI_Part2 = aper.BitString{
			Bytes:   tmp_bs_Ng_5G_S_TMSI_Part2,
			NumBits: uint64(n_Ng_5G_S_TMSI_Part2),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
