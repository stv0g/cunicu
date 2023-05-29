"use strict";(self.webpackChunkcunicu=self.webpackChunkcunicu||[]).push([[7577],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>d});var o=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,o,r=function(e,t){if(null==e)return{};var n,o,r={},i=Object.keys(e);for(o=0;o<i.length;o++)n=i[o],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(o=0;o<i.length;o++)n=i[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=o.createContext({}),c=function(e){var t=o.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},p=function(e){var t=c(e.components);return o.createElement(s.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},m=o.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,s=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),m=c(n),d=r,g=m["".concat(s,".").concat(d)]||m[d]||u[d]||i;return n?o.createElement(g,a(a({ref:t},p),{},{components:n})):o.createElement(g,a({ref:t},p))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,a=new Array(i);a[0]=m;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:r,a[1]=l;for(var c=2;c<i;c++)a[c]=n[c];return o.createElement.apply(null,a)}return o.createElement.apply(null,n)}m.displayName="MDXCreateElement"},6141:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>a,default:()=>u,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var o=n(7462),r=(n(7294),n(3905));const i={sidebar_position:199},a="Development",l={unversionedId:"development/index",id:"development/index",title:"Development",description:"cun\u012bcu is written almost completely in Go and heavily relies on awesome tooling and packages for Golang:",source:"@site/docs/development/index.md",sourceDirName:"development",slug:"/development/",permalink:"/docs/development/",draft:!1,editUrl:"https://github.com/stv0g/cunicu/edit/master/docs/development/index.md",tags:[],version:"current",sidebarPosition:199,frontMatter:{sidebar_position:199},sidebar:"tutorialSidebar",previous:{title:"Comparison",permalink:"/docs/comparison"},next:{title:"Proxying",permalink:"/docs/development/proxying"}},s={},c=[{value:"Testing",id:"testing",level:2}],p={toc:c};function u(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,o.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"development"},"Development"),(0,r.kt)("p",null,"cun\u012bcu is written almost completely in ",(0,r.kt)("a",{parentName:"p",href:"https://go.dev/"},"Go")," and heavily relies on awesome tooling and packages for Golang:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://goreleaser.com/"},"GoReleaser")," for release automation"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://onsi.github.io/ginkgo"},"Ginkgo")," and ",(0,r.kt)("a",{parentName:"li",href:"https://onsi.github.io/gomega"},"Gomega")," for testing"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://github.com/pion"},"Pion")," for its ICE, STUN, TURN implementation"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://github.com/stv0g/gont"},"Gont")," for end-to-end testing in various network topologies")),(0,r.kt)("p",null,"Furthermore use the following services to manage our development:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://github.com/stv0g/cunicu"},"GitHub")," for source code management and CI pipelines"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://app.codacy.com/gh/stv0g/cunicu/dashboard"},"Codacy")," for code quality checks"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://app.codecov.io/gh/stv0g/cunicu"},"Codecov")," for code coverage analysis")),(0,r.kt)("h2",{id:"testing"},"Testing"),(0,r.kt)("p",null,"We aim to maintain a test coverage above 80% of the lines of code.\nPlease make sure that your merge requests include tests which do not lower the coverage percentage."),(0,r.kt)("p",null,"cun\u012bcu's code-base is tested using the Ginkgo/Gomega testing framework.\nUnit tests can be found alongside the code in files with a ",(0,r.kt)("inlineCode",{parentName:"p"},"_test.go")," suffix.\nEnd-to-end (e2e) integration tests can be found in the ",(0,r.kt)("inlineCode",{parentName:"p"},"test/e2e")," directory."),(0,r.kt)("p",null,"The e2e tests use Gont to construct virtual network environment using Linux's ",(0,r.kt)("inlineCode",{parentName:"p"},"net")," namespaces and ",(0,r.kt)("inlineCode",{parentName:"p"},"veth")," point-to-point links.\nThis allows us to test cun\u012bcu in both simple and complex network topologies including, L2 switches, L3 routers, firewalls and NAT boxes."))}u.isMDXComponent=!0}}]);