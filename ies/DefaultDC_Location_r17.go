package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DefaultDC_Location_r17_Choice_nothing uint64 = iota
	DefaultDC_Location_r17_Choice_Ul
	DefaultDC_Location_r17_Choice_Dl
	DefaultDC_Location_r17_Choice_UlAndDL
)

type DefaultDC_Location_r17 struct {
	Choice  uint64
	Ul      *FrequencyComponent_r17
	Dl      *FrequencyComponent_r17
	UlAndDL *FrequencyComponent_r17
}

func (ie *DefaultDC_Location_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DefaultDC_Location_r17_Choice_Ul:
		if err = ie.Ul.Encode(w); err != nil {
			err = utils.WrapError("Encode Ul", err)
		}
	case DefaultDC_Location_r17_Choice_Dl:
		if err = ie.Dl.Encode(w); err != nil {
			err = utils.WrapError("Encode Dl", err)
		}
	case DefaultDC_Location_r17_Choice_UlAndDL:
		if err = ie.UlAndDL.Encode(w); err != nil {
			err = utils.WrapError("Encode UlAndDL", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DefaultDC_Location_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DefaultDC_Location_r17_Choice_Ul:
		ie.Ul = new(FrequencyComponent_r17)
		if err = ie.Ul.Decode(r); err != nil {
			return utils.WrapError("Decode Ul", err)
		}
	case DefaultDC_Location_r17_Choice_Dl:
		ie.Dl = new(FrequencyComponent_r17)
		if err = ie.Dl.Decode(r); err != nil {
			return utils.WrapError("Decode Dl", err)
		}
	case DefaultDC_Location_r17_Choice_UlAndDL:
		ie.UlAndDL = new(FrequencyComponent_r17)
		if err = ie.UlAndDL.Decode(r); err != nil {
			return utils.WrapError("Decode UlAndDL", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
