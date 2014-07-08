Bracket [![Build Status](https://secure.travis-ci.org/Mistobaan/bracket.svg?branch=develop)](http://travis-ci.org/Mistobaan/bracket)

# Bracket Configuration

This project is based heavily on Gnats'd configuration system

```
# Sample config file

port: 4242
net: apcera.me # net interface

http_port: 8222

authorization {
  user:     derek
  password: T0pS3cr3t
  timeout:  1
}

cluster {
  host: '127.0.0.1'
  port: 4244

  authorization {
    user: route_user
    password: T0pS3cr3tT00!
    timeout: 0.5
  }

  # Routes are actively solicited and connected to from this server.
  # Other servers can connect to us if they supply the correct credentials
  # in their routes definitions from above.

  routes = [
    nats-route://user1:pass1@127.0.0.1:4245
    nats-route://user2:pass2@127.0.0.1:4246
  ]
}

# logging options
debug:   false
trace:   true
logtime: false
log_file: "/tmp/gnatsd.log"

#pid file
pid_file: "/tmp/gnatsd.pid"
```


Same toml document in bracket would be
```
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
```



## License

(The MIT License)

Copyright (c) 2012-2014 Apcera Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
