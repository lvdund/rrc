package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16_Choice_nothing uint64 = iota
	SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16_Choice_Sl_OneAntennaPort_r16
	SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16_Choice_Sl_TwoAntennaPort_r16
)

type SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16 struct {
	Choice                uint64
	Sl_OneAntennaPort_r16 aper.BitString `lb:12,ub:12,madatory`
	Sl_TwoAntennaPort_r16 aper.BitString `lb:6,ub:6,madatory`
}

func (ie *SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16_Choice_Sl_OneAntennaPort_r16:
		if err = w.WriteBitString(ie.Sl_OneAntennaPort_r16.Bytes, uint(ie.Sl_OneAntennaPort_r16.NumBits), &aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			err = utils.WrapError("Encode Sl_OneAntennaPort_r16", err)
		}
	case SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16_Choice_Sl_TwoAntennaPort_r16:
		if err = w.WriteBitString(ie.Sl_TwoAntennaPort_r16.Bytes, uint(ie.Sl_TwoAntennaPort_r16.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			err = utils.WrapError("Encode Sl_TwoAntennaPort_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16_Choice_Sl_OneAntennaPort_r16:
		var tmp_bs_Sl_OneAntennaPort_r16 []byte
		var n_Sl_OneAntennaPort_r16 uint
		if tmp_bs_Sl_OneAntennaPort_r16, n_Sl_OneAntennaPort_r16, err = r.ReadBitString(&aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode Sl_OneAntennaPort_r16", err)
		}
		ie.Sl_OneAntennaPort_r16 = aper.BitString{
			Bytes:   tmp_bs_Sl_OneAntennaPort_r16,
			NumBits: uint64(n_Sl_OneAntennaPort_r16),
		}
	case SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16_Choice_Sl_TwoAntennaPort_r16:
		var tmp_bs_Sl_TwoAntennaPort_r16 []byte
		var n_Sl_TwoAntennaPort_r16 uint
		if tmp_bs_Sl_TwoAntennaPort_r16, n_Sl_TwoAntennaPort_r16, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode Sl_TwoAntennaPort_r16", err)
		}
		ie.Sl_TwoAntennaPort_r16 = aper.BitString{
			Bytes:   tmp_bs_Sl_TwoAntennaPort_r16,
			NumBits: uint64(n_Sl_TwoAntennaPort_r16),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
