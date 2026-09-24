package command

import (
	"context"
	"time"

	"github.com/caigee-cmd/cli2api/internal/providers"
)

// Quota reports the account credit balance as usage windows. Command Code does
// not expose the 5h / weekly / monthly subscription windows the plan docs
// describe, so this only surfaces the credit buckets
// (GET /alpha/billing/credits) and never fabricates a subscription window. A
// failure returns a nil snapshot so the console card shows "unknown" rather
// than a wrong number.
func (c *Client) Quota(ctx context.Context, accountID string) (*providers.QuotaInfo, error) {
	credential, err := c.credential(ctx, accountID)
	if err != nil {
		return nil, err
	}
	client, err := c.httpClient(ctx, accountID)
	if err != nil {
		return nil, err
	}
	balance, err := c.fetchCredits(ctx, client, credential)
	if err != nil {
		return nil, err
	}
	return quotaFromCredits(balance, time.Now().UTC()), nil
}

func quotaFromCredits(balance credits, fetchedAt time.Time) *providers.QuotaInfo {
	windows := make([]providers.QuotaWindow, 0, 3)
	appendWindow := func(id, label string, amount float64) {
		if amount == 0 {
			return
		}
		windows = append(windows, providers.QuotaWindow{
			ID:         id,
			Label:      label,
			Used:       0,
			Total:      0,
			Remaining:  amount,
			Percentage: 0,
			Unit:       QuotaUnit,
			Exceeded:   amount <= 0,
		})
	}
	appendWindow("monthly", "Monthly credits", balance.Monthly)
	appendWindow("purchased", "Purchased credits", balance.Purchased)
	appendWindow("free", "Free credits", balance.Free)

	total := balance.total()
	info := &providers.QuotaInfo{
		Used:       0,
		Total:      0,
		Remaining:  total,
		Percentage: 0,
		Unit:       QuotaUnit,
		Exceeded:   total <= 0,
		FetchedAt:  fetchedAt.Format(time.RFC3339),
		ProviderID: "command",
		Windows:    windows,
	}
	return info
}
