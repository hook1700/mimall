package admin

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"mimall/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	userinfo, ok := c.GetSession("userinfo").(models.Manager)
	if ok {
		//获取用户名
		c.Data["username"] = userinfo.Username


		//1、获取角色id
		roleId := userinfo.RoleId

		//2、获取全部的权限 (排序)，忘记的话参考：https://gorm.io/zh_CN/docs/preload.html
		access := []models.Access{}
		models.DB.Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort desc").Where("module_id=?", 0).Find(&access)

		//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", roleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
		for i := 0; i < len(access); i++ {
			if _, ok := roleAccessMap[access[i].Id]; ok {
				access[i].Checked = true
			}
			for j := 0; j < len(access[i].AccessItem); j++ {
				if _, ok := roleAccessMap[access[i].AccessItem[j].Id]; ok {
					access[i].AccessItem[j].Checked = true
				}
			}
		}

		//5、渲染权限数据以及角色 Id
		c.Data["accessList"] = access

		c.Data["isSuper"] = userinfo.IsSuper


	}
	c.TplName = "admin/main/index.html"
}

func (c *MainController) Welcome() {
	c.TplName = "admin/main/welcome.html"
}

