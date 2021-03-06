package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
)

// Sandbox is a an individual sandbox with an application profile
type Sandbox struct {
	SandboxID   string `xml:"sandbox_id,attr"`
	SandboxName string `xml:"sandbox_name,attr"`
	Owner       string `xml:"owner,attr"`
}

// ParseSandboxList parses the getsandboxlist.do API and returns an array of Sandboxes
func ParseSandboxList(credsFile, appID string) ([]Sandbox, error) {
	var sandboxes []Sandbox

	sandboxListAPI, err := sandboxList(credsFile, appID)
	if err != nil {
		return nil, err
	}
	decoder := xml.NewDecoder(bytes.NewReader(sandboxListAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			// Read StartElement and check for flaw
			if se.Name.Local == "sandbox" {
				var sandbox Sandbox
				decoder.DecodeElement(&sandbox, &se)
				sandboxes = append(sandboxes, sandbox)
			}
			if se.Name.Local == "error" {
				return nil, errors.New("api for GetSandboxList returned with an error element")
			}
		}
	}
	return sandboxes, nil
}
