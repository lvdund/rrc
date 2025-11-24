package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFR1 struct {
	Pdcch_MonitoringSingleOccasion          *Phy_ParametersFR1_pdcch_MonitoringSingleOccasion          `optional`
	Scs_60kHz                               *Phy_ParametersFR1_scs_60kHz                               `optional`
	Pdsch_256QAM_FR1                        *Phy_ParametersFR1_pdsch_256QAM_FR1                        `optional`
	Pdsch_RE_MappingFR1_PerSymbol           *Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol           `optional`
	Pdsch_RE_MappingFR1_PerSlot             *Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSlot             `optional,ext-1`
	Pdcch_MonitoringSingleSpanFirst4Sym_r16 *Phy_ParametersFR1_pdcch_MonitoringSingleSpanFirst4Sym_r16 `optional,ext-2`
}

func (ie *Phy_ParametersFR1) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Pdsch_RE_MappingFR1_PerSlot != nil || ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Pdcch_MonitoringSingleOccasion != nil, ie.Scs_60kHz != nil, ie.Pdsch_256QAM_FR1 != nil, ie.Pdsch_RE_MappingFR1_PerSymbol != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcch_MonitoringSingleOccasion != nil {
		if err = ie.Pdcch_MonitoringSingleOccasion.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_MonitoringSingleOccasion", err)
		}
	}
	if ie.Scs_60kHz != nil {
		if err = ie.Scs_60kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_60kHz", err)
		}
	}
	if ie.Pdsch_256QAM_FR1 != nil {
		if err = ie.Pdsch_256QAM_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_256QAM_FR1", err)
		}
	}
	if ie.Pdsch_RE_MappingFR1_PerSymbol != nil {
		if err = ie.Pdsch_RE_MappingFR1_PerSymbol.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_RE_MappingFR1_PerSymbol", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Pdsch_RE_MappingFR1_PerSlot != nil, ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersFR1", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Pdsch_RE_MappingFR1_PerSlot != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdsch_RE_MappingFR1_PerSlot optional
			if ie.Pdsch_RE_MappingFR1_PerSlot != nil {
				if err = ie.Pdsch_RE_MappingFR1_PerSlot.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_RE_MappingFR1_PerSlot", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdcch_MonitoringSingleSpanFirst4Sym_r16 optional
			if ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 != nil {
				if err = ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_MonitoringSingleSpanFirst4Sym_r16", err)
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

func (ie *Phy_ParametersFR1) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_MonitoringSingleOccasionPresent bool
	if Pdcch_MonitoringSingleOccasionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_60kHzPresent bool
	if Scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_256QAM_FR1Present bool
	if Pdsch_256QAM_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_RE_MappingFR1_PerSymbolPresent bool
	if Pdsch_RE_MappingFR1_PerSymbolPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcch_MonitoringSingleOccasionPresent {
		ie.Pdcch_MonitoringSingleOccasion = new(Phy_ParametersFR1_pdcch_MonitoringSingleOccasion)
		if err = ie.Pdcch_MonitoringSingleOccasion.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_MonitoringSingleOccasion", err)
		}
	}
	if Scs_60kHzPresent {
		ie.Scs_60kHz = new(Phy_ParametersFR1_scs_60kHz)
		if err = ie.Scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_60kHz", err)
		}
	}
	if Pdsch_256QAM_FR1Present {
		ie.Pdsch_256QAM_FR1 = new(Phy_ParametersFR1_pdsch_256QAM_FR1)
		if err = ie.Pdsch_256QAM_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_256QAM_FR1", err)
		}
	}
	if Pdsch_RE_MappingFR1_PerSymbolPresent {
		ie.Pdsch_RE_MappingFR1_PerSymbol = new(Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol)
		if err = ie.Pdsch_RE_MappingFR1_PerSymbol.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_RE_MappingFR1_PerSymbol", err)
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pdsch_RE_MappingFR1_PerSlotPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdsch_RE_MappingFR1_PerSlot optional
			if Pdsch_RE_MappingFR1_PerSlotPresent {
				ie.Pdsch_RE_MappingFR1_PerSlot = new(Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSlot)
				if err = ie.Pdsch_RE_MappingFR1_PerSlot.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_RE_MappingFR1_PerSlot", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pdcch_MonitoringSingleSpanFirst4Sym_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdcch_MonitoringSingleSpanFirst4Sym_r16 optional
			if Pdcch_MonitoringSingleSpanFirst4Sym_r16Present {
				ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 = new(Phy_ParametersFR1_pdcch_MonitoringSingleSpanFirst4Sym_r16)
				if err = ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_MonitoringSingleSpanFirst4Sym_r16", err)
				}
			}
		}
	}
	return nil
}
