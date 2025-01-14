#### 1. 获取当前块高信息
```shell
GetLastBlock()

## return：
{
    "blockHash": "0x123456",
    "timestamp": 123456,
}
```

#### 2. 通过块高查询K块信息
```shell
GetBlockByHeighNumber(heighNumber: int)


### return
{
    "blockHash": "0x123456",
    "timestamp": 123456,
}
```

### 3. 获取node info list
```shell
GetAllNodeInfo() ([]NodeInfo, err)

### return
[]NodesRegistry.NodesRegistryNode


### example
./topaiagent nodegovernance getnodeinfo --rpc http://159.135.194.94:28080

```

### 4. ReportProofOfModelWork
```shell
ReportWorkload(
  reporters []string,
  workload, modelId, sessionId, epochID *big.Int

) error

### return
error


```

### 5. GetModelDeploymentMap
```shell
func (c *ConManager) GetModelDeploymentMap(modelId *big.Int) (map[*big.Int][]common.Address, error)

```


### event 监听
```shell

    事件监听功能
    1. 监听节点注册事件
    2. 监听节点注销事件
    3. 监听节点更新事件
    4. 监听节点工作量上报事件
    5. 监听节点工作量结算事件
    6. 监听节点工作量结算结果事件
    7. 监听节点工作量结算失败事件

```
