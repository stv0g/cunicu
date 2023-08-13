"use strict";(self.webpackChunkcunicu=self.webpackChunkcunicu||[]).push([[8671],{3905:(e,n,t)=>{t.d(n,{Zo:()=>c,kt:()=>g});var r=t(67294);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function u(e,n){if(null==e)return{};var t,r,a=function(e,n){if(null==e)return{};var t,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var l=r.createContext({}),p=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},c=function(e){var n=p(e.components);return r.createElement(l.Provider,{value:n},e.children)},s="mdxType",d={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},m=r.forwardRef((function(e,n){var t=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,c=u(e,["components","mdxType","originalType","parentName"]),s=p(t),m=a,g=s["".concat(l,".").concat(m)]||s[m]||d[m]||i;return t?r.createElement(g,o(o({ref:n},c),{},{components:t})):r.createElement(g,o({ref:n},c))}));function g(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var i=t.length,o=new Array(i);o[0]=m;var u={};for(var l in n)hasOwnProperty.call(n,l)&&(u[l]=n[l]);u.originalType=e,u[s]="string"==typeof e?e:a,o[1]=u;for(var p=2;p<i;p++)o[p]=t[p];return r.createElement.apply(null,o)}return r.createElement.apply(null,t)}m.displayName="MDXCreateElement"},52940:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>o,default:()=>d,frontMatter:()=>i,metadata:()=>u,toc:()=>p});var r=t(87462),a=(t(67294),t(3905));const i={title:"cunicu wg pubkey",sidebar_label:"wg pubkey",sidebar_class_name:"command-name",slug:"/usage/man/wg/pubkey",hide_title:!0,keywords:["manpage"]},o=void 0,u={unversionedId:"usage/md/cunicu_wg_pubkey",id:"usage/md/cunicu_wg_pubkey",title:"cunicu wg pubkey",description:"cunicu wg pubkey",source:"@site/docs/usage/md/cunicu_wg_pubkey.md",sourceDirName:"usage/md",slug:"/usage/man/wg/pubkey",permalink:"/docs/usage/man/wg/pubkey",draft:!1,editUrl:"https://github.com/cunicu/cunicu/edit/master/docs/usage/md/cunicu_wg_pubkey.md",tags:[],version:"current",frontMatter:{title:"cunicu wg pubkey",sidebar_label:"wg pubkey",sidebar_class_name:"command-name",slug:"/usage/man/wg/pubkey",hide_title:!0,keywords:["manpage"]},sidebar:"tutorialSidebar",previous:{title:"wg genpsk",permalink:"/docs/usage/man/wg/genpsk"},next:{title:"wg show",permalink:"/docs/usage/man/wg/show"}},l={},p=[{value:"cunicu wg pubkey",id:"cunicu-wg-pubkey",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Examples",id:"examples",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],c={toc:p},s="wrapper";function d(e){let{components:n,...t}=e;return(0,a.kt)(s,(0,r.Z)({},c,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h2",{id:"cunicu-wg-pubkey"},"cunicu wg pubkey"),(0,a.kt)("p",null,"Calculates a public key and prints it in base64 to standard output."),(0,a.kt)("h3",{id:"synopsis"},"Synopsis"),(0,a.kt)("p",null,"Calculates a public key and prints it in base64 to standard output from a corresponding private key (generated with genkey) given in base64 on standard input."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"cunicu wg pubkey [flags]\n")),(0,a.kt)("h3",{id:"examples"},"Examples"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"# A private key and a corresponding public key may be generated at once by calling:\n$ umask 077\n$ wg genkey | tee private.key | wg pubkey > public.key\n")),(0,a.kt)("h3",{id:"options"},"Options"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"  -h, --help   help for pubkey\n")),(0,a.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'  -q, --color string            Enable colorization of output (one of: auto, always, never) (default "auto")\n  -l, --log-file stringArray    path of a file to write logs to\n  -d, --log-level stringArray   log level filter rule (one of: debug, info, warn, error, dpanic, panic, and fatal) (default "info")\n')),(0,a.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/docs/usage/man/wg"},"cunicu wg"),"\t - WireGuard commands")))}d.isMDXComponent=!0}}]);