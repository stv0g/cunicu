"use strict";(self.webpackChunkcunicu=self.webpackChunkcunicu||[]).push([[6435],{3905:(e,t,r)=>{r.d(t,{Zo:()=>u,kt:()=>d});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=n.createContext({}),c=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},u=function(e){var t=c(e.components);return n.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,i=e.originalType,l=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),f=c(r),d=o,y=f["".concat(l,".").concat(d)]||f[d]||p[d]||i;return r?n.createElement(y,a(a({ref:t},u),{},{components:r})):n.createElement(y,a({ref:t},u))}));function d(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=r.length,a=new Array(i);a[0]=f;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:o,a[1]=s;for(var c=2;c<i;c++)a[c]=r[c];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},7746:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>p,frontMatter:()=>i,metadata:()=>s,toc:()=>c});var n=r(7462),o=(r(7294),r(3905));const i={},a="Route Synchronization",s={unversionedId:"features/rtsync",id:"features/rtsync",title:"Route Synchronization",description:"The route synchronization feature keeps the kernel routing table in sync with WireGuard's AllowedIPs setting.",source:"@site/docs/features/rtsync.md",sourceDirName:"features",slug:"/features/rtsync",permalink:"/docs/features/rtsync",draft:!1,editUrl:"https://github.com/stv0g/cunicu/edit/master/docs/features/rtsync.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Pre-shared Key Establishment",permalink:"/docs/features/pske"},next:{title:"Usage",permalink:"/docs/usage/"}},l={},c=[],u={toc:c};function p(e){let{components:t,...r}=e;return(0,o.kt)("wrapper",(0,n.Z)({},u,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"route-synchronization"},"Route Synchronization"),(0,o.kt)("p",null,"The route synchronization feature keeps the kernel routing table in sync with WireGuard's ",(0,o.kt)("em",{parentName:"p"},"AllowedIPs")," setting."),(0,o.kt)("p",null,"This synchronization is bi-directional:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Networks with are found in a Peers AllowedIP list will be installed as a kernel route."),(0,o.kt)("li",{parentName:"ul"},"Kernel routes with the peers link-local IP address as next-hop will be added to the Peers ",(0,o.kt)("em",{parentName:"li"},"AllowedIPs")," list.")),(0,o.kt)("p",null,"This rather simple feature allows user to pair cunicu with a software routing daemon like ",(0,o.kt)("a",{parentName:"p",href:"https://bird.network.cz/"},"Bird2")," while using a single WireGuard interface with multiple peer-to-peer links."))}p.isMDXComponent=!0}}]);