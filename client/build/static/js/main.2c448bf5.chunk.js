(this.webpackJsonpclient=this.webpackJsonpclient||[]).push([[0],{159:function(e,t,n){},160:function(e,t,n){},178:function(e,t,n){"use strict";n.r(t);var a=n(0),i=n.n(a),c=n(54),r=n.n(c),s=(n(159),n(21)),o=(n(160),n(36)),l=n.n(o),j=n(62),d=n(82),h=n(83),b=n(90),u=n(88),O=n(138),x=n(197),m=n(143),p=n(199),g=n(51),f=n(2),v=function(e){Object(b.a)(n,e);var t=Object(u.a)(n);function n(e){return Object(d.a)(this,n),t.call(this,e)}return Object(h.a)(n,[{key:"render",value:function(){return Object(f.jsxs)(x.a,{children:[Object(f.jsxs)(x.a.Content,{children:[Object(f.jsx)(m.a,{floated:"right",size:"mini",src:this.props.image}),Object(f.jsx)(x.a.Header,{children:this.props.name}),Object(f.jsx)(x.a.Meta,{children:this.props.link}),Object(f.jsx)(x.a.Description,{children:this.props.description})]}),Object(f.jsx)(x.a.Content,{extra:!0,centered:!0,children:Object(f.jsxs)(p.a.Group,{widths:"3",children:[Object(f.jsxs)(p.a,{floated:"left",animated:"fade",as:"a",href:"https://"+this.props.link,children:[Object(f.jsx)(p.a.Content,{hidden:!0,children:"View"}),Object(f.jsx)(p.a.Content,{visible:!0,children:Object(f.jsx)(g.a,{name:"eye"})})]}),Object(f.jsxs)(p.a,{animated:"fade",children:[Object(f.jsx)(p.a.Content,{hidden:!0,children:"Edit"}),Object(f.jsx)(p.a.Content,{visible:!0,children:Object(f.jsx)(g.a,{name:"edit"})})]}),Object(f.jsxs)(p.a,{floated:"left",animated:"fade",onClick:this.handleDelete,children:[Object(f.jsx)(p.a.Content,{hidden:!0,children:"Delete"}),Object(f.jsx)(p.a.Content,{visible:!0,children:Object(f.jsx)(g.a,{name:"delete"})})]})]})})]})}}]),n}(a.Component),y=n(191),S=n(200),k=n(192),C=n(201),w=n(195),I=n(202),T=n(203),z=n(193),A=n(196),P=Object(O.createMedia)({breakpoints:{mobile:0,tablet:768,computer:1024}}),L=P.MediaContextProvider,B=P.Media,F=function(e){var t=e.mobile;return Object(f.jsxs)(y.a,{text:!0,children:[Object(f.jsx)(S.a,{as:"h1",content:"Welcome to your bookmark manager",inverted:!0,style:{fontSize:t?"2em":"4em",fontWeight:"normal",marginBottom:0,marginTop:t?"1.5em":"3em"}}),Object(f.jsx)(S.a,{as:"h2",content:"Create, save , view and delete your bookmarks all in one place",inverted:!0,style:{fontSize:t?"1.5em":"1.7em",fontWeight:"normal",marginTop:t?"0.5em":"1.5em"}}),Object(f.jsxs)(p.a,{primary:!0,size:"huge",children:["Get Started",Object(f.jsx)(g.a,{name:"right arrow"})]})]})},M=function(e){Object(b.a)(n,e);var t=Object(u.a)(n);function n(){var e;Object(d.a)(this,n);for(var a=arguments.length,i=new Array(a),c=0;c<a;c++)i[c]=arguments[c];return(e=t.call.apply(t,[this].concat(i))).state={},e.hideFixedMenu=function(){return e.setState({fixed:!1})},e.showFixedMenu=function(){return e.setState({fixed:!0})},e.logout=function(){localStorage.removeItem("token"),window.location.reload(!1)},e}return Object(h.a)(n,[{key:"render",value:function(){var e=this.props.children,t=this.state.fixed;return Object(f.jsxs)(B,{greaterThan:"mobile",children:[Object(f.jsx)(k.a,{once:!1,onBottomPassed:this.showFixedMenu,onBottomPassedReverse:this.hideFixedMenu,children:Object(f.jsxs)(C.a,{inverted:!0,textAlign:"center",style:{minHeight:700,padding:"1em 0em"},vertical:!0,children:[Object(f.jsx)(w.a,{fixed:t?"top":null,inverted:!t,pointing:!t,secondary:!t,size:"large",children:Object(f.jsxs)(y.a,{children:[Object(f.jsx)(w.a.Item,{as:"a",active:!0,children:"Home"}),Object(f.jsx)(w.a.Item,{as:"a",children:"Work"}),Object(f.jsx)(w.a.Item,{as:"a",children:"Company"}),Object(f.jsx)(w.a.Item,{as:"a",children:"Careers"}),Object(f.jsx)(w.a.Item,{position:"right",onClick:this.logout,children:Object(f.jsx)(p.a,{as:"a",inverted:!t,children:"Log Out"})})]})}),Object(f.jsx)(F,{})]})}),e]})}}]),n}(a.Component),D=function(e){Object(b.a)(n,e);var t=Object(u.a)(n);function n(){var e;Object(d.a)(this,n);for(var a=arguments.length,i=new Array(a),c=0;c<a;c++)i[c]=arguments[c];return(e=t.call.apply(t,[this].concat(i))).state={},e.handleSidebarHide=function(){return e.setState({sidebarOpened:!1})},e.handleToggle=function(){return e.setState({sidebarOpened:!0})},e}return Object(h.a)(n,[{key:"render",value:function(){var e=this.props.children,t=this.state.sidebarOpened;return Object(f.jsx)(B,{as:I.a.Pushable,at:"mobile",children:Object(f.jsxs)(I.a.Pushable,{children:[Object(f.jsxs)(I.a,{as:w.a,animation:"overlay",inverted:!0,onHide:this.handleSidebarHide,vertical:!0,visible:t,children:[Object(f.jsx)(w.a.Item,{as:"a",active:!0,children:"Home"}),Object(f.jsx)(w.a.Item,{as:"a",children:"Work"}),Object(f.jsx)(w.a.Item,{as:"a",children:"Company"}),Object(f.jsx)(w.a.Item,{as:"a",children:"Careers"}),Object(f.jsx)(w.a.Item,{as:"a",children:"Log Out"})]}),Object(f.jsxs)(I.a.Pusher,{dimmed:t,children:[Object(f.jsxs)(C.a,{inverted:!0,textAlign:"center",style:{minHeight:350,padding:"1em 0em"},vertical:!0,children:[Object(f.jsx)(y.a,{children:Object(f.jsxs)(w.a,{inverted:!0,pointing:!0,secondary:!0,size:"large",children:[Object(f.jsx)(w.a.Item,{onClick:this.handleToggle,children:Object(f.jsx)(g.a,{name:"sidebar"})}),Object(f.jsxs)(w.a.Item,{position:"right",children:[Object(f.jsx)(p.a,{as:"a",inverted:!0,children:"Log in"}),Object(f.jsx)(p.a,{as:"a",inverted:!0,style:{marginLeft:"0.5em"},children:"Sign Up"})]})]})}),Object(f.jsx)(F,{mobile:!0})]}),e]})]})})}}]),n}(a.Component),H=function(e){var t=e.children;return Object(f.jsxs)(L,{children:[Object(f.jsx)(M,{children:t}),Object(f.jsx)(D,{children:t})]})};var N=function(e){var t=Object(a.useState)([]),n=Object(s.a)(t,2),i=n[0],c=n[1],r=Object(a.useState)(!1),o=Object(s.a)(r,2),d=o[0],h=o[1];return console.log("Billy",e.token),d||function(){var t=Object(j.a)(l.a.mark((function t(){var n;return l.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return n=e.token,t.next=3,fetch("/bookmarks",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({token:n})}).then((function(e){return e.json()})).then((function(e){console.log("Billy",e),c(e),h(!0)}),(function(e){}));case 3:case"end":return t.stop()}}),t)})));return function(){return t.apply(this,arguments)}}()(e),Object(f.jsxs)(H,{children:[Object(f.jsx)(C.a,{style:{padding:"8em 0em"},vertical:!0,children:Object(f.jsx)(T.a,{container:!0,stackable:!0,verticalAlign:"middle",children:Object(f.jsx)(T.a.Row,{children:Object(f.jsx)(T.a.Column,{width:16,children:Object(f.jsx)(x.a.Group,{children:i.map((function(e){return Object(f.jsx)(v,{name:e.Name,link:e.Link,image:e.Image,description:e.Description})}))})})})})}),Object(f.jsx)(C.a,{style:{padding:"0em"},vertical:!0,children:Object(f.jsx)(T.a,{celled:"internally",columns:"equal",stackable:!0,children:Object(f.jsxs)(T.a.Row,{textAlign:"center",children:[Object(f.jsxs)(T.a.Column,{style:{paddingBottom:"5em",paddingTop:"5em"},children:[Object(f.jsx)(S.a,{as:"h3",style:{fontSize:"2em"},children:'"What a Company"'}),Object(f.jsx)("p",{style:{fontSize:"1.33em"},children:"That is what they all say about us"})]}),Object(f.jsxs)(T.a.Column,{style:{paddingBottom:"5em",paddingTop:"5em"},children:[Object(f.jsx)(S.a,{as:"h3",style:{fontSize:"2em"},children:'"I shouldn\'t have gone with their competitor."'}),Object(f.jsxs)("p",{style:{fontSize:"1.33em"},children:[Object(f.jsx)(m.a,{avatar:!0,src:"/images/avatar/large/nan.jpg"}),Object(f.jsx)("b",{children:"Nan"})," Chief Fun Officer Acme Toys"]})]})]})})}),Object(f.jsx)(C.a,{style:{padding:"8em 0em"},vertical:!0,children:Object(f.jsxs)(y.a,{text:!0,children:[Object(f.jsx)(S.a,{as:"h3",style:{fontSize:"2em"},children:"Breaking The Grid, Grabs Your Attention"}),Object(f.jsx)("p",{style:{fontSize:"1.33em"},children:"Instead of focusing on content creation and hard work, we have learned how to master the art of doing nothing by providing massive amounts of whitespace and generic content that can seem massive, monolithic and worth your attention."}),Object(f.jsx)(p.a,{as:"a",size:"large",children:"Read More"}),Object(f.jsx)(z.a,{as:"h4",className:"header",horizontal:!0,style:{margin:"3em 0em",textTransform:"uppercase"},children:Object(f.jsx)("a",{href:"#",children:"Case Studies"})}),Object(f.jsx)(S.a,{as:"h3",style:{fontSize:"2em"},children:"Did We Tell You About Our Bananas?"}),Object(f.jsx)("p",{style:{fontSize:"1.33em"},children:"Yes I know you probably disregarded the earlier boasts as non-sequitur filler content, but it's really true. It took years of gene splicing and combinatory DNA research, but our bananas can really dance."}),Object(f.jsx)(p.a,{as:"a",size:"large",children:"I'm Still Quite Interested"})]})}),Object(f.jsx)(C.a,{inverted:!0,vertical:!0,style:{padding:"5em 0em"},children:Object(f.jsx)(y.a,{children:Object(f.jsx)(T.a,{divided:!0,inverted:!0,stackable:!0,children:Object(f.jsxs)(T.a.Row,{children:[Object(f.jsxs)(T.a.Column,{width:3,children:[Object(f.jsx)(S.a,{inverted:!0,as:"h4",content:"About"}),Object(f.jsxs)(A.a,{link:!0,inverted:!0,children:[Object(f.jsx)(A.a.Item,{as:"a",children:"Sitemap"}),Object(f.jsx)(A.a.Item,{as:"a",children:"Contact Us"}),Object(f.jsx)(A.a.Item,{as:"a",children:"Religious Ceremonies"}),Object(f.jsx)(A.a.Item,{as:"a",children:"Gazebo Plans"})]})]}),Object(f.jsxs)(T.a.Column,{width:3,children:[Object(f.jsx)(S.a,{inverted:!0,as:"h4",content:"Services"}),Object(f.jsxs)(A.a,{link:!0,inverted:!0,children:[Object(f.jsx)(A.a.Item,{as:"a",children:"Banana Pre-Order"}),Object(f.jsx)(A.a.Item,{as:"a",children:"DNA FAQ"}),Object(f.jsx)(A.a.Item,{as:"a",children:"How To Access"}),Object(f.jsx)(A.a.Item,{as:"a",children:"Favorite X-Men"})]})]}),Object(f.jsxs)(T.a.Column,{width:7,children:[Object(f.jsx)(S.a,{as:"h4",inverted:!0,children:"Footer Header"}),Object(f.jsx)("p",{children:"Extra space for a call to action inside the footer that could help re-engage users."})]})]})})})})]})},W=n(194),E=n(198),G=n(10),R=n(38);var J=function(e){var t=e.setToken,n=Object(a.useState)(),i=Object(s.a)(n,2),c=i[0],r=i[1],o=Object(a.useState)(),d=Object(s.a)(o,2),h=d[0],b=d[1],u=Object(a.useState)(!1),O=Object(s.a)(u,2),x=O[0],m=O[1],g=Object(G.f)(),v=function(){var e=Object(j.a)(l.a.mark((function e(n){return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return n.preventDefault(),e.next=3,fetch("/login",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({email:c,password:h})}).then((function(e){if(!e.ok)throw Error(e.statusText);return e.json()})).then((function(e){var n=e.token;console.log(n),g.push("/"),localStorage.setItem("token",n),t(n)}),(function(e){console.log("error"),m(!0)})).catch((function(e){m(!0),console.log("error")}));case 3:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}();return Object(f.jsx)(T.a,{textAlign:"center",style:{height:"100vh"},verticalAlign:"middle",children:Object(f.jsxs)(T.a.Column,{style:{maxWidth:450},children:[Object(f.jsx)(S.a,{as:"h2",color:"teal",textAlign:"center",children:"Log-in to your account"}),Object(f.jsxs)(W.a,{size:"large",onSubmit:v,error:x,children:[Object(f.jsxs)(C.a,{stacked:!0,children:[Object(f.jsx)(W.a.Input,{fluid:!0,icon:"user",iconPosition:"left",placeholder:"E-mail address",onChange:function(e){return r(e.target.value)}}),Object(f.jsx)(W.a.Input,{fluid:!0,icon:"lock",iconPosition:"left",placeholder:"Password",type:"password",onChange:function(e){return b(e.target.value)}}),Object(f.jsx)(p.a,{color:"teal",fluid:!0,size:"large",children:"Login"})]}),Object(f.jsx)(E.a,{error:!0,header:"Login failed",children:"Login failed with an error."})]}),Object(f.jsxs)(E.a,{children:["New to us? ",Object(f.jsx)(R.b,{to:"/signup",children:"Sign Up"})]})]})})};var U=function(){var e=Object(a.useState)(),t=Object(s.a)(e,2),n=t[0],i=t[1],c=Object(a.useState)(),r=Object(s.a)(c,2),o=r[0],d=r[1],h=Object(a.useState)(),b=Object(s.a)(h,2),u=b[0],O=b[1],x=Object(a.useState)(!1),m=Object(s.a)(x,2),g=m[0],v=m[1],y=Object(a.useState)(!1),k=Object(s.a)(y,2),w=k[0],I=k[1],z=function(){var e=Object(j.a)(l.a.mark((function e(t){return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t.preventDefault(),e.next=3,fetch("/signup",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({name:u,email:n,password:o})}).then((function(e){200!=e.status?(I(!0),v(!1)):(I(!1),v(!0))}),(function(e){v(!1),I(!0)}));case 3:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}();return Object(f.jsx)(T.a,{textAlign:"center",style:{height:"100vh"},verticalAlign:"middle",children:Object(f.jsxs)(T.a.Column,{style:{maxWidth:450},children:[Object(f.jsx)(S.a,{as:"h2",color:"teal",textAlign:"center",children:"Create your account"}),Object(f.jsx)(W.a,{size:"large",onSubmit:z,success:g,error:w,children:Object(f.jsxs)(C.a,{stacked:!0,children:[Object(f.jsx)(W.a.Input,{fluid:!0,icon:"id badge",iconPosition:"left",placeholder:"Username",onChange:function(e){return O(e.target.value)}}),Object(f.jsx)(W.a.Input,{fluid:!0,icon:"user",iconPosition:"left",placeholder:"E-mail address",onChange:function(e){return i(e.target.value)}}),Object(f.jsx)(W.a.Input,{fluid:!0,icon:"lock",iconPosition:"left",placeholder:"Password",type:"password",onChange:function(e){return d(e.target.value)}}),Object(f.jsxs)(E.a,{success:!0,header:"Signup Complete",children:["Successfully signed up! ",Object(f.jsx)(R.b,{to:"/login",children:"Log In"}),"."]}),Object(f.jsxs)(E.a,{error:!0,header:"Signup failed",children:["User with the same email already exists. ",Object(f.jsx)(R.b,{to:"/login",children:"Log In"}),"\xa0instead"]}),Object(f.jsx)(p.a,{color:"teal",fluid:!0,size:"large",children:"Signup"})]})}),Object(f.jsxs)(E.a,{children:["Already have an account with us? ",Object(f.jsx)(R.b,{to:"/login",children:"Log In"})]})]})})};n(177);var Y=function(){var e=Object(a.useState)(localStorage.getItem("token")),t=Object(s.a)(e,2),n=t[0],i=t[1];return n?Object(f.jsx)(R.a,{children:Object(f.jsx)(G.c,{children:Object(f.jsx)(G.a,{path:"/",children:Object(f.jsx)(N,{token:n})})})}):Object(f.jsxs)(R.a,{children:[Object(f.jsx)(G.a,{path:"/signup",children:Object(f.jsx)(U,{})}),Object(f.jsx)(G.a,{path:"/",children:Object(f.jsx)(J,{setToken:i})})]})},q=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,204)).then((function(t){var n=t.getCLS,a=t.getFID,i=t.getFCP,c=t.getLCP,r=t.getTTFB;n(e),a(e),i(e),c(e),r(e)}))};r.a.render(Object(f.jsx)(i.a.StrictMode,{children:Object(f.jsx)(Y,{})}),document.getElementById("root")),q()}},[[178,1,2]]]);
//# sourceMappingURL=main.2c448bf5.chunk.js.map