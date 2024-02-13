package server

import "github.com/netbirdio/netbird/management/server/account"

func (am *DefaultAccountManager) UpdateIntegrationApprovalGroups(accountID string, userID string, groups []string) error {
	unlock := am.Store.AcquireAccountLock(accountID)
	defer unlock()

	a, err := am.Store.GetAccountByUser(userID)
	if err != nil {
		return err
	}

	var extra *account.ExtraSettings

	if a.Settings.Extra != nil {
		extra = a.Settings.Extra
	} else {
		extra = &account.ExtraSettings{}
	}
	extra.IntegratedApprovalGroups = groups
	return am.Store.SaveAccount(a)
}

func isPeerAssignedToIntegratedApproval(a *Account, id string) bool {
	if a.Settings.Extra == nil {
		return false
	}

	for _, peerGroup := range a.getPeerGroupsList(id) {
		for _, ig := range a.Settings.Extra.IntegratedApprovalGroups {
			if ig == peerGroup {
				return true
			}
		}
	}
	return false
}
