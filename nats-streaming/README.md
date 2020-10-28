### Question
Based on the info from https://github.com/nats-io/nats-streaming-server, implement three clients in GoLang.

1.  3 clients exchange information through one instance of nats-streaming-server

2.  Client 1 transfer a word into message-bus based on user input

3.  Client 2 change the word and send back to message-bus

4.  Client 3 receive the swap some letter in the changed word and send back to message-bus

5.  Client 1 print the final word on screen

### Implement
#### Thinking
- 在调试过程中遇到的一个问题就是，c1->c2->c3->c1这样一个环形的发布订阅环会导致死锁，三个客户端互相depend，
所以c1的发布用了异步，然后调试一切都ok了。
- 代码跑起来依赖一个nats-streaming-server服务，此服务内嵌了nats服务，于是可以使用客户端go-nats-streaming与他们进行交互。
- 另外设置订阅为Durable以避免客户端关闭丢失发布者在订阅者down过程中发布的消息。

#### Run
```bash
nats-streaming-server &
go run c1.go hello
go run c2.go 
go run c3.go
```
![img](https://github.com/FrankHitman/playGo/blob/master/nats-streaming/nats-streaming-paly.png)
