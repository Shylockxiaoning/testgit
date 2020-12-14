#### 简易图书管理

##### 需求

* 操作级别

* 基础数据维护（类别、单价、库存量）
* 出库信息完整，去向明确
* 可导出为excel文件
* web服务

##### 涉及的语言和第三方库
 ```text
  Go、Gin、Vipper、Zap、Sqlx\Sql、Yaml、SnowFlake、MySQL
  详细说明：
  GO			: 核心语言
  Gin			: web服务框架
  Vipper		: 监控配置信息
  Zap			: 日志库
  Sqlx\Sql		: 链接数据库
  Yaml			: 配置文件信息
  SnowFlake		: 唯一ID
  MySQL			: 数据库
 ```

##### 数据库结构

**分析**

* 需要用户名`username`、密码`password`、职位`position`

* 需要书名`title`、作者`author`、发行时间`issue_date`、库存量`inventory`、简介`synopsis`、单价`price`、类别`category`

```mysql
创建数据库
CREATE DATABASE sap;
创建数据表
用户数据表
create table `admin`(
    -> `id` bigint(20) not null auto_increment,
    -> `username` varchar(64) not null,
    -> `password` varchar(64) not null,
    -> `position` varchar(20) not null,
    -> primary key(`id`)
    -> )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
信息数据表
 create table `godown`(
    ->  `id` bigint(20) not null auto_increment,
    -> `title` varchar(20) not null,
    ->  `author` varchar(20) not null,
    ->  `issue_date` varchar(40) not null,
    -> `inventory` bigint(20) not null,
    -> `synopsis` varchar(100) not null,
    ->  `price` bigint(20) not null,
    ->  primary key (`id`)
    ->  )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

##### 文件

```tex
Controller		:控制部分
Dao				:数据库相关
Logger			:日志相关
Logic			:逻辑控制
Modul			:结构体相关
Router			:路由相关
main			:主函数/程序入口
config			:配置文件
```

##### 流程

* **登陆**
  * 链接数据库，返回web响应
  * 用户是否存在，不存在引导至注册界面，存在校验密码
  * 登陆成功---->返回`cookie`值

* **增删改查**

  * 权限分为:操作员-----组长-----店长

    ```text
    操作员 ---->可进行仓库的录入、查看
    组长   ---->可进行仓库的录入、查看，可进行库存为0的商品删除
    店长	 ---->可进行仓库的录入、查看、修改单价、0库存商品删除
    ```

    ![image-20201210160832406](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20201210160832406.png)

  * 增---->所有用户可进行-------可添加附加功能:高级权限确认是否增加成功-未加入

  ![image-20201210225528217](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20201210225528217.png)
  
  ![image-20201210225656530](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20201210225656530.png)
  
  * 删---->当库存为0时，所有人可进行操作，否则只有更高权限可修改
  * 改---->用户进行改单价时，判断权限