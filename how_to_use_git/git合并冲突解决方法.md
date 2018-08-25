# [git合并冲突解决方法](https://www.cnblogs.com/ZhangRuoXu/p/6706571.html)

- 1、git merge冲突了，根据提示找到冲突的文件，解决冲突

如果文件有冲突，那么会有类似的标记

- 2、修改完之后，执行git add 冲突文件名
- 3、git commit
  `注意`:没有-m选项
  进去类似于vim的操作界面，把conflict相关的行删除掉
- 4、直接push就可以了，因为刚刚已经执行过相关merge操作了

相关的操作如下

- 冲突产生

```
[root@Monitor Demo]# git branch #当前在master分支下
* master
  psr/psr-01
  psr/psr-02
[root@Monitor Demo]# git checkout psr/psr-02 #切换到psr/psr-02分支下
Switched to branch 'psr/psr-02'
[root@Monitor Demo]# git branch
  master
  psr/psr-01
* psr/psr-02
[root@Monitor Demo]# ls
LICENSE  new_remote_branch.txt  psr_psr-02.txt  README.md
[root@Monitor Demo]# vim psr_psr-02.txt #修改psr/psr-02分支上的文件
[root@Monitor Demo]# git add psr_psr-02.txt
[root@Monitor Demo]# git commit -m 'psr_psr-02.txt has changed on psr/psr-02 branch' #提交到暂存区
[psr/psr-02 62ca72c] psr_psr-02.txt has changed on psr/psr-02 branch
 1 files changed, 6 insertions(+), 0 deletions(-)
[root@Monitor Demo]# git checkout master #切换到master分支下
Switched to branch 'master'
[root@Monitor Demo]# vim psr_psr-02.txt #在master分支下也对psr_psr-02.txt进行修改
[root@Monitor Demo]# git add psr_psr-02.txt
[root@Monitor Demo]# git commit -m 'changed this file on master branch'
[master 282fbeb] changed this file on master branch
 1 files changed, 2 insertions(+), 0 deletions(-)
[root@Monitor Demo]# git merge psr/psr-02 #把psr/psr-02分支合并到当前分支，这时提示冲突了
Auto-merging psr_psr-02.txt
CONFLICT (content): Merge conflict in psr_psr-02.txt
Automatic merge failed; fix conflicts and then commit the result.
```

- 冲突解决过程
  冲突文件的格式基本如下

```
<<<<<<<到=======是在当前分支合并之前的文件内容
=======到>>>>>>> psr/psr-02是在其它分支下修改的内容
需要在这个两个版本中选择一个，然后把标记符号也要一起删除
```

```
<<<<<<< HEAD
add some lines on master branch

add some lines on psr/psr-01 branch

2016年12月13日14:43:34 changed after psr/psr-02
=======
1
2
3
4
5
>>>>>>> psr/psr-02
```

冲突文件

```
vim psr_psr-02.txt

<<<<<<< HEAD
add some lines on master branch

add some lines on psr/psr-01 branch

2016年12月13日14:43:34 changed after psr/psr-02
=======
1
2
3
4
5
>>>>>>> psr/psr-02
```

修改冲突文件

```
# vim psr_psr-02.txt
README.md

I'am in new branch psr/psr-02 based on psr/psr-01

add some lines on master branch

add some lines on psr/psr-01 branch

2016年12月13日14:43:34 changed after psr/psr-02
```

- 添加冲突的文件，然后就可以直接push了

```
Merge branch 'psr/psr-02'

Conflicts:
    psr_psr-02.txt
#
# It looks like you may be committing a MERGE.
# If this is not correct, please remove the file
#   .git/MERGE_HEAD
# and try again.
#

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
# On branch master
# Your branch is ahead of 'origin/master' by 1 commit.
#
```

需要删掉下面这两行

```
Conflicts:
    psr_psr-02.txt
```

- 执行git push 操作

```
git push origin master
```

Yesterday is history.
Tomorrow is a mystery. 
But today is a gift. 
That is why it's called the present. 
The old game: give a wolf a taste, then keep him hungry.