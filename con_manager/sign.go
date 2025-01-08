package con_manager

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/max1015070108/topaiagent/con_manager/AIWorkload"
)

// Signature 结构体用于存储签名组件
// type Signature struct {
// 	R string
// 	S string
// 	V uint8
// }

func toBytes32(b []byte) [32]byte {
	var arr [32]byte
	copy(arr[:], b)
	return arr
}

// 生成要签名的消息内容
// func generateContent(
// 	addr common.Address,
// 	workload *big.Int,
// 	param1 *big.Int,
// 	param2 *big.Int,
// 	epochId *big.Int,
// ) ([]byte, error) {
// 	// 定义 ABI 参数类型
// 	arguments := abi.Arguments{
// 		{Type: getAddressType()},
// 		{Type: getUint256Type()},
// 		{Type: getUint256Type()},
// 		{Type: getUint256Type()},
// 		{Type: getUint256Type()},
// 	}

// 	// 进行 ABI 编码
// 	encoded, err := arguments.Pack(
// 		addr,
// 		workload,
// 		param1,
// 		param2,
// 		epochId,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack variables: %v", err)
// 	}

// 	return encoded, nil
// }

// 签名消息并返回签名组件
// func signMessage(privateKey *ecdsa.PrivateKey, message []byte) (*Signature, error) {
// 	// 添加以太坊签名前缀
// 	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
// 	prefixedHash := crypto.Keccak256([]byte(prefixedMessage))

// 	// 签名消息
// 	signature, err := crypto.Sign(prefixedHash, privateKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to sign message: %v", err)
// 	}

// 	// 解析签名组件
// 	r := hexutil.Encode(signature[:32])
// 	s := hexutil.Encode(signature[32:64])
// 	v := signature[64] + 27 // 转换为以太坊的 v 值

// 	return &Signature{
// 		R: r,
// 		S: s,
// 		V: v,
// 	}, nil
// }

// 辅助函数：获取 ABI 类型
func getAddressType() abi.Type {
	t, _ := abi.NewType("address", "", nil)
	return t
}

func getUint256Type() abi.Type {
	t, _ := abi.NewType("uint256", "", nil)
	return t
}

// func (c *ConManager) SignText(
//
//	addr string,
//	workload, param1, param2, epochID *big.Int,
//	privKeys []*ecdsa.PrivateKey,
//
//	) ([]struct {
//		R, S string
//		V    uint8
//	}, error) {
func (c *ConManager) SignText(
	addr string,
	workload, param1, param2, epochID *big.Int,
	privKeys []*ecdsa.PrivateKey,
) ([]AIWorkload.Signature, error) {
	// 准备要签名的数据
	arguments := abi.Arguments{
		{Type: getAddressType()},
		{Type: getUint256Type()},
		{Type: getUint256Type()},
		{Type: getUint256Type()},
		{Type: getUint256Type()},
	}

	// 编码数据
	encoded, err := arguments.Pack(
		// addr,
		common.HexToAddress(addr),
		workload,
		param1,
		param2,
		epochID,
	)
	if err != nil {
		return []AIWorkload.Signature{}, fmt.Errorf("failed to encode data: %v", err)
	}

	// 添加前缀并计算哈希
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(encoded))
	prefixedHash := crypto.Keccak256(append([]byte(prefix), encoded...))

	// 收集所有签名
	// signatures := make([]struct {
	// 	R, S string
	// 	V    uint8
	// }, len(privKeys))
	signatures := make([]AIWorkload.Signature, len(privKeys))
	// 遍历所有私钥进行签名
	for i, privateKey := range privKeys {
		signature, err := crypto.Sign(prefixedHash, privateKey)
		if err != nil {
			return []AIWorkload.Signature{}, fmt.Errorf("failed to sign with key %d: %v", i, err)
		}

		fmt.Printf("signature: %x\n", signature)
		// signatures[i] = AIWorkload.Signature{
		// 	R: toBytes32(signature[:32]),
		// 	S: toBytes32(signature[32:64]),
		// 	V: signature[64] + 27,
		// }
		// 将签名转换为十六进制字符串以便验证
		fullSig := hexutil.Encode(signature) // 0x + r + s + v

		// 解析各部分 - 模拟 JS 的处理方式
		var r, s [32]byte
		rBytes, err := hexutil.Decode(fullSig[:66]) // 0x + 前32字节
		if err != nil {
			return []AIWorkload.Signature{}, fmt.Errorf("failed to decode R: %v", err)
		}
		sBytes, err := hexutil.Decode("0x" + fullSig[66:130]) // 0x + 中间32字节
		if err != nil {
			return []AIWorkload.Signature{}, fmt.Errorf("failed to decode S: %v", err)
		}
		v := signature[64] + 27 // 最后一个字节 + 27

		// fmt.Println("r:", r)
		// fmt.Println("s:", r)
		// fmt.Println("v:", r)
		// 复制到固定大小的数组
		copy(r[:], rBytes)
		copy(s[:], sBytes)

		signatures[i] = AIWorkload.Signature{
			R: r,
			S: s,
			V: v,
		}
		fmt.Printf("+++++signature: %+v\n", signatures[i])
	}

	// fmt.Println("signatures: ", signatures)

	return signatures, nil
}
