package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ProcessingParameters_differentTB_PerSlot struct {
	Upto1 *NumberOfCarriers `optional`
	Upto2 *NumberOfCarriers `optional`
	Upto4 *NumberOfCarriers `optional`
	Upto7 *NumberOfCarriers `optional`
}

func (ie *ProcessingParameters_differentTB_PerSlot) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Upto1 != nil, ie.Upto2 != nil, ie.Upto4 != nil, ie.Upto7 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Upto1 != nil {
		if err = ie.Upto1.Encode(w); err != nil {
			return utils.WrapError("Encode Upto1", err)
		}
	}
	if ie.Upto2 != nil {
		if err = ie.Upto2.Encode(w); err != nil {
			return utils.WrapError("Encode Upto2", err)
		}
	}
	if ie.Upto4 != nil {
		if err = ie.Upto4.Encode(w); err != nil {
			return utils.WrapError("Encode Upto4", err)
		}
	}
	if ie.Upto7 != nil {
		if err = ie.Upto7.Encode(w); err != nil {
			return utils.WrapError("Encode Upto7", err)
		}
	}
	return nil
}

func (ie *ProcessingParameters_differentTB_PerSlot) Decode(r *uper.UperReader) error {
	var err error
	var Upto1Present bool
	if Upto1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Upto2Present bool
	if Upto2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Upto4Present bool
	if Upto4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Upto7Present bool
	if Upto7Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Upto1Present {
		ie.Upto1 = new(NumberOfCarriers)
		if err = ie.Upto1.Decode(r); err != nil {
			return utils.WrapError("Decode Upto1", err)
		}
	}
	if Upto2Present {
		ie.Upto2 = new(NumberOfCarriers)
		if err = ie.Upto2.Decode(r); err != nil {
			return utils.WrapError("Decode Upto2", err)
		}
	}
	if Upto4Present {
		ie.Upto4 = new(NumberOfCarriers)
		if err = ie.Upto4.Decode(r); err != nil {
			return utils.WrapError("Decode Upto4", err)
		}
	}
	if Upto7Present {
		ie.Upto7 = new(NumberOfCarriers)
		if err = ie.Upto7.Decode(r); err != nil {
			return utils.WrapError("Decode Upto7", err)
		}
	}
	return nil
}
