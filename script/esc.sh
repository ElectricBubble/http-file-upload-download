#!/usr/bin/env bash
scriptDir=$(cd `dirname $0`; pwd)
cd ..
projectDir=$(cd `dirname $0`; pwd)
cd ${scriptDir}
assetGoFile=asset.go
internalAssetGoFile=${projectDir}/internal/asset/${assetGoFile}
#删除原来打包好的静态文件
if [[ -f ${internalAssetGoFile} ]]; then
    rm ${internalAssetGoFile}
fi
esc -o=${assetGoFile} -pkg=asset -ignore=".DS_Store" ../asset
#生成在当前脚本路径下，如果直接生成在指定路径下，会多出asset.go文件，精神洁癖，以此避免
mv ${scriptDir}/${assetGoFile} ${internalAssetGoFile}
