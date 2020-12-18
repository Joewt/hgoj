package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "BlogAddGet",
            Router: "/admin/art/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "BlogAddPost",
            Router: "/admin/art/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "BlogList",
            Router: "/admin/art/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:BlogController"],
        beego.ControllerComments{
            Method: "BlogIndex",
            Router: "/article/:artid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:CeinfoController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:CeinfoController"],
        beego.ControllerComments{
            Method: "Ceinfo",
            Router: "/ceinfo/:sid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:CeinfoController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:CeinfoController"],
        beego.ControllerComments{
            Method: "CeinfoContest",
            Router: "/contest/ceinfo/:sid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestPage",
            Router: "/contest/:page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestAddGet",
            Router: "/contest/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestAddPost",
            Router: "/contest/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestCid",
            Router: "/contest/cid/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestList",
            Router: "/contest/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestListPage",
            Router: "/contest/list/:page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestStatus",
            Router: "/contest/status/cid/:cid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestStatusPage",
            Router: "/contest/status/cid/:cid/:page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestUpdatePost",
            Router: "/contest/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestUpdateStatus",
            Router: "/contest/updatestatus",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ContestController"],
        beego.ControllerComments{
            Method: "ContestRank",
            Router: "/contestrank/cid/:cid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Indexs",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexAdmin",
            Router: "/admin",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexContest",
            Router: "/contest",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ContestUpdate",
            Router: "/contest/update/:cid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexFaqs",
            Router: "/faqs",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: "/index",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexUser",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexProblemset",
            Router: "/problemset",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexReg",
            Router: "/reg",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexStatus",
            Router: "/status",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "RejudgePost",
            Router: "/admin/rejudge",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "Rejudge",
            Router: "/admin/rejudge",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemContest",
            Router: "/contest/problem/:id/:cid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "DownloadTestData",
            Router: "/download/testdata",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "Problem",
            Router: "/problem/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemAdd",
            Router: "/problem/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemAddPost",
            Router: "/problem/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemDel",
            Router: "/problem/del",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemEdit",
            Router: "/problem/edit/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemTestDataEdit",
            Router: "/problem/editdata/:pid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ExInport",
            Router: "/problem/exinport",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "Export",
            Router: "/problem/export",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "Fileupload",
            Router: "/problem/fileupload",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "Inport",
            Router: "/problem/inport",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemList",
            Router: "/problem/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemListPage",
            Router: "/problem/list/:page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemUpdate",
            Router: "/problem/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemUpdateStatus",
            Router: "/problem/updatestatus",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "ProblemSetPage",
            Router: "/problemset/:page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:ProblemController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: "/upload/fileupload/:pid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:SolutionController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:SolutionController"],
        beego.ControllerComments{
            Method: "StatusPage",
            Router: "/status/:page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:SolutionController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:SolutionController"],
        beego.ControllerComments{
            Method: "Submit",
            Router: "/submit",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChangePassword",
            Router: "/admin/changepwd",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChangePwd",
            Router: "/admin/changepwd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "PermissionsAdd",
            Router: "/admin/permissions/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "PermissionsAddPost",
            Router: "/admin/permissions/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "PermissionsList",
            Router: "/admin/permissions/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserGen",
            Router: "/admin/user/generate",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserGenPost",
            Router: "/admin/user/generate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserForgotPwd",
            Router: "/forgotpwd",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendEmailForgot",
            Router: "/forgotpwd/sendemail",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "Profile",
            Router: "/profile",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "Profile2",
            Router: "/profile/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserUpdatePost",
            Router: "/profile/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserUpdateGet",
            Router: "/profile/update",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserList",
            Router: "/user/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserListPage",
            Router: "/user/list/:page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yinrenxin/hgoj/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserReg",
            Router: "/user/reg",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
