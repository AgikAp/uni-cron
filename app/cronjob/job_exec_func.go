package cronjob

import (
	"admin_backend/app/helpers"
	"bytes"
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
	Data map[string]string `json:"data"`
}

func (j *JobExecutors) execARI() {
	response := ResponseAri{}

	newPayload := bytes.NewBufferString("type=Link")
	newPayload.WriteString("&linkName=jsdlkfjsdlkffasdlkj")
	newPayload.WriteString("&linkUrl=http://sdfkljdslnkvcsdankl.com")

	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjE0ODMiLCJpZCI6IjFkOWY0NmZkLWQ2NjUtNGVmNC1hYWU4LTI2MDQxYjdhNWIxNSIsInN1YiI6IjFkOWY0NmZkLWQ2NjUtNGVmNC1hYWU4LTI2MDQxYjdhNWIxNSIsInRva2VuU3NvIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SjFibWx4ZFdWZmJtRnRaU0k2SWpFeE1EQXhNalUxSWl3aVpXMWhhV3dpT2lKaGMyVndMblJoY25saGJtRkFZbWx2Wm1GeWJXRXVZMjh1YVdRaUxDSnVZbVlpT2pFM01UYzFOelkyTWprc0ltVjRjQ0k2TVRjeE56VTRNREl5T1N3aWFXRjBJam94TnpFM05UYzJOakk1TENKcGMzTWlPaUpvZEhSd09pOHZiRzlqWVd4b2IzTjBPakU1TkRFMUlpd2lZWFZrSWpvaVFXNTVJbjAuYUhxR19NMWtud2hQWWk4NXdVWDZDZThidXFGOU5HdjRYalRVTzBuOU9EZyIsImNvbXBhbnlDb2RlIjoiQklPRiIsInVzZXJUeXBlIjoiVXNlciBIb2xkaW5nIiwiaWF0IjoxNzE3NTc2NjYyLCJleHAiOjE3MTc2NjMwNjIsImlzcyI6InRlbXBsYXRlLWJhY2tlbmQifQ.TbXiOu0EGxjMZv7ESPO0fSHpBucg4h0d8yWjdxbmn7o")
	helpers.ExecWebService("http://192.168.70.115:3000/api/admin/material-learning", "POST", header, newPayload, &response)
}
