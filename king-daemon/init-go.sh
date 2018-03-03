#!/bin/bash
#Program:
#       創建 vscode golang 項目
#History:
#       2018-02-10 king first release
#       2018-03-03 add git .gitignore file
#Email:
#       zuiwuchang@gmail.com
 
# 獲取 項目路徑作爲 項目 名稱
appName=`pwd`
appName=${appName##*/}
echo create project [ $appName ]
 
# 定義的 各種 輔助 函數
MkDir(){
	mkdir -p "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}
MkOrClear(){
	if test -d "$1";then
		declare path="$1"
		path+="/*"
		rm "$path" -rf
		if [ "$?" != 0 ] ;then
			exit 1
		fi
	else
		MkDir $1
	fi
}
NewFile(){
	echo "$2" > "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}
WriteFile(){
	echo "$2" >> "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}
 
# 創建 bin 目錄
MkDir build/bin
 
# 創建/清空 release 目錄
MkOrClear build/release
 
#創建 vscode 配置目錄
MkDir .vscode
# 創建 vscode tasks.json 檔案
fileName=.vscode/tasks.json
if ! test -f $fileName ;then
	NewFile $fileName	'{'
	WriteFile $fileName	'	"version": "2.0.0",'
	WriteFile $fileName	'	"type": "shell",'
	WriteFile $fileName	'	"tasks": ['
   	WriteFile $fileName	'		/***    設置 build 任務    ***/'
	WriteFile $fileName	'		{'
	WriteFile $fileName	'			"label": "build",'
	WriteFile $fileName	'			"command": "go build -o ${workspaceRoot}/build/release/'"${appName}"'",'
   	WriteFile $fileName	'			"windows": {'
	WriteFile $fileName	'				"command": "go build -o ${workspaceRoot}/build/release/'"${appName}"'.exe",'
	WriteFile $fileName	'			},'
	WriteFile $fileName	'			"problemMatcher":"$go",'
	WriteFile $fileName	'			"group": {'
	WriteFile $fileName	'				"kind": "build",'
	WriteFile $fileName	'				"isDefault": true'
	WriteFile $fileName	'			},'
	WriteFile $fileName	'			"options": { "cwd": "${workspaceRoot}" }'
	WriteFile $fileName	'		},'
	WriteFile $fileName	'		/***    設置 run 任務    ***/'
	WriteFile $fileName	'		{'
	WriteFile $fileName	'			"label": "run",'
	WriteFile $fileName	'			"command": "${workspaceRoot}/build/release/'"${appName}"'",'
	WriteFile $fileName	'			"windows": {'
	WriteFile $fileName	'				"command": "${workspaceRoot}/build/release/'"${appName}"'.exe",'
	WriteFile $fileName	'			},'
	WriteFile $fileName	'			"problemMatcher":"$go",'
	WriteFile $fileName	'			"options": { "cwd": "${workspaceRoot}/build/bin" }'
	WriteFile $fileName	'		},'
	WriteFile $fileName	'	]'
	WriteFile $fileName	'}'
fi
# 創建 vscode launch.json 檔案
fileName=.vscode/launch.json
if ! test -f $fileName ;then
	NewFile $fileName	'{'
	WriteFile $fileName	'	"version": "0.2.0",'
	WriteFile $fileName	'	"configurations": ['
	WriteFile $fileName	'		{'
	WriteFile $fileName	'			"name": "go Launch",'
	WriteFile $fileName	'			"type": "go",'
	WriteFile $fileName	'			"request": "launch",'
	WriteFile $fileName	'			"mode": "exec",'
	WriteFile $fileName	'			"program": "${workspaceRoot}/build/release/'"$appName"'",'
	WriteFile $fileName	'			"windows": {'
	WriteFile $fileName	'				"program": "${workspaceRoot}/build/release/'"$appName"'.exe",'
	WriteFile $fileName	'			},'
	WriteFile $fileName	'			"args": [],'
	WriteFile $fileName	'			"cwd": "${workspaceRoot}/build/bin",'
	WriteFile $fileName	'		}'
	WriteFile $fileName	'	]'
	WriteFile $fileName	'}'
fi
 
# 創建 vscode settings.json
fileName=.vscode/settings.json
if ! test -f $fileName ;then
	NewFile $fileName	'{'
	WriteFile $fileName	'	"files.exclude": {'
	WriteFile $fileName	'		"**/init-go.sh": true,'
	WriteFile $fileName	'		"**/build/release": true,'
	WriteFile $fileName	'	}'
	WriteFile $fileName	'}'
fi

# 創建 git .gitignore
fileName=.gitignore
if ! test -f $fileName ;then
	NewFile $fileName	'# bin'
	WriteFile $fileName	'build/'
fi
