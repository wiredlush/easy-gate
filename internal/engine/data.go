package engine

import (
	"github.com/wiredlush/easy-gate/internal/group"
	"github.com/wiredlush/easy-gate/internal/note"
	"github.com/wiredlush/easy-gate/internal/routine"
	"github.com/wiredlush/easy-gate/internal/service"
)

type data struct {
	Services map[string][]service.Service
	Notes    []note.Note
}

func getData(status *routine.Status, addr string) data {
	return data{
		Services: getServices(status, addr),
		Notes:    getNotes(status, addr),
	}
}

func getServices(status *routine.Status, addr string) map[string][]service.Service {
	services := map[string][]service.Service{}

	for _, statusService := range status.Services {
		if group.IsAllowed(status.Groups, statusService.Groups, addr) {
			service := service.Service{
				Icon:     statusService.Icon,
				Name:     statusService.Name,
				URL:      statusService.URL,
				Category: statusService.Category,
			}

			services[statusService.Category] = append(
				services[statusService.Category],
				service,
			)
		}
	}

	return services
}

func getNotes(status *routine.Status, addr string) []note.Note {
	notes := []note.Note{}
	for _, statusNote := range status.Notes {
		if group.IsAllowed(status.Groups, statusNote.Groups, addr) {
			note := note.Note{
				Name: statusNote.Name,
				Text: statusNote.Text,
			}
			notes = append(notes, note)
		}
	}

	return notes
}
