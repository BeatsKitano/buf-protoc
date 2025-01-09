// Package mail contains the slow query weekly mail sender.
package mail

import (
	"buf-protoc/backend/component/iam"
	"buf-protoc/backend/component/state"
	"buf-protoc/backend/store"

	"context"
	"embed"
	"log/slog"
	"sync"
	"time"
)

var (
	//go:embed templates/for-dba/need_configure.html
	//go:embed templates/for-dba/environment_header.html
	//go:embed templates/for-dba/environment_footer.html
	//go:embed templates/for-dba/environment_no_instance_configured.html
	//go:embed templates/for-dba/environment_no_slow_query.html
	//go:embed templates/for-dba/header.html
	//go:embed templates/for-dba/footer.html
	//go:embed templates/for-dba/database_table_header.html
	//go:embed templates/for-dba/database_table_item.html
	//go:embed templates/for-dba/database_table_footer.html
	//go:embed templates/for-dba/instance_header.html
	//go:embed templates/for-dba/instance_footer.html
	//go:embed templates/for-project-owner/header.html
	//go:embed templates/for-project-owner/footer.html
	//go:embed templates/for-project-owner/environment_header.html
	//go:embed templates/for-project-owner/environment_footer.html
	//go:embed templates/for-project-owner/table_item.html
	emailTemplates embed.FS
)

// NewSender creates a new slow query weekly mail sender.
func NewSender(store *store.Store, stateCfg *state.State, iamManager *iam.Manager) *SlowQueryWeeklyMailSender {
	return &SlowQueryWeeklyMailSender{
		store:      store,
		stateCfg:   stateCfg,
		iamManager: iamManager,
	}
}

// SlowQueryWeeklyMailSender is the slow query weekly mail sender.
type SlowQueryWeeklyMailSender struct {
	store      *store.Store
	stateCfg   *state.State
	iamManager *iam.Manager
}

// Run will run the slow query weekly mail sender.
func (s *SlowQueryWeeklyMailSender) Run(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	defer wg.Done()
	slog.Debug("Slow query weekly mail sender started")
	
	for {
		select {
		case <-ctx.Done():
			slog.Debug("Slow query weekly mail sender received context cancellation")
			return
		case <-ticker.C:
			slog.Debug("Slow query weekly mail sender received tick")
			now := time.Now()
			// Send email every Saturday in 00:00 ~ 00:59.
			if now.Weekday() == time.Saturday && now.Hour() == 0 {

			}
		}
	}
}
