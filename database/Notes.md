这些方法提供了以下功能：

1. `StoreBlockInfo`: 存储新的区块信息
2. `GetBlockInfoByHash`: 通过区块哈希查询区块信息
3. `GetBlockInfoByNumber`: 通过区块号查询区块信息
4. `UpdateBlockInfo`: 更新已存在的区块信息
5. `DeleteBlockInfo`: 删除区块信息
6. `GetLatestBlockInfo`: 获取最新的区块信息
7. `GetBlockInfoInRange`: 获取指定区块号范围内的所有区块信息

你可以根据需要使用这些方法来管理区块信息。每个方法都有适当的错误处理，并返回相应的错误信息。

示例使用：
```go
// 存储区块信息
blockInfo := BlockInfo{
    BlockHash: "0x...",
    Timestamp: time.Now().Unix(),
    Number:    12345,
}
err := store.StoreBlockInfo(blockInfo)

// 查询区块信息
block, err := store.GetBlockInfoByHash("0x...")
