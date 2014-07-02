package bracket

import (
	"reflect"
	"testing"
	"time"
)

// Test to make sure we get what we expect.

func test(t *testing.T, data string, ex map[string]interface{}) {
	m, err := Parse(data)
	if err != nil {
		t.Fatalf("Received err: %v\n", err)
	}
	if m == nil {
		t.Fatal("Received nil map")
	}

	if !reflect.DeepEqual(m, ex) {
		t.Logf("%+v\n", m)
		t.Fatalf("Not Equal:\nReceived: '%#v'\nExpected: '%#v'\n", m, ex)
	}
}

func TestSimpleTopLevel(t *testing.T) {
	ex := map[string]interface{}{
		"foo": "1",
		"bar": float64(2.2),
		"baz": true,
		"boo": int64(22),
	}
	test(t, "foo='1'; bar=2.2; baz=true; boo=22", ex)
}

var sample1 = `
foo  {
  host {
    ip   = '127.0.0.1'
    port = 4242
  }
  servers = [ "a.com", "b.com", "c.com"]
}
`

func TestSample1(t *testing.T) {
	ex := map[string]interface{}{
		"foo": map[string]interface{}{
			"host": map[string]interface{}{
				"ip":   "127.0.0.1",
				"port": int64(4242),
			},
			"servers": []interface{}{"a.com", "b.com", "c.com"},
		},
	}
	test(t, sample1, ex)
}

var cluster = `
cluster {
  port: 4244

  authorization {
    user: route_user
    password: top_secret
    timeout: 1
  }

  # Routes are actively solicited and connected to from this server.
  # Other servers can connect to us if they supply the correct credentials
  # in their routes definitions from above.

  // Test both styles of comments

  routes = [
    nats-route://foo:bar@apcera.me:4245
    nats-route://foo:bar@apcera.me:4246
  ]
}
`

func TestSample2(t *testing.T) {
	ex := map[string]interface{}{
		"cluster": map[string]interface{}{
			"port": int64(4244),
			"authorization": map[string]interface{}{
				"user":     "route_user",
				"password": "top_secret",
				"timeout":  int64(1),
			},
			"routes": []interface{}{
				"nats-route://foo:bar@apcera.me:4245",
				"nats-route://foo:bar@apcera.me:4246",
			},
		},
	}

	test(t, cluster, ex)
}

var sample3 = `
foo  {
  expr = '(true == "false")'
  text = 'This is a multi-line
text block.'
}
`

func TestSample3(t *testing.T) {
	ex := map[string]interface{}{
		"foo": map[string]interface{}{
			"expr": "(true == \"false\")",
			"text": "This is a multi-line\ntext block.",
		},
	}
	test(t, sample3, ex)
}

var sample4 = `
  array [
    { abc: 123 }
    { xyz: "word" }
  ]
`

func TestSample4(t *testing.T) {
	ex := map[string]interface{}{
		"array": []interface{}{
			map[string]interface{}{"abc": int64(123)},
			map[string]interface{}{"xyz": "word"},
		},
	}
	test(t, sample4, ex)
}

var complexCase = `
# This is a BRKT document. Boom.


database {
  server: "192.168.1.1"
  ports: [ 8001, 8001, 8002 ]
  connection_max: 5000
  enabled: true
}


title = "Braket Example"

owner {
  name: "Tom Preston-Werner"
  organization: "GitHub"
  bio: "GitHub Cofounder & CEO\nLikes tater tots and beer."
  dob: 0001-01-01T00:00:00Z # First class dates? Why not?
}


servers {

  # You can indent as you please. Tabs or spaces. Bracket don't care.
  alpha {
    ip: "10.0.0.1"
    dc: "eqdc10"
  }

  beta {
   ip: "10.0.0.2"
   dc: "eqdc10"
  }

  clients { data : [ ["gamma", "delta"], [1, 2] ] }

  # Line breaks are OK when inside arrays
  hosts : [
     "alpha",
     "omega"
  ]
}
`

func xxxTestComplexCase(t *testing.T) {

	date, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Fail()
	}

	ex := map[string]interface{}{
		"title": "Braket Example",

		"owner": map[string]interface{}{
			"organization": "GitHub",
			"bio":          "GitHub Cofounder & CEO\\nLikes tater tots and beer.",
			"dob":          date,
			"name":         "Tom Preston-Werner",
		},

		/*
			"database": map[string]interface{}{"server": "192.168.1.1", "ports": []interface{}{8001, 8001, 8002}, "connection_max": 5000, "enabled": true},
		*/
		"servers": map[string]interface{}{

			"alpha": map[string]interface{}{
				"ip": "10.0.0.1",
				"dc": "eqdc10",
			},

			"beta": map[string]interface{}{
				"ip": "10.0.0.2",
				"dc": "eqdc10",
			},

			"clients": map[string]interface{}{
				"data": []interface{}{[]interface{}{
					"gamma", "delta"}, []int{1, 2}}},
			"hosts": []string{"alpha", "omega"},
		},
	}

	test(t, complexCase, ex)
}
