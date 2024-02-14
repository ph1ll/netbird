package integrated_approval

import (
	"github.com/netbirdio/netbird/management/server/account"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

// IntegratedApproval interface exists to avoid the circle dependencies
type IntegratedApproval interface {
	PreparePeer(peer *nbpeer.Peer, extraSettings *account.ExtraSettings) *nbpeer.Peer
	ValidatePeer(peer *nbpeer.Peer) (bool, error)
}
