"use strict";(self.webpackChunkcunicu=self.webpackChunkcunicu||[]).push([[7943],{3905:(e,t,n)=>{n.d(t,{Zo:()=>s,kt:()=>m});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var u=a.createContext({}),c=function(e){var t=a.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},s=function(e){var t=c(e.components);return a.createElement(u.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,u=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),d=c(n),m=r,h=d["".concat(u,".").concat(m)]||d[m]||p[m]||i;return n?a.createElement(h,o(o({ref:t},s),{},{components:n})):a.createElement(h,o({ref:t},s))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,o=new Array(i);o[0]=d;var l={};for(var u in t)hasOwnProperty.call(t,u)&&(l[u]=t[u]);l.originalType=e,l.mdxType="string"==typeof e?e:r,o[1]=l;for(var c=2;c<i;c++)o[c]=n[c];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}d.displayName="MDXCreateElement"},7520:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>o,default:()=>p,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var a=n(7462),r=(n(7294),n(3905));const i={sidebar_position:5},o="Installation",l={unversionedId:"install",id:"install",title:"Installation",description:"This guide shows how to install cun\u012bcu.",source:"@site/docs/install.md",sourceDirName:".",slug:"/install",permalink:"/docs/install",draft:!1,editUrl:"https://github.com/stv0g/cunicu/edit/master/docs/install.md",tags:[],version:"current",sidebarPosition:5,frontMatter:{sidebar_position:5},sidebar:"tutorialSidebar",previous:{title:"Welcome",permalink:"/docs/"},next:{title:"Configuration",permalink:"/docs/config"}},u={},c=[{value:"From the Binary Releases",id:"from-the-binary-releases",level:2},{value:"By Hand",id:"by-hand",level:2},{value:"From Script",id:"from-script",level:2},{value:"Through Package Managers",id:"through-package-managers",level:2},{value:"From Apt (Debian, Ubuntu)",id:"from-apt-debian-ubuntu",level:3},{value:"From Yum (Redhat, Fedora, RockyLinux)",id:"from-yum-redhat-fedora-rockylinux",level:3},{value:"From Homebrew (macOS)",id:"from-homebrew-macos",level:3},{value:"From Archlinux User Repository (AUR)",id:"from-archlinux-user-repository-aur",level:3},{value:"From Source (all)",id:"from-source-all",level:2},{value:"Conclusion",id:"conclusion",level:2}],s={toc:c};function p(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,a.Z)({},s,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"installation"},"Installation"),(0,r.kt)("p",null,"This guide shows how to install cun\u012bcu.\ncun\u012bcu can be installed either from source, or from pre-built binary releases."),(0,r.kt)("h2",{id:"from-the-binary-releases"},"From the Binary Releases"),(0,r.kt)("p",null,"Every release of cun\u012bcu provides binary releases for a variety of OSes.\nThese binary versions can be manually downloaded and installed."),(0,r.kt)("h2",{id:"by-hand"},"By Hand"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("a",{parentName:"li",href:"https://github.com/stv0g/cunicu/releases"},"Download your desired version")),(0,r.kt)("li",{parentName:"ol"},"Unzip it: ",(0,r.kt)("inlineCode",{parentName:"li"},"gunzip cunicu_0.0.1_linux_amd64.tar.gz")),(0,r.kt)("li",{parentName:"ol"},"Move the unzipped binary to its desired destination: ",(0,r.kt)("inlineCode",{parentName:"li"},"mv cunicu /usr/local/bin/cunicu")),(0,r.kt)("li",{parentName:"ol"},"Make it executable: ",(0,r.kt)("inlineCode",{parentName:"li"},"chmod +x /usr/local/bin/cunicu")),(0,r.kt)("li",{parentName:"ol"},"From there, you should be able to run the client and add the stable repo: ",(0,r.kt)("inlineCode",{parentName:"li"},"cunicu help"),".")),(0,r.kt)("admonition",{type:"note"},(0,r.kt)("p",{parentName:"admonition"},"cun\u012bcu automated tests are performed for Linux, macOS and Windows on x86_64, ARMv6, ARMv8 amd ARM64 architectures.\nTesting of other OSes are the responsibility of the community requesting cun\u012bcu for the OS in question.")),(0,r.kt)("h2",{id:"from-script"},"From Script"),(0,r.kt)("p",null,"cun\u012bcu also has an installer script that will automatically grab the latest version of cun\u012bcu and install it locally."),(0,r.kt)("p",null,"You can fetch that script, and then execute it locally.\nIt's well documented so that you can read through it and understand what it is doing before you run it."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"curl -fsSL -o get_cunicu.sh get.cunicu.li\nchmod 700 get_cunicu.sh\n./get_cunicu.sh\n")),(0,r.kt)("p",null,"Yes, you can ",(0,r.kt)("inlineCode",{parentName:"p"},"curl -fsSL get.cunicu.li | bash")," if you want to live on the edge."),(0,r.kt)("h2",{id:"through-package-managers"},"Through Package Managers"),(0,r.kt)("p",null,"cun\u012bcu provides the ability to install via operating system package managers."),(0,r.kt)("h3",{id:"from-apt-debian-ubuntu"},"From Apt (Debian, Ubuntu)"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},'sudo apt-get install apt-transport-https --yes\necho "deb [arch=$(dpkg --print-architecture) trusted=yes] https://packages.cunicu.li/apt/ /" | sudo tee /etc/apt/sources.list.d/cunicu.list\nsudo apt-get update\nsudo apt-get install cunicu\n')),(0,r.kt)("h3",{id:"from-yum-redhat-fedora-rockylinux"},"From Yum (Redhat, Fedora, RockyLinux)"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"echo '[cunicu]\nname=cunicu\nbaseurl=https://packages.cunicu.li/yum/\nenabled=1\ngpgcheck=0' | sudo tee /etc/yum.repos.d/cunicu.repo\nsudo yum install cunicu\n")),(0,r.kt)("h3",{id:"from-homebrew-macos"},"From Homebrew (macOS)"),(0,r.kt)("p",null,"A formulae for cun\u012bcu is available in our Homebrew Tap: ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/stv0g/homebrew-cunicu"},"https://github.com/stv0g/homebrew-cunicu"),"."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"brew tap stv0g/cunicu\nbrew install cunicu\n")),(0,r.kt)("h3",{id:"from-archlinux-user-repository-aur"},"From Archlinux User Repository (AUR)"),(0,r.kt)("p",null,"cun\u012bcu is available in the Archlinux User Repository: ",(0,r.kt)("a",{parentName:"p",href:"https://aur.archlinux.org/packages/cunicu-bin"},"https://aur.archlinux.org/packages/cunicu-bin"),"."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash",metastring:'title="via Yaourt"',title:'"via','Yaourt"':!0},"yaourt -S cunicu-bin\n")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash",metastring:'title="or via Packer"',title:'"or',via:!0,'Packer"':!0},"packer -S cunicu-bin\n")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash",metastring:'title="or via yay"',title:'"or',via:!0,'yay"':!0},"yay cunicu-bin\n")),(0,r.kt)("h2",{id:"from-source-all"},"From Source (all)"),(0,r.kt)("p",null,"Building cun\u012bcu is fairly easy and allows you to install the latest unreleased version."),(0,r.kt)("p",null,"You must have a working Go environment."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"go install github.com/stv0g/cunicu/cmd/cunicu@latest\n")),(0,r.kt)("p",null,"If required, it will fetch the dependencies and cache them, and validate configuration.\nIt will then compile cun\u012bcu and place it in ",(0,r.kt)("inlineCode",{parentName:"p"},"${GOPATH}/bin/cunicu"),"."),(0,r.kt)("h2",{id:"conclusion"},"Conclusion"),(0,r.kt)("p",null,"In most cases, installation is as simple as getting a pre-built cun\u012bcu binary.\nThis document covers additional cases for those who want to do more sophisticated things with cun\u012bcu."),(0,r.kt)("p",null,"Once you have cun\u012bcu successfully installed, you can move on to ",(0,r.kt)("a",{parentName:"p",href:"/docs/usage/"},"using cun\u012bcu")," to setup your mesh VPN network."))}p.isMDXComponent=!0}}]);