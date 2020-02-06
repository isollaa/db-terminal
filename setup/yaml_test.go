package setup

import "testing"

const case1 = `
driver:
    mongo:
        host:        "localhost"
        port:        27017
        username:    ""
        password:    ""
        dbname:      "xsaas_ctms"
        collection:  "relationship"
beauty: false
prompt: false
`

const case2 = `
driver:
    mongo:
        host:        "localhost"
        port:        "27017"
        username:    ""
        password:    ""
        dbname:      "xsaas_ctms"
        collection:  "relationship"
beauty: false
prompt: false
`

const case3 = `
driver:
    mongo:
        host:        "localhost"
        port:        27017
        username:    ""
        password:    12345
        dbname:      "xsaas_ctms"
        collection:  "relationship"
beauty: false
prompt: false
`

const case4 = `
driver:
    mongo:
        host:        "localhost"
        port:        27017
        username:    ""
        password:    ""
        dbname:      "xsaas_ctms"
        collection:  "relationship"
beauty: "false"
prompt: false
`

const case5 = `
driver:
    mongo:
        host:        "localhost"
        port:        27017
        username:    ""
        password:    ""
        dbname:      "xsaas_ctms"
        collection:  "relationship"
beauty: false
prompt: "false"
`

type testCase struct {
	Raw   string
	Valid bool
}

var testCases = []testCase{
	{
		Raw:   case1,
		Valid: true,
	},
	{
		Raw:   case2,
		Valid: false,
	},
	{
		Raw:   case3,
		Valid: true,
	},
	{
		Raw:   case4,
		Valid: false,
	},
	{
		Raw:   case5,
		Valid: false,
	},
}

func TestGetYamlConfig(t *testing.T) {
	for i, v := range testCases {
		_, err := GetYamlConfig([]byte(v.Raw))
		valid := err == nil
		println(valid)
		if valid != v.Valid {
			t.Error(i, err)
		}
	}
}
