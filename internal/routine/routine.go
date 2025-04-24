package routine

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/wiredlush/easy-gate/internal/config"
	"github.com/wiredlush/easy-gate/internal/engine/static"
)

// Routine - Routine struct
type Routine struct {
	Error        error
	Status       *Status
	Client       *http.Client
	FilePath     string
	LastChecksum string
	Interval     time.Duration
	sync.Mutex
}

// NewRoutine - Create new config routine
func NewRoutine(filePath string, interval time.Duration) (*Routine, error) {
	routine := Routine{
		FilePath: filePath,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Interval: interval,
		Status:   &Status{},
	}

	cfg, checksum, err := config.Load(filePath)
	if err != nil {
		return nil, err
	}

	routine.Status = routine.updateStatus(cfg)
	routine.LastChecksum = checksum

	styleData, err := static.StaticFS.ReadFile("public/styles/style.css")
	if err != nil {
		return nil, err
	}

	if cfg.Theme.CustomCSS != "" {
		styleData, err = os.ReadFile(cfg.Theme.CustomCSS)
		if err != nil {
			return nil, err
		}

		fileInfo, err := os.Stat(cfg.Theme.CustomCSS)
		if err != nil {
			return nil, err
		}
		routine.Status.CSSLastEdit = fileInfo.ModTime().Unix()
	}

	routine.Status.CSSData = string(styleData)

	return &routine, nil
}

// GetStatus - Get current status
func (r *Routine) GetStatus() (*Status, error) {
	defer r.Unlock()
	r.Lock()
	return r.Status, r.Error
}

// Start - Start config routine
func (r *Routine) Start() {
	for {
		cfg, checksum, err := config.Load(r.FilePath)
		if err != nil {
			r.Lock()
			r.Error = err
			r.Unlock()
			continue
		}

		r.Lock()
		r.Error = nil
		if checksum != r.LastChecksum {
			log.Println("Detected configuration change, reloading...")
			r.Status = r.updateStatus(cfg)
		}
		r.LastChecksum = checksum

		if cfg.Theme.CustomCSS != "" {
			fileInfo, err := os.Stat(cfg.Theme.CustomCSS)
			if err != nil {
				log.Println("Error in getting custom css file info:", err)
				r.Unlock()
				continue
			}

			if fileInfo.ModTime().Unix() != r.Status.CSSLastEdit {
				data, err := os.ReadFile(cfg.Theme.CustomCSS)
				if err != nil {
					log.Println("Error in reading custom css file:", err)
					r.Unlock()
					continue
				}

				r.Status.CSSData = string(data)
				r.Status.CSSLastEdit = fileInfo.ModTime().Unix()
			}
		}

		r.Unlock()
		time.Sleep(r.Interval)
	}
}
