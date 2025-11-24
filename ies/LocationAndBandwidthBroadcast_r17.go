package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LocationAndBandwidthBroadcast_r17_Choice_nothing uint64 = iota
	LocationAndBandwidthBroadcast_r17_Choice_SameAsSib1ConfiguredLocationAndBW
	LocationAndBandwidthBroadcast_r17_Choice_LocationAndBandwidth
)

type LocationAndBandwidthBroadcast_r17 struct {
	Choice                            uint64
	SameAsSib1ConfiguredLocationAndBW uper.NULL `madatory`
	LocationAndBandwidth              int64     `lb:0,ub:37949,madatory`
}

func (ie *LocationAndBandwidthBroadcast_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationAndBandwidthBroadcast_r17_Choice_SameAsSib1ConfiguredLocationAndBW:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode SameAsSib1ConfiguredLocationAndBW", err)
		}
	case LocationAndBandwidthBroadcast_r17_Choice_LocationAndBandwidth:
		if err = w.WriteInteger(int64(ie.LocationAndBandwidth), &uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			err = utils.WrapError("Encode LocationAndBandwidth", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *LocationAndBandwidthBroadcast_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationAndBandwidthBroadcast_r17_Choice_SameAsSib1ConfiguredLocationAndBW:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode SameAsSib1ConfiguredLocationAndBW", err)
		}
	case LocationAndBandwidthBroadcast_r17_Choice_LocationAndBandwidth:
		var tmp_int_LocationAndBandwidth int64
		if tmp_int_LocationAndBandwidth, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			return utils.WrapError("Decode LocationAndBandwidth", err)
		}
		ie.LocationAndBandwidth = tmp_int_LocationAndBandwidth
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
