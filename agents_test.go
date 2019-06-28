package wazuh_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mrtc0/wazuh"
)

func TestGetAllAgents(t *testing.T) {
	data := []wazuh.AgentInformation{
		wazuh.AgentInformation{
			Status:        "Active",
			Name:          "wazuh-manager",
			NodeName:      "node01",
			DateAdd:       "2019-01-05 08:35:27",
			Version:       "Wazuh v3.7.2",
			LastKeepAlive: "9999-12-31 23:59:59",
			Os: wazuh.Os{
				Major:    "16",
				Name:     "Ubuntu",
				Uname:    "Linux |redis |4.15.0-36-generic |#39~16.04.1-Ubuntu SMP Tue Sep 25 08:59:23 UTC 2018 |x86_64",
				Platform: "ubuntu",
				Version:  "16.04.5 LTS",
				Codename: "Xenial Xerus",
				Arch:     "x86_64",
				Minor:    "04",
			},
			ID:        "004",
			ConfigSum: "ab73af41699f13fdd81903b5f23d8d00",
			Group: []string{
				"default",
			},
			MergedSum: "f1a9e24e02ba4cc5ea80a9d3feb3bb9a",
		},
	}

	response := wazuh.GetAllAgentsResponse{
		Error: 0,
		Data: wazuh.AgentInformationData{
			TotalItems: 1,
			Items:      data,
		},
	}

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		t.Fatalf("want no err, but has error %#v", err)
	}

	once.Do(startServer)

	http.HandleFunc("/agents", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, string(b)) })

	client, err := wazuh.New(endpoint)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err.Error())
	}
	agents, err := client.GetAllAgents()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err.Error())
	}

	expected := "Active"
	agent := (*agents)[0]

	if agent.Status != expected {
		t.Errorf("expected %s, but got %s\n", expected, agent.Status)
	}
}

func TestGetGroups(t *testing.T) {
	once.Do(startServer)
	http.HandleFunc("/agents/groups", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"error":0,"data":{"totalItems":1,"items":[{"count":4,"mergedSum":"345bce156ecdbf8cfaee550d003341a9","configSum":"ab73af41699f13fdd81903b5f23d8d00","name":"default"}]}}`)
	})

	wazuh, err := wazuh.New(endpoint)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err.Error())
	}

	groups, err := wazuh.GetGroups()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err.Error())
	}

	expected := "default"
	group := (*groups)[0]

	if group.Name != expected {
		t.Errorf("expected %s, but got %s\n", expected, group.Name)
	}
}
