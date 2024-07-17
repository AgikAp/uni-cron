package cronjob

import (
	"admin_backend/app/helpers"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JobExecutors struct {
}

func NewJobExecutors() *JobExecutors {
	return &JobExecutors{}
}

type ResponseAri interface{}

type RequestAri struct {
	ELearningParticipantId *string             `json:"eLearningParticipantId"`
	ELearningSectionId     *string             `json:"eLearningSectionId"`
	MaterialActivity       *[]MaterialActivity `json:"materialActivity"`
}

type MaterialActivity struct {
	ID                         *string `json:"id"`
	ELearningParticipantId     *string `json:"eLearningParticipantId"`
	ELearningSectionActivityId *string `json:"eLearningSectionActivityId"`
	ELearningMaterialId        *string `json:"eLearningMaterialId"`
	Status                     *string `json:"status"`
	Progress                   *int    `json:"progress"`
}

var url = "http://192.168.1.19:3000/api/e-learning/change-material-progress"
var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjIzMTkiLCJpZCI6Ijk3OTNlOTI2LWZkMDQtNDEwYy1iNTczLTA1N2RmMTg4NmU3YyIsInN1YiI6Ijk3OTNlOTI2LWZkMDQtNDEwYy1iNTczLTA1N2RmMTg4NmU3YyIsInRva2VuU3NvIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SjFibWx4ZFdWZmJtRnRaU0k2SWpFeE1EQXhPVFF5SWl3aVpXMWhhV3dpT2lKaGNtbHpkR1YxY3k1ellXeGxjSEJoYm1kQVltbHZabUZ5YldFdVkyOHVhV1FpTENKdVltWWlPakUzTWpFeE9UYzVPRGNzSW1WNGNDSTZNVGN5TVRJd09EYzROeXdpYVdGMElqb3hOekl4TVRrM09UZzNMQ0pwYzNNaU9pSm9kSFJ3T2k4dmJHOWpZV3hvYjNOME9qRTVOREUxSWl3aVlYVmtJam9pUVc1NUluMC40UjlHS2c5UU5pUHZleUNCaFYxZU5NM3ZSQjMzelpFdElPSW5zVnRNWmFRIiwiY29tcGFueUNvZGUiOiJCSU9GIiwidXNlclR5cGUiOiJVc2VyIEhvbGRpbmciLCJpYXQiOjE3MjExOTgwMDYsImV4cCI6MTcyMTI4NDQwNiwiaXNzIjoidGVtcGxhdGUtYmFja2VuZCJ9.7StSPge0EddyXv4TvNLNRqD6gSNN1fEENZD7ZLinjS8"
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
						"status": "Done",
						"progress": 1
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

	fmt.Println("REQUEST ", bytes.NewBuffer(fixPayload).String())
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
