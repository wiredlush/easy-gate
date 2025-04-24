package routine

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/wiredlush/easy-gate/internal/config"
)

var cfgFilePath string

func TestMain(m *testing.M) {
	configContent := `{
"addr": "127.0.0.1:8080",
"use_tls": false,
"cert_file": "",
"key_file": "",
"behind_proxy": false,
"title": "Easy Gate",
"theme": {
	"background": "#FFFFFF",
	"foreground": "#000000",	
},
"groups": [],
"services": [],
"notes": []
}`
	cfgFile, err := os.CreateTemp(".", "easy_gate_test_")
	if err != nil {
		log.Fatal("Unable to write tmp files for test")
	}
	cfgFile.WriteString(configContent)
	cfgFilePath = cfgFile.Name()

	exitCode := m.Run()

	os.Remove(cfgFilePath)

	os.Exit(exitCode)
}

func TestRoutine(t *testing.T) {
	_, err := NewRoutine("", 1*time.Millisecond)
	if err == nil {
		t.Fatal()
	}

	testRoutine, err := NewRoutine(cfgFilePath, 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	_, err = testRoutine.GetStatus()
	if err != nil {
		t.Fatal()
	}
}

func TestGetServices(t *testing.T) {
	testRoutine := Routine{
		Client: http.DefaultClient,
	}

	cfg := config.Config{
		Services: []config.Service{
			{
				Icon: "",
				Name: "Test 1",
				URL:  "https://xxxxxxxx.xxxxxx",
			},
		},
	}

	services := testRoutine.getServices(&cfg)
	if services[0].Name != "Test 1" {
		t.Fatal()
	}
}

func TestGetNotes(t *testing.T) {
	testRoutine := Routine{
		Client: http.DefaultClient,
	}

	cfg := config.Config{
		Notes: []config.Note{
			{
				Name: "Test 1",
				Text: "...",
			},
		},
	}

	notes := testRoutine.getNotes(&cfg)
	if notes[0].Name != "Test 1" {
		t.Fatal()
	}
}

func TestIcon(t *testing.T) {
	testRoutine := Routine{
		Client: http.DefaultClient,
	}

	service := config.Service{
		Icon: "data:image/png;base64,TEST",
	}
	icon := testRoutine.getIconData(service)
	if icon != "data:image/png;base64,TEST" {
		t.Fail()
	}

	service = config.Service{
		Icon: "data:XXXXX",
	}
	icon = testRoutine.getIconData(service)
	if icon != "" {
		t.Fail()
	}

	service = config.Service{
		Icon: "https://xxxxxxxx.xxxxxx",
	}
	icon = testRoutine.getIconData(service)
	if icon != "" {
		t.Fail()
	}
	service = config.Service{
		URL: "https://xxxxxxxx.xxxxxx",
	}
	icon = testRoutine.getIconData(service)
	if icon != "" {
		t.Fail()
	}
}
