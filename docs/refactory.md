1. monitor数据

2. TOPAI-Node
    a. 公共模块: 请求签名、链上事件监听、调用合约作为公共模块、钱包签名 - Harry

    c. 区块链节点进程: Harry

1. proxy需要从链上拉取所有的Node,需要包含节点的各种元信息，比如GPU,CPU,MEM,DISK,NET等，以及这些Node的公钥。通过合约拉取？还是？
    a. GetAllNodeInfo() ([]NodeInfo, err)
    b. NodeInfo包含address

    ```shell
      ###
      {
        "address": "0x123456",
        "gpu": "NVIDIA",
        "cpu": "Intel",
        "mem": "16G",
        "disk": "1T",
        "net": "1Gbps",
        "publicKey": "0x123456"
      }
    ```
2. 客户端需要拉取到拉取到最新的链上epoch的id以及此 epoch中链上的时间戳以及随机数。通过合约拉取？还是？（开个接口）
    a. GetLastEpoch() (timestamp, blockHash)
3. proxy需要根据用户提交的epoch的id拉取对应的链上的时间戳以及随机数。通过合约拉取？还是？
    a. GetEpochByEpochId(epochid: str) (timestamp, blockHash)
4. proxy需要拉取模型分布图。这个需要调用合约吧？
    a. GetModelDeploymentMap() (map, err)
5. proxy需要可以上报模型执行的审计结果数据。这个需要调用合约吧？
    a. ReportProofOfModelWork(...)



recordModelUpload 上传模型， reportDeployment 模型的部署实例

ReportWorkload


### 1. mode
模型id作为key 地址list作为value的json格式
 mapping(uint256 => address[]) public modelDistribution;
