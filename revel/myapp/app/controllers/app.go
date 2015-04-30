package controllers//控制器

import "github.com/revel/revel"

type App struct {//导入 *revel.Controller//导入一个控制器
	*revel.Controller
}

func (c App) Index() revel.Result {//为控制器注册Index 方法 名称必须是导入的 即为大写首字母
	return c.Render()//控制器方法 返回一个 revel.Result  c.Rander 渲染一个模板 并且可以 传递参数到模板
}
