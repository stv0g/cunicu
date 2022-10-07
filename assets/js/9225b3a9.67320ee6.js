"use strict";(self.webpackChunkwice=self.webpackChunkwice||[]).push([[4446],{5142:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>r,default:()=>m,frontMatter:()=>o,metadata:()=>l,toc:()=>u});var a=t(7462),i=(t(7294),t(3905));t(1839);const o={sidebar_position:7},r="Configuration",l={unversionedId:"config",id:"config",title:"Configuration",description:"This page describes the ways of configuring the cun\u012bcu daemon (cunicu daemon).",source:"@site/docs/config.md",sourceDirName:".",slug:"/config",permalink:"/docs/config",draft:!1,editUrl:"https://github.com/stv0g/cunicu/edit/master/docs/config.md",tags:[],version:"current",sidebarPosition:7,frontMatter:{sidebar_position:7},sidebar:"tutorialSidebar",previous:{title:"Installation",permalink:"/docs/install"},next:{title:"Configuration Reference",permalink:"/docs/config-reference"}},c={},u=[{value:"Command Line Flags",id:"command-line-flags",level:2},{value:"Configuration File",id:"configuration-file",level:2},{value:"Environment Variables",id:"environment-variables",level:2},{value:"At Runtime",id:"at-runtime",level:2},{value:"DNS Auto-configuration",id:"dns-auto-configuration",level:2},{value:"Remote Configuration Files",id:"remote-configuration-files",level:2},{value:"Auto-reload",id:"auto-reload",level:2}],s={toc:u};function m(e){let{components:n,...t}=e;return(0,i.kt)("wrapper",(0,a.Z)({},s,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"configuration"},"Configuration"),(0,i.kt)("p",null,"This page describes the ways of configuring the cun\u012bcu daemon (",(0,i.kt)("inlineCode",{parentName:"p"},"cunicu daemon"),")."),(0,i.kt)("h2",{id:"command-line-flags"},"Command Line Flags"),(0,i.kt)("p",null,"Basic options of ",(0,i.kt)("inlineCode",{parentName:"p"},"cunicu daemon")," can be configured by passing command line arguments.\nA full overview is available in its ",(0,i.kt)("a",{parentName:"p",href:"/docs/usage/man/daemon"},"manpage"),"."),(0,i.kt)("h2",{id:"configuration-file"},"Configuration File"),(0,i.kt)("p",null,"For more advanced setups, a configuration file can be used for a persistent configuration:"),(0,i.kt)("p",null,"Please have a look at the ",(0,i.kt)("a",{parentName:"p",href:"/docs/config-reference"},"example configuration file")," for a full reference of all available settings."),(0,i.kt)("h2",{id:"environment-variables"},"Environment Variables"),(0,i.kt)("p",null,"All the settings from the configuration file can also be passed via environment variables by following the following rules:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Convert the setting name to uppercase"),(0,i.kt)("li",{parentName:"ul"},"Prefixing the setting name with ",(0,i.kt)("inlineCode",{parentName:"li"},"CUNICU_")),(0,i.kt)("li",{parentName:"ul"},"Nested settings are separated by underscores")),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Example:")," The setting ",(0,i.kt)("inlineCode",{parentName:"p"},"ice.max_binding_requests")," can be set by the environment variable ",(0,i.kt)("inlineCode",{parentName:"p"},"CUNICU_ICE_MAX_BINDING_REQUESTS")),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("p",{parentName:"admonition"},"Setting lists such as ",(0,i.kt)("inlineCode",{parentName:"p"},"ice.urls")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"backends")," can currently not be set via environment variables.")),(0,i.kt)("h2",{id:"at-runtime"},"At Runtime"),(0,i.kt)("p",null,"cun\u012bcu's configuration can also be updated at runtime, elevating the need to restart the daemon to avoid interruption of connectivity."),(0,i.kt)("p",null,"Please have a look at the ",(0,i.kt)("a",{parentName:"p",href:"/docs/usage/man/config"},(0,i.kt)("inlineCode",{parentName:"a"},"cunicu config"))," commands."),(0,i.kt)("h2",{id:"dns-auto-configuration"},"DNS Auto-configuration"),(0,i.kt)("p",null,"cun\u012bcu als supports retrieving parts of the configuration via DNS lookups.\nThis is useful for corporate environments in which a fleet of cun\u012bcu daemon need to be configured centrally."),(0,i.kt)("p",null,"In this case ",(0,i.kt)("inlineCode",{parentName:"p"},"cunicu daemon")," is started one or more ",(0,i.kt)("inlineCode",{parentName:"p"},"--domain example.com")," parameters to look for the following DNS records to obtain its configuration."),(0,i.kt)("p",null,"STUN and TURN servers used for ICE are retrieved by SVR lookups and other cun\u012bcu settings are retrieved via ",(0,i.kt)("inlineCode",{parentName:"p"},"SRV")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"TXT")," lookups: "),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'_stun._udp.example.com.  3600 IN SRV 10 0 3478 stun.example.com.\n_stuns._tcp.example.com. 3600 IN SRV 10 0 3478 stun.example.com.\n_turn._udp.example.com.  3600 IN SRV 10 0 3478 turn.example.com.\n_turn._tcp.example.com.  3600 IN SRV 10 0 3478 turn.example.com.\n_turns._tcp.example.com. 3600 IN SRV 10 0 5349 turn.example.com.\n\nexample.com.             3600 IN TXT "cunicu-config=https://example.com/cunicu.yaml"\nexample.com.             3600 IN TXT "cunicu-backend=grpc://signal.example.com:443"\nexample.com.             3600 IN TXT "cunicu-community=my-community-password"\nexample.com.             3600 IN TXT "cunicu-ice-username=user1"\nexample.com.             3600 IN TXT "cunicu-ice-password=pass1"\n')),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("p",{parentName:"admonition"},"The ",(0,i.kt)("inlineCode",{parentName:"p"},"cunicu-backend")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"cunicu-config")," TXT records can be provided multiple times. Others not.")),(0,i.kt)("h2",{id:"remote-configuration-files"},"Remote Configuration Files"),(0,i.kt)("p",null,"When ",(0,i.kt)("inlineCode",{parentName:"p"},"cunicu daemon")," can be started with ",(0,i.kt)("inlineCode",{parentName:"p"},"--config")," options pointing to HTTPS URIs:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"cunicu daemon --config http://example.com/cunicu.yaml\n")),(0,i.kt)("p",null,"cun\u012bcu will download all configuration files in the order they are specified on the command line and merge them subsequently."),(0,i.kt)("p",null,"This feature can be combined with the DNS auto-configuration method by providing a TXT record pointing to the configuration file:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'example.com. 3600 IN TXT "cunicu-config=https://example.com/cunicu.yaml"\n')),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("p",{parentName:"admonition"},"Remote configuration files must be fetched via HTTPS if they are not hosted locally and required a trusted server certificate.")),(0,i.kt)("h2",{id:"auto-reload"},"Auto-reload"),(0,i.kt)("p",null,"cun\u012bcu watches local and remote files as well as the DNS configuration for changes and automatically reloads its configuration from them whenever a change has been detected."),(0,i.kt)("p",null,"For local files the change is detected by ",(0,i.kt)("a",{parentName:"p",href:"https://man7.org/linux/man-pages/man7/inotify.7.html"},"inotify(7)"),".\nFor remote sources, cun\u012bcu periodically checks the ",(0,i.kt)("inlineCode",{parentName:"p"},"Last-Modified")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"Etag")," headers in case of HTTP files or the DNS zone's ",(0,i.kt)("a",{parentName:"p",href:"https://en.wikipedia.org/wiki/SOA_record#Structure"},"SOA serial number")," to detect changes without request the full remote source."),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("p",{parentName:"admonition"},"Configuration file distributed via ",(0,i.kt)("inlineCode",{parentName:"p"},"conicu-config")," DNS TXT record are not yet monitored for changes.")))}m.isMDXComponent=!0}}]);