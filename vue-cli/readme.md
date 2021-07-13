想要快速开发vue原型，可以用`vue serve`

```sh
npm install -g @vue/cli-service-global

vue serve
```

注意，单独安装无效，`bash: vue: command not found`，需要先安装了`vue-cli`

```
npm install -g @vue/cli
```

 问题来了，这是两个库合并使用还是每次只下载了库里的一部分呢？

