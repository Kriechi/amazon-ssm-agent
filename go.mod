module github.com/aws/amazon-ssm-agent

go 1.19

replace github.com/aws/aws-sdk-go => ./extra/aws-sdk-go

replace github.com/nightlyone/lockfile => ./extra/lockfile

require (
	github.com/Jeffail/gabs v1.0.0
	github.com/Workiva/go-datastructures v1.0.53
	github.com/aws/aws-sdk-go v1.44.261
	github.com/carlescere/scheduler v0.0.0-20150615230211-9b78eac89dfb
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/coreos/go-semver v0.2.0
	github.com/creack/pty v1.1.11
	github.com/digitalocean/go-smbios v0.0.0-20180907143718-390a4f403a8e
	github.com/fsnotify/fsnotify v1.5.1
	github.com/go-git/go-git/v5 v5.11.0
	github.com/google/go-github v0.0.0-20170604025028-a117bb2595a5 // this corresponds to v8
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/gorilla/websocket v1.4.2
	github.com/hectane/go-acl v0.0.0-20151103031024-7f56832555fc // Don't update -- breaks
	github.com/mitchellh/go-ps v1.0.0
	github.com/nightlyone/lockfile v0.0.0
	github.com/pborman/ansi v1.0.0
	github.com/stretchr/testify v1.8.4
	github.com/twinj/uuid v0.0.0-20151029044442-89173bcdda19 // Don't update -- breaks
	github.com/xtaci/smux v1.5.15
	go.nanomsg.org/mangos/v3 v3.3.0
	golang.org/x/crypto v0.17.0
	golang.org/x/net v0.19.0
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
	golang.org/x/sync v0.3.0
	golang.org/x/sys v0.15.0
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/yaml.v2 v2.3.0
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230828082145-3c4c8a2d2371 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/skeema/knownhosts v1.2.1 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
