package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSIB_TypeAndInfo_r16_Choice_nothing uint64 = iota
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_1_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_2_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_3_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_4_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_5_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_6_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_7_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_8_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_1_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_2_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_3_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_4_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_5_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_6_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_7_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_8_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_9_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_10_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_11_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_12_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_13_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_14_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_15_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_16_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_17_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_18_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_19_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_20_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_21_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_22_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_23_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib3_1_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib4_1_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib5_1_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib6_1_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib6_2_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib6_3_r16
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_9_v1700
	PosSIB_TypeAndInfo_r16_Choice_PosSib1_10_v1700
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_24_v1700
	PosSIB_TypeAndInfo_r16_Choice_PosSib2_25_v1700
	PosSIB_TypeAndInfo_r16_Choice_PosSib6_4_v1700
	PosSIB_TypeAndInfo_r16_Choice_PosSib6_5_v1700
	PosSIB_TypeAndInfo_r16_Choice_PosSib6_6_v1700
)

type PosSIB_TypeAndInfo_r16 struct {
	Choice           uint64
	PosSib1_1_r16    *SIBpos_r16
	PosSib1_2_r16    *SIBpos_r16
	PosSib1_3_r16    *SIBpos_r16
	PosSib1_4_r16    *SIBpos_r16
	PosSib1_5_r16    *SIBpos_r16
	PosSib1_6_r16    *SIBpos_r16
	PosSib1_7_r16    *SIBpos_r16
	PosSib1_8_r16    *SIBpos_r16
	PosSib2_1_r16    *SIBpos_r16
	PosSib2_2_r16    *SIBpos_r16
	PosSib2_3_r16    *SIBpos_r16
	PosSib2_4_r16    *SIBpos_r16
	PosSib2_5_r16    *SIBpos_r16
	PosSib2_6_r16    *SIBpos_r16
	PosSib2_7_r16    *SIBpos_r16
	PosSib2_8_r16    *SIBpos_r16
	PosSib2_9_r16    *SIBpos_r16
	PosSib2_10_r16   *SIBpos_r16
	PosSib2_11_r16   *SIBpos_r16
	PosSib2_12_r16   *SIBpos_r16
	PosSib2_13_r16   *SIBpos_r16
	PosSib2_14_r16   *SIBpos_r16
	PosSib2_15_r16   *SIBpos_r16
	PosSib2_16_r16   *SIBpos_r16
	PosSib2_17_r16   *SIBpos_r16
	PosSib2_18_r16   *SIBpos_r16
	PosSib2_19_r16   *SIBpos_r16
	PosSib2_20_r16   *SIBpos_r16
	PosSib2_21_r16   *SIBpos_r16
	PosSib2_22_r16   *SIBpos_r16
	PosSib2_23_r16   *SIBpos_r16
	PosSib3_1_r16    *SIBpos_r16
	PosSib4_1_r16    *SIBpos_r16
	PosSib5_1_r16    *SIBpos_r16
	PosSib6_1_r16    *SIBpos_r16
	PosSib6_2_r16    *SIBpos_r16
	PosSib6_3_r16    *SIBpos_r16
	PosSib1_9_v1700  *SIBpos_r16
	PosSib1_10_v1700 *SIBpos_r16
	PosSib2_24_v1700 *SIBpos_r16
	PosSib2_25_v1700 *SIBpos_r16
	PosSib6_4_v1700  *SIBpos_r16
	PosSib6_5_v1700  *SIBpos_r16
	PosSib6_6_v1700  *SIBpos_r16
}

