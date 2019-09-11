# snapscale-explorer-middleware
middleware of snapscale explorer  
![](https://img.shields.io/badge/version-1.0.0-brightgreen) ![](https://img.shields.io/badge/author-Miguel-blue)

## Before Everything
- Add Index (mongodb)
    - transactions | createdAt
    ```
    db.transactions.createIndex({"createdAt":1}{"background":true})
    ```
- Net
    Snapscale-explorer | Snapscale-explorer-middleware should under same network group or provide 8089/8090 ports for explorer.
    
## Build && Run
```bash
./build
```

## Configuration
```bash
  docker run -itd \
  -p 8089:8089 \
  -p 8089:8089/udp \
  -p 8090:8090 \
  -e ApiBase='http://192.168.1.201:30132/v1/'
  -e HttpPort='8090'
  -e WsPort='8089'
  -e MongoConfig='mongodb://xeniro:N0password@192.168.1.201:30017/?authSource=admin'
  --name $NAME $NAME
```

## Log (optional)
Run script for log support
```bash
mkdir /file/output/log
```