lua learning

项目观察
    git:https://github.com/CNSRE/ABTestingGateway
    项目结构
        按接口分包
        公共模块分包

        层级结构
            入口 层
            model 数据操作层 + 策略 + 公用模块




数据结构
    万能的 table  --

special
    1.数组长度操作# nil会作为终止条件 
    2. a%b = a - floor(a/b)*b
    3. and or 短路求值(short_cur evaluation)  类三目运算  不可完全挪用
    4. 函数 为第一类型   可以善加利用 尾调用消除特性 
    5. error pcall
    6.字符串为不可变类型 拼接用数据存储 table.concat统一处理
    7.全局变量 即为 全局table _G 存储  rawset 可 跳过 __newindex error 抛出 设置全局变量
    0. 协程    todo
