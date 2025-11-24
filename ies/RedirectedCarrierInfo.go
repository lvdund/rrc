package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RedirectedCarrierInfo_Choice_nothing uint64 = iota
	RedirectedCarrierInfo_Choice_Nr
	RedirectedCarrierInfo_Choice_Eutra
)

type RedirectedCarrierInfo struct {
	Choice uint64
	Nr     *CarrierInfoNR
	Eutra  *RedirectedCarrierInfo_EUTRA
}

func (ie *RedirectedCarrierInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RedirectedCarrierInfo_Choice_Nr:
		if err = ie.Nr.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr", err)
		}
	case RedirectedCarrierInfo_Choice_Eutra:
		if err = ie.Eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode Eutra", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RedirectedCarrierInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RedirectedCarrierInfo_Choice_Nr:
		ie.Nr = new(CarrierInfoNR)
		if err = ie.Nr.Decode(r); err != nil {
			return utils.WrapError("Decode Nr", err)
		}
	case RedirectedCarrierInfo_Choice_Eutra:
		ie.Eutra = new(RedirectedCarrierInfo_EUTRA)
		if err = ie.Eutra.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
