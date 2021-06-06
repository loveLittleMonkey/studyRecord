

[Learn the Command Line](https://www.codecademy.com/learn/learn-the-command-line)

# Command Line Interface[CLI]

```sh
# step1 don't want global, use step3
npm install rollup --global # or `npm i rollup -g` for short

# step2
mkdir -p my-rollup-project
cd my-rollup-project

# step3
npm install rollup --save-dev # then the next 'rollup ...' should be 'npx rollup ...'

# step4
rollup # equals rollup --help

mkdir src
cd src
touch main.js
open main.js # the code is below

touch foo.js
open foo.js # the code is below

cd ../

rollup src/main.js -f cjs # print in stdout

rollup src/main.js -o dist/bundle.js -f cjs

node # run the bundle in node to see the result
> var myBundle = require('./dist/bundle.js');
> myBundle();
```

```js
// src/main.js
import foo from './foo.js';
export default function () {
    console.log(foo);
}
```

```js
// src/foo.js
export default 'hello world!';
```

# Using config files

```sh
touch rollup.config.js # the code is below

rollup -c

# or
rollup -c -o ./dist/bundle-2.js # `-o` is short for `--output.file`

# default rollup.config.js, if want to do some change
# rollup --config rollup.config.dev.js
# rollup --config rollup.config.prod.js
```

```js
// rollup.config.js
export default {
  input: 'src/main.js',
  output: {
    file: './dist/bundle-2.js',
    format: 'cjs'
  }
};
```

# Using plugins

```sh
npm install --save-dev rollup-plugin-json
# no --save becasuse the code is no dependence on it, just package use it

cd src
touch main2.js # the code is below
touch package.json # the code is below
cd ..

touch rollup.config.main2.js # the code is below

rollup --config rollup.config.main2.js 
```

```js
// src/main2.js
import { version } from './package.json';

export default function () {
  console.log('version ' + version);
}
```

```json
{
    "author": {
      "name": "Rich Harris"
    },
    "bin": {
      "rollup": "dist/bin/rollup"
    },
    "bugs": {
      "url": "https://github.com/rollup/rollup/issues"
    },
    "bundleDependencies": false,
    "dependencies": {
      "fsevents": "~2.3.1"
    },
    "deprecated": false,
    "description": "Next-generation ES module bundler",
    "engines": {
      "node": ">=10.0.0"
    },
    "exports": {
      ".": {
        "node": {
          "require": "./dist/rollup.js",
          "import": "./dist/es/rollup.js"
        },
        "default": "./dist/es/rollup.browser.js"
      },
      "./dist/": "./dist/"
    },
    "files": [
      "dist/**/*.js",
      "dist/*.d.ts",
      "dist/bin/rollup",
      "dist/rollup.browser.js.map"
    ],
    "homepage": "https://rollupjs.org/",
    "keywords": [
      "modules",
      "bundler",
      "bundling",
      "es6",
      "optimizer"
    ],
    "license": "MIT",
    "main": "dist/rollup.js",
    "module": "dist/es/rollup.js",
    "name": "rollup",
    "optionalDependencies": {
      "fsevents": "~2.3.1"
    },
    "repository": {
      "type": "git",
      "url": "git+https://github.com/rollup/rollup.git"
    },
    "version": "2.50.6"
  }
```

```js
// rollup.config.main2.js
import json from 'rollup-plugin-json';

export default {
  input: 'src/main2.js',
  output: {
    file: './dist/bundle-3.js',
    format: 'cjs'
  },
  plugins: [ json() ]
};
```

```js
// bundle-3.js tree-shaking is effective
'use strict';

var version = "2.50.6";

var main = function () {
  console.log('version ' + version);
};

module.exports = main;
```



# 参考文献

[rollup.js 中文文档](https://www.rollupjs.com/)

[rollup.js 英文文档](https://rollupjs.org/guide/en/) -> 更全面