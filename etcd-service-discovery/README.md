### Question
how to register a miro-server in etcd and gateway(master) discovery it

### Implement

- mico-server and master confirm a etcd key
- mico-server put the key self's server information with TTL and repeated, TTL is use to delete key before restart gateway and mico-server
- gateway(master) watch the key and distinguish the action is "put" or "delete" and to do respond something 

### [Reference](https://gitbook.cn/books/5bb037728f7d8b7e900ff2d7/index.html)
#### 服务发现（Service Discovery）要解决的是分布式系统中最常见的问题之一，即在同一个分布式集群中的进程或服务如何才能找到对方并建立连接。服务发现的实现原理如下：

1. 存在一个高可靠、高可用的中心配置节点：基于 Ralf 算法的 Etcd 天然支持，不必多解释。
2. 服务提供方会持续的向配置节点注册服务：用户可以在 Etcd 中注册服务，并且对注册的服务配置租约，定时续约以达到维持服务的目的（一旦停止续约，对应的服务就会失效）。
3. 服务的调用方会持续的读取中心配置节点的配置并修改本机配置，然后 reload 服务：服务提供方在 Etcd 指定的目录（前缀机制支持）下注册的服务，服务调用方在对应的目录下查服务。通过 Watch 机制，服务调用方还可以监测服务的变化。