日前和同事有一个讨论一个问题，代码如下。当a被重新赋值’123‘的时候，没有触发到watch的事件，那么字符串的赋值与比较是如何进行的？当变量被赋值同一个值的时候，是不是js底层先进行了数值的比较，发现值是一样的，所以没有重新开辟内存空间来放置，直接栈引用的复用了呢？

```js
export default {
	data() {
        return {
            a: '123'
        }
    },
    watch: {
        a(newVal, oldVal) {
            if(newVal != oldVal){
                console.log('a is change')
            }
        }
    },
    methods: {
        changeA() {
            this.a = '123'
        }
    }
}
```

突然被这么一问就比较懵。心里的第一想法是字符串是js中的基础数据类型，但是说到这个内存分配的问题，得查一下资料或者具体代码跑一下给论证一下。好像代码论证有点难，一时不知道怎么调试浏览器的内存。但资料查找好像没有很详细的资料说明这个js的数据类型和内存分配的关系。

自己通过查阅的资料，得到我方的观点，js这边的字符串变量创建是放一个栈空间，然后 `var a = "字符串"; var b = a;` 此刻是在栈里开辟两个不同的地方储存的。当a、b进行比较的时候，通过不可能是通过判断内存地址来判断大小的。同时，js中字符串的大小比较，是转成了每个字符的ascii码，然后再做比较。字符串比较大小的方式就是通过ascii码的数值来确定的。如此这般，好像可以反证上面的问题，但好像还不够强有力。

有个ios的同事丢了这么个网络查询资料图片

![网络查询资料图片](C:\Users\chenzhenghan\Desktop\studyRecord\JS-GC\the-answer-find-in-network.png)

what！怎么感觉好像说得也对的样子？这很难受。有没有办法再继续论证下？扩大讨论的范围，在大神群里高呼，程老师半夜跑了个代码来论证，通过[node的vm]([vm 虚拟机 | Node.js API 文档 (nodejs.cn)](http://nodejs.cn/api/vm.html))来模拟v8，代码大概如下

```sh
node main.mjs # code is below

# 结果打印
第1次GC
(node:13720) ExperimentalWarning: vm.measureMemory is an experimental feature. This feature could change at any time
(Use `node --trace-warnings ...` to show where the warning was created)
135736
第2次GC
135328
第3次GC
135328
执行代码: var a = 1;
第4次GC
135432
执行代码: var b = 1;
第5次GC
135472
执行代码: var c = b;
第6次GC
135576
执行代码: var b = 2;
第7次GC
135576
执行代码: var d = 1;
第8次GC
135616
```



```js
// main.mjs
#!/usr/bin/env node

import { measureMemory, Script, createContext} from 'vm';
const context = createContext();

async function showMemory() {
    console.log(
        (await measureMemory({ mode: 'detailed', execution: 'eager'})).other[0].jsMemoryEstimate
    )
}

function exec(code) {
    console.log('执行代码:', code);
    new Script(code).runInNewContext(context);
}

console.log('第1次GC')
await showMemory();
console.log('第2次GC')
await showMemory();
console.log('第3次GC')
await showMemory();
exec('var a = 1;')
console.log('第4次GC')
await showMemory();
exec('var b = 1;')
console.log('第5次GC')
await showMemory();
exec('var c = b;')
console.log('第6次GC')
await showMemory();
exec('var b = 2;')
console.log('第7次GC')
await showMemory();
exec('var d = 1;')
console.log('第8次GC')
await showMemory();
```

！通过打印的GC可以获悉，ios同事搜索到的结果是不正确的。



最后，应该我方观点正确，但其中无法正面论证，`var a = '123'; a = '123';`，第二步代码是重新赋值一个123，还是旧的123的重用。求各路大佬的指点！

