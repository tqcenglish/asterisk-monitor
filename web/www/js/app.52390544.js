(function(t){function e(e){for(var a,c,o=e[0],s=e[1],l=e[2],f=0,d=[];f<o.length;f++)c=o[f],Object.prototype.hasOwnProperty.call(r,c)&&r[c]&&d.push(r[c][0]),r[c]=0;for(a in s)Object.prototype.hasOwnProperty.call(s,a)&&(t[a]=s[a]);u&&u(e);while(d.length)d.shift()();return i.push.apply(i,l||[]),n()}function n(){for(var t,e=0;e<i.length;e++){for(var n=i[e],a=!0,o=1;o<n.length;o++){var s=n[o];0!==r[s]&&(a=!1)}a&&(i.splice(e--,1),t=c(c.s=n[0]))}return t}var a={},r={app:0},i=[];function c(e){if(a[e])return a[e].exports;var n=a[e]={i:e,l:!1,exports:{}};return t[e].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.m=t,c.c=a,c.d=function(t,e,n){c.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},c.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},c.t=function(t,e){if(1&e&&(t=c(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var a in t)c.d(n,a,function(e){return t[e]}.bind(null,a));return n},c.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return c.d(e,"a",e),e},c.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},c.p="";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],s=o.push.bind(o);o.push=e,o=o.slice();for(var l=0;l<o.length;l++)e(o[l]);var u=s;i.push([0,"chunk-vendors"]),n()})({0:function(t,e,n){t.exports=n("56d7")},"05ea":function(t,e,n){"use strict";n("1473")},1119:function(t,e,n){"use strict";n("2047")},1395:function(t,e,n){},1473:function(t,e,n){},2047:function(t,e,n){},"2cc5":function(t,e,n){"use strict";n("bfe8")},"3b37":function(t,e,n){},"56d7":function(t,e,n){"use strict";n.r(e);n("cadf"),n("551c"),n("f751"),n("097d");var a=n("2b0e"),r=function(){var t=this,e=t._self._c;return e("div",{attrs:{id:"app"}},[e("datav")],1)},i=[],c=function(){var t=this,e=t._self._c;return e("div",{attrs:{id:"data-view"}},[e("dv-full-screen-container",[e("div",{staticClass:"main-header"},[e("div",{staticClass:"mh-left"},[e("dv-border-box-2",{staticStyle:{width:"140px",height:"50px","line-height":"50px","text-align":"center","margin-left":"200px"}},[e("span",{on:{click:t.fullScreen}},[t._v("大数据平台")])])],1)]),e("dv-border-box-1",{staticClass:"main-container"},[e("dv-border-box-3",{staticClass:"left-chart-container"},[e("System-Info"),e("Storage"),e("Network")],1),e("div",{staticClass:"right-main-container"},[e("div",{staticClass:"rmc-top-container"},[e("dv-border-box-3",{staticClass:"rmctc-left-container"},[e("Center-Cmp")],1),e("div",{staticClass:"rmctc-right-container"},[e("dv-border-box-3",{staticClass:"rmctc-chart-1"},[e("Trunk")],1),e("dv-border-box-4",{staticClass:"rmctc-chart-2",attrs:{reverse:!0}},[e("Extension")],1)],1)],1),e("dv-border-box-4",{staticClass:"rmc-bottom-container"},[e("Queue")],1)],1)],1)],1)],1)},o=[],s=(n("7f7f"),function(){var t=this,e=t._self._c;return e("div",{staticClass:"left-chart-1"},[e("div",{staticClass:"lc1-header"},[t._v("设备信息")]),e("div",{staticClass:"lc1-details"},[t._v("无故障运行时间"),e("span",[t._v(t._s(t.maxTime)+" 小时 ")])]),e("dv-capsule-chart",{staticClass:"lc1-chart",attrs:{config:t.config}}),e("dv-decoration-2",{staticStyle:{height:"10px"}})],1)}),l=[],u=(n("8e6e"),n("ac6a"),n("456d"),n("bd86")),f=window.location.origin;function d(t){return fetch("".concat(f,"/api/calllog")).then((function(t){return t.json().then((function(t){return t}))}))}function v(t){return fetch("".concat(f,"/api/storageinfo")).then((function(t){return t.json().then((function(t){return t}))}))}function p(t){return fetch("".concat(f,"/api/systeminfo")).then((function(t){return t.json().then((function(t){return t}))}))}function b(t){return fetch("".concat(f,"/api/queuestatus")).then((function(t){return t.json().then((function(t){return t}))}))}function g(t){return fetch("".concat(f,"/api/extensionstatus")).then((function(t){return t.json().then((function(t){return t}))}))}function h(t){return fetch("".concat(f,"/api/trunkstatus")).then((function(t){return t.json().then((function(t){return t}))}))}function m(t){return fetch("".concat(f,"/api/networkinfo")).then((function(t){return t.json().then((function(t){return t}))}))}function O(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function y(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?O(Object(n),!0).forEach((function(e){Object(u["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):O(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var C={name:"SystemInfo",data:function(){return{config:{data:[{name:"Web 服务时间",value:10},{name:"Asterisk 服务时间",value:10},{name:"系统启动时间",value:10}],colors:["#00baff","#3de7c9","#ffc530"],unit:"小时"},maxTime:0}},created:function(){var t=this;p().then((function(e){var n=t.config;n.data[0].value=e.webUptime,n.data[1].value=e.voipUptime,n.data[2].value=e.systemUptime,t.config=y({},n),t.maxTime=e.systemUptime}))}},j=C,w=(n("1119"),n("2877")),_=Object(w["a"])(j,s,l,!1,null,null,null),P=_.exports,x=function(){var t=this,e=t._self._c;return e("div",{staticClass:"left-chart-2"},[e("div",{staticClass:"lc2-header"},[t._v("存储信息")]),t._m(0),e("dv-charts",{staticClass:"lc2-chart",attrs:{option:t.option}}),e("dv-decoration-2",{staticStyle:{height:"10px"}})],1)},S=[function(){var t=this,e=t._self._c;return e("div",{staticClass:"lc2-details"},[t._v("设备空间"),e("span",[t._v("245 GB")])])}];function D(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function k(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?D(Object(n),!0).forEach((function(e){Object(u["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):D(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var I={name:"LeftChart2",data:function(){return{option:{series:[{type:"pie",data:[{name:"通话录音",value:93},{name:"系统",value:32},{name:"语音邮件",value:65},{name:"音乐文件",value:44},{name:"其他",value:52}],radius:["45%","65%"],insideLabel:{show:!1},outsideLabel:{labelLineEndLength:10,formatter:"{percent}%\n{name}",style:{fontSize:14,fill:"#fff"}}}],color:["#00baff","#3de7c9","#fff","#ffc530","#469f4b"]}}},created:function(){var t=this;v().then((function(e){var n=t.option;n.series[0].data[0].value=e.record,n.series[0].data[1].value=e.system,n.series[0].data[2].value=e.voicemail,n.series[0].data[3].value=e.music,n.series[0].data[4].value=e.other,t.option=k({},n)}))}},E=I,L=(n("d681"),Object(w["a"])(E,x,S,!1,null,null,null)),T=L.exports,q=function(){var t=this,e=t._self._c;return e("div",{staticClass:"left-chart-3"},[e("div",{staticClass:"lc3-header"},[t._v("网络信息")]),t._m(0),e("dv-capsule-chart",{staticClass:"lc3-chart",attrs:{config:t.config}})],1)},R=[function(){var t=this,e=t._self._c;return e("div",{staticClass:"lc3-details"},[t._v("阻止攻击"),e("span",[t._v("215 次")])])}];function z(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function F(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?z(Object(n),!0).forEach((function(e){Object(u["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):z(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var U={name:"Network",data:function(){return{config:{data:[{name:"发送流量",value:1},{name:"接收流量",value:1}],colors:["#00baff","#3de7c9","#fff","#ffc530","#469f4b"],unit:"Gb"}}},created:function(){var t=this;this.getData(),this.timer=setInterval((function(){t.getData()}),6e4)},beforeDestroy:function(){clearInterval(this.timer)},methods:{getData:function(){var t=this;m().then((function(e){var n=t.config;n.data[0].value=e["tx"],n.data[1].value=e["rx"],console.log(e),t.config=F({},n)}))}}},M=U,N=(n("bbb5"),Object(w["a"])(M,q,R,!1,null,null,null)),V=N.exports,B=(n("6b54"),function(){var t=this,e=t._self._c;return e("div",{staticClass:"center-cmp"},[e("div",{staticClass:"cc-header"},[e("dv-decoration-1",{staticStyle:{width:"200px",height:"50px"}}),e("div",[t._v("今日通话统计")]),e("dv-decoration-1",{staticStyle:{width:"200px",height:"50px"}})],1),e("div",{staticClass:"cc-details"},[e("div",[t._v("通话总数")]),t._l(t.total.toString(),(function(n){return e("div",{key:n,staticClass:"card"},[t._v("\n      "+t._s(n)+"\n    ")])}))],2),e("div",{staticClass:"cc-main-container"},[e("div",{staticClass:"ccmc-left"},[e("div",{staticClass:"station-info"},[t._v("\n        呼入"),e("span",{staticStyle:{width:"40px"}},[t._v(t._s(t.incomingCallCount))])]),e("div",{staticClass:"station-info"},[t._v("\n        呼出"),e("span",{staticStyle:{width:"40px"}},[t._v(t._s(t.outgoingCallCount))])]),e("div",{staticClass:"station-info"},[t._v("\n        内部"),e("span",{staticStyle:{width:"40px"}},[t._v(t._s(t.internalCallCount))])])]),e("dv-active-ring-chart",{staticClass:"ccmc-middle",attrs:{config:t.config}}),e("LabelTag",{attrs:{config:t.labelConfig}})],1)])}),Q=[],G=function(){var t=this,e=t._self._c;return e("div",{staticClass:"label-tag"},[t.mergedConfig?t._l(t.mergedConfig.data,(function(n,a){return e("div",{key:n,staticClass:"label-item"},[t._v("\n      "+t._s(n)+" "),e("div",{style:"background-color: ".concat(t.mergedConfig.colors[a%t.mergedConfig.colors.length],";")})])})):t._e()],2)},J=[],W=n("becb"),A=n("5557"),$={name:"LabelTag",props:{config:{type:Object,default:function(){return[]}}},data:function(){return{defaultConfig:{data:[],colors:["#00baff","#3de7c9","#fff","#ffc530","#469f4b"]},mergedConfig:null}},watch:{config:function(){var t=this.mergeConfig;t()}},methods:{mergeConfig:function(){var t=this.config,e=this.defaultConfig;this.mergedConfig=Object(W["deepMerge"])(Object(A["deepClone"])(e,!0),t||{})}},mounted:function(){var t=this.mergeConfig;t()}},H=$,K=(n("05ea"),Object(w["a"])(H,G,J,!1,null,null,null)),X=K.exports;function Y(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function Z(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?Y(Object(n),!0).forEach((function(e){Object(u["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):Y(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var tt={name:"CenterCmp",components:{LabelTag:X},data:function(){return{config:{data:[{name:"呼入",value:10},{name:"呼出",value:10},{name:"内部",value:10}],color:["#00baff","#3de7c9","#fff","#ffc530","#469f4b"],lineWidth:30,radius:"55%",activeRadius:"60%"},labelConfig:{data:["呼入","呼出","内部"]},total:0,incomingCallCount:0,outgoingCallCount:0,internalCallCount:0}},created:function(){var t=this;this.getData(),this.timer=setInterval((function(){t.getData()}),6e4)},beforeDestroy:function(){clearInterval(this.timer)},methods:{getData:function(){var t=this;d().then((function(e){var n=t.config;n.data[0].value=e.incomingCallCount,n.data[1].value=e.outgoingCallCount,n.data[2].value=e.internalCallCount,t.config=Z({},n),t.incomingCallCount=e.incomingCallCount,t.outgoingCallCount=e.outgoingCallCount,t.internalCallCount=e.internalCallCount,t.total=e.internalCallCount+e.incomingCallCount+e.outgoingCallCount}))}}},et=tt,nt=(n("748d"),Object(w["a"])(et,B,Q,!1,null,null,null)),at=nt.exports,rt=function(){var t=this,e=t._self._c;return e("div",{staticClass:"right-chart-1"},[e("div",{staticClass:"rc1-header"},[t._v("中继累计通话")]),e("div",{staticClass:"rc1-chart-container"},[e("div",{staticClass:"left"},[e("div",{staticClass:"number"},[t._v(t._s(t.total))]),e("div",[t._v("线路总数")])]),e("dv-capsule-chart",{staticClass:"right",attrs:{config:t.config}})],1)])},it=[];function ct(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function ot(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?ct(Object(n),!0).forEach((function(e){Object(u["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):ct(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var st={name:"Trunk",data:function(){return{config:{data:[{name:"告警线路",value:25},{name:"客服线路",value:66}]},total:0}},created:function(){var t=this;this.getData(),this.timer=setInterval((function(){t.getData()}),1e4)},beforeDestroy:function(){clearInterval(this.timer)},methods:{getData:function(){var t=this;h().then((function(e){var n=Object.keys(e);t.total=n.length,t.config.data=n.map((function(t){return{name:t,value:e[t]}})),t.config=ot({},t.config)}))}}},lt=st,ut=(n("d740"),Object(w["a"])(lt,rt,it,!1,null,null,null)),ft=ut.exports,dt=function(){var t=this,e=t._self._c;return e("div",{staticClass:"right-chart-2"},[e("div",{staticClass:"rc1-header"},[t._v("分机状态")]),e("div",{staticClass:"rc1-chart-container"},[e("div",{staticClass:"left"},[e("div",{staticClass:"number"},[t._v(t._s(t.total))]),e("div",[t._v("分机总数")])]),e("dv-charts",{staticClass:"right",attrs:{option:t.option}})],1)])},vt=[];function pt(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function bt(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?pt(Object(n),!0).forEach((function(e){Object(u["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):pt(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var gt={name:"Extension",data:function(){return{option:{series:[{type:"pie",data:[{name:"在线",value:1},{name:"通话中",value:1},{name:"离线",value:1}],radius:["45%","65%"],insideLabel:{show:!1},outsideLabel:{labelLineEndLength:20,formatter:"{value}\n{name}",style:{fontSize:14,fill:"#fff"}}}],color:["#00baff","#3de7c9","#ffc530","#469f4b"]},total:0}},created:function(){var t=this;this.getData(),this.timer=setInterval((function(){t.getData()}),1e4)},beforeDestroy:function(){clearInterval(this.timer)},methods:{getData:function(){var t=this;g().then((function(e){t.option.series[0].data[0].value=e["Not in use"],t.option.series[0].data[1].value=e["Busy"]+e["In use"]+e["Ringing"],t.option.series[0].data[2].value=e["Unavailable"],t.option=bt({},t.option),t.total=e["Not in use"]+e["Busy"]+e["In use"]+e["Ringing"]+e["Unavailable"],console.log(t.option)}))}}},ht=gt,mt=(n("fa80"),Object(w["a"])(ht,dt,vt,!1,null,null,null)),Ot=mt.exports,yt=function(){var t=this,e=t._self._c;return e("div",{staticClass:"bottom-charts"},[e("div",{staticClass:"bc-chart-item"},[e("div",{staticClass:"bcci-header"},[t._v(t._s(t.config1.name))]),e("dv-active-ring-chart",{attrs:{config:t.config1}}),e("Label-Tag",{attrs:{config:t.labelConfig}})],1),e("dv-decoration-2",{staticClass:"decoration-1",staticStyle:{width:"5px"},attrs:{reverse:!0}}),e("div",{staticClass:"bc-chart-item"},[e("div",{staticClass:"bcci-header"},[t._v(t._s(t.config2.name))]),e("dv-active-ring-chart",{attrs:{config:t.config2}}),e("Label-Tag",{attrs:{config:t.labelConfig}})],1),e("dv-decoration-2",{staticClass:"decoration-1",staticStyle:{width:"5px"},attrs:{reverse:!0}}),e("div",{staticClass:"bc-chart-item"},[e("div",{staticClass:"bcci-header"},[t._v(t._s(t.config3.name))]),e("dv-active-ring-chart",{attrs:{config:t.config3}}),e("Label-Tag",{attrs:{config:t.labelConfig}})],1),e("dv-decoration-2",{staticClass:"decoration-1",staticStyle:{width:"5px"},attrs:{reverse:!0}}),e("div",{staticClass:"bc-chart-item"},[e("div",{staticClass:"bcci-header"},[t._v(t._s(t.config4.name))]),e("dv-active-ring-chart",{attrs:{config:t.config4}}),e("Label-Tag",{attrs:{config:t.labelConfig}})],1),e("dv-decoration-2",{staticClass:"decoration-1",staticStyle:{width:"5px"},attrs:{reverse:!0}})],1)},Ct=[];function jt(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function wt(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?jt(Object(n),!0).forEach((function(e){Object(u["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):jt(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var _t={name:"Queue",components:{LabelTag:X},data:function(){return{config1:{radius:"70%",activeRadius:"75%",data:[{name:"中止",value:1},{name:"通话中",value:1},{name:"已完成",value:1}],color:["#00baff","#3de7c9","#fff","#ffc530","#469f4b"],digitalFlopStyle:{fontSize:20},showOriginValue:!0},config2:{radius:"70%",activeRadius:"75%",data:[{name:"中止",value:1},{name:"通话中",value:1},{name:"已完成",value:1}],color:["#00baff","#3de7c9","#fff","#ffc530","#469f4b"],digitalFlopStyle:{fontSize:20},showOriginValue:!0},config3:{radius:"70%",activeRadius:"75%",data:[{name:"中止",value:1},{name:"通话中",value:1},{name:"已完成",value:1}],digitalFlopStyle:{fontSize:20},showOriginValue:!0},config4:{radius:"70%",activeRadius:"75%",data:[{name:"中止",value:1},{name:"通话中",value:1},{name:"已完成",value:1}],digitalFlopStyle:{fontSize:20},showOriginValue:!0},labelConfig:{data:["中止","通话中","已完成"]}}},created:function(){var t=this;this.getData(),this.timer=setInterval((function(){t.getData()}),1e4)},beforeDestroy:function(){clearInterval(this.timer)},methods:{getData:function(){var t=this;b().then((function(e){e.forEach((function(e,n){if(!(n>3))switch(n){case 0:var a=t.config1;if(a.name="".concat(e.name,"(").concat(e.queue,")"),a.data[0].value=parseInt(e.abandoned),a.data[1].value=parseInt(e.calls),a.data[2].value=parseInt(e.completed),0===a.data[0].value&&0===a.data[1].value&&0===a.data[2].value)break;t.config1=wt({},a);break;case 1:var r=t.config2;if(r.name="".concat(e.name,"(").concat(e.queue,")"),r.data[0].value=parseInt(e.abandoned),r.data[1].value=parseInt(e.calls),r.data[2].value=parseInt(e.completed),0===r.data[0].value&&0===r.data[1].value&&0===r.data[2].value)break;t.config2=wt({},r);break;case 2:var i=t.config3;if(i.name="".concat(e.name,"(").concat(e.queue,")"),i.data[0].value=parseInt(e.abandoned),i.data[1].value=parseInt(e.calls),i.data[2].value=parseInt(e.completed),0===i.data[0].value&&0===i.data[1].value&&0===i.data[2].value)break;t.config3=wt({},i);break;case 3:var c=t.config4;if(c.name="".concat(e.name,"(").concat(e.queue,")"),c.data[0].value=parseInt(e.abandoned),c.data[1].value=parseInt(e.calls),c.data[2].value=parseInt(e.completed),0===c.data[0].value&&0===c.data[1].value&&0===c.data[2].value)break;t.config4=wt({},c);break;default:break}}))}))}}},Pt=_t,xt=(n("6452"),Object(w["a"])(Pt,yt,Ct,!1,null,null,null)),St=xt.exports,Dt={name:"DataView",components:{SystemInfo:P,Storage:T,Network:V,CenterCmp:at,Trunk:ft,Extension:Ot,Queue:St},data:function(){return{}},methods:{fullScreen:function(){var t=document.querySelector("#app");document.fullscreenElement?document.exitFullscreen():t.requestFullscreen().catch((function(t){alert("Error attempting to enable fullscreen mode: ".concat(t.message," (").concat(t.name,")"))}))}}},kt=Dt,It=(n("2cc5"),Object(w["a"])(kt,c,o,!1,null,null,null)),Et=It.exports,Lt={name:"app",components:{datav:Et},data:function(){return{}}},Tt=Lt,qt=(n("99f2"),Object(w["a"])(Tt,r,i,!1,null,null,null)),Rt=qt.exports,zt=(n("1395"),n("6c29"));a["a"].config.productionTip=!1,a["a"].use(zt["a"]),new a["a"]({render:function(t){return t(Rt)}}).$mount("#app")},6452:function(t,e,n){"use strict";n("3b37")},6786:function(t,e,n){},"748d":function(t,e,n){"use strict";n("901a")},"901a":function(t,e,n){},9891:function(t,e,n){},"99f2":function(t,e,n){"use strict";n("9891")},"9d11":function(t,e,n){},bbb5:function(t,e,n){"use strict";n("9d11")},bfe8:function(t,e,n){},d681:function(t,e,n){"use strict";n("f0e7")},d740:function(t,e,n){"use strict";n("ed15")},ed15:function(t,e,n){},f0e7:function(t,e,n){},fa80:function(t,e,n){"use strict";n("6786")}});
//# sourceMappingURL=app.52390544.js.map