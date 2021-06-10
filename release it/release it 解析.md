```sh
# release it 执行解析

git diff --quiet HEAD # 查询工作目录是否干净

git rev-parse --abbrev-ref HEAD # 获取当前分支名称
# feature/new_stage

git config --get branch.feature/new_stage.remote # 获取当前分支的远程仓库的源
# origin

git remote get-url origin # 获取当前分支的远程仓库的地址
# remote url

git fetch # 拉取代码

git describe --tags --abbrev=0 # 获取最后一个tag
# 2.15.14

git symbolic-ref HEAD # 查看工作树所在分支的参数
# refs/heads/feature/new_stage

git for-each-ref --format="%(upstream:short)" refs/heads/feature/new_stage # 分支的信息输出
# origin/feature/new_stage

git log --pretty=format:"* %s (%h)" v2.12.14...HEAD # 指定两个git提交区间内代码提交的message         
# * message (hash)
# * message (hash)

# inquirer version的自增

npm version 30.3.3 --no-git-tag-version # package文件的自增，且不提交记录

git status --short --untracked-files=no # 查看工作区变更文件
# 执行结果：
#  M package-lock.json
#  M package.json

execa command: git
[ 'add', '.', '--update' ]

# inquirer 确认提交的commit信息

execa command: git
[ 'commit', '--message', 'chore: release 30.3.3' ]

# inquirer 确认tag信息

execa command: git
[ 'tag', '--annotate', '--message', 'Release 30.3.3', 'v30.3.3' ]

# inquirer 推代码

execa command: git
[ 'push', '--follow-tags' ]

# inquirer 推npm仓库
```



使用 shelljs.js 来处理shell脚本执行

用参数可以使用execa库 用的node的process的子进程的方式



[git指南]([Git 中文参考 - 《Git 中文参考》 - 书栈网 · BookStack](https://www.bookstack.cn/read/git-doc-zh/README.md))





# release it 解析

```
插件有的属性
namespace
options
global
context
config 也是个类函数
log 也是个类函数
shell 也是个类函数
spinner 也是个类函数
prompt 也是个类函数
debug 也是个类函数

插件拥有的函数
getInitialOptions
=== 无具体实现 start
init
getName
getLatestVersion
getChangelog
getIncrementedVersionCI
getIncrementedVersion
beforeBump
bump
beforeRelease
release
afterRelease
=== 无具体实现 end
getContext
setContext
exec
registerPrompts
showPrompt
step


一共有的插件
VantCliReleasePlugin vantcli新增的
npm
Git
Version

VantCliReleasePlugin：
beforeRelease

npm：
isEnabled
init
getName
getLatestVersion
bump
release
isRegistryUp
isAuthenticated
getRegistry
getLatestRegistryVersion
getRegistryPreReleaseTags
getPackageUrl
guessPreReleaseTag
resolveTag
publish
afterRelease

git：
getInitialOptions
init
getName
getLatestVersion
getChangelog
bump
isRemoteName
getRemoteUrl
getRemote
getBranchName
getRemoteForBranch
fetch
getLatestTagName
getSecondLatestTagName
isEnabled
beforeRelease
release
isRequiredBranch
hasUpstreamBranch
tagExists
isWorkingDirClean
getCommitsSinceLatestTag
stage
stageDir
reset
rollback
status
commit
tag
push
afterRelease


Version：
getLatestVersion
getIncrementedVersionCI
getIncrementedVersion
promptIncrementVersion
isPreRelease
isValid
incrementVersion

配置信息
{
  plugins: {
    'F:\\codes\\vant\\node_modules\\@vant\\cli\\lib\\compiler\\vant-cli-release-plugin.js': {}
  },
  git: {
    tagName: 'v${version}',
    commitMessage: 'chore: release ${version}',
    changelog: 'git log --pretty=format:"* %s (%h)" ${from}...${to}',
    requireCleanWorkingDir: true,
    requireBranch: false,
    requireUpstream: true,
    requireCommits: false,
    addUntrackedFiles: false,
    commit: true,
    commitArgs: [],
    tag: true,
    tagAnnotation: 'Release ${version}',
    tagArgs: [],
    push: true,
    pushArgs: [ '--follow-tags' ],
    pushRepo: ''
  },
  ci: undefined,
  hooks: {},
  npm: {
    publish: true,
    publishPath: '.',
    tag: null,
    otp: null,
    ignoreVersion: false,
    skipChecks: false,
    timeout: 10
  },
  github: {
    release: false,
    releaseName: 'Release ${version}',
    releaseNotes: null,
    preRelease: false,
    draft: false,
    tokenRef: 'GITHUB_TOKEN',
    assets: null,
    host: null,
    timeout: 0,
    proxy: null,
    skipChecks: false
  },
  gitlab: {
    release: false,
    releaseName: 'Release ${version}',
    releaseNotes: null,
    tokenRef: 'GITLAB_TOKEN',
    assets: null,
    origin: null,
    skipChecks: false
  },
  isUpdate: false,
  version: {
    increment: undefined,
    isPreRelease: false,
    preReleaseId: undefined
  }
}
```











