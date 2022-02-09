module github.com/crowdsecurity/cs-firewall-bouncer

go 1.14

require (
	github.com/ahmetb/dlog v0.0.0-20170105205344-4fb5f8204f26 // indirect
	github.com/antonmedv/expr v1.9.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf
	github.com/crowdsecurity/crowdsec v1.3.0
	github.com/crowdsecurity/go-cs-bouncer v0.0.0-20220209104231-631313ec9d39
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/runtime v0.23.0 // indirect
	github.com/go-openapi/strfmt v0.21.2 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/google/nftables v0.0.0-20220206174406-91d3b4571db1
	github.com/hashicorp/go-version v1.4.0 // indirect
	github.com/logrusorgru/grokky v0.0.0-20180829062225-47edf017d42c // indirect
	github.com/mdlayher/netlink v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/vishvananda/netns v0.0.0-20190625233234-7109fa855b0f // indirect
	go.mongodb.org/mongo-driver v1.8.3 // indirect
	golang.org/x/sys v0.0.0-20220207234003-57398862261d
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
	gopkg.in/yaml.v2 v2.4.0
)

exclude github.com/mattn/go-sqlite3 v2.0.3+incompatible

exclude github.com/mattn/go-sqlite3 v2.0.1+incompatible

replace github.com/koneu/natend => ./koneu/natend
