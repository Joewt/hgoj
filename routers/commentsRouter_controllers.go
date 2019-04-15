package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "BlogAddGet",
            Router: `/article/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "ContestAdd",
            Router: `/article/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "BlogIndex",
            Router: `/article/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "ContestList",
            Router: `/article/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:CeinfoController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:CeinfoController"],
        beego.ControllerComments{
            Method: "Ceinfo",
            Router: `/ceinfo/:k`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestAddGet",
            Router: `/contest/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestAddPost",
            Router: `/contest/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestCid",
            Router: `/contest/cid/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestList",
            Router: `/contest/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexAdmin",
            Router: `/admin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexContest",
            Router: `/contest`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexFaqs",
            Router: `/faqs`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexUser",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexProblemset",
            Router: `/problemset`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexReg",
            Router: `/reg`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexSkinConfig",
            Router: `/skin-config.html`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexStatus",
            Router: `/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemContest",
            Router: `/contest/problem/:id/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "Problem",
            Router: `/problem/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemAdd",
            Router: `/problem/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemAddPost",
            Router: `/problem/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemDel",
            Router: `/problem/del`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemEdit",
            Router: `/problem/edit/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemList",
            Router: `/problem/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemUpdate",
            Router: `/problem/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:SolutionController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:SolutionController"],
        beego.ControllerComments{
            Method: "Submit",
            Router: `/submit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "Profile",
            Router: `/profile/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserList",
            Router: `/user/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserReg",
            Router: `/user/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
