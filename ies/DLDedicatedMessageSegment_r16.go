package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DLDedicatedMessageSegment_r16 struct {
	CriticalExtensions DLDedicatedMessageSegment_r16_CriticalExtensions `madatory`
}

func (ie *DLDedicatedMessageSegment_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CriticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode CriticalExtensions", err)
	}
	return nil
}

func (ie *DLDedicatedMessageSegment_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CriticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode CriticalExtensions", err)
	}
	return nil
}
