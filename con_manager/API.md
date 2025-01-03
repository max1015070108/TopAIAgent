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
