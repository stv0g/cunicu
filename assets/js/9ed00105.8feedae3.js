"use strict";(self.webpackChunkwice=self.webpackChunkwice||[]).push([[9004],{9733:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>r,default:()=>d,frontMatter:()=>o,metadata:()=>s,toc:()=>c});var i=t(7462),a=(t(7294),t(3905));t(1839);const o={sidebar_position:7},r="Configuration",s={unversionedId:"configuration",id:"configuration",title:"Configuration",description:"This page describes the ways of configuring the cunicu daemon (cunicu daemon).",source:"@site/docs/configuration.md",sourceDirName:".",slug:"/configuration",permalink:"/docs/configuration",draft:!1,editUrl:"https://github.com/stv0g/cunicu/tree/master/website/docs/configuration.md",tags:[],version:"current",sidebarPosition:7,frontMatter:{sidebar_position:7},sidebar:"tutorialSidebar",previous:{title:"Installation",permalink:"/docs/installation"},next:{title:"Auto-configuration",permalink:"/docs/features/autocfg"}},l={},c=[{value:"Command Line Flags",id:"command-line-flags",level:2},{value:"Configuration File",id:"configuration-file",level:2},{value:"Environment Variables",id:"environment-variables",level:2},{value:"DNS Auto-configuration",id:"dns-auto-configuration",level:2},{value:"Remote Configuration File",id:"remote-configuration-file",level:2}],u={toc:c};function d(e){let{components:n,...t}=e;return(0,a.kt)("wrapper",(0,i.Z)({},u,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"configuration"},"Configuration"),(0,a.kt)("p",null,"This page describes the ways of configuring the cunicu daemon (",(0,a.kt)("inlineCode",{parentName:"p"},"cunicu daemon"),")."),(0,a.kt)("h2",{id:"command-line-flags"},"Command Line Flags"),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"cunicu daemon")," can almost fully be configured by passing command line arguments.\nA full overview is available in its ",(0,a.kt)("a",{parentName:"p",href:"/docs/usage/man/daemon"},"manpage"),"."),(0,a.kt)("h2",{id:"configuration-file"},"Configuration File"),(0,a.kt)("p",null,"Alternatively a configuration file can be used for a persistent configuration:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml",metastring:'title="cunicu.yaml"',title:'"cunicu.yaml"'},'watch_interval: 1s\n\nbackends:\n- grpc://localhost:8080?insecure=true&skip_verify=true\n- k8s:///path/to/your/kubeconfig.yaml?namespace=default\n\n# WireGuard settings\nwireguard:  \n  # Create WireGuard interfaces using bundled wireguard-go Userspace implementation\n  # This will be the default if there is no WireGuard kernel module present.\n  userspace: false\n\n  # Ignore WireGuard interface which do not match this regular expression\n  interface_filter: .*\n\n  # A list of WireGuard interfaces which should be configured\n  interfaces:\n  - wg-vpn\n\n  # Port range for ListenPort setting of newly created WireGuard interfaces\n  # cun\u012bcu will select the first available port in this range.\n  port:\n    min: 52820\n    max: 65535\n\n# Control socket settings\nsocket:\n  path: /var/run/cunicu.sock\n\n  # Start of cun\u012bcu daemon will block until its unblocked via the control socket\n  # Mostly useful for testing automation\n  wait: false\n\n# Synchronize WireGuard interface configurations with wg(8) config-files.\nconfig_sync:\n  enabled: false\n  \n  # Directory where Wireguard configuration files are located.\n  # We expect the same format as used by wg(8) and wg-quick(8).\n  # Filenames must match the interface name with a \'.conf\' suffix.\n  path: /etc/wireguard\n\n  # Watch the configuration files for changes and apply them accordingly.\n  watch: false\n  \n# Synchronize WireGuard AllowedIPs with Kernel routing table\nroute_sync:\n  enabled: true\n\n  table: main\n\n# Discovery of other WireGuard peers\npeer_disc:\n  enabled: true\n\n  # A list of WireGuard public keys which are accepted peers\n  whitelist:\n  - coNsGPwVPdpahc8U+dbbWGzTAdCd6+1BvPIYg10wDCI=\n  - AOZzBaNsoV7P8vo0D5UmuIJUQ7AjMbHbGt2EA8eAuEc=\n\n  # A passphrase shared among all peers of the same community\n  community: "some-common-password"\n\n# Discovery of WireGuard endpoint addressesendpoint_disc:\n  enabled: true\n\n  # Interactive Connectivity Establishment parameters\n  ice:\n    # A list of STUN and TURN servers used by ICE\n    urls:\n    - stun:stun.l.google.com:19302\n\n    # Credentials for STUN/TURN servers configured above\n    username: ""\n    password: ""\n\n    # Allow connections to STUNS/TURNS servers for which\n    # we cant validate their TLS certificates\n    insecure_skip_verify: false\n\n    # Limit available network and candidate types\n    network_types: [udp4, udp6, tcp4, tcp6]\n    candidate_types: [host, srflx, prflx ,relay]\n\n    # Regular expression whitelist of interfaces which are used to gather ICE candidates.\n    interface_filter: .*\n\n    # Lite agents do not perform connectivity check and only provide host candidates.\n    lite: false\n\n    # Attempt to find candidates via mDNS discovery\n    mdns: false\n\n    # Sets the max amount of binding requests the agent will send over a candidate pair for validation or nomination.\n    # If after the the configured number, the candidate is yet to answer a binding request or a nomination we set the pair as failed.\n    max_binding_requests: 7\n\n    # SetNAT1To1IPs sets a list of external IP addresses of 1:1 (D)NAT and a candidate type for which the external IP address is used.\n    # This is useful when you are host a server using Pion on an AWS EC2 instance which has a private address, behind a 1:1 DNAT with a public IP (e.g. Elastic IP).\n    # In this case, you can give the public IP address so that Pion will use the public IP address in its candidate instead of the private IP address.\n    nat_1to1_ips: []\n\n    # Limit the port range used by ICE\n    port:\n        min: 49152\n        max: 65535\n\n    # The check interval controls how often our task loop runs when in the connecting state.\n    check_interval: 200ms\n    \n    # If the duration is 0, the ICE Agent will never go to disconnected\n    disconnected_timeout: 5s\n\n    # If the duration is 0, we will never go to failed.\n    failed_timeout: 5s\n    restart_timeout: 5s\n\n    # Determines how often should we send ICE keepalives (should be less then connection timeout above).\n    # A keepalive interval of 0 means we never send keepalive packets\n    keepalive_interval: 2s\n')),(0,a.kt)("h2",{id:"environment-variables"},"Environment Variables"),(0,a.kt)("p",null,"All the settings from the configuration file can also be passed via environment variables by following the following rules:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Convert the setting name to uppercase"),(0,a.kt)("li",{parentName:"ul"},"Prefixing the setting name with ",(0,a.kt)("inlineCode",{parentName:"li"},"CUNICU_")),(0,a.kt)("li",{parentName:"ul"},"Nested settings are separated by underscores")),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Example:")," The setting ",(0,a.kt)("inlineCode",{parentName:"p"},"endpoint_disc.ice.max_binding_requests")," can be set by the environment variable ",(0,a.kt)("inlineCode",{parentName:"p"},"CUNICU_ENDPOINT_DISC_ICE_MAX_BINDING_REQUESTS")),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Note:")," Setting lists such as ",(0,a.kt)("inlineCode",{parentName:"p"},"endpoint_disc.ice.urls")," or ",(0,a.kt)("inlineCode",{parentName:"p"},"backends")," can currently not be set via environment variables."),(0,a.kt)("h2",{id:"dns-auto-configuration"},"DNS Auto-configuration"),(0,a.kt)("p",null,"cun\u012bcu als supports retrieving parts of the configuration via DNS lookups."),(0,a.kt)("p",null,"When ",(0,a.kt)("inlineCode",{parentName:"p"},"cunicu daemon")," is started with a ",(0,a.kt)("inlineCode",{parentName:"p"},"--domain example.com")," parameter it will look for the following DNS records to obtain its configuration."),(0,a.kt)("p",null,"STUN and TURN servers used for ICE are retrieved by SVR lookups and other cun\u012bcu settings are retrieved via TXT lookups: "),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},'_stun._udp.example.com.  3600 IN SRV 10 0 3478 stun.example.com.\n_stuns._tcp.example.com. 3600 IN SRV 10 0 3478 stun.example.com.\n_turn._udp.example.com.  3600 IN SRV 10 0 3478 turn.example.com.\n_turn._tcp.example.com.  3600 IN SRV 10 0 3478 turn.example.com.\n_turns._tcp.example.com. 3600 IN SRV 10 0 5349 turn.example.com.\n\nexample.com.             3600 IN TXT "cunicu-backend=p2p"\nexample.com.             3600 IN TXT "cunicu-peer-disc-community=my-community-password"\nexample.com.             3600 IN TXT "cunicu-endpoint-disc-ice-username=user1"\nexample.com.             3600 IN TXT "cunicu-endpoint-disc-ice-password=pass1"\nexample.com.             3600 IN TXT "cunicu-config=https://example.com/cunicu.yaml"\n')),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Note:")," The ",(0,a.kt)("inlineCode",{parentName:"p"},"cunicu-backend")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"cunicu-config")," TXT records can be provided multiple times. Others not."),(0,a.kt)("h2",{id:"remote-configuration-file"},"Remote Configuration File"),(0,a.kt)("p",null,"When ",(0,a.kt)("inlineCode",{parentName:"p"},"cunicu daemon")," can be started with ",(0,a.kt)("inlineCode",{parentName:"p"},"--config")," options pointing to HTTPS URIs.\ncun\u012bcu will download all configuration files in the order they are specified on the command line and merge them subsequently."),(0,a.kt)("p",null,"This feature can be combined with the DNS auto-configuration method by providing a TXT record pointing to the configuration file:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},'example.com.             3600 IN TXT "cunicu-config=https://example.com/cunicu.yaml"\n')),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Note:")," Remote configuration files must be fetched via HTTPS if they are not hosted locally and required a trusted server certificate."))}d.isMDXComponent=!0}}]);