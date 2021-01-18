// Copyright 2020 Orange SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.package apis

package config

import (
	"fmt"

	"github.com/Orange-OpenSource/nifikop/api/v1alpha1"
	"github.com/Orange-OpenSource/nifikop/pkg/util/nifi"
	"github.com/go-logr/logr"
)

var NifiPropertiesTemplate = `# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Core Properties #
nifi.flow.configuration.file=../data/flow.xml.gz
nifi.flow.configuration.archive.enabled=true
nifi.flow.configuration.archive.dir=../data/archive/
nifi.flow.configuration.archive.max.time=30 days
nifi.flow.configuration.archive.max.storage=500 MB
nifi.flow.configuration.archive.max.count=
nifi.flowcontroller.autoResumeState=true
nifi.flowcontroller.graceful.shutdown.period=10 sec
nifi.flowservice.writedelay.interval=500 ms
nifi.administrative.yield.duration=30 sec
# If a component has no work to do (is "bored"), how long should we wait before checking again for work?
nifi.bored.yield.duration=10 millis

nifi.authorizer.configuration.file=./conf/authorizers.xml
nifi.login.identity.provider.configuration.file=./conf/login-identity-providers.xml
nifi.templates.directory=../data/templates
nifi.ui.banner.text=
nifi.ui.autorefresh.interval=30 sec
nifi.nar.library.directory=./lib
nifi.nar.working.directory=./work/nar/
nifi.documentation.working.directory=./work/docs/components

####################
# State Management #
####################
nifi.state.management.configuration.file=./conf/state-management.xml
# The ID of the local state provider
nifi.state.management.provider.local=local-provider
# The ID of the cluster-wide state provider. This will be ignored if NiFi is not clustered but must be populated if running in a cluster.
nifi.state.management.provider.cluster=zk-provider
# Specifies whether or not this instance of NiFi should run an embedded ZooKeeper server
nifi.state.management.embedded.zookeeper.start=false
# Properties file that provides the ZooKeeper properties to use if <nifi.state.management.embedded.zookeeper.start> is set to true
nifi.state.management.embedded.zookeeper.properties=./conf/zookeeper.properties


# H2 Settings
nifi.database.directory=../data/database_repository
nifi.h2.url.append=;LOCK_TIMEOUT=25000;WRITE_DELAY=0;AUTO_SERVER=FALSE

# FlowFile Repository
nifi.flowfile.repository.implementation=org.apache.nifi.controller.repository.WriteAheadFlowFileRepository
nifi.flowfile.repository.directory=../flowfile_repository
nifi.flowfile.repository.partitions=256
nifi.flowfile.repository.checkpoint.interval=2 mins
nifi.flowfile.repository.always.sync=false

nifi.swap.manager.implementation=org.apache.nifi.controller.FileSystemSwapManager
nifi.queue.swap.threshold=20000
nifi.swap.in.period=5 sec
nifi.swap.in.threads=1
nifi.swap.out.period=5 sec
nifi.swap.out.threads=4

# Content Repository
nifi.content.repository.implementation=org.apache.nifi.controller.repository.FileSystemRepository
nifi.content.claim.max.appendable.size=1 MB
nifi.content.claim.max.flow.files=100
nifi.content.repository.directory.default=../content_repository
nifi.content.repository.archive.max.retention.period=3 days
nifi.content.repository.archive.max.usage.percentage=85%
nifi.content.repository.archive.enabled=true
nifi.content.repository.always.sync=false
nifi.content.viewer.url=/nifi-content-viewer/

# Provenance Repository Properties
nifi.provenance.repository.implementation=org.apache.nifi.provenance.WriteAheadProvenanceRepository
nifi.provenance.repository.debug.frequency=1_000_000
nifi.provenance.repository.encryption.key.provider.implementation=
nifi.provenance.repository.encryption.key.provider.location=
nifi.provenance.repository.encryption.key.id=
nifi.provenance.repository.encryption.key=

# Persistent Provenance Repository Properties
nifi.provenance.repository.directory.default=../provenance_repository
nifi.provenance.repository.max.storage.time=10 days
nifi.provenance.repository.max.storage.size={{ .ProvenanceStorage }}
nifi.provenance.repository.rollover.time=30 secs
nifi.provenance.repository.rollover.size=100 MB
nifi.provenance.repository.query.threads=2
nifi.provenance.repository.index.threads=2
nifi.provenance.repository.compress.on.rollover=true
nifi.provenance.repository.always.sync=false
nifi.provenance.repository.journal.count=16
# Comma-separated list of fields. Fields that are not indexed will not be searchable. Valid fields are: 
# EventType, FlowFileUUID, Filename, TransitURI, ProcessorID, AlternateIdentifierURI, Relationship, Details
nifi.provenance.repository.indexed.fields=EventType, FlowFileUUID, Filename, ProcessorID, Relationship
# FlowFile Attributes that should be indexed and made searchable.  Some examples to consider are filename, uuid, mime.type
nifi.provenance.repository.indexed.attributes=
# Large values for the shard size will result in more Java heap usage when searching the Provenance Repository
# but should provide better performance
nifi.provenance.repository.index.shard.size=500 MB
# Indicates the maximum length that a FlowFile attribute can be when retrieving a Provenance Event from
# the repository. If the length of any attribute exceeds this value, it will be truncated when the event is retrieved.
nifi.provenance.repository.max.attribute.length=65536

# Volatile Provenance Respository Properties
nifi.provenance.repository.buffer.size=100000

# Component Status Repository
nifi.components.status.repository.implementation=org.apache.nifi.controller.status.history.VolatileComponentStatusRepository
nifi.components.status.repository.buffer.size=1440
nifi.components.status.snapshot.frequency=1 min

# Site to Site properties
nifi.remote.input.secure={{ .SiteToSiteSecure}}
nifi.remote.input.http.enabled=true
nifi.remote.input.http.transaction.ttl=30 sec

# web properties #
nifi.web.war.directory=./lib
nifi.web.proxy.host={{ .WebProxyHosts }}
{{ .ListenerConfig }}

nifi.web.http.network.interface.default=
nifi.web.https.network.interface.default=
nifi.web.jetty.working.directory=./work/jetty
nifi.web.jetty.threads=200

# security properties #
nifi.sensitive.props.key=
nifi.sensitive.props.key.protected=
nifi.sensitive.props.algorithm=PBEWITHMD5AND256BITAES-CBC-OPENSSL
nifi.sensitive.props.provider=BC
nifi.sensitive.props.additional.keys=

{{ if .NifiCluster.Spec.ListenersConfig.SSLSecrets }}
nifi.security.keystore={{ .ServerKeystorePath }}/{{ .KeystoreFile }}
nifi.security.keystoreType=JKS
nifi.security.keystorePasswd={{ .ServerKeystorePassword }}
nifi.security.keyPasswd={{ .ServerKeystorePassword }}
nifi.security.truststore={{ .ServerKeystorePath }}/{{ .KeystoreFile }}
nifi.security.truststoreType=JKS
nifi.security.truststorePasswd={{ .ServerKeystorePassword }}
{{ end }}
nifi.security.needClientAuth={{ .NeedClientAuth }}
nifi.security.user.authorizer={{ .Authorizer }}
{{if .LdapConfiguration.Enabled}}
nifi.security.user.login.identity.provider=ldap-provider
{{else}}
nifi.security.user.login.identity.provider=
{{end}}
nifi.security.ocsp.responder.url=
nifi.security.ocsp.responder.certificate=

# OpenId Connect SSO Properties #
nifi.security.user.oidc.discovery.url=
nifi.security.user.oidc.connect.timeout=5 secs
nifi.security.user.oidc.read.timeout=5 secs
nifi.security.user.oidc.client.id=
nifi.security.user.oidc.client.secret=
nifi.security.user.oidc.preferred.jwsalgorithm=

# Apache Knox SSO Properties #
nifi.security.user.knox.url=
nifi.security.user.knox.publicKey=
nifi.security.user.knox.cookieName=hadoop-jwt
nifi.security.user.knox.audiences=

# Identity Mapping Properties #
# These properties allow normalizing user identities such that identities coming from different identity providers
# (certificates, LDAP, Kerberos) can be treated the same internally in NiFi. The following example demonstrates normalizing
# DNs from certificates and principals from Kerberos into a common identity string:
#
# nifi.security.identity.mapping.pattern.dn=^CN=(.*?), OU=(.*?), O=(.*?), L=(.*?), ST=(.*?), C=(.*?)$
# nifi.security.identity.mapping.value.dn=$1@$2
# nifi.security.identity.mapping.pattern.kerb=^(.*?)/instance@(.*?)$
# nifi.security.identity.mapping.value.kerb=$1@$2

# cluster common properties (all nodes must have same values) #
nifi.cluster.protocol.heartbeat.interval=5 sec
nifi.cluster.protocol.is.secure={{ .ClusterSecure }}

# cluster node properties (only configure for cluster nodes) #

nifi.cluster.is.node={{.IsNode}}
nifi.cluster.node.protocol.threads=10
nifi.cluster.node.protocol.max.threads=50
nifi.cluster.node.event.history.size=25
nifi.cluster.node.connection.timeout=5 sec
nifi.cluster.node.read.timeout=5 sec
nifi.cluster.node.max.concurrent.requests=100
nifi.cluster.firewall.file=
nifi.cluster.flow.election.max.wait.time=1 mins
nifi.cluster.flow.election.max.candidates=

# zookeeper properties, used for cluster management #
nifi.zookeeper.connect.string= {{ .ZookeeperConnectString }}
nifi.zookeeper.connect.timeout=3 secs
nifi.zookeeper.session.timeout=3 secs
nifi.zookeeper.root.node={{ .ZookeeperPath }}

# Zookeeper properties for the authentication scheme used when creating acls on znodes used for cluster management
# Values supported for nifi.zookeeper.auth.type are "default", which will apply world/anyone rights on znodes
# and "sasl" which will give rights to the sasl/kerberos identity used to authenticate the nifi node
# The identity is determined using the value in nifi.kerberos.service.principal and the removeHostFromPrincipal
# and removeRealmFromPrincipal values (which should align with the kerberos.removeHostFromPrincipal and kerberos.removeRealmFromPrincipal
# values configured on the zookeeper server).
nifi.zookeeper.auth.type=
nifi.zookeeper.kerberos.removeHostFromPrincipal=
nifi.zookeeper.kerberos.removeRealmFromPrincipal=

# kerberos #
nifi.kerberos.krb5.file=

# kerberos service principal #
nifi.kerberos.service.principal=
nifi.kerberos.service.keytab.location=

# kerberos spnego principal #
nifi.kerberos.spnego.principal=
nifi.kerberos.spnego.keytab.location=
nifi.kerberos.spnego.authentication.expiration=12 hours

# external properties files for variable registry
# supports a comma delimited list of file locations
nifi.variable.registry.properties=
`

