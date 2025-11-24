package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ZP_CSI_RS_Resource struct {
	Zp_CSI_RS_ResourceId ZP_CSI_RS_ResourceId              `madatory`
	ResourceMapping      CSI_RS_ResourceMapping            `madatory`
	PeriodicityAndOffset *CSI_ResourcePeriodicityAndOffset `optional`
}

func (ie *ZP_CSI_RS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PeriodicityAndOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Zp_CSI_RS_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode Zp_CSI_RS_ResourceId", err)
	}
	if err = ie.ResourceMapping.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceMapping", err)
	}
	if ie.PeriodicityAndOffset != nil {
		if err = ie.PeriodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicityAndOffset", err)
		}
	}
	return nil
}

func (ie *ZP_CSI_RS_Resource) Decode(r *uper.UperReader) error {
	var err error
	var PeriodicityAndOffsetPresent bool
	if PeriodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Zp_CSI_RS_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode Zp_CSI_RS_ResourceId", err)
	}
	if err = ie.ResourceMapping.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceMapping", err)
	}
	if PeriodicityAndOffsetPresent {
		ie.PeriodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
		if err = ie.PeriodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicityAndOffset", err)
		}
	}
	return nil
}
