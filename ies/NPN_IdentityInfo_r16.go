package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NPN_IdentityInfo_r16 struct {
	Npn_IdentityList_r16           []NPN_Identity_r16                                  `lb:1,ub:maxNPN_r16,madatory`
	TrackingAreaCode_r16           TrackingAreaCode                                    `madatory`
	Ranac_r16                      *RAN_AreaCode                                       `optional`
	CellIdentity_r16               CellIdentity                                        `madatory`
	CellReservedForOperatorUse_r16 NPN_IdentityInfo_r16_cellReservedForOperatorUse_r16 `madatory`
	Iab_Support_r16                *NPN_IdentityInfo_r16_iab_Support_r16               `optional`
	GNB_ID_Length_r17              *int64                                              `lb:22,ub:32,optional,ext-1`
}

func (ie *NPN_IdentityInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.GNB_ID_Length_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Ranac_r16 != nil, ie.Iab_Support_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_Npn_IdentityList_r16 := utils.NewSequence[*NPN_Identity_r16]([]*NPN_Identity_r16{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.Npn_IdentityList_r16 {
		tmp_Npn_IdentityList_r16.Value = append(tmp_Npn_IdentityList_r16.Value, &i)
	}
	if err = tmp_Npn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Npn_IdentityList_r16", err)
	}
	if err = ie.TrackingAreaCode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TrackingAreaCode_r16", err)
	}
	if ie.Ranac_r16 != nil {
		if err = ie.Ranac_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ranac_r16", err)
		}
	}
	if err = ie.CellIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CellIdentity_r16", err)
	}
	if err = ie.CellReservedForOperatorUse_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CellReservedForOperatorUse_r16", err)
	}
	if ie.Iab_Support_r16 != nil {
		if err = ie.Iab_Support_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Iab_Support_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.GNB_ID_Length_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap NPN_IdentityInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.GNB_ID_Length_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode GNB_ID_Length_r17 optional
			if ie.GNB_ID_Length_r17 != nil {
				if err = extWriter.WriteInteger(*ie.GNB_ID_Length_r17, &aper.Constraint{Lb: 22, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode GNB_ID_Length_r17", err)
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

func (ie *NPN_IdentityInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ranac_r16Present bool
	if Ranac_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Iab_Support_r16Present bool
	if Iab_Support_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_Npn_IdentityList_r16 := utils.NewSequence[*NPN_Identity_r16]([]*NPN_Identity_r16{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn_Npn_IdentityList_r16 := func() *NPN_Identity_r16 {
		return new(NPN_Identity_r16)
	}
	if err = tmp_Npn_IdentityList_r16.Decode(r, fn_Npn_IdentityList_r16); err != nil {
		return utils.WrapError("Decode Npn_IdentityList_r16", err)
	}
	ie.Npn_IdentityList_r16 = []NPN_Identity_r16{}
	for _, i := range tmp_Npn_IdentityList_r16.Value {
		ie.Npn_IdentityList_r16 = append(ie.Npn_IdentityList_r16, *i)
	}
	if err = ie.TrackingAreaCode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TrackingAreaCode_r16", err)
	}
	if Ranac_r16Present {
		ie.Ranac_r16 = new(RAN_AreaCode)
		if err = ie.Ranac_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ranac_r16", err)
		}
	}
	if err = ie.CellIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CellIdentity_r16", err)
	}
	if err = ie.CellReservedForOperatorUse_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CellReservedForOperatorUse_r16", err)
	}
	if Iab_Support_r16Present {
		ie.Iab_Support_r16 = new(NPN_IdentityInfo_r16_iab_Support_r16)
		if err = ie.Iab_Support_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Iab_Support_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			GNB_ID_Length_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode GNB_ID_Length_r17 optional
			if GNB_ID_Length_r17Present {
				var tmp_int_GNB_ID_Length_r17 int64
				if tmp_int_GNB_ID_Length_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 22, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode GNB_ID_Length_r17", err)
				}
				ie.GNB_ID_Length_r17 = &tmp_int_GNB_ID_Length_r17
			}
		}
	}
	return nil
}
