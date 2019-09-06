# snapscale-explorer-middleware
middleware of snapscale explorer  
![](https://img.shields.io/badge/version-1.0.0-brightgreen) ![](https://img.shields.io/badge/author-Miguel-blue)

## Before Everything
- ###Add Index (mongodb)
    - transactions | createdAt
    ```
    db.transactions.createIndex({"createdAt":1}{"background":true})
    ```
- ###Net
    Snapscale-explorer | Snapscale-explorer-middleware should under same network group or provide 8089/8090 ports for explorer.
    
## Build && Run
```bash
./build
```

## Configuration
- DB config  
/config/config.go

- Docker config  
./Dockerfile

## Log (optional)
Run script for log support
```bash
mkdir /file/output/log
```