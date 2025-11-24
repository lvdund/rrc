package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_occasions struct {
	Rach_ConfigGeneric   RACH_ConfigGeneric                   `madatory`
	Ssb_perRACH_Occasion *CFRA_occasions_ssb_perRACH_Occasion `optional`
}

func (ie *CFRA_occasions) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_perRACH_Occasion != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Rach_ConfigGeneric.Encode(w); err != nil {
		return utils.WrapError("Encode Rach_ConfigGeneric", err)
	}
	if ie.Ssb_perRACH_Occasion != nil {
		if err = ie.Ssb_perRACH_Occasion.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_perRACH_Occasion", err)
		}
	}
	return nil
}

func (ie *CFRA_occasions) Decode(r *uper.UperReader) error {
	var err error
	var Ssb_perRACH_OccasionPresent bool
	if Ssb_perRACH_OccasionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Rach_ConfigGeneric.Decode(r); err != nil {
		return utils.WrapError("Decode Rach_ConfigGeneric", err)
	}
	if Ssb_perRACH_OccasionPresent {
		ie.Ssb_perRACH_Occasion = new(CFRA_occasions_ssb_perRACH_Occasion)
		if err = ie.Ssb_perRACH_Occasion.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_perRACH_Occasion", err)
		}
	}
	return nil
}
