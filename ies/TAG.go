package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TAG struct {
	Tag_Id             TAG_Id             `madatory`
	TimeAlignmentTimer TimeAlignmentTimer `madatory`
}

func (ie *TAG) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Tag_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Tag_Id", err)
	}
	if err = ie.TimeAlignmentTimer.Encode(w); err != nil {
		return utils.WrapError("Encode TimeAlignmentTimer", err)
	}
	return nil
}

func (ie *TAG) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Tag_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Tag_Id", err)
	}
	if err = ie.TimeAlignmentTimer.Decode(r); err != nil {
		return utils.WrapError("Decode TimeAlignmentTimer", err)
	}
	return nil
}
