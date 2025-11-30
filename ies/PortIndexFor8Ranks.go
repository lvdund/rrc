package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PortIndexFor8Ranks_Choice_nothing uint64 = iota
	PortIndexFor8Ranks_Choice_PortIndex8
	PortIndexFor8Ranks_Choice_PortIndex4
	PortIndexFor8Ranks_Choice_PortIndex2
	PortIndexFor8Ranks_Choice_PortIndex1
)

type PortIndexFor8Ranks struct {
	Choice     uint64
	PortIndex8 *PortIndexFor8Ranks_portIndex8
	PortIndex4 *PortIndexFor8Ranks_portIndex4
	PortIndex2 *PortIndexFor8Ranks_portIndex2
	PortIndex1 aper.NULL `madatory`
}

func (ie *PortIndexFor8Ranks) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PortIndexFor8Ranks_Choice_PortIndex8:
		if err = ie.PortIndex8.Encode(w); err != nil {
			err = utils.WrapError("Encode PortIndex8", err)
		}
	case PortIndexFor8Ranks_Choice_PortIndex4:
		if err = ie.PortIndex4.Encode(w); err != nil {
			err = utils.WrapError("Encode PortIndex4", err)
		}
	case PortIndexFor8Ranks_Choice_PortIndex2:
		if err = ie.PortIndex2.Encode(w); err != nil {
			err = utils.WrapError("Encode PortIndex2", err)
		}
	case PortIndexFor8Ranks_Choice_PortIndex1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode PortIndex1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PortIndexFor8Ranks) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PortIndexFor8Ranks_Choice_PortIndex8:
		ie.PortIndex8 = new(PortIndexFor8Ranks_portIndex8)
		if err = ie.PortIndex8.Decode(r); err != nil {
			return utils.WrapError("Decode PortIndex8", err)
		}
	case PortIndexFor8Ranks_Choice_PortIndex4:
		ie.PortIndex4 = new(PortIndexFor8Ranks_portIndex4)
		if err = ie.PortIndex4.Decode(r); err != nil {
			return utils.WrapError("Decode PortIndex4", err)
		}
	case PortIndexFor8Ranks_Choice_PortIndex2:
		ie.PortIndex2 = new(PortIndexFor8Ranks_portIndex2)
		if err = ie.PortIndex2.Decode(r); err != nil {
			return utils.WrapError("Decode PortIndex2", err)
		}
	case PortIndexFor8Ranks_Choice_PortIndex1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode PortIndex1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
