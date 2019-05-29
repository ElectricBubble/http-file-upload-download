#!/usr/bin/env bash
scriptDir=$(cd `dirname $0`; pwd)
cd ..
projectDir=$(cd `dirname $0`; pwd)
cd ${scriptDir}
templateDir=${projectDir}/template
internalTmpltDir=${projectDir}/internal/template
#删除内部模版文件夹内的全部.go文件，应对如果模版文件名字有修改的情况
rm -rf ${internalTmpltDir}/*
hero -source="${templateDir}" -dest="${internalTmpltDir}" -extensions=".html,.vue,.tmplt"
