module DailyServer

require (
	github.com/SAP/go-hdb v0.13.2 // indirect
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/Unknwon/goconfig v0.0.0-20180308125533-ef1e4c783f8f
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190130090545-97fb8dcd4a63 // indirect
	github.com/araddon/gou v0.0.0-20190110011759-c797efecbb61 // indirect
	github.com/aws/aws-sdk-go v1.16.36 // indirect
	github.com/boombuler/barcode v1.0.0 // indirect
	github.com/briankassouf/jose v0.9.1 // indirect
	github.com/bsm/sarama-cluster v2.1.15+incompatible // indirect
	github.com/centrify/cloud-golang-sdk v0.0.0-20180119173102-7c97cc6fde16 // indirect
	github.com/chrismalek/oktasdk-go v0.0.0-20181212195951-3430665dfaa0 // indirect
	github.com/circonus-labs/circonus-gometrics v2.2.6+incompatible // indirect
	github.com/containerd/continuity v0.0.0-20181203112020-004b46473808 // indirect
	github.com/coredns/coredns v1.3.1 // indirect
	github.com/coreos/bbolt v1.3.2 // indirect
	github.com/coreos/etcd v3.3.12+incompatible // indirect
	github.com/coreos/go-oidc v2.0.0+incompatible // indirect
	github.com/dancannon/gorethink v4.0.0+incompatible // indirect
	github.com/digitalocean/godo v1.1.3 // indirect
	github.com/dimchansky/utfbom v1.1.0 // indirect
	github.com/duosecurity/duo_api_golang v0.0.0-20190107154727-539434bf0d45 // indirect
	github.com/envoyproxy/go-control-plane v0.6.7 // indirect
	github.com/fullsailor/pkcs7 v0.0.0-20180613152042-8306686428a5 // indirect
	github.com/gammazero/deque v0.0.0-20190130191400-2afb3858e9c7 // indirect
	github.com/gammazero/workerpool v0.0.0-20181230203049-86a96b5d5d92 // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/gin-contrib/multitemplate v0.0.0-20190301062633-f9896279eead
	github.com/gin-gonic/gin v1.3.0
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-ldap/ldap v3.0.1+incompatible // indirect
	github.com/go-ole/go-ole v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/builder v0.3.4 // indirect
	github.com/go-xorm/core v0.6.2
	github.com/go-xorm/xorm v0.7.1
	github.com/gocql/gocql v0.0.0-20190126123547-8516aabb0f99 // indirect
	github.com/gogo/googleapis v1.1.0 // indirect
	github.com/gogo/protobuf v1.1.1
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75 // indirect
	github.com/gorilla/websocket v1.4.0

	github.com/grpc-ecosystem/grpc-gateway v1.6.4 // indirect
	github.com/hashicorp/consul v1.4.3 // indirect
	github.com/hashicorp/go-discover v0.0.0-20190117190025-e88f86e24f50 // indirect
	github.com/hashicorp/go-gcp-common v0.0.0-20180425173946-763e39302965 // indirect
	github.com/hashicorp/go-hclog v0.0.0-20190109152822-4783caec6f2e // indirect
	github.com/hashicorp/go-plugin v0.0.0-20190129155509-362c99b11937 // indirect
	github.com/hashicorp/go-retryablehttp v0.5.2 // indirect
	github.com/hashicorp/go-rootcerts v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-version v1.1.0 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/hashicorp/hil v0.0.0-20190129155652-59d7c1fee952 // indirect
	github.com/hashicorp/net-rpc-msgpackrpc v0.0.0-20151116020338-a14192a58a69 // indirect
	github.com/hashicorp/nomad v0.8.7 // indirect
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea // indirect
	github.com/hashicorp/serf v0.8.2 // indirect
	github.com/hashicorp/vault v1.0.3 // indirect
	github.com/hashicorp/vault-plugin-auth-alicloud v0.0.0-20181109180636-f278a59ca3e8 // indirect
	github.com/hashicorp/vault-plugin-auth-azure v0.0.0-20181207232528-4c0b46069a22 // indirect
	github.com/hashicorp/vault-plugin-auth-centrify v0.0.0-20180816201131-66b0a34a58bf // indirect
	github.com/hashicorp/vault-plugin-auth-gcp v0.0.0-20181210200133-4d63bbfe6fcf // indirect
	github.com/hashicorp/vault-plugin-auth-jwt v0.0.0-20190128234440-a608a5ad1c24 // indirect
	github.com/hashicorp/vault-plugin-auth-kubernetes v0.0.0-20181130162533-091d9e5d5fab // indirect
	github.com/hashicorp/vault-plugin-secrets-ad v0.0.0-20181109182834-540c0b6f1f11 // indirect
	github.com/hashicorp/vault-plugin-secrets-alicloud v0.0.0-20181109181453-2aee79cc5cbf // indirect
	github.com/hashicorp/vault-plugin-secrets-azure v0.0.0-20181207232500-0087bdef705a // indirect
	github.com/hashicorp/vault-plugin-secrets-gcp v0.0.0-20180921173200-d6445459e80c // indirect
	github.com/hashicorp/vault-plugin-secrets-gcpkms v0.0.0-20190116164938-d6b25b0b4a39 // indirect
	github.com/hashicorp/vault-plugin-secrets-kv v0.0.0-20190115203747-edbfe287c5d9 // indirect
	github.com/influxdata/influxdb v1.7.4 // indirect
	github.com/influxdata/platform v0.0.0-20190117200541-d500d3cf5589 // indirect
	github.com/jeffchao/backoff v0.0.0-20140404060208-9d7fd7aa17f2 // indirect
	github.com/jefferai/jsonx v1.0.0 // indirect
	github.com/json-iterator/go v1.1.5
	github.com/keybase/go-crypto v0.0.0-20181127160227-255a5089e85a // indirect
	github.com/lyft/protoc-gen-validate v0.0.13 // indirect
	github.com/mattbaird/elastigo v0.0.0-20170123220020-2fe47fd29e4b // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/michaelklishin/rabbit-hole v1.4.0 // indirect
	github.com/micro/cli v0.1.0
	github.com/micro/go-config v0.13.3
	github.com/micro/go-micro v1.0.0
	github.com/micro/go-plugins v0.22.0
	github.com/micro/go-web v1.0.0
	github.com/miekg/dns v1.1.6 // indirect
	github.com/ory-am/common v0.4.0 // indirect
	github.com/ory/dockertest v3.3.4+incompatible // indirect
	github.com/pascaldekloe/goe v0.1.0 // indirect
	github.com/pborman/uuid v0.0.0-20180827223501-4c1ecd6722e8 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/pquerna/otp v1.1.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
	github.com/vmware/vic v1.4.3 // indirect
	go.etcd.io/bbolt v1.3.2 // indirect
	go.uber.org/zap v1.9.1
	golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a // indirect
	golang.org/x/net v0.0.0-20190320064053-1272bf9dcd53 // indirect
	golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6 // indirect
	golang.org/x/sys v0.0.0-20190318195719-6c81ef8f67ca // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/gorethink/gorethink.v4 v4.1.0 // indirect
	gopkg.in/ory-am/dockertest.v2 v2.2.3 // indirect
	gopkg.in/square/go-jose.v2 v2.2.2 // indirect
	k8s.io/api v0.0.0-20190313115550-3c12c96769cc // indirect
	k8s.io/client-go v10.0.0+incompatible // indirect
	k8s.io/klog v0.2.0 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
	layeh.com/radius v0.0.0-20190118135028-0f678f039617 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)