//
func GenerateListenerSpecificConfig(
	l *v1alpha1.ListenersConfig,
	nodeId int32,
	namespace,
	crName string,
	headlessServiceEnabled bool,
	clusterDomain string,
	useExternalDNS bool,
	log logr.Logger) string {

	var nifiConfig string

	hostListener := nifi.ComputeNodeHostname(nodeId, crName, namespace,
		headlessServiceEnabled, clusterDomain, useExternalDNS)

	if !headlessServiceEnabled {
		hostListener = nifi.ComputeNodeAllNodeHostname(nodeId, crName, namespace,
			headlessServiceEnabled, clusterDomain, useExternalDNS)
	}

	clusterPortConfig := "nifi.cluster.node.protocol.port=\n"
	httpPortConfig := "nifi.web.http.port=\n"
	httpHostConfig := "nifi.web.http.host=\n"
	httpsPortConfig := "nifi.web.https.port=\n"
	httpsHostConfig := "nifi.web.https.host=\n"
	s2sPortConfig := "nifi.remote.input.socket.port=\n"

	for _, iListener := range l.InternalListeners {
		switch iListener.Type {
		case v1alpha1.ClusterListenerType:
			clusterPortConfig = fmt.Sprintf("nifi.cluster.node.protocol.port=%d", iListener.ContainerPort) + "\n"
		case v1alpha1.HttpListenerType:
			httpPortConfig = fmt.Sprintf("nifi.web.http.port=%d", iListener.ContainerPort) + "\n"
			httpHostConfig = fmt.Sprintf("nifi.web.http.host=%s", hostListener) + "\n"
		case v1alpha1.HttpsListenerType:
			httpsPortConfig = fmt.Sprintf("nifi.web.https.port=%d", iListener.ContainerPort) + "\n"
			httpsHostConfig = fmt.Sprintf("nifi.web.https.host=%s", hostListener) + "\n"
		case v1alpha1.S2sListenerType:
			s2sPortConfig = fmt.Sprintf("nifi.remote.input.socket.port=%d", iListener.ContainerPort) + "\n"
		}
	}
	nifiConfig = nifiConfig +
		clusterPortConfig +
		httpPortConfig +
		httpHostConfig +
		httpsPortConfig +
		httpsHostConfig +
		s2sPortConfig

	nifiConfig = nifiConfig + fmt.Sprintf("nifi.remote.input.host=%s", hostListener) + "\n"
	nifiConfig = nifiConfig + fmt.Sprintf("nifi.cluster.node.address=%s", hostListener) + "\n"
	return nifiConfig
}