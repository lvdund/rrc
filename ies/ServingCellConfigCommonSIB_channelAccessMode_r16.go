package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_nothing uint64 = iota
	ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_Dynamic
	ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_SemiStatic
)

type ServingCellConfigCommonSIB_channelAccessMode_r16 struct {
	Choice     uint64
	Dynamic    uper.NULL `madatory`
	SemiStatic *SemiStaticChannelAccessConfig_r16
}

func (ie *ServingCellConfigCommonSIB_channelAccessMode_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_Dynamic:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Dynamic", err)
		}
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_SemiStatic:
		if err = ie.SemiStatic.Encode(w); err != nil {
			err = utils.WrapError("Encode SemiStatic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ServingCellConfigCommonSIB_channelAccessMode_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_Dynamic:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Dynamic", err)
		}
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_SemiStatic:
		ie.SemiStatic = new(SemiStaticChannelAccessConfig_r16)
		if err = ie.SemiStatic.Decode(r); err != nil {
			return utils.WrapError("Decode SemiStatic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
