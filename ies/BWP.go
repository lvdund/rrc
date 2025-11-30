package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BWP struct {
	LocationAndBandwidth int64             `lb:0,ub:37949,madatory`
	SubcarrierSpacing    SubcarrierSpacing `madatory`
	CyclicPrefix         *BWP_cyclicPrefix `optional`
}

func (ie *BWP) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CyclicPrefix != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.LocationAndBandwidth, &aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("WriteInteger LocationAndBandwidth", err)
	}
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing", err)
	}
	if ie.CyclicPrefix != nil {
		if err = ie.CyclicPrefix.Encode(w); err != nil {
			return utils.WrapError("Encode CyclicPrefix", err)
		}
	}
	return nil
}

func (ie *BWP) Decode(r *aper.AperReader) error {
	var err error
	var CyclicPrefixPresent bool
	if CyclicPrefixPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_LocationAndBandwidth int64
	if tmp_int_LocationAndBandwidth, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("ReadInteger LocationAndBandwidth", err)
	}
	ie.LocationAndBandwidth = tmp_int_LocationAndBandwidth
	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing", err)
	}
	if CyclicPrefixPresent {
		ie.CyclicPrefix = new(BWP_cyclicPrefix)
		if err = ie.CyclicPrefix.Decode(r); err != nil {
			return utils.WrapError("Decode CyclicPrefix", err)
		}
	}
	return nil
}
