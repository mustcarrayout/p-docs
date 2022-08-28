- vim 在行尾新增一行
- ```shell
  :%s/$/\0\r/g
  ```
- [[/etc/profile、/etc/bashrc、~/.bashrc、~/.bash_profile之间的区别]]
	- /etc/profile contains Linux system wide environment and startup programs. It is used by all users with bash, ksh, sh shell. Usually used to set PATH variable, user limits, and other settings for user. It only runs for login shell. If you wanted to make large changes or application specific changes use /etc/profile.d directory.
	- 系统变量=/etc/profile,影响所有用户
	- 用户变量=~/.bashrc 然后加source
	- /etc/profile做了什么事情?
		- 登录shell之后执行**/etc/profile.d/*.sh**脚本
- ![vim](https://www.runoob.com/wp-content/uploads/2015/10/vi-vim-cheat-sheet-sch.gif)



# mac

- ```shell
  # 允许任意来源的软件安装
  sudo spctl --master-disable
  # 关闭允许任意来源的软件安装
  sudo spctl --master-enable
  ```