package ies

// DeltaFList-SPUCCH-r15-setup ::= SEQUENCE
// Extensible
type DeltaflistSpucchR15Setup struct {
	DeltafSlotspucchFormat1R15         *DeltaflistSpucchR15SetupDeltafSlotspucchFormat1R15
	DeltafSlotspucchFormat1aR15        *DeltaflistSpucchR15SetupDeltafSlotspucchFormat1aR15
	DeltafSlotspucchFormat1bR15        *DeltaflistSpucchR15SetupDeltafSlotspucchFormat1bR15
	DeltafSlotspucchFormat3R15         *DeltaflistSpucchR15SetupDeltafSlotspucchFormat3R15
	DeltafSlotspucchRmFormat4R15       *DeltaflistSpucchR15SetupDeltafSlotspucchRmFormat4R15
	DeltafSlotspucchTbccFormat4R15     *DeltaflistSpucchR15SetupDeltafSlotspucchTbccFormat4R15
	DeltafSubslotspucchFormat1and1aR15 *DeltaflistSpucchR15SetupDeltafSubslotspucchFormat1and1aR15
	DeltafSubslotspucchFormat1bR15     *DeltaflistSpucchR15SetupDeltafSubslotspucchFormat1bR15
	DeltafSubslotspucchRmFormat4R15    *DeltaflistSpucchR15SetupDeltafSubslotspucchRmFormat4R15
	DeltafSubslotspucchTbccFormat4R15  *DeltaflistSpucchR15SetupDeltafSubslotspucchTbccFormat4R15
}
