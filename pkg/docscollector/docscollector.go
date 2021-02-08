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

	mdfile, err := os.Create(file + ".md")
	if err != nil {
		log.Fatalf("Error Creating Markdown Files : %s", err)
	}
	t := template.Must(template.New("tmpl").Parse(docsTemplateText))
	plan, _ := ioutil.ReadFile(file + ".json")
	var names []Docsdb
	err = json.Unmarshal(plan, &names)
	if err != nil {
		log.Fatalf("Error Parsing the Metadata :%s", err)
	}

	t.Execute(mdfile, names)

}
