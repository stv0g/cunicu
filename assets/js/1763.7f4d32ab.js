"use strict";(self.webpackChunkcunicu=self.webpackChunkcunicu||[]).push([[1763],{43349:(e,t,n)=>{n.d(t,{a:()=>l});var r=n(96225);function l(e,t){var n=e.append("foreignObject").attr("width","100000"),l=n.append("xhtml:div");l.attr("xmlns","http://www.w3.org/1999/xhtml");var o=t.label;switch(typeof o){case"function":l.insert(o);break;case"object":l.insert((function(){return o}));break;default:l.html(o)}r.bg(l,t.labelStyle),l.style("display","inline-block"),l.style("white-space","nowrap");var a=l.node().getBoundingClientRect();return n.attr("width",a.width).attr("height",a.height),n}},96225:(e,t,n)=>{n.d(t,{$p:()=>d,O1:()=>a,WR:()=>p,bF:()=>o,bg:()=>c});var r=n(37514),l=n(73234);function o(e,t){return!!e.children(t).length}function a(e){return i(e.v)+":"+i(e.w)+":"+i(e.name)}var s=/:/g;function i(e){return e?String(e).replace(s,"\\:"):""}function c(e,t){t&&e.attr("style",t)}function d(e,t,n){t&&e.attr("class",t).attr("class",n+" "+e.attr("class"))}function p(e,t){var n=t.graph();if(r.Z(n)){var o=n.transition;if(l.Z(o))return o(e)}return e}},1763:(e,t,n)=>{n.d(t,{diagram:()=>i});var r=n(88955),l=(n(45625),n(97934));n(36690),n(17452),n(3688),n(70870),n(41644),n(96225);n(43349);n(66749),n(74379);n(61666);l.c_6;var o=n(21358);n(27484),n(17967),n(27856),n(39354);const a={},s=function(e){const t=Object.keys(e);for(const n of t)a[n]=e[n]},i={parser:r.p,db:r.f,renderer:o.f,styles:o.a,init:e=>{e.flowchart||(e.flowchart={}),e.flowchart.arrowMarkerAbsolute=e.arrowMarkerAbsolute,s(e.flowchart),r.f.clear(),r.f.setGen("gen-1")}}},21358:(e,t,n)=>{n.d(t,{a:()=>h,f:()=>w});var r=n(45625),l=n(97934),o=n(36690),a=n(87936),s=n(43349),i=n(61691),c=n(71610);const d=(e,t)=>i.Z.lang.round(c.Z.parse(e)[t]);var p=n(51117);const b={},f=function(e,t,n,r,l,a){const i=r.select(`[id="${n}"]`);Object.keys(e).forEach((function(n){const r=e[n];let c="default";r.classes.length>0&&(c=r.classes.join(" ")),c+=" flowchart-label";const d=(0,o.k)(r.styles);let p,b=void 0!==r.text?r.text:r.id;if(o.l.info("vertex",r,r.labelType),"markdown"===r.labelType)o.l.info("vertex",r,r.labelType);else if((0,o.m)((0,o.c)().flowchart.htmlLabels)){const e={label:b.replace(/fa[blrs]?:fa-[\w-]+/g,(e=>`<i class='${e.replace(":"," ")}'></i>`))};p=(0,s.a)(i,e).node(),p.parentNode.removeChild(p)}else{const e=l.createElementNS("http://www.w3.org/2000/svg","text");e.setAttribute("style",d.labelStyle.replace("color:","fill:"));const t=b.split(o.e.lineBreakRegex);for(const n of t){const t=l.createElementNS("http://www.w3.org/2000/svg","tspan");t.setAttributeNS("http://www.w3.org/XML/1998/namespace","xml:space","preserve"),t.setAttribute("dy","1em"),t.setAttribute("x","1"),t.textContent=n,e.appendChild(t)}p=e}let f=0,u="";switch(r.type){case"round":f=5,u="rect";break;case"square":case"group":default:u="rect";break;case"diamond":u="question";break;case"hexagon":u="hexagon";break;case"odd":case"odd_right":u="rect_left_inv_arrow";break;case"lean_right":u="lean_right";break;case"lean_left":u="lean_left";break;case"trapezoid":u="trapezoid";break;case"inv_trapezoid":u="inv_trapezoid";break;case"circle":u="circle";break;case"ellipse":u="ellipse";break;case"stadium":u="stadium";break;case"subroutine":u="subroutine";break;case"cylinder":u="cylinder";break;case"doublecircle":u="doublecircle"}t.setNode(r.id,{labelStyle:d.labelStyle,shape:u,labelText:b,labelType:r.labelType,rx:f,ry:f,class:c,style:d.style,id:r.id,link:r.link,linkTarget:r.linkTarget,tooltip:a.db.getTooltip(r.id)||"",domId:a.db.lookUpDomId(r.id),haveCallback:r.haveCallback,width:"group"===r.type?500:void 0,dir:r.dir,type:r.type,props:r.props,padding:(0,o.c)().flowchart.padding}),o.l.info("setNode",{labelStyle:d.labelStyle,labelType:r.labelType,shape:u,labelText:b,rx:f,ry:f,class:c,style:d.style,id:r.id,domId:a.db.lookUpDomId(r.id),width:"group"===r.type?500:void 0,type:r.type,dir:r.dir,props:r.props,padding:(0,o.c)().flowchart.padding})}))},u=function(e,t,n){o.l.info("abc78 edges = ",e);let r,a,s=0,i={};if(void 0!==e.defaultStyle){const t=(0,o.k)(e.defaultStyle);r=t.style,a=t.labelStyle}e.forEach((function(n){s++;const c="L-"+n.start+"-"+n.end;void 0===i[c]?(i[c]=0,o.l.info("abc78 new entry",c,i[c])):(i[c]++,o.l.info("abc78 new entry",c,i[c]));let d=c+"-"+i[c];o.l.info("abc78 new link id to be used is",c,d,i[c]);const p="LS-"+n.start,f="LE-"+n.end,u={style:"",labelStyle:""};switch(u.minlen=n.length||1,"arrow_open"===n.type?u.arrowhead="none":u.arrowhead="normal",u.arrowTypeStart="arrow_open",u.arrowTypeEnd="arrow_open",n.type){case"double_arrow_cross":u.arrowTypeStart="arrow_cross";case"arrow_cross":u.arrowTypeEnd="arrow_cross";break;case"double_arrow_point":u.arrowTypeStart="arrow_point";case"arrow_point":u.arrowTypeEnd="arrow_point";break;case"double_arrow_circle":u.arrowTypeStart="arrow_circle";case"arrow_circle":u.arrowTypeEnd="arrow_circle"}let w="",h="";switch(n.stroke){case"normal":w="fill:none;",void 0!==r&&(w=r),void 0!==a&&(h=a),u.thickness="normal",u.pattern="solid";break;case"dotted":u.thickness="normal",u.pattern="dotted",u.style="fill:none;stroke-width:2px;stroke-dasharray:3;";break;case"thick":u.thickness="thick",u.pattern="solid",u.style="stroke-width: 3.5px;fill:none;";break;case"invisible":u.thickness="invisible",u.pattern="solid",u.style="stroke-width: 0;fill:none;"}if(void 0!==n.style){const e=(0,o.k)(n.style);w=e.style,h=e.labelStyle}u.style=u.style+=w,u.labelStyle=u.labelStyle+=h,void 0!==n.interpolate?u.curve=(0,o.n)(n.interpolate,l.c_6):void 0!==e.defaultInterpolate?u.curve=(0,o.n)(e.defaultInterpolate,l.c_6):u.curve=(0,o.n)(b.curve,l.c_6),void 0===n.text?void 0!==n.style&&(u.arrowheadStyle="fill: #333"):(u.arrowheadStyle="fill: #333",u.labelpos="c"),u.labelType=n.labelType,u.label=n.text.replace(o.e.lineBreakRegex,"\n"),void 0===n.style&&(u.style=u.style||"stroke: #333; stroke-width: 1.5px;fill:none;"),u.labelStyle=u.labelStyle.replace("color:","fill:"),u.id=d,u.classes="flowchart-link "+p+" "+f,t.setEdge(n.start,n.end,u,s)}))},w={setConf:function(e){const t=Object.keys(e);for(const n of t)b[n]=e[n]},addVertices:f,addEdges:u,getClasses:function(e,t){return t.db.getClasses()},draw:async function(e,t,n,s){o.l.info("Drawing flowchart");let i=s.db.getDirection();void 0===i&&(i="TD");const{securityLevel:c,flowchart:d}=(0,o.c)(),p=d.nodeSpacing||50,b=d.rankSpacing||50;let w;"sandbox"===c&&(w=(0,l.Ys)("#i"+t));const h="sandbox"===c?(0,l.Ys)(w.nodes()[0].contentDocument.body):(0,l.Ys)("body"),g="sandbox"===c?w.nodes()[0].contentDocument:document,y=new r.k({multigraph:!0,compound:!0}).setGraph({rankdir:i,nodesep:p,ranksep:b,marginx:0,marginy:0}).setDefaultEdgeLabel((function(){return{}}));let k;const x=s.db.getSubGraphs();o.l.info("Subgraphs - ",x);for(let r=x.length-1;r>=0;r--)k=x[r],o.l.info("Subgraph - ",k),s.db.addVertex(k.id,{text:k.title,type:k.labelType},"group",void 0,k.classes,k.dir);const v=s.db.getVertices(),m=s.db.getEdges();o.l.info("Edges",m);let S=0;for(S=x.length-1;S>=0;S--){k=x[S],(0,l.td_)("cluster").append("text");for(let e=0;e<k.nodes.length;e++)o.l.info("Setting up subgraphs",k.nodes[e],k.id),y.setParent(k.nodes[e],k.id)}f(v,y,t,h,g,s),u(m,y);const T=h.select(`[id="${t}"]`),_=h.select("#"+t+" g");if(await(0,a.r)(_,y,["point","circle","cross"],"flowchart",t),o.u.insertTitle(T,"flowchartTitleText",d.titleTopMargin,s.db.getDiagramTitle()),(0,o.o)(y,T,d.diagramPadding,d.useMaxWidth),s.db.indexNodes("subGraph"+S),!d.htmlLabels){const e=g.querySelectorAll('[id="'+t+'"] .edgeLabel .label');for(const t of e){const e=t.getBBox(),n=g.createElementNS("http://www.w3.org/2000/svg","rect");n.setAttribute("rx",0),n.setAttribute("ry",0),n.setAttribute("width",e.width),n.setAttribute("height",e.height),t.insertBefore(n,t.firstChild)}}Object.keys(v).forEach((function(e){const n=v[e];if(n.link){const r=(0,l.Ys)("#"+t+' [id="'+e+'"]');if(r){const e=g.createElementNS("http://www.w3.org/2000/svg","a");e.setAttributeNS("http://www.w3.org/2000/svg","class",n.classes.join(" ")),e.setAttributeNS("http://www.w3.org/2000/svg","href",n.link),e.setAttributeNS("http://www.w3.org/2000/svg","rel","noopener"),"sandbox"===c?e.setAttributeNS("http://www.w3.org/2000/svg","target","_top"):n.linkTarget&&e.setAttributeNS("http://www.w3.org/2000/svg","target",n.linkTarget);const t=r.insert((function(){return e}),":first-child"),l=r.select(".label-container");l&&t.append((function(){return l.node()}));const o=r.select(".label");o&&t.append((function(){return o.node()}))}}}))}},h=e=>`.label {\n    font-family: ${e.fontFamily};\n    color: ${e.nodeTextColor||e.textColor};\n  }\n  .cluster-label text {\n    fill: ${e.titleColor};\n  }\n  .cluster-label span,p {\n    color: ${e.titleColor};\n  }\n\n  .label text,span,p {\n    fill: ${e.nodeTextColor||e.textColor};\n    color: ${e.nodeTextColor||e.textColor};\n  }\n\n  .node rect,\n  .node circle,\n  .node ellipse,\n  .node polygon,\n  .node path {\n    fill: ${e.mainBkg};\n    stroke: ${e.nodeBorder};\n    stroke-width: 1px;\n  }\n  .flowchart-label text {\n    text-anchor: middle;\n  }\n  // .flowchart-label .text-outer-tspan {\n  //   text-anchor: middle;\n  // }\n  // .flowchart-label .text-inner-tspan {\n  //   text-anchor: start;\n  // }\n\n  .node .label {\n    text-align: center;\n  }\n  .node.clickable {\n    cursor: pointer;\n  }\n\n  .arrowheadPath {\n    fill: ${e.arrowheadColor};\n  }\n\n  .edgePath .path {\n    stroke: ${e.lineColor};\n    stroke-width: 2.0px;\n  }\n\n  .flowchart-link {\n    stroke: ${e.lineColor};\n    fill: none;\n  }\n\n  .edgeLabel {\n    background-color: ${e.edgeLabelBackground};\n    rect {\n      opacity: 0.5;\n      background-color: ${e.edgeLabelBackground};\n      fill: ${e.edgeLabelBackground};\n    }\n    text-align: center;\n  }\n\n  /* For html labels only */\n  .labelBkg {\n    background-color: ${((e,t)=>{const n=d,r=n(e,"r"),l=n(e,"g"),o=n(e,"b");return p.Z(r,l,o,t)})(e.edgeLabelBackground,.5)};\n    // background-color: \n  }\n\n  .cluster rect {\n    fill: ${e.clusterBkg};\n    stroke: ${e.clusterBorder};\n    stroke-width: 1px;\n  }\n\n  .cluster text {\n    fill: ${e.titleColor};\n  }\n\n  .cluster span,p {\n    color: ${e.titleColor};\n  }\n  /* .cluster div {\n    color: ${e.titleColor};\n  } */\n\n  div.mermaidTooltip {\n    position: absolute;\n    text-align: center;\n    max-width: 200px;\n    padding: 2px;\n    font-family: ${e.fontFamily};\n    font-size: 12px;\n    background: ${e.tertiaryColor};\n    border: 1px solid ${e.border2};\n    border-radius: 2px;\n    pointer-events: none;\n    z-index: 100;\n  }\n\n  .flowchartTitleText {\n    text-anchor: middle;\n    font-size: 18px;\n    fill: ${e.textColor};\n  }\n`}}]);