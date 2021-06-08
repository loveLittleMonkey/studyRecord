# Inquirer.js

### 基础知识点

```sh
npm install inquirer

node main.js # the code is below
```

```js
// main.js
var inquirer = require('inquirer');
inquirer
  .prompt([
    /* Pass your questions in here */
  ])
  .then((answers) => {
    // Use user feedback for... whatever!!
  })
  .catch((error) => {
    if (error.isTtyError) {
      // Prompt couldn't be rendered in the current environment
    } else {
      // Something else went wrong
    }
  });

// or
var prompt = inquirer.createPromptModule();
prompt(questions).then(/* ... */);

// or
inquirer.registerPrompt(name, prompt)
```

```js
// example pizza.js
var inquirer = require('inquirer');

console.log('Hi, welcome to Node Pizza');

const questions = [
  {
    type: 'confirm',
    name: 'toBeDelivered',
    message: 'Is this for delivery?',
    default: false,
  },
  {
    type: 'input',
    name: 'phone',
    message: "What's your phone number?",
    validate(value) {
      const pass = value.match(
        /^([01]{1})?[-.\s]?\(?(\d{3})\)?[-.\s]?(\d{3})[-.\s]?(\d{4})\s?((?:#|ext\.?\s?|x\.?\s?){1}(?:\d+)?)?$/i
      );
      if (pass) {
        return true;
      }

      return 'Please enter a valid phone number';
    },
  },
  {
    type: 'list',
    name: 'size',
    message: 'What size do you need?',
    choices: ['Large', 'Medium', 'Small'],
    filter(val) {
      return val.toLowerCase();
    },
  },
  {
    type: 'input',
    name: 'quantity',
    message: 'How many do you need?',
    validate(value) {
      const valid = !isNaN(parseFloat(value));
      return valid || 'Please enter a number';
    },
    filter: Number,
  },
  {
    type: 'expand',
    name: 'toppings',
    message: 'What about the toppings?',
    choices: [
      {
        key: 'p',
        name: 'Pepperoni and cheese',
        value: 'PepperoniCheese',
      },
      {
        key: 'a',
        name: 'All dressed',
        value: 'alldressed',
      },
      {
        key: 'w',
        name: 'Hawaiian',
        value: 'hawaiian',
      },
    ],
  },
  {
    type: 'rawlist',
    name: 'beverage',
    message: 'You also get a free 2L beverage',
    choices: ['Pepsi', '7up', 'Coke'],
  },
  {
    type: 'input',
    name: 'comments',
    message: 'Any comments on your purchase experience?',
    default: 'Nope, all good!',
  },
  {
    type: 'list',
    name: 'prize',
    message: 'For leaving a comment, you get a freebie',
    choices: ['cake', 'fries'],
    when(answers) {
      return answers.comments !== 'Nope, all good!';
    },
  },
];

inquirer.prompt(questions).then((answers) => {
  console.log('\nOrder receipt:');
  console.log(JSON.stringify(answers, null, '  '));
});
```

### question 配置项目解析

1. **type**[字符串]: 提问的类型，默认值是`input`，可以有的值:
   * input: 手输模式
   * number: 感觉是input类型默认转成number类型返回
   * confirm: 是与否选择
   * list: 列表项目单选选择
   * rawlist: 列表项目单选，可手输选项下标
   * expand: 列表单选，提供的是简要的key选项，手输key值会有对应提示，enter后可看到所有选项提示再选择
   * checkbox: 列表多选
   * password: 密码输入选项，默认输入不可见，可以通过mask属性控制显示***
   * editor：吊起文本输入ide输入内容
2. **name**[字符串]：存储答案对应的key值。如果包含了英文句号，将表示为一个嵌套的对象
3. **message**[字符串、函数]: 提问的问题，如果没有设置将使用name作为显示。如果设置为函数，函数的第一个参数为当前的答案[测试感觉是空对象？]
4.  **default**[字符串、数字、布尔值、数组、函数]：如果用户没有操作时的默认的答案，如果设置为函数，函数的第一个参数为当前的答案
5. **choices**[数组，函数]：选项数组，或者一个返回选项数组的函数。如果设置为函数，函数的第一个参数为当前的答案。选项数组里的值可以是简单的数字，字符串或者对象[对象中包含name，value，short的key值，name用来展示，short用来选择后的简单提示，value用来做答案]。还可以包含一个 Separator 做分割作用。
6. **validate**[函数]： 用户的输入的数据的合法性校验，如果合法返回true，否则返回一个提示错误的信息。如果返回false，模式的错误信息将会展示。
7. **filter**[函数]：用户的输入的数据进行一定的过滤处理，处理后的返回才是答案
8. transformer[函数]：个人感觉没啥乱用的感觉，实例代码用了一个chalkPipe的库进行实时答案展示的颜色变换，不影响最终结果的展示。函数的第一个参数是用户的输入，第二个参数是答案数组，第三个参数是一个操作标志[不知道哪里来的]，返回一个转换后的值显示实时答案。
9. **when**[函数，布尔值]：接收当前的答案数组然后返回真假，决定当前问题是否需要被询问。也可以是简单的布尔值直接设置。
10. **pageSize**[数字]: 当是list，rawlist，expand和checkbox类型的时候，展示的列表长度
11. prefix，suffix[字符串]： 问题展示的前缀和后缀，没啥用感觉
12. askAnswered[布尔值]：如果答案已经存在，则强制提示问题。不知道干啥用
13. loop[布尔值]：list类型能否循环，模式为true

### Answers 答案对象

key：每个问题的name值

value：如果是confirm类型，是布尔值

​			如果是input类型，返回用户输入，如果定义了过滤器则会过滤，字符串类型

​            如果是number类型，返回用户输入，如果定义了过滤器则会过滤，数字类型

​           如果是rawlist，list类型，返回用户选择的值，字符串类型

### Separator 

用法： `new inquirer.Separator()` or `new inquirer.Separator('输入一段文字做分割字符，默认是 -------- ')`



# 参考文献

[官方github教程]([SBoudrias/Inquirer.js: A collection of common interactive command line user interfaces. (github.com)](https://github.com/SBoudrias/Inquirer.js)) 本文就是个中文翻译使用指南,github崩了看同目录下的源代码压缩包吧哈哈
