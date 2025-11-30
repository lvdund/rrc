package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkParametersNR_r16 struct {
	Rlc_ParametersSidelink_r16           *RLC_ParametersSidelink_r16                    `optional`
	Mac_ParametersSidelink_r16           *MAC_ParametersSidelink_r16                    `optional`
	Fdd_Add_UE_Sidelink_Capabilities_r16 *UE_SidelinkCapabilityAddXDD_Mode_r16          `optional`
	Tdd_Add_UE_Sidelink_Capabilities_r16 *UE_SidelinkCapabilityAddXDD_Mode_r16          `optional`
	SupportedBandListSidelink_r16        []BandSidelink_r16                             `lb:1,ub:maxBands,optional`
	RelayParameters_r17                  *RelayParameters_r17                           `optional,ext-1`
	P0_OLPC_Sidelink_r17                 *SidelinkParametersNR_r16_p0_OLPC_Sidelink_r17 `optional,ext-2`
}

func (ie *SidelinkParametersNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.RelayParameters_r17 != nil || ie.P0_OLPC_Sidelink_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Rlc_ParametersSidelink_r16 != nil, ie.Mac_ParametersSidelink_r16 != nil, ie.Fdd_Add_UE_Sidelink_Capabilities_r16 != nil, ie.Tdd_Add_UE_Sidelink_Capabilities_r16 != nil, len(ie.SupportedBandListSidelink_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rlc_ParametersSidelink_r16 != nil {
		if err = ie.Rlc_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_ParametersSidelink_r16", err)
		}
	}
	if ie.Mac_ParametersSidelink_r16 != nil {
		if err = ie.Mac_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersSidelink_r16", err)
		}
	}
	if ie.Fdd_Add_UE_Sidelink_Capabilities_r16 != nil {
		if err = ie.Fdd_Add_UE_Sidelink_Capabilities_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if ie.Tdd_Add_UE_Sidelink_Capabilities_r16 != nil {
		if err = ie.Tdd_Add_UE_Sidelink_Capabilities_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if len(ie.SupportedBandListSidelink_r16) > 0 {
		tmp_SupportedBandListSidelink_r16 := utils.NewSequence[*BandSidelink_r16]([]*BandSidelink_r16{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.SupportedBandListSidelink_r16 {
			tmp_SupportedBandListSidelink_r16.Value = append(tmp_SupportedBandListSidelink_r16.Value, &i)
		}
		if err = tmp_SupportedBandListSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandListSidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.RelayParameters_r17 != nil, ie.P0_OLPC_Sidelink_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SidelinkParametersNR_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.RelayParameters_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RelayParameters_r17 optional
			if ie.RelayParameters_r17 != nil {
				if err = ie.RelayParameters_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RelayParameters_r17", err)
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
			optionals_ext_2 := []bool{ie.P0_OLPC_Sidelink_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode P0_OLPC_Sidelink_r17 optional
			if ie.P0_OLPC_Sidelink_r17 != nil {
				if err = ie.P0_OLPC_Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode P0_OLPC_Sidelink_r17", err)
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

func (ie *SidelinkParametersNR_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlc_ParametersSidelink_r16Present bool
	if Rlc_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_ParametersSidelink_r16Present bool
	if Mac_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdd_Add_UE_Sidelink_Capabilities_r16Present bool
	if Fdd_Add_UE_Sidelink_Capabilities_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_Add_UE_Sidelink_Capabilities_r16Present bool
	if Tdd_Add_UE_Sidelink_Capabilities_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandListSidelink_r16Present bool
	if SupportedBandListSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rlc_ParametersSidelink_r16Present {
		ie.Rlc_ParametersSidelink_r16 = new(RLC_ParametersSidelink_r16)
		if err = ie.Rlc_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rlc_ParametersSidelink_r16", err)
		}
	}
	if Mac_ParametersSidelink_r16Present {
		ie.Mac_ParametersSidelink_r16 = new(MAC_ParametersSidelink_r16)
		if err = ie.Mac_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersSidelink_r16", err)
		}
	}
	if Fdd_Add_UE_Sidelink_Capabilities_r16Present {
		ie.Fdd_Add_UE_Sidelink_Capabilities_r16 = new(UE_SidelinkCapabilityAddXDD_Mode_r16)
		if err = ie.Fdd_Add_UE_Sidelink_Capabilities_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if Tdd_Add_UE_Sidelink_Capabilities_r16Present {
		ie.Tdd_Add_UE_Sidelink_Capabilities_r16 = new(UE_SidelinkCapabilityAddXDD_Mode_r16)
		if err = ie.Tdd_Add_UE_Sidelink_Capabilities_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if SupportedBandListSidelink_r16Present {
		tmp_SupportedBandListSidelink_r16 := utils.NewSequence[*BandSidelink_r16]([]*BandSidelink_r16{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_SupportedBandListSidelink_r16 := func() *BandSidelink_r16 {
			return new(BandSidelink_r16)
		}
		if err = tmp_SupportedBandListSidelink_r16.Decode(r, fn_SupportedBandListSidelink_r16); err != nil {
			return utils.WrapError("Decode SupportedBandListSidelink_r16", err)
		}
		ie.SupportedBandListSidelink_r16 = []BandSidelink_r16{}
		for _, i := range tmp_SupportedBandListSidelink_r16.Value {
			ie.SupportedBandListSidelink_r16 = append(ie.SupportedBandListSidelink_r16, *i)
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

			RelayParameters_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RelayParameters_r17 optional
			if RelayParameters_r17Present {
				ie.RelayParameters_r17 = new(RelayParameters_r17)
				if err = ie.RelayParameters_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode RelayParameters_r17", err)
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

			P0_OLPC_Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode P0_OLPC_Sidelink_r17 optional
			if P0_OLPC_Sidelink_r17Present {
				ie.P0_OLPC_Sidelink_r17 = new(SidelinkParametersNR_r16_p0_OLPC_Sidelink_r17)
				if err = ie.P0_OLPC_Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode P0_OLPC_Sidelink_r17", err)
				}
			}
		}
	}
	return nil
}
