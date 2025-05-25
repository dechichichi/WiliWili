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

# 监控与可观测性（Prometheus/Alloy/Grafana/Loki）

## QPS 业务指标
- user 服务已埋点 QPS 指标（user_handler_qps），通过 OpenTelemetry metrics 导出。
- Alloy/otel-collector 采集后，Prometheus 可抓取，Grafana 可视化。
- 推荐在 Grafana 添加 Prometheus 数据源，搜索 `user_handler_qps`，可按 method 维度展示。

## 日志采集
- user 服务日志已标准化为 json 格式，Loki/promtail 可直接采集。
- 在 Grafana 添加 Loki 数据源后，可通过 `{job="user"}` 检索日志。

## 参考
- 可根据 docker-compose.yml 启动 Prometheus、Grafana、Loki、Alloy。
- 其他服务可参考 user 服务埋点方式扩展。

  

