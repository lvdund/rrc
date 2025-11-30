package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_moreThanOneRLC_primaryPath struct {
	CellGroup      *CellGroupId            `optional`
	LogicalChannel *LogicalChannelIdentity `optional`
}

func (ie *PDCP_Config_moreThanOneRLC_primaryPath) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CellGroup != nil, ie.LogicalChannel != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CellGroup != nil {
		if err = ie.CellGroup.Encode(w); err != nil {
			return utils.WrapError("Encode CellGroup", err)
		}
	}
	if ie.LogicalChannel != nil {
		if err = ie.LogicalChannel.Encode(w); err != nil {
			return utils.WrapError("Encode LogicalChannel", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_moreThanOneRLC_primaryPath) Decode(r *aper.AperReader) error {
	var err error
	var CellGroupPresent bool
	if CellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LogicalChannelPresent bool
	if LogicalChannelPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CellGroupPresent {
		ie.CellGroup = new(CellGroupId)
		if err = ie.CellGroup.Decode(r); err != nil {
			return utils.WrapError("Decode CellGroup", err)
		}
	}
	if LogicalChannelPresent {
		ie.LogicalChannel = new(LogicalChannelIdentity)
		if err = ie.LogicalChannel.Decode(r); err != nil {
			return utils.WrapError("Decode LogicalChannel", err)
		}
	}
	return nil
}
