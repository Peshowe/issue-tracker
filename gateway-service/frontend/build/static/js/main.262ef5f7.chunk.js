(this["webpackJsonpfrontend-client"]=this["webpackJsonpfrontend-client"]||[]).push([[0],{56:function(e,t,n){},57:function(e,t,n){},77:function(e,t,n){"use strict";n.r(t);n(51);var c,s,r,i,o,a=n(1),u=n.n(a),j=n(16),l=n.n(j),b=(n(56),n(7)),d=(n.p,n(57),n(3)),O=n(44),f=n(22),p=n(5),h=n(0),x=p.default.ul(c||(c=Object(b.a)(["\n  list-style-type: none;\n  margin: 0;\n  padding: 0;\n"]))),g=p.default.li(s||(s=Object(b.a)(["\n  display: block;\n  margin: 0;\n  padding: 0;\n"]))),m=p.default.a(r||(r=Object(b.a)(["\n  border-bottom: 1px solid #efefef;\n  color: ",";\n  display: block;\n  font-family: Helvetica, Arial, sans-serif;\n  font-size: 16px;\n  font-weight: 700;\n  margin-left: 18px;\n  padding: 18px 18px 18px 5px;\n  text-decoration: none;\n  text-transform: uppercase;\n  transition: background 200ms ease-in-out;\n  &:hover {\n    color: rgba(218, 135, 196);\n  }\n"])),(function(e){return e.active?"rgba(218, 135, 196)":"#333"})),v=function(){return Object(h.jsx)(h.Fragment,{children:Object(h.jsxs)(x,{children:[Object(h.jsx)(g,{children:Object(h.jsx)(m,{active:!0,href:"#",children:"Home"})}),Object(h.jsx)(g,{children:Object(h.jsx)(m,{href:"#",children:"About"})}),Object(h.jsx)(g,{children:Object(h.jsx)(m,{href:"#",children:"Work"})}),Object(h.jsx)(g,{children:Object(h.jsx)(m,{href:"#",children:"Contact"})})]})})},y=n(38),S=function(e){var t=e.onClose;return Object(h.jsx)("div",{style:{padding:"10px 10px 20px 10px",textAlign:"right"},children:Object(h.jsx)(y.a,{onClick:t,size:36,style:{cursor:"pointer",textAlign:"center"}})})},C=n(41),k=Object(p.default)(C.a)(i||(i=Object(b.a)(["\n  ","\n  color: #fff;\n  ","\n  ","\n  transition: transform 200ms ease-in-out;\n  &:hover {\n    transform: scale(1.2);\n  }\n"])),"","",""),w=function(){return Object(h.jsx)("a",{href:"https://github.com/Peshowe/issue-tracker/",target:"_blank",rel:"noopener noreferrer",children:Object(h.jsx)(k,{size:36})})},N=n(42),I=Object(p.default)(N.a)(o||(o=Object(b.a)(["\n&:hover {\n    transform: scale(1.2);\n  }\n"])));var _,P=function(e){return Object(h.jsx)(I,{size:48,style:{borderRadius:"6px",padding:"8px",cursor:"pointer"},onClick:function(){"/"!=window.location.pathname&&window.history.back()}})},E=Object(p.default)(O.a)(_||(_=Object(b.a)(["\n&:hover {\n    transform: scale(1.2);\n  }\n"])));var T=function(){var e=Object(a.useState)(!1),t=Object(d.a)(e,2),n=t[0],c=t[1];return Object(h.jsx)("div",{children:Object(h.jsxs)(f.StyledOffCanvas,{menuBackground:"black",isOpen:n,onClose:function(){return c(!1)},children:[Object(h.jsx)(P,{}),Object(h.jsx)(E,{size:48,style:{borderRadius:"6px",padding:"8px",cursor:"pointer"},onClick:function(){c((function(e){return!e}))}}),Object(h.jsx)(w,{}),Object(h.jsx)("span",{style:{padding:"8px","font-size":"1.7em","margin-left":"1em"},children:"Parvus JIRA"}),Object(h.jsx)(f.Menu,{children:Object(h.jsxs)(h.Fragment,{children:[Object(h.jsx)(S,{onClose:function(){return c(!1)}}),Object(h.jsx)(v,{})]})}),Object(h.jsx)(f.Overlay,{})]})})},A=n(21),B=n(4),D=n(45);var F,z=function(){return Object(h.jsx)(D.FingerprintSpinner,{color:"red"})},H=n(13),J=n.n(H),L=J.a.styled(F||(F=Object(b.a)(["\n  width: 30rem;\n  height: 20rem;\n  display: flex;\n  align-items: center;\n  justify-content: center;\n  background-color: grey;\n  opacity: ",";\n  transition : all 0.3s ease-in-out;"])),(function(e){return e.opacity}));var M=function(e){var t=Object(a.useState)(!1),n=Object(d.a)(t,2),c=n[0],s=n[1],r=Object(a.useState)(0),i=Object(d.a)(r,2),o=i[0],u=i[1],j=Object(a.useState)(""),l=Object(d.a)(j,2),b=l[0],O=l[1],f=Object(a.useState)(!1),p=Object(d.a)(f,2),x=p[0],g=p[1],m=Object(a.useState)(null),v=Object(d.a)(m,2),y=v[0],S=v[1];function C(e){u(0),s(!c)}function k(e){O(e.target.value)}function w(){g(!0),fetch("/v1/projects",{method:"POST",headers:{Accept:"application/json","Content-Type":"application/json"},body:JSON.stringify({name:b})}).then((function(t){g(!1),s(!1),e.onCreate()}),(function(e){g(!1),S(e)}))}return Object(h.jsxs)("div",{style:{padding:"1.5em"},children:[Object(h.jsx)("button",{onClick:C,className:"btn btn-primary",children:"New Project"}),Object(h.jsx)(L,{isOpen:c,afterOpen:function(){setTimeout((function(){u(1)}),100)},beforeClose:function(){return new Promise((function(e){u(0),setTimeout(e,300)}))},onBackgroundClick:C,onEscapeKeydown:C,opacity:o,backgroundProps:{opacity:o},children:x?Object(h.jsx)(z,{}):y?Object(h.jsxs)("div",{children:["Error: ",y.message]}):Object(h.jsxs)("div",{children:[Object(h.jsxs)("form",{children:[Object(h.jsx)("label",{children:"Project Name: "}),Object(h.jsx)("input",{type:"text",value:b,onChange:k})]}),Object(h.jsx)("button",{onClick:w,className:"btn btn-primary",children:"Create"})]})})]})};var R=function(){var e=Object(a.useState)(null),t=Object(d.a)(e,2),n=t[0],c=t[1],s=Object(a.useState)(!1),r=Object(d.a)(s,2),i=r[0],o=r[1],u=Object(a.useState)([]),j=Object(d.a)(u,2),l=j[0],b=j[1];function O(){fetch("/v1/projects").then((function(e){return e.json()})).then((function(e){o(!0),b(e.projects)}),(function(e){o(!0),c(e)}))}return Object(a.useEffect)((function(){O()}),[]),n?Object(h.jsxs)("div",{children:["Error: ",n.message]}):i?Object(h.jsxs)("div",{children:[Object(h.jsx)(M,{onCreate:O}),l.map((function(e){return Object(h.jsx)("div",{className:"listItem",children:Object(h.jsx)(A.b,{to:"/projects/".concat(e.id),className:"btn btn-primary",children:e.name})})}))]}):Object(h.jsx)(z,{})},K=n(47),U=n(23),X=n(48),W=n.n(X);var q,G=function(e){var t=Object(a.useState)({}),n=Object(d.a)(t,2),c=n[0],s=n[1],r=Object(a.useState)(""),i=Object(d.a)(r,2),o=i[0],u=i[1],j=Object(a.useState)(""),l=Object(d.a)(j,2),b=l[0],O=l[1],f=Object(a.useState)("feature"),p=Object(d.a)(f,2),x=p[0],g=p[1],m=Object(a.useState)(""),v=Object(d.a)(m,2),y=v[0],S=v[1];return Object(a.useEffect)((function(){null!=e.issue&&(s(e.issue),"name"in e.issue&&u(e.issue.name),"desc"in e.issue&&O(e.issue.desc),"issue_type"in e.issue&&g(e.issue.issue_type),"bug_trace"in e.issue&&S(e.issue.bug_trace))}),[]),Object(h.jsxs)("div",{style:{},children:[Object(h.jsxs)("form",{children:[Object(h.jsxs)("div",{children:[Object(h.jsx)("label",{children:"Name: "}),Object(h.jsx)("input",{type:"text",value:o,onChange:function(e){return u(e.target.value)}}),Object(h.jsx)("label",{children:"Issue type: "}),Object(h.jsxs)("select",{value:x,onChange:function(e){return g(e.target.value)},children:[Object(h.jsx)("option",{value:"bug",children:"Bug"}),Object(h.jsx)("option",{value:"feature",children:"Feature"}),Object(h.jsx)("option",{value:"adhoc",children:"Ad Hoc"})]})]}),Object(h.jsxs)("div",{children:[Object(h.jsx)("label",{children:"Description: "}),Object(h.jsx)("br",{}),Object(h.jsx)("textarea",{className:"inputArea",value:b,onChange:function(e){return O(e.target.value)}})]}),Object(h.jsxs)("div",{children:[Object(h.jsx)("label",{children:"Bug trace: "}),Object(h.jsx)("br",{}),Object(h.jsx)("textarea",{className:"inputArea",value:y,onChange:function(e){return S(e.target.value)}})]})]}),Object(h.jsx)("button",{onClick:function(){c.name=o,c.desc=b,c.issue_type=x,c.bug_trace=y,console.log(c),console.log(b),s(c),e.buttonHandler(c)},className:"btn btn-primary",children:e.buttonLabel})]})},Q=J.a.styled(q||(q=Object(b.a)(["\n  width: 70rem;\n  height: 40rem;\n  display: flex;\n  align-items: center;\n  justify-content: center;\n  background-color: grey;\n  opacity: ",";\n  transition : all 0.3s ease-in-out;"])),(function(e){return e.opacity}));var V=function(e){var t=Object(a.useState)(!1),n=Object(d.a)(t,2),c=n[0],s=n[1],r=Object(a.useState)(0),i=Object(d.a)(r,2),o=i[0],u=i[1],j=Object(a.useState)(!1),l=Object(d.a)(j,2),b=l[0],O=l[1],f=Object(a.useState)(null),p=Object(d.a)(f,2),x=p[0],g=p[1];function m(e){e.stopPropagation(),u(0),s(!c)}function v(t){O(!0),console.log("Heeloo"),e.onSubmit(t).then((function(t){O(!1),s(!1),e.onDone()}),(function(e){O(!1),g(e)}))}var y=null==e.issue?"Create Issue":"Update Issue",S=null==e.issue?"New Issue":e.issue.name;return Object(h.jsxs)("div",{children:[Object(h.jsx)("button",{onClick:m,className:"btn btn-primary",children:S}),Object(h.jsx)(Q,{isOpen:c,afterOpen:function(){setTimeout((function(){u(1)}),100)},beforeClose:function(){return new Promise((function(e){u(0),setTimeout(e,300)}))},onBackgroundClick:m,onEscapeKeydown:m,opacity:o,backgroundProps:{opacity:o},children:b?Object(h.jsx)(z,{}):x?Object(h.jsxs)("div",{children:["Error: ",x.message]}):Object(h.jsx)(G,{buttonHandler:v,buttonLabel:y,issue:e.issue})})]})};var Y=function(e){var t=Object(a.useState)(null),n=Object(d.a)(t,2),c=n[0],s=n[1],r=Object(a.useState)(!1),i=Object(d.a)(r,2),o=i[0],u=i[1],j=Object(a.useState)([]),l=Object(d.a)(j,2),b=l[0],O=l[1],f=Object(a.useState)([]),p=Object(d.a)(f,2),x=p[0],g=p[1],m=Object(a.useState)([]),v=Object(d.a)(m,2),y=v[0],S=v[1],C=["todo_col","inprogress_col","done_col"],k=["to do","in progress","done"];function w(e){return console.log(e),fetch("/v1/issues",{method:"PUT",headers:{Accept:"application/json","Content-Type":"application/json"},body:JSON.stringify({issue:e})})}function N(e,t){!function(e){"to do"==e.status?O(b.filter((function(t){return t.id!=e.id}))):"in progress"==e.status?g(x.filter((function(t){return t.id!=e.id}))):"done"==e.status&&S(y.filter((function(t){return t.id!=e.id})))}(e),e.status=t,function(e){"to do"==e.status?O([].concat(Object(U.a)(b),[e])):"in progress"==e.status?g([].concat(Object(U.a)(x),[e])):"done"==e.status&&S([].concat(Object(U.a)(y),[e]))}(e),u(!1),w(e).then((function(e){u(!0)}),(function(e){u(!0),s(e)}))}function I(){fetch("/v1/issues/byproject/".concat(e.projectId)).then((function(e){return e.json()})).then((function(e){if(u(!0),void 0!=e.issues){var t=e.issues;t.sort((function(e,t){return e.last_modified_on-t.last_modified_on})),O(t.filter((function(e){return"to do"==e.status}))),g(t.filter((function(e){return"in progress"==e.status}))),S(t.filter((function(e){return"done"==e.status})))}}),(function(e){u(!0),s(e)}))}function _(e){var t,n=[],c=Object(K.a)(e.entries());try{var s=function(){var e=Object(d.a)(t.value,2),c=(e[0],e[1]);n.push(Object(h.jsx)(W.a,{bounds:"#board",onDrag:function(e){return function(e,t){var n;for(e.stopPropagation(),n=0;n<C.length;n++){var c=document.getElementById(C[n]).getBoundingClientRect();if(e.clientX>c.left&&e.clientX<c.right){console.log("In "+k[n]),t.status==k[n]?(O(b),console.log("no change")):N(t,k[n]);break}}}(e,c)},children:Object(h.jsx)("div",{className:"listItem",children:Object(h.jsx)(V,{issue:c,onSubmit:w,onDone:function(){return""}})})}))};for(c.s();!(t=c.n()).done;)s()}catch(r){c.e(r)}finally{c.f()}return n}if(Object(a.useEffect)((function(){I()}),[]),c)return Object(h.jsxs)("div",{children:["Error: ",c.message]});if(o){var P=_(b),E=_(x),T=_(y);return Object(h.jsxs)("div",{style:{padding:"1.5em"},children:[Object(h.jsx)(V,{projectId:e.projectId,onSubmit:function(t){return fetch("/v1/issues",{method:"POST",headers:{Accept:"application/json","Content-Type":"application/json"},body:JSON.stringify({issue:{name:t.name,desc:t.desc,issue_type:t.issue_type,bug_trace:t.bug_trace,status:"to do",project:e.projectId}})})},issue:null,onDone:I}),Object(h.jsx)("br",{}),Object(h.jsxs)("div",{id:"board",className:"row",style:{position:"relative",overflow:"auto",padding:"0"},children:[Object(h.jsxs)("div",{id:C[0],className:"column",children:["To Do",P]}),Object(h.jsxs)("div",{id:C[1],className:"column",children:["In progress",E]}),Object(h.jsxs)("div",{id:C[2],className:"column",children:["Done",T]})]})]})}return Object(h.jsx)(z,{})};var Z,$=function(e){var t=Object(a.useState)(null),n=Object(d.a)(t,2),c=n[0],s=n[1],r=Object(a.useState)(!1),i=Object(d.a)(r,2),o=i[0],u=i[1],j=Object(a.useState)(null),l=Object(d.a)(j,2),b=l[0],O=l[1];return Object(a.useEffect)((function(){fetch("/v1/projects/byid/".concat(e.match.params.projectId)).then((function(e){return e.json()})).then((function(e){O(e),u(!0)}),(function(e){u(!0),s(e)}))}),[]),c?Object(h.jsxs)("div",{children:["Error: ",c.message]}):o?Object(h.jsxs)("div",{children:[Object(h.jsx)("h2",{style:{"padding-top":"0.5em","padding-left":"1em"},children:b.name}),Object(h.jsx)(Y,{projectId:b.id})]}):Object(h.jsx)(z,{})},ee=n(49),te=Object(p.default)(ee.a)(Z||(Z=Object(b.a)(["\n  color: #fff;\n  transition: transform 200ms ease-in-out;\n  &:hover {\n    transform: scale(1.2);\n  }\n"])));var ne,ce=function(e){return Object(h.jsxs)("div",{children:[Object(h.jsx)("h3",{children:"Login: "}),Object(h.jsx)("a",{href:window.location.origin+"/auth/google",children:Object(h.jsx)(te,{size:36,style:{cursor:"pointer",textAlign:"center"}})})]})},se=Object(p.default)(H.BaseModalBackground)(ne||(ne=Object(b.a)(["\n  opacity: ",";\n  transition: all 0.3s ease-in-out;\n"])),(function(e){return e.opacity})),re=function(){return Object(h.jsx)("div",{children:Object(h.jsxs)(H.ModalProvider,{backgroundComponent:se,children:[Object(h.jsx)(T,{}),Object(h.jsx)(A.a,{children:Object(h.jsxs)(B.c,{children:[Object(h.jsx)(B.a,{path:"/projects/:projectId",component:$}),Object(h.jsx)(B.a,{path:"/login",component:ce}),Object(h.jsx)(B.a,{path:"/",children:Object(h.jsx)(R,{})})]})})]})})},ie=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,78)).then((function(t){var n=t.getCLS,c=t.getFID,s=t.getFCP,r=t.getLCP,i=t.getTTFB;n(e),c(e),s(e),r(e),i(e)}))};l.a.render(Object(h.jsx)(u.a.StrictMode,{children:Object(h.jsx)(re,{})}),document.getElementById("root")),ie()}},[[77,1,2]]]);
//# sourceMappingURL=main.262ef5f7.chunk.js.map