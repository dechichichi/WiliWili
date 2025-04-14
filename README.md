# 目录结构

```
├─app
│  ├─comment
│  │  ├─controllers
│  │  │  └─rpc
│  │  │      └─pack
│  │  ├─domain
│  │  │  ├─model
│  │  │  ├─repository
│  │  │  └─service
│  │  ├─infrastructure
│  │  │  ├─mysql
│  │  │  └─redis
│  │  └─usecase
│  ├─gateway
│  │  ├─handler
│  │  │  └─api
│  │  │      ├─comment
│  │  │      ├─like
│  │  │      ├─user
│  │  │      └─video
│  │  ├─model
│  │  │  └─api
│  │  │      ├─comment
│  │  │      ├─like
│  │  │      ├─user
│  │  │      └─video
│  │  ├─mw
│  │  ├─pack
│  │  ├─router
│  │  │  └─api
│  │  │      ├─comment
│  │  │      ├─like
│  │  │      ├─user
│  │  │      └─video
│  │  └─rpc
│  ├─like
│  │  ├─controllers
│  │  │  └─rpc
│  │  │      └─pack
│  │  ├─domain
│  │  │  ├─model
│  │  │  ├─repository
│  │  │  └─service
│  │  ├─infrastructure
│  │  │  ├─mysql
│  │  │  └─redis
│  │  └─usecase
│  ├─user
│  │  ├─controllers
│  │  │  └─rpc
│  │  │      └─pack
│  │  ├─domain
│  │  │  ├─model
│  │  │  ├─repository
│  │  │  └─service
│  │  ├─infrastructure
│  │  │  ├─mysql
│  │  │  └─redis
│  │  └─usecase
│  └─video
│      ├─controllers
│      │  └─rpc
│      │      └─pack
│      ├─domain
│      │  ├─model
│      │  ├─repository
│      │  └─service
│      ├─infrastructure
│      │  ├─mysql
│      │  └─redis
│      └─usecase
├─cmd
│  ├─chat
│  ├─comment
│  ├─gateway
│  ├─like
│  ├─user
│  └─video
├─config
│  └─sql
├─docker
│  ├─env
│  └─script
├─idl
│  └─api
├─kitex_gen
│  ├─comment
│  │  └─commentservice
│  ├─like
│  │  └─likeservice
│  ├─model
│  ├─user
│  │  └─userservice
│  └─video
│      └─videoservice
├─output
│  ├─gateway
│  └─user
└─pkg
    ├─base
    │  ├─client
    │  └─context
    ├─constants
    ├─errno
    ├─middleware
    └─utils
```

# 项目部署

- 更改config\config.yaml文件

- 在linux或者mac环境下运行

- 初始化环境：在命令行输入（在此之前保证docker启动）

  ```
  make env-up
  ```

- 启动项目

  ```
  make +相应项目名字（在cmd文件夹下）
  ```

  值得注意的是 chat模块并不能直接make chat 

  若要启动chat功能 应该

  ```
  cd cmd
  cd chat
  go run ./
  ```

  

