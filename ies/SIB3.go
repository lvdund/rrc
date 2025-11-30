package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB3 struct {
	IntraFreqNeighCellList          *IntraFreqNeighCellList            `optional`
	IntraFreqExcludedCellList       *IntraFreqExcludedCellList         `optional`
	LateNonCriticalExtension        *[]byte                            `optional`
	IntraFreqNeighCellList_v1610    *IntraFreqNeighCellList_v1610      `optional,ext-1`
	IntraFreqAllowedCellList_r16    *IntraFreqAllowedCellList_r16      `optional,ext-1`
	IntraFreqCAG_CellList_r16       []IntraFreqCAG_CellListPerPLMN_r16 `lb:1,ub:maxPLMN,optional,ext-1`
	IntraFreqNeighHSDN_CellList_r17 *IntraFreqNeighHSDN_CellList_r17   `optional,ext-2`
	IntraFreqNeighCellList_v1710    *IntraFreqNeighCellList_v1710      `optional,ext-2`
	ChannelAccessMode2_r17          *SIB3_channelAccessMode2_r17       `optional,ext-3`
}

func (ie *SIB3) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.IntraFreqNeighCellList_v1610 != nil || ie.IntraFreqAllowedCellList_r16 != nil || len(ie.IntraFreqCAG_CellList_r16) > 0 || ie.IntraFreqNeighHSDN_CellList_r17 != nil || ie.IntraFreqNeighCellList_v1710 != nil || ie.ChannelAccessMode2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.IntraFreqNeighCellList != nil, ie.IntraFreqExcludedCellList != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IntraFreqNeighCellList != nil {
		if err = ie.IntraFreqNeighCellList.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqNeighCellList", err)
		}
	}
	if ie.IntraFreqExcludedCellList != nil {
		if err = ie.IntraFreqExcludedCellList.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqExcludedCellList", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.IntraFreqNeighCellList_v1610 != nil || ie.IntraFreqAllowedCellList_r16 != nil || len(ie.IntraFreqCAG_CellList_r16) > 0, ie.IntraFreqNeighHSDN_CellList_r17 != nil || ie.IntraFreqNeighCellList_v1710 != nil, ie.ChannelAccessMode2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB3", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.IntraFreqNeighCellList_v1610 != nil, ie.IntraFreqAllowedCellList_r16 != nil, len(ie.IntraFreqCAG_CellList_r16) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode IntraFreqNeighCellList_v1610 optional
			if ie.IntraFreqNeighCellList_v1610 != nil {
				if err = ie.IntraFreqNeighCellList_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraFreqNeighCellList_v1610", err)
				}
			}
			// encode IntraFreqAllowedCellList_r16 optional
			if ie.IntraFreqAllowedCellList_r16 != nil {
				if err = ie.IntraFreqAllowedCellList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraFreqAllowedCellList_r16", err)
				}
			}
			// encode IntraFreqCAG_CellList_r16 optional
			if len(ie.IntraFreqCAG_CellList_r16) > 0 {
				tmp_IntraFreqCAG_CellList_r16 := utils.NewSequence[*IntraFreqCAG_CellListPerPLMN_r16]([]*IntraFreqCAG_CellListPerPLMN_r16{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
				for _, i := range ie.IntraFreqCAG_CellList_r16 {
					tmp_IntraFreqCAG_CellList_r16.Value = append(tmp_IntraFreqCAG_CellList_r16.Value, &i)
				}
				if err = tmp_IntraFreqCAG_CellList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraFreqCAG_CellList_r16", err)
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
			optionals_ext_2 := []bool{ie.IntraFreqNeighHSDN_CellList_r17 != nil, ie.IntraFreqNeighCellList_v1710 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode IntraFreqNeighHSDN_CellList_r17 optional
			if ie.IntraFreqNeighHSDN_CellList_r17 != nil {
				if err = ie.IntraFreqNeighHSDN_CellList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraFreqNeighHSDN_CellList_r17", err)
				}
			}
			// encode IntraFreqNeighCellList_v1710 optional
			if ie.IntraFreqNeighCellList_v1710 != nil {
				if err = ie.IntraFreqNeighCellList_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraFreqNeighCellList_v1710", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.ChannelAccessMode2_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ChannelAccessMode2_r17 optional
			if ie.ChannelAccessMode2_r17 != nil {
				if err = ie.ChannelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelAccessMode2_r17", err)
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

func (ie *SIB3) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFreqNeighCellListPresent bool
	if IntraFreqNeighCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFreqExcludedCellListPresent bool
	if IntraFreqExcludedCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if IntraFreqNeighCellListPresent {
		ie.IntraFreqNeighCellList = new(IntraFreqNeighCellList)
		if err = ie.IntraFreqNeighCellList.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqNeighCellList", err)
		}
	}
	if IntraFreqExcludedCellListPresent {
		ie.IntraFreqExcludedCellList = new(IntraFreqExcludedCellList)
		if err = ie.IntraFreqExcludedCellList.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqExcludedCellList", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			IntraFreqNeighCellList_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntraFreqAllowedCellList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntraFreqCAG_CellList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode IntraFreqNeighCellList_v1610 optional
			if IntraFreqNeighCellList_v1610Present {
				ie.IntraFreqNeighCellList_v1610 = new(IntraFreqNeighCellList_v1610)
				if err = ie.IntraFreqNeighCellList_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraFreqNeighCellList_v1610", err)
				}
			}
			// decode IntraFreqAllowedCellList_r16 optional
			if IntraFreqAllowedCellList_r16Present {
				ie.IntraFreqAllowedCellList_r16 = new(IntraFreqAllowedCellList_r16)
				if err = ie.IntraFreqAllowedCellList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraFreqAllowedCellList_r16", err)
				}
			}
			// decode IntraFreqCAG_CellList_r16 optional
			if IntraFreqCAG_CellList_r16Present {
				tmp_IntraFreqCAG_CellList_r16 := utils.NewSequence[*IntraFreqCAG_CellListPerPLMN_r16]([]*IntraFreqCAG_CellListPerPLMN_r16{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
				fn_IntraFreqCAG_CellList_r16 := func() *IntraFreqCAG_CellListPerPLMN_r16 {
					return new(IntraFreqCAG_CellListPerPLMN_r16)
				}
				if err = tmp_IntraFreqCAG_CellList_r16.Decode(extReader, fn_IntraFreqCAG_CellList_r16); err != nil {
					return utils.WrapError("Decode IntraFreqCAG_CellList_r16", err)
				}
				ie.IntraFreqCAG_CellList_r16 = []IntraFreqCAG_CellListPerPLMN_r16{}
				for _, i := range tmp_IntraFreqCAG_CellList_r16.Value {
					ie.IntraFreqCAG_CellList_r16 = append(ie.IntraFreqCAG_CellList_r16, *i)
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

			IntraFreqNeighHSDN_CellList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntraFreqNeighCellList_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode IntraFreqNeighHSDN_CellList_r17 optional
			if IntraFreqNeighHSDN_CellList_r17Present {
				ie.IntraFreqNeighHSDN_CellList_r17 = new(IntraFreqNeighHSDN_CellList_r17)
				if err = ie.IntraFreqNeighHSDN_CellList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraFreqNeighHSDN_CellList_r17", err)
				}
			}
			// decode IntraFreqNeighCellList_v1710 optional
			if IntraFreqNeighCellList_v1710Present {
				ie.IntraFreqNeighCellList_v1710 = new(IntraFreqNeighCellList_v1710)
				if err = ie.IntraFreqNeighCellList_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraFreqNeighCellList_v1710", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			ChannelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ChannelAccessMode2_r17 optional
			if ChannelAccessMode2_r17Present {
				ie.ChannelAccessMode2_r17 = new(SIB3_channelAccessMode2_r17)
				if err = ie.ChannelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelAccessMode2_r17", err)
				}
			}
		}
	}
	return nil
}
