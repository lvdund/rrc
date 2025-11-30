package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SSB_InfoNcell_r16 struct {
	PhysicalCellId_r16    PhysCellId             `madatory`
	Ssb_IndexNcell_r16    *SSB_Index             `optional`
	Ssb_Configuration_r16 *SSB_Configuration_r16 `optional`
}

func (ie *SSB_InfoNcell_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_IndexNcell_r16 != nil, ie.Ssb_Configuration_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysicalCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PhysicalCellId_r16", err)
	}
	if ie.Ssb_IndexNcell_r16 != nil {
		if err = ie.Ssb_IndexNcell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_IndexNcell_r16", err)
		}
	}
	if ie.Ssb_Configuration_r16 != nil {
		if err = ie.Ssb_Configuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_Configuration_r16", err)
		}
	}
	return nil
}

func (ie *SSB_InfoNcell_r16) Decode(r *aper.AperReader) error {
	var err error
	var Ssb_IndexNcell_r16Present bool
	if Ssb_IndexNcell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_Configuration_r16Present bool
	if Ssb_Configuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysicalCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PhysicalCellId_r16", err)
	}
	if Ssb_IndexNcell_r16Present {
		ie.Ssb_IndexNcell_r16 = new(SSB_Index)
		if err = ie.Ssb_IndexNcell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_IndexNcell_r16", err)
		}
	}
	if Ssb_Configuration_r16Present {
		ie.Ssb_Configuration_r16 = new(SSB_Configuration_r16)
		if err = ie.Ssb_Configuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Configuration_r16", err)
		}
	}
	return nil
}
