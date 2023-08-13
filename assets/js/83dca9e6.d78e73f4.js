"use strict";(self.webpackChunkcunicu=self.webpackChunkcunicu||[]).push([[822],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>f});var o=n(67294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,o,r=function(e,t){if(null==e)return{};var n,o,r={},a=Object.keys(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var c=o.createContext({}),s=function(e){var t=o.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=s(e.components);return o.createElement(c.Provider,{value:t},e.children)},p="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},d=o.forwardRef((function(e,t){var n=e.components,r=e.mdxType,a=e.originalType,c=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),p=s(n),d=r,f=p["".concat(c,".").concat(d)]||p[d]||m[d]||a;return n?o.createElement(f,i(i({ref:t},u),{},{components:n})):o.createElement(f,i({ref:t},u))}));function f(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var a=n.length,i=new Array(a);i[0]=d;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l[p]="string"==typeof e?e:r,i[1]=l;for(var s=2;s<a;s++)i[s]=n[s];return o.createElement.apply(null,i)}return o.createElement.apply(null,n)}d.displayName="MDXCreateElement"},55149:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>i,default:()=>m,frontMatter:()=>a,metadata:()=>l,toc:()=>s});var o=n(87462),r=(n(67294),n(3905));const a={title:"cunicu completion bash",sidebar_label:"completion bash",sidebar_class_name:"command-name",slug:"/usage/man/completion/bash",hide_title:!0,keywords:["manpage"]},i=void 0,l={unversionedId:"usage/md/cunicu_completion_bash",id:"usage/md/cunicu_completion_bash",title:"cunicu completion bash",description:"cunicu completion bash",source:"@site/docs/usage/md/cunicu_completion_bash.md",sourceDirName:"usage/md",slug:"/usage/man/completion/bash",permalink:"/docs/usage/man/completion/bash",draft:!1,editUrl:"https://github.com/cunicu/cunicu/edit/master/docs/usage/md/cunicu_completion_bash.md",tags:[],version:"current",frontMatter:{title:"cunicu completion bash",sidebar_label:"completion bash",sidebar_class_name:"command-name",slug:"/usage/man/completion/bash",hide_title:!0,keywords:["manpage"]},sidebar:"tutorialSidebar",previous:{title:"completion",permalink:"/docs/usage/man/completion"},next:{title:"completion fish",permalink:"/docs/usage/man/completion/fish"}},c={},s=[{value:"cunicu completion bash",id:"cunicu-completion-bash",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Linux:",id:"linux",level:4},{value:"macOS:",id:"macos",level:4},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],u={toc:s},p="wrapper";function m(e){let{components:t,...n}=e;return(0,r.kt)(p,(0,o.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h2",{id:"cunicu-completion-bash"},"cunicu completion bash"),(0,r.kt)("p",null,"Generate the autocompletion script for bash"),(0,r.kt)("h3",{id:"synopsis"},"Synopsis"),(0,r.kt)("p",null,"Generate the autocompletion script for the bash shell."),(0,r.kt)("p",null,"This script depends on the 'bash-completion' package.\nIf it is not installed already, you can install it via your OS's package manager."),(0,r.kt)("p",null,"To load completions in your current shell session:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"source <(cunicu completion bash)\n")),(0,r.kt)("p",null,"To load completions for every new session, execute once:"),(0,r.kt)("h4",{id:"linux"},"Linux:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"cunicu completion bash > /etc/bash_completion.d/cunicu\n")),(0,r.kt)("h4",{id:"macos"},"macOS:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"cunicu completion bash > $(brew --prefix)/etc/bash_completion.d/cunicu\n")),(0,r.kt)("p",null,"You will need to start a new shell for this setup to take effect."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"cunicu completion bash\n")),(0,r.kt)("h3",{id:"options"},"Options"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},"  -h, --help              help for bash\n      --no-descriptions   disable completion descriptions\n")),(0,r.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre"},'  -q, --color string            Enable colorization of output (one of: auto, always, never) (default "auto")\n  -l, --log-file stringArray    path of a file to write logs to\n  -d, --log-level stringArray   log level filter rule (one of: debug, info, warn, error, dpanic, panic, and fatal) (default "info")\n')),(0,r.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"/docs/usage/man/completion"},"cunicu completion"),"\t - Generate the autocompletion script for the specified shell")))}m.isMDXComponent=!0}}]);