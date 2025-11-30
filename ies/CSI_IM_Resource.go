package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_IM_Resource struct {
	Csi_IM_ResourceId             CSI_IM_ResourceId                              `madatory`
	Csi_IM_ResourceElementPattern *CSI_IM_Resource_csi_IM_ResourceElementPattern `lb:0,ub:12,optional`
	FreqBand                      *CSI_FrequencyOccupation                       `optional`
	PeriodicityAndOffset          *CSI_ResourcePeriodicityAndOffset              `optional`
}

func (ie *CSI_IM_Resource) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Csi_IM_ResourceElementPattern != nil, ie.FreqBand != nil, ie.PeriodicityAndOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Csi_IM_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_IM_ResourceId", err)
	}
	if ie.Csi_IM_ResourceElementPattern != nil {
		if err = ie.Csi_IM_ResourceElementPattern.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_IM_ResourceElementPattern", err)
		}
	}
	if ie.FreqBand != nil {
		if err = ie.FreqBand.Encode(w); err != nil {
			return utils.WrapError("Encode FreqBand", err)
		}
	}
	if ie.PeriodicityAndOffset != nil {
		if err = ie.PeriodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicityAndOffset", err)
		}
	}
	return nil
}

func (ie *CSI_IM_Resource) Decode(r *aper.AperReader) error {
	var err error
	var Csi_IM_ResourceElementPatternPresent bool
	if Csi_IM_ResourceElementPatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FreqBandPresent bool
	if FreqBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PeriodicityAndOffsetPresent bool
	if PeriodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Csi_IM_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_IM_ResourceId", err)
	}
	if Csi_IM_ResourceElementPatternPresent {
		ie.Csi_IM_ResourceElementPattern = new(CSI_IM_Resource_csi_IM_ResourceElementPattern)
		if err = ie.Csi_IM_ResourceElementPattern.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_IM_ResourceElementPattern", err)
		}
	}
	if FreqBandPresent {
		ie.FreqBand = new(CSI_FrequencyOccupation)
		if err = ie.FreqBand.Decode(r); err != nil {
			return utils.WrapError("Decode FreqBand", err)
		}
	}
	if PeriodicityAndOffsetPresent {
		ie.PeriodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
		if err = ie.PeriodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicityAndOffset", err)
		}
	}
	return nil
}
