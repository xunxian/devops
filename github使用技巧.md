# github切换token鉴权
```
# 使用用户名密码进行push/pull时爆如下错误：
remote: Support for password authentication was removed on August 13, 2021. Please use a personal access token instead.
remote: Please see https://github.blog/2020-12-15-token-authentication-requirements-for-git-operations/ for more information.
fatal: 无法访问 'https://github.com/xunxian/devops.git/'：The requested URL returned error: 403
原因：github推行使用token进行登录，设置方法如下：
1、删除本地缓存的用户名及密码：git config --system --unset credential.helper
2、登录github增加token：
	头像 -> Settings -> Developer settings -> Personal access tokens
	Generate new token
	copy token
3、在本地使用git push origin main
	输入用户名：xxxxx@xx.xxx
	输入密码：粘贴上面复制的token，回车即可。
永久保存：
git config --edit
[remote "origin"]
        url = https://github_username:github_token@github.com/github_username/repository_name
        fetch = +refs/heads/*:refs/remotes/origin/*
```

