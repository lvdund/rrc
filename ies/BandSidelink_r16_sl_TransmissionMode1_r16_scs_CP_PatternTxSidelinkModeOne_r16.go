package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16_Choice_nothing uint64 = iota
	BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16_Choice_Fr1_r16
	BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16_Choice_Fr2_r16
)

type BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16 struct {
	Choice  uint64
	Fr1_r16 *Fr1_r16
	Fr2_r16 *Fr2_r16
}

func (ie *BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16_Choice_Fr1_r16:
		if err = ie.Fr1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Fr1_r16", err)
		}
	case BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16_Choice_Fr2_r16:
		if err = ie.Fr2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Fr2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16_Choice_Fr1_r16:
		ie.Fr1_r16 = new(Fr1_r16)
		if err = ie.Fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_r16", err)
		}
	case BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16_Choice_Fr2_r16:
		ie.Fr2_r16 = new(Fr2_r16)
		if err = ie.Fr2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
