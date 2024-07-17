package cronjob

import (
	"admin_backend/app/helpers"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type JobExecutors struct {
}

func NewJobExecutors() *JobExecutors {
	return &JobExecutors{}
}

func (j *JobExecutors) JobItems() {
	j.execARI()
}

type ResponseAri struct {
	Data map[string]interface{} `json:"data"`
}

type RequestAri struct {
	ELearningParticipantId *string           `json:"eLearningParticipantId"`
	ELearningSectionId     *string           `json:"eLearningSectionId"`
	MaterialActivity       *MaterialActivity `json:"materialActivity"`
}

type MaterialActivity struct {
	ID                         *string `json:"id"`
	ELearningParticipantId     *string `json:"eLearningParticipantId"`
	ELearningSectionActivityId *string `json:"eLearningSectionActivityId"`
	ELearningMaterialId        *string `json:"eLearningMaterialId"`
	Status                     *string `json:"status"`
	Progress                   *string `json:"progress"`
}

var url = ""
var token = ""
var stringPayload = `
{
    "eLearningParticipantId": "e2105ccd-c39a-4a6d-9f47-6d5d8c4e3e7d",
    "eLearningSectionId": "0196fe0f-866e-4d12-a264-ac6a6f7c0df2",
    "materialActivity": [
        {
            "id": "f0674a16-1e94-4a83-b864-cf95249f53a8",
            "eLearningParticipantId": "e2105ccd-c39a-4a6d-9f47-6d5d8c4e3e7d",
            "eLearningSectionActivityId": "0196fe0f-866e-4d12-a264-ac6a6f7c0df2",
            "eLearningMaterialId": "37b8cb82-1d74-4c46-8cdc-a16a82ce308e",
            "status": null,
            "progress": null
        }
    ]
}`

func (j *JobExecutors) execARI() {
	var paylaod RequestAri
	err := json.Unmarshal([]byte(stringPayload), &paylaod)
	if err != nil {
		log.Fatalln(err)
	}

	fixPayload, err := json.Marshal(paylaod)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var response ResponseAri

	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", "Bearer "+token)
	helpers.ExecWebService(url, "POST", header, bytes.NewBuffer(fixPayload), &response)

	// response := ResponseAri{}

	// newPayload := bytes.NewBufferString("type=Link")
	// newPayload.WriteString("&linkName=jsdlkfjsdlkffasdlkj")
	// newPayload.WriteString("&linkUrl=http://sdfkljdslnkvcsdankl.com")

	// header := http.Header{}
	// header.Set("Content-Type", "application/x-www-form-urlencoded")
	// header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjE0ODMiLCJpZCI6IjFkOWY0NmZkLWQ2NjUtNGVmNC1hYWU4LTI2MDQxYjdhNWIxNSIsInN1YiI6IjFkOWY0NmZkLWQ2NjUtNGVmNC1hYWU4LTI2MDQxYjdhNWIxNSIsInRva2VuU3NvIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SjFibWx4ZFdWZmJtRnRaU0k2SWpFeE1EQXhNalUxSWl3aVpXMWhhV3dpT2lKaGMyVndMblJoY25saGJtRkFZbWx2Wm1GeWJXRXVZMjh1YVdRaUxDSnVZbVlpT2pFM01UYzFOelkyTWprc0ltVjRjQ0k2TVRjeE56VTRNREl5T1N3aWFXRjBJam94TnpFM05UYzJOakk1TENKcGMzTWlPaUpvZEhSd09pOHZiRzlqWVd4b2IzTjBPakU1TkRFMUlpd2lZWFZrSWpvaVFXNTVJbjAuYUhxR19NMWtud2hQWWk4NXdVWDZDZThidXFGOU5HdjRYalRVTzBuOU9EZyIsImNvbXBhbnlDb2RlIjoiQklPRiIsInVzZXJUeXBlIjoiVXNlciBIb2xkaW5nIiwiaWF0IjoxNzE3NTc2NjYyLCJleHAiOjE3MTc2NjMwNjIsImlzcyI6InRlbXBsYXRlLWJhY2tlbmQifQ.TbXiOu0EGxjMZv7ESPO0fSHpBucg4h0d8yWjdxbmn7o")
	// helpers.ExecWebService("http://192.168.70.115:3000/api/admin/material-learning", "POST", header, newPayload, &response)
}
