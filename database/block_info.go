package database

import "fmt"

var LatestBlockNumber = 0

// BlockInfo 结构体用于表示区块信息
type BlockInfo struct {
	BlockHash string
	Timestamp int64
	Number    int64
}

// StoreBlockInfo 存储区块信息
func (s *EventStore) StoreBlockInfo(blockInfo BlockInfo) error {
	stmt := `
	INSERT OR REPLACE INTO block_info (
            id, block_hash, timestamp, number
        ) VALUES (1, ?, ?, ?)
    `
	_, err := s.db.Exec(stmt,
		blockInfo.BlockHash,
		blockInfo.Timestamp,
		blockInfo.Number,
	)
	return err
}

// GetBlockInfoByHash 通过区块哈希获取区块信息
func (s *EventStore) GetBlockInfoByHash(blockHash string) (*BlockInfo, error) {
	var blockInfo BlockInfo
	stmt := `
        SELECT block_hash, timestamp, number
        FROM block_info
        WHERE block_hash = ?
    `
	err := s.db.QueryRow(stmt, blockHash).Scan(
		&blockInfo.BlockHash,
		&blockInfo.Timestamp,
		&blockInfo.Number,
	)
	if err != nil {
		return nil, err
	}
	return &blockInfo, nil
}

// GetBlockInfoByNumber 通过区块号获取区块信息
func (s *EventStore) GetBlockInfoByNumber(number uint64) (*BlockInfo, error) {
	var blockInfo BlockInfo
	stmt := `
        SELECT block_hash, timestamp, number
        FROM block_info
        WHERE number = ?
    `
	err := s.db.QueryRow(stmt, number).Scan(
		&blockInfo.BlockHash,
		&blockInfo.Timestamp,
		&blockInfo.Number,
	)
	if err != nil {
		return nil, err
	}
	return &blockInfo, nil
}

// UpdateBlockInfo 更新区块信息
func (s *EventStore) UpdateBlockInfo(blockInfo BlockInfo) error {
	stmt := `
        UPDATE block_info
        SET timestamp = ?, number = ?
        WHERE block_hash = ?
    `
	_, err := s.db.Exec(stmt,
		blockInfo.Timestamp,
		blockInfo.Number,
		blockInfo.BlockHash,
	)
	return err
}

// DeleteBlockInfo 删除区块信息
func (s *EventStore) DeleteBlockInfo(blockHash string) error {
	stmt := `DELETE FROM block_info WHERE block_hash = ?`
	_, err := s.db.Exec(stmt, blockHash)
	return err
}

// GetLatestBlockInfo 获取最新的区块信息
func (s *EventStore) GetLatestBlockInfo() (*BlockInfo, error) {
	var blockInfo BlockInfo
	stmt := `
        SELECT block_hash, timestamp, number
        FROM block_info
        ORDER BY number DESC
        LIMIT 1
    `
	err := s.db.QueryRow(stmt).Scan(
		&blockInfo.BlockHash,
		&blockInfo.Timestamp,
		&blockInfo.Number,
	)
	if err != nil {
		return nil, err
	}
	return &blockInfo, nil
}

// GetBlockInfoInRange 获取指定区块号范围内的区块信息
func (s *EventStore) GetBlockInfoInRange(startBlock, endBlock uint64) ([]BlockInfo, error) {
	stmt := `
        SELECT block_hash, timestamp, number
        FROM block_info
        WHERE number BETWEEN ? AND ?
        ORDER BY number ASC
    `
	rows, err := s.db.Query(stmt, startBlock, endBlock)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blocks []BlockInfo
	for rows.Next() {
		var block BlockInfo
		err := rows.Scan(
			&block.BlockHash,
			&block.Timestamp,
			&block.Number,
		)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

type InfoChan struct {
	Channel chan interface{}
	Done    chan struct{}
	Err     chan error
}

// BlockInfoChannel 返回一个封装了通道和控制信号的结构体
func (s *EventStore) EventDataStoreChannel(bufferSize int) *InfoChan {
	infoChan := &InfoChan{
		Channel: make(chan interface{}, bufferSize),
		Done:    make(chan struct{}),
		Err:     make(chan error, 1),
	}

	// 启动后台处理协程
	fmt.Println("start something")
	go func() {
		defer close(infoChan.Channel)
		defer close(infoChan.Err)

		for {
			select {
			case <-infoChan.Done:
				return

			case infoData, ok := <-infoChan.Channel:
				fmt.Println("start something:", infoData, " ok:", ok)
				if !ok {
					return
				}

				switch infoData.(type) {
				case BlockInfo:
					// Normal BlockInfo type, continue with default processing
					if err := s.StoreBlockInfo(infoData.(BlockInfo)); err != nil {
						infoChan.Err <- err
					}

				case string:
					// Handle string type - could be JSON
					// Add JSON parsing logic if needed
					continue
				default:
					infoChan.Err <- fmt.Errorf("unsupported block info type: %T", infoData)
					continue
				}

			}
		}
	}()

	return infoChan
}

// Close 关闭通道和相关资源
func (b *InfoChan) Close() {
	close(b.Done)
}
