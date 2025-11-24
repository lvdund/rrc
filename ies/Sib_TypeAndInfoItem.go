package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Sib_TypeAndInfoItem_Choice_nothing uint64 = iota
	Sib_TypeAndInfoItem_Choice_Sib2
	Sib_TypeAndInfoItem_Choice_Sib3
	Sib_TypeAndInfoItem_Choice_Sib4
	Sib_TypeAndInfoItem_Choice_Sib5
	Sib_TypeAndInfoItem_Choice_Sib6
	Sib_TypeAndInfoItem_Choice_Sib7
	Sib_TypeAndInfoItem_Choice_Sib8
	Sib_TypeAndInfoItem_Choice_Sib9
	Sib_TypeAndInfoItem_Choice_Sib10_v1610
	Sib_TypeAndInfoItem_Choice_Sib11_v1610
	Sib_TypeAndInfoItem_Choice_Sib12_v1610
	Sib_TypeAndInfoItem_Choice_Sib13_v1610
	Sib_TypeAndInfoItem_Choice_Sib14_v1610
	Sib_TypeAndInfoItem_Choice_Sib15_v1700
	Sib_TypeAndInfoItem_Choice_Sib16_v1700
	Sib_TypeAndInfoItem_Choice_Sib17_v1700
	Sib_TypeAndInfoItem_Choice_Sib18_v1700
	Sib_TypeAndInfoItem_Choice_Sib19_v1700
	Sib_TypeAndInfoItem_Choice_Sib20_v1700
	Sib_TypeAndInfoItem_Choice_Sib21_v1700
)

type Sib_TypeAndInfoItem struct {
	Choice      uint64
	Sib2        *SIB2
	Sib3        *SIB3
	Sib4        *SIB4
	Sib5        *SIB5
	Sib6        *SIB6
	Sib7        *SIB7
	Sib8        *SIB8
	Sib9        *SIB9
	Sib10_v1610 *SIB10_r16
	Sib11_v1610 *SIB11_r16
	Sib12_v1610 *SIB12_r16
	Sib13_v1610 *SIB13_r16
	Sib14_v1610 *SIB14_r16
	Sib15_v1700 *SIB15_r17
	Sib16_v1700 *SIB16_r17
	Sib17_v1700 *SIB17_r17
	Sib18_v1700 *SIB18_r17
	Sib19_v1700 *SIB19_r17
	Sib20_v1700 *SIB20_r17
	Sib21_v1700 *SIB21_r17
}

func (ie *Sib_TypeAndInfoItem) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Sib_TypeAndInfoItem_Choice_Sib2:
		if err = ie.Sib2.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib2", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib3:
		if err = ie.Sib3.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib3", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib4:
		if err = ie.Sib4.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib4", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib5:
		if err = ie.Sib5.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib5", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib6:
		if err = ie.Sib6.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib6", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib7:
		if err = ie.Sib7.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib7", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib8:
		if err = ie.Sib8.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib8", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib9:
		if err = ie.Sib9.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib9", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib10_v1610:
		if err = ie.Sib10_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib10_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib11_v1610:
		if err = ie.Sib11_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib11_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib12_v1610:
		if err = ie.Sib12_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib12_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib13_v1610:
		if err = ie.Sib13_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib13_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib14_v1610:
		if err = ie.Sib14_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib14_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib15_v1700:
		if err = ie.Sib15_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib15_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib16_v1700:
		if err = ie.Sib16_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib16_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib17_v1700:
		if err = ie.Sib17_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib17_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib18_v1700:
		if err = ie.Sib18_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib18_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib19_v1700:
		if err = ie.Sib19_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib19_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib20_v1700:
		if err = ie.Sib20_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib20_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib21_v1700:
		if err = ie.Sib21_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode Sib21_v1700", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *Sib_TypeAndInfoItem) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Sib_TypeAndInfoItem_Choice_Sib2:
		ie.Sib2 = new(SIB2)
		if err = ie.Sib2.Decode(r); err != nil {
			return utils.WrapError("Decode Sib2", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib3:
		ie.Sib3 = new(SIB3)
		if err = ie.Sib3.Decode(r); err != nil {
			return utils.WrapError("Decode Sib3", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib4:
		ie.Sib4 = new(SIB4)
		if err = ie.Sib4.Decode(r); err != nil {
			return utils.WrapError("Decode Sib4", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib5:
		ie.Sib5 = new(SIB5)
		if err = ie.Sib5.Decode(r); err != nil {
			return utils.WrapError("Decode Sib5", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib6:
		ie.Sib6 = new(SIB6)
		if err = ie.Sib6.Decode(r); err != nil {
			return utils.WrapError("Decode Sib6", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib7:
		ie.Sib7 = new(SIB7)
		if err = ie.Sib7.Decode(r); err != nil {
			return utils.WrapError("Decode Sib7", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib8:
		ie.Sib8 = new(SIB8)
		if err = ie.Sib8.Decode(r); err != nil {
			return utils.WrapError("Decode Sib8", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib9:
		ie.Sib9 = new(SIB9)
		if err = ie.Sib9.Decode(r); err != nil {
			return utils.WrapError("Decode Sib9", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib10_v1610:
		ie.Sib10_v1610 = new(SIB10_r16)
		if err = ie.Sib10_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Sib10_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib11_v1610:
		ie.Sib11_v1610 = new(SIB11_r16)
		if err = ie.Sib11_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Sib11_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib12_v1610:
		ie.Sib12_v1610 = new(SIB12_r16)
		if err = ie.Sib12_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Sib12_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib13_v1610:
		ie.Sib13_v1610 = new(SIB13_r16)
		if err = ie.Sib13_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Sib13_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib14_v1610:
		ie.Sib14_v1610 = new(SIB14_r16)
		if err = ie.Sib14_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Sib14_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib15_v1700:
		ie.Sib15_v1700 = new(SIB15_r17)
		if err = ie.Sib15_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sib15_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib16_v1700:
		ie.Sib16_v1700 = new(SIB16_r17)
		if err = ie.Sib16_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sib16_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib17_v1700:
		ie.Sib17_v1700 = new(SIB17_r17)
		if err = ie.Sib17_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sib17_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib18_v1700:
		ie.Sib18_v1700 = new(SIB18_r17)
		if err = ie.Sib18_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sib18_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib19_v1700:
		ie.Sib19_v1700 = new(SIB19_r17)
		if err = ie.Sib19_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sib19_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib20_v1700:
		ie.Sib20_v1700 = new(SIB20_r17)
		if err = ie.Sib20_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sib20_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_Sib21_v1700:
		ie.Sib21_v1700 = new(SIB21_r17)
		if err = ie.Sib21_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Sib21_v1700", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
