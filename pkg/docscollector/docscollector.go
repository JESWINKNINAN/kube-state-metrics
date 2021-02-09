/*
Copyright 2019 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package docscollector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

//Docsdb is struct to store the decoded docs metadata
type Docsdb struct {
	MetricName    string `json:"MetricName"`
	MetricType    string `json:"MetricType"`
	Description   string `json:"Description"`
	Unit          string `json:"Unit"`
	LabelsAndTags string `json:"LabelsAndTags"`
	Status        string `json:"Status"`
}

//docsTemplateText is template text for Docs
const docsTemplateText = `| Metric name | Metric Type | Description | Unit (where applicable) | Labels/tags | Status |
| ----------- | ----------- | ----------- | ----------------------- | ----------- | ------ |{{range .}}
|{{.MetricName }}|{{.MetricType}}|{{.Description}}|{{.Unit}}|{{.LabelsAndTags}}|{{.Status}}|{{end}}`

//DocsCreate creates md files automatically
func DocsCreate(file string) {

	mdfile, err := os.Create("../../docs/" + file + ".md")
	if err != nil {
		log.Fatalf("Error Creating Markdown Files : %s", err)
	}
	t := template.Must(template.New("tmpl").Parse(docsTemplateText))
	getDetails, _ := ioutil.ReadFile(file + ".json")
	var docsMetaData []Docsdb
	err = json.Unmarshal(getDetails, &docsMetaData)
	if err != nil {
		log.Fatalf("Error Parsing the Metadata :%s", err)
	}

	t.Execute(mdfile, docsMetaData)

}
