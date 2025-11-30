package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_Identity struct {
	Mcc *MCC `optional`
	Mnc MNC  `madatory`
}

func (ie *PLMN_Identity) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mcc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mcc != nil {
		if err = ie.Mcc.Encode(w); err != nil {
			return utils.WrapError("Encode Mcc", err)
		}
	}
	if err = ie.Mnc.Encode(w); err != nil {
		return utils.WrapError("Encode Mnc", err)
	}
	return nil
}

func (ie *PLMN_Identity) Decode(r *aper.AperReader) error {
	var err error
	var MccPresent bool
	if MccPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MccPresent {
		ie.Mcc = new(MCC)
		if err = ie.Mcc.Decode(r); err != nil {
			return utils.WrapError("Decode Mcc", err)
		}
	}
	if err = ie.Mnc.Decode(r); err != nil {
		return utils.WrapError("Decode Mnc", err)
	}
	return nil
}
