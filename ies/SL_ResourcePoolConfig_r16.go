package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_ResourcePoolConfig_r16 struct {
	Sl_ResourcePoolID_r16 SL_ResourcePoolID_r16 `madatory`
	Sl_ResourcePool_r16   *SL_ResourcePool_r16  `optional`
}

func (ie *SL_ResourcePoolConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ResourcePool_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_ResourcePoolID_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ResourcePoolID_r16", err)
	}
	if ie.Sl_ResourcePool_r16 != nil {
		if err = ie.Sl_ResourcePool_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ResourcePool_r16", err)
		}
	}
	return nil
}

func (ie *SL_ResourcePoolConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_ResourcePool_r16Present bool
	if Sl_ResourcePool_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_ResourcePoolID_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ResourcePoolID_r16", err)
	}
	if Sl_ResourcePool_r16Present {
		ie.Sl_ResourcePool_r16 = new(SL_ResourcePool_r16)
		if err = ie.Sl_ResourcePool_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ResourcePool_r16", err)
		}
	}
	return nil
}
