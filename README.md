## 目录结构

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

