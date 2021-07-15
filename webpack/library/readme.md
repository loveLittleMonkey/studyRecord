webpack打包组件和基础库

实现功能：
大整数加法
需要打包压缩版和非压缩版
支持AMD，CJS，ESM模块引入

具体分析需求：
1. 未压缩版 large-number.js
   压缩版 large-number.min.js
2. 目录
+ |- /dist
+    |- large-number.js
+    |- large-number.min.js
+ |- webpack.config.js
+ |- package.json
+ |- index.js
+ |- /src
+    |- index.js 
3. 支持的使用方式
ESM：
```javascript
import * as largeNumber from 'large-number'
// ...
largeNumber.add('999','1')
```
CJS：
```javascript
const largeNumber = require('large-number')
// ...
largeNumber.add('999','1')
```
AMD：
```javascript
require(['large-number'], function(large-number) {
// ...
largeNumber.add('999','1')
})
```
script直接引入:
```html
<script src="../large-number"></script>
<script>
  largeNumber.add('999','1')
  window.largeNumber.add('999','1')  
</script>
```

如何把库暴露出去
library: 制定库的全局变量
libraryTarget: 支持库引入的方式

有趣的知识点，应该可以追溯源码看看实现
没有设置mode，将output设置为 [name].ls
出错后可以看到没有被压缩，结论应该是默认会压缩js文件

使用
```javascript
import largeNumber from 'large-number'
export default {
    created() {
        var sum = largeNumber('999', '1')
        console.log('和：'+sum)
    }
}
```