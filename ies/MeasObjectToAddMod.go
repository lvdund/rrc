package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectToAddMod struct {
	MeasObjectId MeasObjectId                  `madatory`
	MeasObject   MeasObjectToAddMod_measObject `madatory`
}

func (ie *MeasObjectToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MeasObjectId.Encode(w); err != nil {
		return utils.WrapError("Encode MeasObjectId", err)
	}
	if err = ie.MeasObject.Encode(w); err != nil {
		return utils.WrapError("Encode MeasObject", err)
	}
	return nil
}

func (ie *MeasObjectToAddMod) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MeasObjectId.Decode(r); err != nil {
		return utils.WrapError("Decode MeasObjectId", err)
	}
	if err = ie.MeasObject.Decode(r); err != nil {
		return utils.WrapError("Decode MeasObject", err)
	}
	return nil
}
