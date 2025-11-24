package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RegisteredAMF struct {
	Plmn_Identity  *PLMN_Identity `optional`
	Amf_Identifier AMF_Identifier `madatory`
}

func (ie *RegisteredAMF) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Plmn_Identity != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Plmn_Identity != nil {
		if err = ie.Plmn_Identity.Encode(w); err != nil {
			return utils.WrapError("Encode Plmn_Identity", err)
		}
	}
	if err = ie.Amf_Identifier.Encode(w); err != nil {
		return utils.WrapError("Encode Amf_Identifier", err)
	}
	return nil
}

func (ie *RegisteredAMF) Decode(r *uper.UperReader) error {
	var err error
	var Plmn_IdentityPresent bool
	if Plmn_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Plmn_IdentityPresent {
		ie.Plmn_Identity = new(PLMN_Identity)
		if err = ie.Plmn_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_Identity", err)
		}
	}
	if err = ie.Amf_Identifier.Decode(r); err != nil {
		return utils.WrapError("Decode Amf_Identifier", err)
	}
	return nil
}
