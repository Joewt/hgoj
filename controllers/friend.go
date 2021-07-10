package controllers
type FriendController struct {
	BaseController

}
func (c *FriendController) Get() {
	c.TplName = "friend.html"
}
