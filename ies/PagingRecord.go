package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PagingRecord struct {
	Ue_Identity PagingUE_Identity        `madatory`
	AccessType  *PagingRecord_accessType `optional`
}

func (ie *PagingRecord) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AccessType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ue_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode Ue_Identity", err)
	}
	if ie.AccessType != nil {
		if err = ie.AccessType.Encode(w); err != nil {
			return utils.WrapError("Encode AccessType", err)
		}
	}
	return nil
}

func (ie *PagingRecord) Decode(r *aper.AperReader) error {
	var err error
	var AccessTypePresent bool
	if AccessTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ue_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode Ue_Identity", err)
	}
	if AccessTypePresent {
		ie.AccessType = new(PagingRecord_accessType)
		if err = ie.AccessType.Decode(r); err != nil {
			return utils.WrapError("Decode AccessType", err)
		}
	}
	return nil
}
