
<a href="#chinese">点击此处直达中文版 </a>


# SnapScale Explorer Middleware

SnapScale Explorer Middleware is a Snapscale Explorer gateway.

This software provides substantial statistical data and dimensions compared to EOS API.

It is part of Snapscale Explorer, for which it serves as a gateway. Deploying SnapScale Explorer Middleware is mandatory for anyone looking to run SnapScale Explorer.

## Prerequisites

Start a SnapScale node and make sure that SnapScale is using a copy of Mongo save history. Create collection and index.

```sh
db.createCollection("daily");
db.transactions.createIndex({"createdAt":1},{"background":true});
db.daily.ensureIndex({xid:1})
```

**Note: If your database is very large, createIndex may take a long time to execute, please be patient**

## Build

```sh
./build build
```

## Configuration

1.	SNAPSCALE_EXPLORER_MID_API_BASE: CHAIN API ENDPOINT
2.	SNAPSCALE_EXPLORER_MID_HTTP_PORT: Middle-ware http service port
3.	SNAPSCALE_EXPLORER_MID_WS_PORT: Middleware websocket service port
4.	SNAPSCALE_EXPLORER_MID_MONGO_CONFIG: Mongo address and certification
5.	SNAPSCALE_EXPLORER_MID_LOG_PATH: Middleware log path

```sh
docker run -itd \
  -p 8089:8089 \
  -p 8089:8089/udp \
  -p 8090:8090 \
  -e SNAPSCALE_EXPLORER_MID_API_BASE='http://192.168.1.201:30132/v1/'
  -e SNAPSCALE_EXPLORER_MID_HTTP_PORT='8090'
  -e SNAPSCALE_EXPLORER_MID_WS_PORT='8089'
  -e SNAPSCALE_EXPLORER_MID_MONGO_CONFIG='mongodb://username:password@192.168.1.1:30017/?authSource=admin'
  -e SNAPSCALE_EXPLORER_MID_LOG_PATH='/log'
  --name $NAME $NAME
```

## License

SnapScale Explorer is released under the Apache 2.0 license and is offered “AS IS” without warranty of any kind, express or implied. Any security provided by the SnapScale software depends in part on how it is used, configured, and deployed. SnapScale is built upon many third-party libraries such as WABT (Apache License) and WAVM (BSD 3-clause) which are also provided “AS IS” without warranty of any kind.

## Important

See [LICENSE](./LICENSE) for copyright and license terms.

All repositories and other materials are provided subject to the terms of this [IMPORTANT](./IMPORTANT.md) notice and you must familiarize yourself with its terms.  The notice contains important information, limitations and restrictions relating to our software, publications, trademarks, third-party resources, and forward-looking statements.  By accessing any of our repositories and other materials, you accept and agree to the terms of the notice. <br>

<br>

---

<a id="chinese"></a><br>
# SnapScale 浏览器中间件

SnapScale 浏览器中间件是SnapScale 浏览器的网关。

与EOS API相比，该浏览器中间件提供了大量数据统计的维度。

它是SnapScale 浏览器的一部分，可以用作网关。 运行SnapScale 浏览器的任何人都必须部署该中间件。

## 先决条件

启动一个SnapScale节点，并确保SnapScale 使用的是Mongo 保存历史记录的副本。创建集合和索引
```sh
db.createCollection("daily");
db.transactions.createIndex({"createdAt":1},{"background":true});
db.daily.ensureIndex({xid:1})
```

**注意：如果您的数据库很大，则createIndex 可能需要较长时间才能执行，请耐心等待**

## 创建

```sh
./build build
```

## 配置

1.	SNAPSCALE_EXPLORER_MID_API_BASE: CHAIN API ENDPOINT
2.	SNAPSCALE_EXPLORER_MID_HTTP_PORT: Middle-ware http service port
3.	SNAPSCALE_EXPLORER_MID_WS_PORT: Middleware websocket service port
4.	SNAPSCALE_EXPLORER_MID_MONGO_CONFIG: Mongo address and certification
5.	SNAPSCALE_EXPLORER_MID_LOG_PATH: Middleware log path

```sh
docker run -itd \
  -p 8089:8089 \
  -p 8089:8089/udp \
  -p 8090:8090 \
  -e SNAPSCALE_EXPLORER_MID_API_BASE='http://192.168.1.201:30132/v1/'
  -e SNAPSCALE_EXPLORER_MID_HTTP_PORT='8090'
  -e SNAPSCALE_EXPLORER_MID_WS_PORT='8089'
  -e SNAPSCALE_EXPLORER_MID_MONGO_CONFIG='mongodb://username:password@192.168.1.1:30017/?authSource=admin'
  -e SNAPSCALE_EXPLORER_MID_LOG_PATH='/log'
  --name $NAME $NAME
```

## 许可协议

SnapScale 遵循Apache 2.0 开源协议发布，按“原样”提供，没有任何明示或暗示的担保。SnapScale 软件提供的任何安全性部分取决于它的使用，配置和部署方式。 SnapScale 建立在许多第三方库上，如Binaryen（Apache许可证）和WAVM（BSD 3-clause），它们也是“按现状”提供的，没有任何形式的保证。


## 重要事项

有关版权和许可条款，请参考[许可协议](./LICENSE) 。

我们提供的所有信息都受限于本[重要通知](./IMPORTANT.md) 您必须熟悉其中的条款。该通知包含了关于我们的软件、发布、商标、第三方资源和前瞻性声明的重要信息、条件和限制。您获取任何材料的时候就意味着您接受并同意该通知的条款。
