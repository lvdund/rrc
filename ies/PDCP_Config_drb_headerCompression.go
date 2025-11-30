package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_drb_headerCompression_Choice_nothing uint64 = iota
	PDCP_Config_drb_headerCompression_Choice_NotUsed
	PDCP_Config_drb_headerCompression_Choice_Rohc
	PDCP_Config_drb_headerCompression_Choice_UplinkOnlyROHC
)

type PDCP_Config_drb_headerCompression struct {
	Choice         uint64
	NotUsed        aper.NULL `madatory`
	Rohc           *PDCP_Config_drb_headerCompression_rohc
	UplinkOnlyROHC *PDCP_Config_drb_headerCompression_uplinkOnlyROHC
}

func (ie *PDCP_Config_drb_headerCompression) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCP_Config_drb_headerCompression_Choice_NotUsed:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode NotUsed", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_Rohc:
		if err = ie.Rohc.Encode(w); err != nil {
			err = utils.WrapError("Encode Rohc", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_UplinkOnlyROHC:
		if err = ie.UplinkOnlyROHC.Encode(w); err != nil {
			err = utils.WrapError("Encode UplinkOnlyROHC", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDCP_Config_drb_headerCompression) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCP_Config_drb_headerCompression_Choice_NotUsed:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode NotUsed", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_Rohc:
		ie.Rohc = new(PDCP_Config_drb_headerCompression_rohc)
		if err = ie.Rohc.Decode(r); err != nil {
			return utils.WrapError("Decode Rohc", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_UplinkOnlyROHC:
		ie.UplinkOnlyROHC = new(PDCP_Config_drb_headerCompression_uplinkOnlyROHC)
		if err = ie.UplinkOnlyROHC.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkOnlyROHC", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