func (ie *PosSIB_TypeAndInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 44, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_1_r16:
		if err = ie.PosSib1_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_2_r16:
		if err = ie.PosSib1_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_3_r16:
		if err = ie.PosSib1_3_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_4_r16:
		if err = ie.PosSib1_4_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_5_r16:
		if err = ie.PosSib1_5_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_6_r16:
		if err = ie.PosSib1_6_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_7_r16:
		if err = ie.PosSib1_7_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_8_r16:
		if err = ie.PosSib1_8_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_1_r16:
		if err = ie.PosSib2_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_2_r16:
		if err = ie.PosSib2_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_3_r16:
		if err = ie.PosSib2_3_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_4_r16:
		if err = ie.PosSib2_4_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_5_r16:
		if err = ie.PosSib2_5_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_6_r16:
		if err = ie.PosSib2_6_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_7_r16:
		if err = ie.PosSib2_7_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_8_r16:
		if err = ie.PosSib2_8_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_9_r16:
		if err = ie.PosSib2_9_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_9_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_10_r16:
		if err = ie.PosSib2_10_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_10_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_11_r16:
		if err = ie.PosSib2_11_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_11_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_12_r16:
		if err = ie.PosSib2_12_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_12_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_13_r16:
		if err = ie.PosSib2_13_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_13_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_14_r16:
		if err = ie.PosSib2_14_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_14_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_15_r16:
		if err = ie.PosSib2_15_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_15_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_16_r16:
		if err = ie.PosSib2_16_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_16_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_17_r16:
		if err = ie.PosSib2_17_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_17_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_18_r16:
		if err = ie.PosSib2_18_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_18_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_19_r16:
		if err = ie.PosSib2_19_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_19_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_20_r16:
		if err = ie.PosSib2_20_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_20_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_21_r16:
		if err = ie.PosSib2_21_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_21_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_22_r16:
		if err = ie.PosSib2_22_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_22_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_23_r16:
		if err = ie.PosSib2_23_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_23_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib3_1_r16:
		if err = ie.PosSib3_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib3_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib4_1_r16:
		if err = ie.PosSib4_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib4_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib5_1_r16:
		if err = ie.PosSib5_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib5_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_1_r16:
		if err = ie.PosSib6_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib6_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_2_r16:
		if err = ie.PosSib6_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib6_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_3_r16:
		if err = ie.PosSib6_3_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib6_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_9_v1700:
		if err = ie.PosSib1_9_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_9_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_10_v1700:
		if err = ie.PosSib1_10_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib1_10_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_24_v1700:
		if err = ie.PosSib2_24_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_24_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_25_v1700:
		if err = ie.PosSib2_25_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib2_25_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_4_v1700:
		if err = ie.PosSib6_4_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib6_4_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_5_v1700:
		if err = ie.PosSib6_5_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib6_5_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_6_v1700:
		if err = ie.PosSib6_6_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSib6_6_v1700", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PosSIB_TypeAndInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(44, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_1_r16:
		ie.PosSib1_1_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_2_r16:
		ie.PosSib1_2_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_3_r16:
		ie.PosSib1_3_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_4_r16:
		ie.PosSib1_4_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_5_r16:
		ie.PosSib1_5_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_6_r16:
		ie.PosSib1_6_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_6_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_7_r16:
		ie.PosSib1_7_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_7_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_8_r16:
		ie.PosSib1_8_r16 = new(SIBpos_r16)
		if err = ie.PosSib1_8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_1_r16:
		ie.PosSib2_1_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_2_r16:
		ie.PosSib2_2_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_3_r16:
		ie.PosSib2_3_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_4_r16:
		ie.PosSib2_4_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_5_r16:
		ie.PosSib2_5_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_6_r16:
		ie.PosSib2_6_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_6_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_7_r16:
		ie.PosSib2_7_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_7_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_8_r16:
		ie.PosSib2_8_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_9_r16:
		ie.PosSib2_9_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_9_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_9_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_10_r16:
		ie.PosSib2_10_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_10_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_10_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_11_r16:
		ie.PosSib2_11_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_11_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_11_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_12_r16:
		ie.PosSib2_12_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_12_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_12_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_13_r16:
		ie.PosSib2_13_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_13_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_13_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_14_r16:
		ie.PosSib2_14_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_14_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_14_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_15_r16:
		ie.PosSib2_15_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_15_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_15_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_16_r16:
		ie.PosSib2_16_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_16_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_16_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_17_r16:
		ie.PosSib2_17_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_17_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_17_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_18_r16:
		ie.PosSib2_18_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_18_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_18_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_19_r16:
		ie.PosSib2_19_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_19_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_19_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_20_r16:
		ie.PosSib2_20_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_20_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_20_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_21_r16:
		ie.PosSib2_21_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_21_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_21_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_22_r16:
		ie.PosSib2_22_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_22_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_22_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_23_r16:
		ie.PosSib2_23_r16 = new(SIBpos_r16)
		if err = ie.PosSib2_23_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_23_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib3_1_r16:
		ie.PosSib3_1_r16 = new(SIBpos_r16)
		if err = ie.PosSib3_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib3_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib4_1_r16:
		ie.PosSib4_1_r16 = new(SIBpos_r16)
		if err = ie.PosSib4_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib4_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib5_1_r16:
		ie.PosSib5_1_r16 = new(SIBpos_r16)
		if err = ie.PosSib5_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib5_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_1_r16:
		ie.PosSib6_1_r16 = new(SIBpos_r16)
		if err = ie.PosSib6_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib6_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_2_r16:
		ie.PosSib6_2_r16 = new(SIBpos_r16)
		if err = ie.PosSib6_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib6_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_3_r16:
		ie.PosSib6_3_r16 = new(SIBpos_r16)
		if err = ie.PosSib6_3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib6_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_9_v1700:
		ie.PosSib1_9_v1700 = new(SIBpos_r16)
		if err = ie.PosSib1_9_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_9_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib1_10_v1700:
		ie.PosSib1_10_v1700 = new(SIBpos_r16)
		if err = ie.PosSib1_10_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib1_10_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_24_v1700:
		ie.PosSib2_24_v1700 = new(SIBpos_r16)
		if err = ie.PosSib2_24_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_24_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib2_25_v1700:
		ie.PosSib2_25_v1700 = new(SIBpos_r16)
		if err = ie.PosSib2_25_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib2_25_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_4_v1700:
		ie.PosSib6_4_v1700 = new(SIBpos_r16)
		if err = ie.PosSib6_4_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib6_4_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_5_v1700:
		ie.PosSib6_5_v1700 = new(SIBpos_r16)
		if err = ie.PosSib6_5_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib6_5_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_PosSib6_6_v1700:
		ie.PosSib6_6_v1700 = new(SIBpos_r16)
		if err = ie.PosSib6_6_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PosSib6_6_v1700", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
