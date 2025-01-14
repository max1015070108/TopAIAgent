package database

import (
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/max1015070108/TopAIAgent/con_manager/AIModels"
	"github.com/max1015070108/TopAIAgent/con_manager/AIWorkload"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistry"
)

// StoreModelEvent stores an AIModels contract event
func (s *EventStore) StoreModelEvent(event interface{}, blockNumber uint64, txHash string) error {
	stmt := ""
	var args []interface{}

	switch e := event.(type) {
	case *AIModels.AIModelsModelDeployed:
		stmt = `
			INSERT INTO model_events (
				event_type, model_id, node, block_number, tx_hash, timestamp
			) VALUES (?, ?, ?, ?, ?, ?)
		`
		args = []interface{}{
			"ModelDeployed",
			e.ModelId.Uint64(),
			e.Node.Hex(),
			blockNumber,
			txHash,
			time.Now().Unix(),
		}

	case *AIModels.AIModelsModelRemoved:
		stmt = `
			INSERT INTO model_events (
				event_type, model_id, node, block_number, tx_hash, timestamp
			) VALUES (?, ?, ?, ?, ?, ?)
		`
		args = []interface{}{
			"ModelRemoved",
			e.ModelId.Uint64(),
			e.Node.Hex(),
			blockNumber,
			txHash,
			time.Now().Unix(),
		}

	case *AIModels.AIModelsUploadModeled:
		stmt = `
			INSERT INTO model_events (
				event_type, model_id, uploader, model_name, model_version, model_info,
				block_number, tx_hash, timestamp
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
		args = []interface{}{
			"UploadModeled",
			e.ModelId.Uint64(),
			e.Uploader.Hex(),
			e.ModelName,
			e.ModelVersion,
			e.ModelInfo,
			blockNumber,
			txHash,
			time.Now().Unix(),
		}
	}

	if stmt != "" {
		if _, err := s.db.Exec(stmt, args...); err != nil {
			return err
		}
	}

	return nil
}

// StoreWorkloadEvent stores an AIWorkload contract event
func (s *EventStore) StoreWorkloadEvent(event interface{}, blockNumber uint64, txHash string) error {
	switch e := event.(type) {
	case *AIWorkload.AIWorkloadWorkloadReported:
		stmt := `
			INSERT INTO workload_events (
				event_type, session_id, reporter, worker, epoch_id, workload, model_id,
				block_number, tx_hash, timestamp
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
		_, err := s.db.Exec(stmt,
			"WorkloadReported",
			e.SessionId.Uint64(),
			e.Reporter.Hex(),
			e.Worker.Hex(),
			e.EpochId.Uint64(),
			e.Workload.Uint64(),
			e.ModelId.Uint64(),
			blockNumber,
			txHash,
			time.Now().Unix(),
		)
		return err
	}
	return nil
}

// StoreNodeRegistryEvent stores a NodesRegistry contract event
func (s *EventStore) StoreNodeRegistryEvent(event interface{}, blockNumber uint64, txHash string) error {
	stmt := ""
	var args []interface{}

	switch e := event.(type) {
	case *NodesRegistry.NodesRegistryNodeRegistered:
		stmt = `
			INSERT INTO node_registry_events (
				event_type, wallet, identifier, time, alias_identifier,
				block_number, tx_hash, timestamp
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`
		args = []interface{}{
			"NodeRegistered",
			e.Wallet.Hex(),
			e.Identifier.Hex(),
			e.Time.Uint64(),
			e.AliasIdentifier,
			blockNumber,
			txHash,
			time.Now().Unix(),
		}

	case *NodesRegistry.NodesRegistryNodeActived:
		stmt = `
			INSERT INTO node_registry_events (
				event_type, wallet, identifier, time, alias_identifier,
				block_number, tx_hash, timestamp
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`
		args = []interface{}{
			"NodeActived",
			e.Wallet.Hex(),
			e.Identifier.Hex(),
			e.Time.Uint64(),
			e.AliasIdentifier,
			blockNumber,
			txHash,
			time.Now().Unix(),
		}

	case *NodesRegistry.NodesRegistryNodeDeregistered:
		stmt = `
			INSERT INTO node_registry_events (
				event_type, identifier, time, alias_identifier,
				block_number, tx_hash, timestamp
			) VALUES (?, ?, ?, ?, ?, ?, ?)
		`
		args = []interface{}{
			"NodeDeregistered",
			e.Identifier.Hex(),
			e.Time.Uint64(),
			e.AliasIdentifier,
			blockNumber,
			txHash,
			time.Now().Unix(),
		}
	}

	if stmt != "" {
		if _, err := s.db.Exec(stmt, args...); err != nil {
			return err
		}
	}

	return nil
}

// StoreNodesRegistryNode stores a NodesRegistryNode record
func (s *EventStore) StoreNodesRegistryNode(node NodesRegistry.NodesRegistryNode) error {
	gpusJSON, err := json.Marshal(node.Gpus)
	if err != nil {
		return err
	}

	stmt := `
		INSERT INTO nodes_registry_nodes (
			identifier, alias_identifier, registration_time, active, gpus, wallet, stake
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err = s.db.Exec(stmt,
		node.Identifier.Hex(),
		node.AliasIdentifier,
		node.RegistrationTime.Uint64(),
		node.Active,
		string(gpusJSON),
		node.Wallet.Hex(),
		node.Stake.String(),
	)
	return err
}

// GetNodesRegistryNode retrieves a NodesRegistryNode by identifier
func (s *EventStore) GetNodesRegistryNode(identifier common.Address) (*NodesRegistry.NodesRegistryNode, error) {
	var node NodesRegistry.NodesRegistryNode
	var gpusJSON string
	var stake string

	stmt := `
		SELECT identifier, alias_identifier, registration_time, active, gpus, wallet, stake
		FROM nodes_registry_nodes
		WHERE identifier = ?
	`
	err := s.db.QueryRow(stmt, identifier.Hex()).Scan(
		&node.Identifier,
		&node.AliasIdentifier,
		&node.RegistrationTime,
		&node.Active,
		&gpusJSON,
		&node.Wallet,
		&stake,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(gpusJSON), &node.Gpus)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

// UpdateNodesRegistryNode updates a NodesRegistryNode record
func (s *EventStore) UpdateNodesRegistryNode(node NodesRegistry.NodesRegistryNode) error {
	gpusJSON, err := json.Marshal(node.Gpus)
	if err != nil {
		return err
	}

	stmt := `
		UPDATE nodes_registry_nodes
		SET alias_identifier = ?, registration_time = ?, active = ?, gpus = ?, wallet = ?, stake = ?
		WHERE identifier = ?
	`
	_, err = s.db.Exec(stmt,
		node.AliasIdentifier,
		node.RegistrationTime.Uint64(),
		node.Active,
		string(gpusJSON),
		node.Wallet.Hex(),
		node.Stake.String(),
		node.Identifier.Hex(),
	)
	return err
}

// DeleteNodesRegistryNode deletes a NodesRegistryNode by identifier
func (s *EventStore) DeleteNodesRegistryNode(identifier common.Address) error {
	stmt := `DELETE FROM nodes_registry_nodes WHERE identifier = ?`
	_, err := s.db.Exec(stmt, identifier.Hex())
	return err
}
