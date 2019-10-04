module github.com/rishi-anand/k8s-dynamic-client-CRDs

go 1.12

require (
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/spectrocloud/palette v0.0.0-20191001200450-d2d1b816efd1
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20191001170739-f9e2070545dc // indirect
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0 // indirect
	gopkg.in/yaml.v2 v2.2.3 // indirect
	k8s.io/apimachinery v0.0.0-20191001195453-082230a5ffdd
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20190923111123-69764acb6e8e // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190704095032-f4ca3d3bdf1d
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190704094733-8f6ac2502e51
	sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v0.2.1
)
