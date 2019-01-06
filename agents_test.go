package wazuh

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllAgents(t *testing.T) {
	data := []AgentInformation{
		AgentInformation{
			Status:        "Active",
			Name:          "wazuh-manager",
			NodeName:      "node01",
			DateAdd:       "2019-01-05 08:35:27",
			Version:       "Wazuh v3.7.2",
			LastKeepAlive: "9999-12-31 23:59:59",
			Os: Os{
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

	response := GetAllAgentsResponse{
		Error: 0,
		Data: AgentInformationData{
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
	APIURL := "http://" + serverAddr + "/"

	http.HandleFunc("/agents", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, string(b)) })

	client := New(APIURL)
	client.SetBasicAuth("user", "pass")
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
