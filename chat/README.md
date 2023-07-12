# 星火认知大模型服务


## 关于Web接口的说明:

- 必须符合 websocket 协议规范（rfc6455）。
- websocket握手成功后用户在60秒内没有发送请求数据，服务侧会主动断开。
- 本接口默认采用短链接的模式，即接口每次将结果完整返回给用户后会主动断开链接，用户在下次发送请求的时候需要重新握手链接。

ref: 
- https://www.xfyun.cn/doc/spark/%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E.html
- https://www.xfyun.cn/doc/spark/Web.html

