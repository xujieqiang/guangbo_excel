go语言开发
为广播系统生成 excel 方便导入。


# 主要实现的功能：

* 管理所有的铃声模板
* 管理铃声分组区域

* 管理方案
* 生成铃声任务列表和管理

* 计算JobData
* 铃声终端管理

生成excel表格，拷贝和生成原来方案的铃声任务，集中式管理，便捷易用。



---



# 项目采用的技术

golang+gin

框架是mcv的模式

routers集中式管理，

models 模型

db 是数据库操作  ，用gorm.io 来操作数据库  

db/guangbo.db 是sqlite3数据库

methods 是相当于控制器，分别对每个表采用一个文件来代表  alltask.go.....

static 静态文件 js  css  pic 等

templates 是views

生成的文件是  data.xlsx


> 程序范例

```

func main(){
	fmt.Println("hello")
}
```
