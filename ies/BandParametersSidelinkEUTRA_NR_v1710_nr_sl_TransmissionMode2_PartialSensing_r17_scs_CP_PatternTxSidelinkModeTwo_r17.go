package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_nothing uint64 = iota
	BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_Fr1_r17
	BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_Fr2_r17
)

type BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17 struct {
	Choice  uint64
	Fr1_r17 *Fr1_r17
	Fr2_r17 *Fr2_r17
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_Fr1_r17:
		if err = ie.Fr1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Fr1_r17", err)
		}
	case BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_Fr2_r17:
		if err = ie.Fr2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Fr2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_Fr1_r17:
		ie.Fr1_r17 = new(Fr1_r17)
		if err = ie.Fr1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_r17", err)
		}
	case BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_Fr2_r17:
		ie.Fr2_r17 = new(Fr2_r17)
		if err = ie.Fr2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
