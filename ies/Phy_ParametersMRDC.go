package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersMRDC struct {
	Naics_Capability_List             []NAICS_Capability_Entry                              `lb:1,ub:maxNrofNAICS_Entries,optional`
	SpCellPlacement                   *CarrierAggregationVariant                            `optional,ext-1`
	Tdd_PCellUL_TX_AllUL_Subframe_r16 *Phy_ParametersMRDC_tdd_PCellUL_TX_AllUL_Subframe_r16 `optional,ext-2`
	Fdd_PCellUL_TX_AllUL_Subframe_r16 *Phy_ParametersMRDC_fdd_PCellUL_TX_AllUL_Subframe_r16 `optional,ext-2`
}

func (ie *Phy_ParametersMRDC) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.SpCellPlacement != nil || ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 != nil || ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.Naics_Capability_List) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Naics_Capability_List) > 0 {
		tmp_Naics_Capability_List := utils.NewSequence[*NAICS_Capability_Entry]([]*NAICS_Capability_Entry{}, aper.Constraint{Lb: 1, Ub: maxNrofNAICS_Entries}, false)
		for _, i := range ie.Naics_Capability_List {
			tmp_Naics_Capability_List.Value = append(tmp_Naics_Capability_List.Value, &i)
		}
		if err = tmp_Naics_Capability_List.Encode(w); err != nil {
			return utils.WrapError("Encode Naics_Capability_List", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.SpCellPlacement != nil, ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 != nil || ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersMRDC", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SpCellPlacement != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SpCellPlacement optional
			if ie.SpCellPlacement != nil {
				if err = ie.SpCellPlacement.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpCellPlacement", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 != nil, ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Tdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 != nil {
				if err = ie.Tdd_PCellUL_TX_AllUL_Subframe_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tdd_PCellUL_TX_AllUL_Subframe_r16", err)
				}
			}
			// encode Fdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 != nil {
				if err = ie.Fdd_PCellUL_TX_AllUL_Subframe_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Fdd_PCellUL_TX_AllUL_Subframe_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *Phy_ParametersMRDC) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Naics_Capability_ListPresent bool
	if Naics_Capability_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Naics_Capability_ListPresent {
		tmp_Naics_Capability_List := utils.NewSequence[*NAICS_Capability_Entry]([]*NAICS_Capability_Entry{}, aper.Constraint{Lb: 1, Ub: maxNrofNAICS_Entries}, false)
		fn_Naics_Capability_List := func() *NAICS_Capability_Entry {
			return new(NAICS_Capability_Entry)
		}
		if err = tmp_Naics_Capability_List.Decode(r, fn_Naics_Capability_List); err != nil {
			return utils.WrapError("Decode Naics_Capability_List", err)
		}
		ie.Naics_Capability_List = []NAICS_Capability_Entry{}
		for _, i := range tmp_Naics_Capability_List.Value {
			ie.Naics_Capability_List = append(ie.Naics_Capability_List, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SpCellPlacementPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SpCellPlacement optional
			if SpCellPlacementPresent {
				ie.SpCellPlacement = new(CarrierAggregationVariant)
				if err = ie.SpCellPlacement.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpCellPlacement", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Tdd_PCellUL_TX_AllUL_Subframe_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Fdd_PCellUL_TX_AllUL_Subframe_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Tdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if Tdd_PCellUL_TX_AllUL_Subframe_r16Present {
				ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 = new(Phy_ParametersMRDC_tdd_PCellUL_TX_AllUL_Subframe_r16)
				if err = ie.Tdd_PCellUL_TX_AllUL_Subframe_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tdd_PCellUL_TX_AllUL_Subframe_r16", err)
				}
			}
			// decode Fdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if Fdd_PCellUL_TX_AllUL_Subframe_r16Present {
				ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 = new(Phy_ParametersMRDC_fdd_PCellUL_TX_AllUL_Subframe_r16)
				if err = ie.Fdd_PCellUL_TX_AllUL_Subframe_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Fdd_PCellUL_TX_AllUL_Subframe_r16", err)
				}
			}
		}
	}
	return nil
}
