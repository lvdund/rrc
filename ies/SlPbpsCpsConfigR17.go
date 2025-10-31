package ies

import "rrc/utils"

// SL-PBPS-CPS-Config-r17 ::= SEQUENCE
// Extensible
type SlPbpsCpsConfigR17 struct {
	SlAllowedresourceselectionconfigR17 *SlPbpsCpsConfigR17SlAllowedresourceselectionconfigR17
	SlMinnumcandidateslotsperiodicR17   *utils.INTEGER   `lb:0,ub:32`
	SlPbpsOccasionreserveperiodlistR17  *[]utils.INTEGER `lb:1,ub:16`
	SlAdditionalPbpsOccasionR17         *SlPbpsCpsConfigR17SlAdditionalPbpsOccasionR17
	SlCpsWindowperiodicR17              *utils.INTEGER `lb:0,ub:30`
	SlMinnumcandidateslotsaperiodicR17  *utils.INTEGER `lb:0,ub:32`
	SlMinnumrssimeasurementslotsR17     *utils.INTEGER `lb:0,ub:800`
	SlDefaultcbrRandomselectionR17      *utils.INTEGER `lb:0,ub:100`
	SlDefaultcbrPartialsensingR17       *utils.INTEGER `lb:0,ub:100`
	SlCpsWindowaperiodicR17             *utils.INTEGER `lb:0,ub:30`
	SlPartialsensinginactivetimeR17     *SlPbpsCpsConfigR17SlPartialsensinginactivetimeR17
}
