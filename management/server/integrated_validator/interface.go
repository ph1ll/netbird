package integrated_validator

import (
	"github.com/netbirdio/netbird/management/server/account"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedValidator interface {
	PreparePeer(peer *nbpeer.Peer, extraSettings *account.ExtraSettings) *nbpeer.Peer
	ValidatePeer(peer *nbpeer.Peer) (bool, error)
}
