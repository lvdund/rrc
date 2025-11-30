package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PerRAInfo_r16_Choice_nothing uint64 = iota
	PerRAInfo_r16_Choice_PerRASSBInfoList_r16
	PerRAInfo_r16_Choice_PerRACSI_RSInfoList_r16
)

type PerRAInfo_r16 struct {
	Choice                  uint64
	PerRASSBInfoList_r16    *PerRASSBInfo_r16
	PerRACSI_RSInfoList_r16 *PerRACSI_RSInfo_r16
}

func (ie *PerRAInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PerRAInfo_r16_Choice_PerRASSBInfoList_r16:
		if err = ie.PerRASSBInfoList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PerRASSBInfoList_r16", err)
		}
	case PerRAInfo_r16_Choice_PerRACSI_RSInfoList_r16:
		if err = ie.PerRACSI_RSInfoList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PerRACSI_RSInfoList_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PerRAInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PerRAInfo_r16_Choice_PerRASSBInfoList_r16:
		ie.PerRASSBInfoList_r16 = new(PerRASSBInfo_r16)
		if err = ie.PerRASSBInfoList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PerRASSBInfoList_r16", err)
		}
	case PerRAInfo_r16_Choice_PerRACSI_RSInfoList_r16:
		ie.PerRACSI_RSInfoList_r16 = new(PerRACSI_RSInfo_r16)
		if err = ie.PerRACSI_RSInfoList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PerRACSI_RSInfoList_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
