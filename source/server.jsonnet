//權限 定義
local modeRoot = 1; //管理員 權限
local modeUser = 2; //一般用戶 權限 (可以 瀏覽軟件包 下載 更新 軟件)
{
    // 服務器 工作地址
	Addr:":19000",

    //ssl證書 位置
    Crt:"source/test.crt",
    //ssl證書key 位置
    Key:"source/test.key",

    //用戶 定義
    Users:[
        //定義一個 管理員
        {
            Name:"king",
            Pwd:"cerberus is an idea",
            Mode:modeRoot, //modeRoot 設置 優先於 modeUser
        },

        //定義一個 不用密碼的 用戶
        {
            Name:"anyone",   //注意 如果 Name Mode 重複了 則 使用 第一個 配置
            Pwd:"",
            Mode:modeUser,
        },     
    ],
    //軟件包 存放 位置
    Folders:[
        "source/t0",
        "source/t1",
    ],
    //運行 日誌
    Log:{
        // 需要打印的 日誌
        Logs:[
            "TRACE",
            "DEBUG",
            "INFO",
            "WARN",
            "ERROR",
            "FAULT",
        ],
        //是否 顯示 短 檔案名
        Short:true,
        //是否 顯示 長 檔案名
        Long:false,
    },
 }