package routine

import (
	"html/template"

	"github.com/wiredlush/easy-gate/internal/config"
	"github.com/wiredlush/easy-gate/internal/service"
)

func (r *Routine) getServices(cfg *config.Config) []service.Service {
	type serviceWrapper struct {
		service service.Service
		index   int
	}

	servicePChan := make(chan serviceWrapper)
	for index, cfgService := range cfg.Services {
		go func(index int, cfgService config.Service) {
			servicePChan <- serviceWrapper{
				service: service.Service{
					Icon:     r.getIconData(cfgService),
					Name:     cfgService.Name,
					URL:      template.URL(cfgService.URL),
					Category: cfgService.Category,
					Groups:   cfgService.Groups,
				},
				index: index,
			}
		}(index, cfgService)
	}

	processedServices := map[int]service.Service{}
	for i := 1; i <= len(cfg.Services); i++ {
		processedService := <-servicePChan
		processedServices[processedService.index] = processedService.service
	}

	services := []service.Service{}
	for index := range cfg.Services {
		services = append(services, processedServices[index])
	}

	return services
}
