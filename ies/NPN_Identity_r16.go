package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NPN_Identity_r16_Choice_nothing uint64 = iota
	NPN_Identity_r16_Choice_Pni_npn_r16
	NPN_Identity_r16_Choice_Snpn_r16
)

type NPN_Identity_r16 struct {
	Choice      uint64
	Pni_npn_r16 *NPN_Identity_r16_pni_npn_r16
	Snpn_r16    *NPN_Identity_r16_snpn_r16
}

func (ie *NPN_Identity_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NPN_Identity_r16_Choice_Pni_npn_r16:
		if err = ie.Pni_npn_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Pni_npn_r16", err)
		}
	case NPN_Identity_r16_Choice_Snpn_r16:
		if err = ie.Snpn_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Snpn_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NPN_Identity_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NPN_Identity_r16_Choice_Pni_npn_r16:
		ie.Pni_npn_r16 = new(NPN_Identity_r16_pni_npn_r16)
		if err = ie.Pni_npn_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pni_npn_r16", err)
		}
	case NPN_Identity_r16_Choice_Snpn_r16:
		ie.Snpn_r16 = new(NPN_Identity_r16_snpn_r16)
		if err = ie.Snpn_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Snpn_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
