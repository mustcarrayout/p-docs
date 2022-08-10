
# git

## 什么是git？

### 看图理解

1. 从中看出了哪些问题？
2. 从中知道了哪些问题的答案？

## .gitignore

1. 新规则只考虑新的文件**不向后兼容**
2. .gitignore配置文件是按行从上到下进行规则匹配的
3. 该文件会影响repo的所有人，如果只是个人习惯问题建议操作**.git/info/exclude**，语法和效果一样的

```

```



## git操作

### 养成良好的习惯

1. 上班前pull|fetch
2. 下班前commit|push
3. 多用branch、patch
4. 不建议用git管理二进制文件(超过5M的)；

### check out

场景：
* 取消本地文件更改
* 取消本地分支更改

```shell
git checkout /patch/src/todo/
git checkout [current branch]
```

### git rm

场景：
* 删除tracked的文件
* --cached 只删除暂存区,不删除本地;不加删除本地和缓存区的

```shell
git rm -rf --cached /path/src/todo/
# 设置忽略跟踪(1)，忽略远程的变更，只保留本地文件；
git update-index --assume-unchanged /path/file
# 设置忽略跟踪(2)，这个本地会同步远程的，本地的变更不会推送给远程
git update-index --skip-worktree /path/file
#恢复跟踪
git update-index --no-assume-unchanged /path/to/files

# (1)assume-unchanged：这个会关闭文件与远程仓库的跟踪，认为这个文件远程仓库是不会修改，所以每次pull都是本地的文件
# (2)skip-worktree：这个不会关闭文件与远程仓库的跟踪，只是告诉Git不要跟踪对本地文件/文件夹的更改。如果远端仓库内容有变化，pull时会拉取最新的变化，并提示冲突，但因为没有跟踪本地更改，所以需要no-skip-worktree再合并最新的变化。

```