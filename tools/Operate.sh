## pre-requisite

# npm install -g solc / brew install solidity
#
# go install github.com/ethereum/go-ethereum/cmd/abigen@latest


# CONT_PATH=../contract/src/TopAIAgent.sol
# solc --overwrite --abi --bin $CONT_PATH -o build
# # path
# Package_PATH=../topaiagent/con_manager
# abigen --bin=build/TopAIAgent.bin --abi=build/TopAIAgent.abi --pkg=token --out=$Package_PATH/manager.go


## AIModels
Package_PATH=../con_manager
CON_SOURCE=../../TOP-crosschain-contracts/contracts/contracts/AI
CON_ABI_SOURCE=../../TOP-crosschain-contracts/artifacts/contracts/contracts/AI

for file in `ls $CON_SOURCE`
do
    echo $file
    # solc --overwrite --abi --bin $CON_SOURCE/$file -o build

    filename=${file%.*}
    echo $filename

    jq -r .abi $CON_ABI_SOURCE/$filename.sol/$filename.json > build/$filename.abi
    jq -r .bytecode $CON_ABI_SOURCE/$filename.sol/$filename.json | sed 's/0x//' > build/$filename.bin

    mkdir -p $Package_PATH/$filename

    # abigen --combined-json $CON_ABI_SOURCE/$filename.sol/$filename.json --pkg=con_manager --out=$Package_PATH/$filename".go"
    abigen --bin=build/$filename.bin --abi=build/$filename.abi --pkg=$filename --out=$Package_PATH/$filename/$filename".go"

done
