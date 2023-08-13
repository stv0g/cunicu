"use strict";(self.webpackChunkcunicu=self.webpackChunkcunicu||[]).push([[1503],{3905:(e,n,t)=>{t.d(n,{Zo:()=>u,kt:()=>m});var r=t(67294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function a(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?a(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function s(e,n){if(null==e)return{};var t,r,i=function(e,n){if(null==e)return{};var t,r,i={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var l=r.createContext({}),c=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},u=function(e){var n=c(e.components);return r.createElement(l.Provider,{value:n},e.children)},p="mdxType",d={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},f=r.forwardRef((function(e,n){var t=e.components,i=e.mdxType,a=e.originalType,l=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),p=c(t),f=i,m=p["".concat(l,".").concat(f)]||p[f]||d[f]||a;return t?r.createElement(m,o(o({ref:n},u),{},{components:t})):r.createElement(m,o({ref:n},u))}));function m(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var a=t.length,o=new Array(a);o[0]=f;var s={};for(var l in n)hasOwnProperty.call(n,l)&&(s[l]=n[l]);s.originalType=e,s[p]="string"==typeof e?e:i,o[1]=s;for(var c=2;c<a;c++)o[c]=t[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,t)}f.displayName="MDXCreateElement"},58548:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>o,default:()=>d,frontMatter:()=>a,metadata:()=>s,toc:()=>c});var r=t(87462),i=(t(67294),t(3905));const a={title:"cunicu wg show",sidebar_label:"wg show",sidebar_class_name:"command-name",slug:"/usage/man/wg/show",hide_title:!0,keywords:["manpage"]},o=void 0,s={unversionedId:"usage/md/cunicu_wg_show",id:"usage/md/cunicu_wg_show",title:"cunicu wg show",description:"cunicu wg show",source:"@site/docs/usage/md/cunicu_wg_show.md",sourceDirName:"usage/md",slug:"/usage/man/wg/show",permalink:"/docs/usage/man/wg/show",draft:!1,editUrl:"https://github.com/cunicu/cunicu/edit/master/docs/usage/md/cunicu_wg_show.md",tags:[],version:"current",frontMatter:{title:"cunicu wg show",sidebar_label:"wg show",sidebar_class_name:"command-name",slug:"/usage/man/wg/show",hide_title:!0,keywords:["manpage"]},sidebar:"tutorialSidebar",previous:{title:"wg pubkey",permalink:"/docs/usage/man/wg/pubkey"},next:{title:"wg showconf",permalink:"/docs/usage/man/wg/showconf"}},l={},c=[{value:"cunicu wg show",id:"cunicu-wg-show",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],u={toc:c},p="wrapper";function d(e){let{components:n,...t}=e;return(0,i.kt)(p,(0,r.Z)({},u,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"cunicu-wg-show"},"cunicu wg show"),(0,i.kt)("p",null,"Shows current WireGuard configuration and runtime information of specified ","[interface]","."),(0,i.kt)("h3",{id:"synopsis"},"Synopsis"),(0,i.kt)("p",null,"Shows current WireGuard configuration and runtime information of specified ","[interface]","."),(0,i.kt)("p",null,"If no ","[interface]"," is specified, ","[interface]"," defaults to 'all'."),(0,i.kt)("p",null,"If 'interfaces' is specified, prints a list of all WireGuard interfaces, one per line, and quits."),(0,i.kt)("p",null,"If no options are given after the interface specification, then prints a list of all attributes in a visually pleasing way meant for the terminal.\nOtherwise, prints specified information grouped by newlines and tabs, meant to be used in scripts."),(0,i.kt)("p",null,"For this script-friendly display, if 'all' is specified, then the first field for all categories of information is the interface name."),(0,i.kt)("p",null,"If 'dump' is specified, then several lines are printed; the first contains in order separated by tab: private-key, public-key, listen-port, fwmark.\nSubsequent lines are printed for each peer and contain in order separated by tab: public-key, preshared-key, endpoint, allowed-ips, latest-handshake, transfer-rx, transfer-tx, persistent-keepalive."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"cunicu wg show { interface-name | all | interfaces } [{ public-key | private-key | listen-port | fwmark | peers | preshared-keys | endpoints | allowed-ips | latest-handshakes | transfer | persistent-keepalive | dump }] [flags]\n")),(0,i.kt)("h3",{id:"options"},"Options"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},'  -h, --help                help for show\n  -s, --rpc-socket string   Unix control and monitoring socket (default "/var/run/cunicu.sock")\n')),(0,i.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},'  -q, --color string            Enable colorization of output (one of: auto, always, never) (default "auto")\n  -l, --log-file stringArray    path of a file to write logs to\n  -d, --log-level stringArray   log level filter rule (one of: debug, info, warn, error, dpanic, panic, and fatal) (default "info")\n')),(0,i.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/docs/usage/man/wg"},"cunicu wg"),"\t - WireGuard commands")))}d.isMDXComponent=!0}}]);