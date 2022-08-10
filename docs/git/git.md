
# git

## 什么是git？

### 看图理解

1. 从中看出了哪些问题？
2. 从中知道了哪些问题的答案？

## .gitignore

1. 新规则只考虑新的文件**不向后兼容**
2. .gitignore配置文件是按行从上到下进行规则匹配的

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
```