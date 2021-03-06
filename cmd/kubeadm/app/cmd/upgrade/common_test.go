/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package upgrade

import (
	"bytes"
	"testing"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
)

func TestPrintConfiguration(t *testing.T) {
	var tests = []struct {
		cfg           *kubeadmapi.MasterConfiguration
		buf           *bytes.Buffer
		expectedBytes []byte
	}{
		{
			cfg:           nil,
			expectedBytes: []byte(""),
		},
		{
			cfg: &kubeadmapi.MasterConfiguration{
				KubernetesVersion: "v1.7.1",
			},
			expectedBytes: []byte(`[upgrade/config] Configuration used:
	api:
	  advertiseAddress: ""
	  bindPort: 0
	  controlPlaneEndpoint: ""
	apiVersion: kubeadm.k8s.io/v1alpha2
	auditPolicy:
	  logDir: ""
	  path: ""
	certificatesDir: ""
	cloudProvider: ""
	etcd:
	  caFile: ""
	  certFile: ""
	  dataDir: ""
	  endpoints: null
	  image: ""
	  keyFile: ""
	imageRepository: ""
	kind: MasterConfiguration
	kubeProxy: {}
	kubeletConfiguration: {}
	kubernetesVersion: v1.7.1
	networking:
	  dnsDomain: ""
	  podSubnet: ""
	  serviceSubnet: ""
	nodeName: ""
	privilegedPods: false
	token: ""
	unifiedControlPlaneImage: ""
`),
		},
		{
			cfg: &kubeadmapi.MasterConfiguration{
				KubernetesVersion: "v1.7.1",
				Networking: kubeadmapi.Networking{
					ServiceSubnet: "10.96.0.1/12",
				},
			},
			expectedBytes: []byte(`[upgrade/config] Configuration used:
	api:
	  advertiseAddress: ""
	  bindPort: 0
	  controlPlaneEndpoint: ""
	apiVersion: kubeadm.k8s.io/v1alpha2
	auditPolicy:
	  logDir: ""
	  path: ""
	certificatesDir: ""
	cloudProvider: ""
	etcd:
	  caFile: ""
	  certFile: ""
	  dataDir: ""
	  endpoints: null
	  image: ""
	  keyFile: ""
	imageRepository: ""
	kind: MasterConfiguration
	kubeProxy: {}
	kubeletConfiguration: {}
	kubernetesVersion: v1.7.1
	networking:
	  dnsDomain: ""
	  podSubnet: ""
	  serviceSubnet: 10.96.0.1/12
	nodeName: ""
	privilegedPods: false
	token: ""
	unifiedControlPlaneImage: ""
`),
		},
		{
			cfg: &kubeadmapi.MasterConfiguration{
				KubernetesVersion: "v1.7.1",
				Etcd: kubeadmapi.Etcd{
					SelfHosted: &kubeadmapi.SelfHostedEtcd{
						CertificatesDir:    "/var/foo",
						ClusterServiceName: "foo",
						EtcdVersion:        "v0.1.0",
						OperatorVersion:    "v0.1.0",
					},
				},
			},
			expectedBytes: []byte(`[upgrade/config] Configuration used:
	api:
	  advertiseAddress: ""
	  bindPort: 0
	  controlPlaneEndpoint: ""
	apiVersion: kubeadm.k8s.io/v1alpha2
	auditPolicy:
	  logDir: ""
	  path: ""
	certificatesDir: ""
	cloudProvider: ""
	etcd:
	  caFile: ""
	  certFile: ""
	  dataDir: ""
	  endpoints: null
	  image: ""
	  keyFile: ""
	  selfHosted:
	    certificatesDir: /var/foo
	    clusterServiceName: foo
	    etcdVersion: v0.1.0
	    operatorVersion: v0.1.0
	imageRepository: ""
	kind: MasterConfiguration
	kubeProxy: {}
	kubeletConfiguration: {}
	kubernetesVersion: v1.7.1
	networking:
	  dnsDomain: ""
	  podSubnet: ""
	  serviceSubnet: ""
	nodeName: ""
	privilegedPods: false
	token: ""
	unifiedControlPlaneImage: ""
`),
		},
	}
	for _, rt := range tests {
		rt.buf = bytes.NewBufferString("")
		printConfiguration(rt.cfg, rt.buf)
		actualBytes := rt.buf.Bytes()
		if !bytes.Equal(actualBytes, rt.expectedBytes) {
			t.Errorf(
				"failed PrintConfiguration:\n\texpected: %q\n\t  actual: %q",
				string(rt.expectedBytes),
				string(actualBytes),
			)
		}
	}
}
