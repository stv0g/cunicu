"use strict";(self.webpackChunkwice=self.webpackChunkwice||[]).push([[5734],{5558:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>o,contentTitle:()=>i,default:()=>c,frontMatter:()=>d,metadata:()=>r,toc:()=>u});var n=a(7462),t=(a(7294),a(3905));a(1839);const d={title:"cunicu addresses",sidebar_label:"addresses",sidebar_class_name:"command-name",slug:"/usage/man/addresses",hide_title:!0,keywords:["manpage"]},i=void 0,r={unversionedId:"usage/md/cunicu_addresses",id:"usage/md/cunicu_addresses",title:"cunicu addresses",description:"cunicu addresses",source:"@site/docs/usage/md/cunicu_addresses.md",sourceDirName:"usage/md",slug:"/usage/man/addresses",permalink:"/docs/usage/man/addresses",draft:!1,editUrl:"https://github.com/stv0g/cunicu/edit/master/docs/usage/md/cunicu_addresses.md",tags:[],version:"current",frontMatter:{title:"cunicu addresses",sidebar_label:"addresses",sidebar_class_name:"command-name",slug:"/usage/man/addresses",hide_title:!0,keywords:["manpage"]},sidebar:"tutorialSidebar",previous:{title:"cunicu",permalink:"/docs/usage/man/"},next:{title:"completion",permalink:"/docs/usage/man/completion"}},o={},u=[{value:"cunicu addresses",id:"cunicu-addresses",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Examples",id:"examples",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],l={toc:u};function c(e){let{components:s,...a}=e;return(0,t.kt)("wrapper",(0,n.Z)({},l,a,{components:s,mdxType:"MDXLayout"}),(0,t.kt)("h2",{id:"cunicu-addresses"},"cunicu addresses"),(0,t.kt)("p",null,"Calculate link-local IPv4 and IPv6 addresses from a WireGuard X25519 public key"),(0,t.kt)("h3",{id:"synopsis"},"Synopsis"),(0,t.kt)("p",null,"cun\u012bcu auto-configuration feature derives and assigns link-local IPv4 and IPv6 addresses based on the public key of the WireGuard interface.\nThis sub-command accepts a WireGuard public key on the standard input and prints out the calculated IP addresses on the standard output."),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"cunicu addresses [flags]\n")),(0,t.kt)("h3",{id:"examples"},"Examples"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"$ wg genkey | wg pubkey | cunicu addresses\nfc2f:9a4d:777f:7a97:8197:4a5d:1d1b:ed79\n10.237.119.127\n")),(0,t.kt)("h3",{id:"options"},"Options"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},"  -h, --help   help for addresses\n  -m, --mask   Print CIDR mask\n")),(0,t.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,t.kt)("pre",null,(0,t.kt)("code",{parentName:"pre"},'  -q, --color string       Enable colorization of output (one of: auto, always, never) (default "auto")\n  -l, --log-file string    path of a file to write logs to\n  -d, --log-level string   log level (one of: debug, info, warn, error, dpanic, panic, and fatal) (default "info")\n  -v, --verbose int        verbosity level\n')),(0,t.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,t.kt)("ul",null,(0,t.kt)("li",{parentName:"ul"},(0,t.kt)("a",{parentName:"li",href:"/docs/usage/man/"},"cunicu"),"\t - cun\u012bcu is a user-space daemon managing WireGuard\xae interfaces to establish peer-to-peer connections in harsh network environments.")))}c.isMDXComponent=!0}}]);