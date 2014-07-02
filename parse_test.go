package conf

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
		t.Fatalf("Not Equal:\nReceived: '%+v'\nExpected: '%+v'\n", m, ex)
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

title = "Braket Example"

owner {
  name: "Tom Preston-Werner"
  organization: "GitHub"
  bio: "GitHub Cofounder & CEO\nLikes tater tots and beer."
  dob: 1979-05-27T07:32:00Z # First class dates? Why not?
}

database {
  server: "192.168.1.1"
  ports: [ 8001, 8001, 8002 ]
  connection_max: 5000
  enabled: true
}

servers{

  # You can indent as you please. Tabs or spaces. TOML don't care.
  alpha {
    ip: "10.0.0.1"
    dc: "eqdc10"
  }

  beta {
   ip: "10.0.0.2"
   dc: "eqdc10"
  }

clients{ data : [ ["gamma", "delta"], [1, 2] ] }

  # Line breaks are OK when inside arrays
  hosts : [
     "alpha",
     "omega"
  ]
}`

type mm map[string]interface{}

func TestComplexCase(t *testing.T) {
	ex := mm{
		"title": "Bracket Example",
		"owner": mm{"name", "Tom Preston-Werner", "organization": "GitHub", "bio": "Github", "dob": time.Time()},
	}
	test(t, sample4, ex)
}
