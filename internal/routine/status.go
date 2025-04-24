package routine

import (
	"github.com/wiredlush/easy-gate/internal/config"
	"github.com/wiredlush/easy-gate/internal/group"
	"github.com/wiredlush/easy-gate/internal/note"
	"github.com/wiredlush/easy-gate/internal/service"
	"github.com/wiredlush/easy-gate/internal/theme"
)

// Status - Status struct
type Status struct {
	Theme       theme.Theme
	CSSLastEdit int64
	CSSData     string
	Addr        string
	Title       string
	CertFile    string
	KeyFile     string
	Groups      []group.Group
	Services    []service.Service
	Notes       []note.Note
	BehindProxy bool
	UseTLS      bool
}

func (r *Routine) updateStatus(cfg *config.Config) *Status {
	return &Status{
		Theme:       cfg.Theme,
		CSSLastEdit: r.Status.CSSLastEdit,
		CSSData:     r.Status.CSSData,
		Addr:        cfg.Addr,
		Title:       cfg.Title,
		CertFile:    cfg.CertFile,
		KeyFile:     cfg.KeyFile,
		Groups:      cfg.Groups,
		Services:    r.getServices(cfg),
		Notes:       r.getNotes(cfg),
		BehindProxy: cfg.BehindProxy,
		UseTLS:      cfg.UseTLS,
	}
}
