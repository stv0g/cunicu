"use strict";(self.webpackChunkwice=self.webpackChunkwice||[]).push([[5439],{1669:(e,n,t)=>{t.d(n,{Z:()=>a});var i=t(7294),s=t(814);function a(e){let n={...e};n.language||(n.language="yaml"),n.title="/etc/cunicu.yaml";let t='# An interval at which cun\u012bcu will periodically check for added,\n# removed or modified WireGuard interfaces.\nwatch_interval: 1s\n\n## Signaling backends\n#\n# These backends are used for exchanging control-plane messages\n# between the peers.\n# E.g. ICE candidates, Peer information\nbackends:\n- grpc://signal.cunicu.li:443\n# - grpc://localhost:8080?insecure=true&skip_verify=true\n# - k8s:///path/to/your/kubeconfig.yaml?namespace=default\n\n\n# RPC control socket settings\nrpc:\n  # Path to a Unix socket for management\n  # and monitoring of the cunicu daemon.\n  socket: /var/run/cunicu.sock\n\n  # Start of cun\u012bcu daemon will block until\n  # its unblocked via the control socket.\n  # Mostly useful for testing automation\n  wait: false\n\n\n#### Interface settings start here\n# The following settings can be overwritten for each interface\n# using the \'interfaces\' settings (see below).\n# The following settings will be used as default.\n\n## WireGuard interface settings\n#\n# These settings configure WireGuard specific settings\n# of the interface.\n\n# A base64 private key generated by wg genkey.\n# Will be automatically generated if not provided.\nprivate_key: KLoqDLKgoqaUkwctTd+Ov3pfImOfadkkvTdPlXsuLWM=\n\n# Create WireGuard interfaces using bundled wireguard-go\n# user space implementation. This will be the default\n# if there is no WireGuard kernel module present.\nuserspace: false\n\n# A range constraint for an automatically assigned\n# selected listen port.\n# If the interface has no listen port specified, cun\u012bcu\n# will use the first available port from this range.\nlisten_port_range:\n  min: 52820\n  max: 65535\n\n# A 16-bit port for listening. Optional;\n# If not specified, first available port from listen_port_range\n# will be used.\nlisten_port: 51825\n\n# A 32-bit fwmark for outgoing packets which can be used\n# for Netfilter or TC classification.\n# If set to 0 or "off", this option is disabled.\n# May be specified in hexadecimal by prepending "0x". Optional.\nfwmark: 0x1000\n\n# The remote WireGuard peers provided as a dictionary\n# The keys of this dictionary are used as names for the peers\npeers:  \n  test:\n    # A base64 public key calculated by wg pubkey from a private key,\n    # and usually transmitted out of band\n    # to the author of the configuration file.\n    public_key: FlKHqqQQx+bTAq7+YhwEECwWRg2Ih7NQ48F/SeOYRH8=\n\n    # A base64 pre-shared key generated by wg genpsk.\n    # Optional, and may be omitted.\n    # This option adds an additional layer of symmetric-key\n    # cryptography to be mixed into the already existing\n    # public-key cryptography, for post-quantum resistance.\n    preshared_key: zu86NBVsWOU3cx4UKOQ6MgNj3gv8GXsV9ATzSemdqlI=\n\n    # A pre-shared passphrase which is used to derive a preshared key.\n    # cun\u012bcu is using Argon2id as the key derivation function.\n    preshared_key_passphrase: some-shared-passphrase\n\n    # An endpoint IP or hostname, followed by a colon,\n    # and then a port number. This endpoint will be updated\n    # automatically to the most recent source IP address and\n    # port of correctly authenticated packets from the peer.\n    # If provided, no endpoint discovery will be performed.\n    endpoint: vpn.example.com:51820\n\n    # A time duration, between 1 and 65535s inclusive, of how\n    # often to send an authenticated empty packet to the peer\n    # for the purpose of keeping a stateful firewall or NAT mapping\n    # valid persistently. For example, if the interface very rarely\n    # sends traffic, but it might at anytime receive traffic from a\n    # peer, and it is behind NAT, the interface might benefit from\n    # having a persistent keepalive interval of 25 seconds.\n    # If set to zero, this option is disabled.\n    # By default or when unspecified, this option is off.\n    # Most users will not need this. Optional.\n    persistent_keepalive: 120s\n\n    # A comma-separated list of IP (v4 or v6) addresses with\n    # CIDR masks from which incoming traffic for this peer is\n    # allowed and to which outgoing  traffic for this peer is directed.\n    # The catch-all 0.0.0.0/0 may be specified for matching\n    # all IPv4 addresses, and ::/0 may be specified for matching\n    # all IPv6 addresses. May be specified multiple times.\n    allowed_ips:\n    - 192.168.5.0/24\n\n## Basic interface settings\n#\n\n# The Maximum Transfer Unit of the WireGuard interface.\n# If not specified, the MTU is automatically determined from\n# the endpoint addresses or the system default route,\n# which is usually a sane choice.\n# However, to manually specify an MTU to override this\n# automatic discovery, this value may be specified explicitly.\nmtu: 1420\n\n# A list of IP (v4 or v6) addresses (optionally with CIDR masks)\n# to be assigned to the interface.\n# May be specified multiple times.\naddresses:\n- 10.10.0.1/24\n\n# A list of prefixes which cunicu uses to derive local addresses\n# from the interfaces public key\nprefixes:\n- fc2f:9a4d::/32\n- 10.237.0.0/16\n\n# A list of IP (v4 or v6) addresses to be set as the interface\'s\n# DNS servers, or non-IP hostnames to be set as the interface\'s\n# DNS search domains.\n# May be specified multiple times.\n# Upon bringing the interface up, this runs `resolvconf -a tun.INTERFACE -m 0 -x`\n# and upon bringing it down, this runs `resolvconf -d tun.INTERFACE`.\n# If these particular invocations of resolvconf(8) are undesirable,\n# custom hooks can be used instead.\ndns:\n- 1.1.1.1\n\n\n## Config synchronization\n#\n# Synchronize local WireGuard interface configuration with wg(8) config-files.\n\n# Enable config synchronization\nsync_config: false\n\n# Keep watching for changes in the configuration and apply them on-the-fly\nwatch_config: false\n\n## Route Synchronization\n#\n# Synchronize the kernel routing table with WireGuard\'s AllowedIPs setting\n# \n# It checks for routes in the kernel routing table which have a peers address\n# as next-hop and adds those routes to the AllowedIPs setting of the respective peer.\n#\n# In reverse, also networks listed in a peers AllowedIPs setting will be installed as a\n# kernel route with the peers address as the routes next-hop. \n\n# Enable route synchronization\nsync_routes: true\n\n# Kernel routing table which is used\n# On Linux, see /etc/iproute2/rt_tables for table ids and names\nrouting_table: 254\n\n# Keep watching the for changes in the kernel routing table via netlink multicast group.\nwatch_routes: true\n\n\n## /etc/hosts synchronization\n#\n# Synchronizes the local /etc/hosts file with host names and addresses of connected peers. \n\n# Enable hosts file synchronization\nsync_hosts: true\n\n# The domain name which is appended to each of the peer host names\ndomain: wg-local\n\n\n## Peer discovery\n#\n# Peer discovery finds new peers within the same community and adds them to the respective interface\n\n# Enable/disable peer discovery\ndiscover_peers: true\n\n# The hostname which gets advertised to remote peers\nhostname: my-node\n\n# A passphrase shared among all peers of the same community\ncommunity: "some-common-password"\n\n# Networks which are reachable via this peer and get advertised to remote peers\n# These will be part of this interfaces AllowedIPs at the remote peers.\nnetworks:\n- 192.168.1.0/24\n- 10.2.0.0/24\n\n# A list of WireGuard public keys which are accepted peers\n# If not configured, all peers will be accepted.\nwhitelist:\n- coNsGPwVPdpahc8U+dbbWGzTAdCd6+1BvPIYg10wDCI=\n\n# A list if WireGuard public keys which are rejected as peers\nblacklist:\n- AOZzBaNsoV7P8vo0D5UmuIJUQ7AjMbHbGt2EA8eAuEc=\n\n\n## Endpoint discovery\n#\n# Endpoint discovery uses Interactive Connectivity Establishment (ICE) as used by WebRTC to\n# gather a list of candidate endpoints and performs connectivity checks to find a suitable\n# endpoint address which can be used by WireGuard\n\n# Enable/disable endpoint discovery\ndiscover_endpoints: true\n\n# Interactive Connectivity Establishment (ICE) parameters\nice:\n  # A list of STUN and TURN servers used by ICE.\n  urls:\n  # Community provided STUN/TURN servers\n  - grpc://relay.cunicu.li\n\n  # Public STUN servers\n  - stun:stun3.l.google.com:19302\n  - stun:relay.webwormhole.io\n  - stun:stun.sipgate.net\n  - stun:stun.ekiga.net\n  - stun:stun.services.mozilla.com\n\n  # Caution: OpenRelay servers are located in Ontario, Canada.\n  # Beware of the latency!\n  # See also: https://www.metered.ca/tools/openrelay/\n  # - turn:openrelayproject:openrelayproject@openrelay.metered.ca:80\n  # - turn:openrelayproject:openrelayproject@openrelay.metered.ca:443\n  # - turn:openrelayproject:openrelayproject@openrelay.metered.ca:443?transport=tcp\n\n  # Credentials for STUN/TURN servers configured above.\n  username: ""\n  password: ""\n\n  # Allow connections to STUNS/TURNS servers for which we can not validate TLS certificates.\n  insecure_skip_verify: false\n\n  # Limit available network and candidate types.\n  # network_types: [udp4, udp6, tcp4, tcp6]\n  # candidate_types: [host, srflx, prflx, relay]\n\n  # A glob(7) pattern to match interfaces against which are used to gather ICE candidates (e.g. \\"eth[0-9]\\").\n  interface_filter: "*"\n\n  # Lite agents do not perform connectivity check and only provide host candidates.\n  lite: false\n\n  # Enable local Multicast DNS discovery.\n  mdns: false\n\n  # Sets the max amount of binding requests the agent will send over a candidate pair for validation or nomination.\n  # If after the the configured number, the candidate is yet to answer a binding request or a nomination we set the pair as failed.\n  max_binding_requests: 7\n\n  # SetNAT1To1IPs sets a list of external IP addresses of 1:1 (D)NAT and a candidate type for which the external IP address is used.\n  # This is useful when you are host a server using Pion on an AWS EC2 instance which has a private address, behind a 1:1 DNAT with a public IP (e.g. Elastic IP).\n  # In this case, you can give the public IP address so that Pion will use the public IP address in its candidate instead of the private IP address.\n  # nat_1to1_ips:\n  # - 10.10.2.3\n\n  # Limit the port range used by ICE\n  port_range:\n      # Minimum port for allocation policy for ICE sockets (range: 0-65535)\n      min: 49152\n\n      # Maximum port for allocation policy for ICE sockets (range: 0-65535)\n      max: 65535\n\n  # Interval at which the agent performs candidate checks in the connecting phase\n  check_interval: 200ms\n    \n  # Time until an Agent transitions disconnected.\n  # If the duration is 0, the ICE Agent will never go to disconnected\n  disconnected_timeout: 5s\n\n  # Time until an Agent transitions to failed after disconnected\n  # If the duration is 0, we will never go to failed.\n  failed_timeout: 5s\n\n  # Time to wait before ICE restart\n  restart_timeout: 5s\n\n  # Interval between STUN keepalives (should be less then connection timeout above).\n  # Af the interval is 0, we never send keepalive packets\n  keepalive_interval: 2s\n\n\n## Hook callbacks\n#\n# Hook callback can be used to invoke subprocesses\n# or web-hooks on certain events within cun\u012bcu.\nhooks:\n\n  # An \'exec\' hook spawn a subprocess for each event.\n  - type: exec\n    command: ../../scripts/hook.sh\n  \n    # Prepend additional arguments\n    args: []\n  \n    # Pass JSON object via Stdin to command\n    stdin: true\n  \n    # Set environment variables for invocation\n    env:\n      COLOR: "1"\n  \n  # A \'web\' hook performs HTTP requests for each event.\n  - type: web\n  \n    # URL of the webhook endpoint\n    url: https://my-webhook-endpoint.com/api/v1/webhook\n    \n    # HTTP method of the request\n    method: POST\n  \n    # Additional HTTP headers which are used for the requests\n    headers:\n      User-Agent: ahoi\n      Authorization: Bearer XXXXXX\n\n\n## Interface specific settings / overwrites.\n#\n# Most of the top-level settings of this configuration file can be customized\n# for specific interfaces.\n# \n# The keys of the \'interfaces\' dictionary are glob(7) patterns which will be\n# matched against the interface names.\n# Settings are overlayed in the order in which the keys are provided in the\n# interface map.\n#\n# Keys which are not a glob(8) pattern, will be created as new interfaces if\n# they do not exist already in the system.\ninterfaces:\n  # A simple interface specific setting\n  # cunicu will set the private key of interface \'wg0\' to the provided value.\n  wg0:\n    discover_endpoints: false\n\n  # No settings are overwritten. But since this is not a glob pattern,\n  # A new interface named \'wg1\' will be created if it does not exist yet.\n  # The same applies to the previous interface \'wg0\'\n  wg1: {}\n\n  # Create a new interface using the wireguard-go user-space implementation.\n  wg2:\n    userspace: true\n\n  # This pattern configuration will be applied to all interfaces which match the pattern.\n  # This rule will not create any new interfaces.\n  wg-work-*:\n    community: "mysecret-pass" \n    \n    ice:\n      urls:\n      - turn:mysecret.turn-server.com\n\n  # Multiple patterns are supported and evaluated in the order they a defined in the configuration file.\n  # \n  wg-work-external-*:\n    ice:\n      network_types: [ udp6 ]\n';if(n.section){const e=t.split("\n");let i=[],s=[],a=!1;for(let t of e){let e=!1,o=!1,r=t.startsWith("#"),c=""===t.trim(),d=t.match(/^([a-zA-z]+):/);null!==d&&(e=d[1]==n.section,o=d[1]!=n.section),r&&(a=!1,i.push(t)),e&&(a=!0,s.push(...i),i=[]),o&&(a=!1),c&&(i=[]),a&&s.push(t)}""==s[s.length-1]&&(s=s.slice(0,-1)),t=s.join("\n"),n.title=`Section "${n.section}" of ${n.title}`}return i.createElement(s.Z,n,t)}},4721:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>d,contentTitle:()=>r,default:()=>p,frontMatter:()=>o,metadata:()=>c,toc:()=>l});var i=t(7462),s=(t(7294),t(3905)),a=(t(1839),t(1669));const o={title:"Endpoint Discovery"},r="Endpoint Discovery",c={unversionedId:"features/epdisc",id:"features/epdisc",title:"Endpoint Discovery",description:"The endpoint discovery finds usable WireGuard endpoint addresses for remote peers using Interactive Connectivity Establishment (ICE).",source:"@site/docs/features/epdisc.md",sourceDirName:"features",slug:"/features/epdisc",permalink:"/docs/features/epdisc",draft:!1,editUrl:"https://github.com/stv0g/cunicu/edit/master/docs/features/epdisc.md",tags:[],version:"current",frontMatter:{title:"Endpoint Discovery"},sidebar:"tutorialSidebar",previous:{title:"Config-file Synchronization",permalink:"/docs/features/cfgsync"},next:{title:"Hooks",permalink:"/docs/features/hooks"}},d={},l=[{value:"Configuration",id:"configuration",level:2}],h={toc:l};function p(e){let{components:n,...t}=e;return(0,s.kt)("wrapper",(0,i.Z)({},h,t,{components:n,mdxType:"MDXLayout"}),(0,s.kt)("h1",{id:"endpoint-discovery"},"Endpoint Discovery"),(0,s.kt)("p",null,"The endpoint discovery finds usable WireGuard endpoint addresses for remote peers using ",(0,s.kt)("a",{parentName:"p",href:"https://en.wikipedia.org/wiki/Interactive_Connectivity_Establishment"},"Interactive Connectivity Establishment (ICE)"),"."),(0,s.kt)("h2",{id:"configuration"},"Configuration"),(0,s.kt)(a.Z,{section:"epdisc",mdxType:"ExampleConfig"}))}p.isMDXComponent=!0}}]);