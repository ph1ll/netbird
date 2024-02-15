package server

import "github.com/netbirdio/netbird/management/server/account"

// UpdateIntegratedApprovalGroups updates the integrated approval groups for a specified account.
// It retrieves the account associated with the provided userID, then updates the integrated approval groups
// with the provided list of group ids. The updated account is then saved.
//
// Parameters:
//   - accountID: The ID of the account for which integrated approval groups are to be updated.
//   - userID: The ID of the user whose account is being updated.
//   - groups: A slice of strings representing the ids of integrated approval groups to be updated.
//
// Returns:
//   - error: An error if any occurred during the process, otherwise returns nil
func (am *DefaultAccountManager) UpdateIntegratedApprovalGroups(accountID string, userID string, groups []string) error {
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
