# DetectDee
<font color="red">*严正声明：本文与本工具仅限于技术讨论与分享，严禁用于非法途径。</font>

神探狄仁杰: 在[社交网络](site.md)上通过**用户名，电子邮件或电话**搜索社交媒体账户
![example.gif](https://s2.loli.net/2023/04/30/FZ1QtKoGud4xVPW.gif)
## 特性
- 集成**网络安全从业者**常用网站
- 通过**用户名，电子邮件或电话**查找社交媒体账户
- 可拓展，简单易用的模板
- 集成社交网站的移动端版本
## 安装
### 直接下载运行(推荐)
https://github.com/piaolin/DetectDee/releases
### 编译运行
```shell
git clone https://github.com/piaolin/DetectDee.git
cd DetectDee
go mod tidy
go run .
```
## 用法
[English](README.md)

[中文文档](README_ZH.md)
### 探测
```text
Usage:
  DetectDee detect [flags]

Flags:
  -c, --check          self-check
  -f, --file string    Site data file (default "data.json")
  -h, --help           help for detect
  -n, --name strings   name[s]
      --nsfw           Include checking of NSFW sites from default list.        
      --precisely      Check precisely
  -p, --proxy string   Make requests over a proxy. e.g. socks5://127.0.0.1:1080 
  -r, --retry int      Retry times after request failed (default 3)
  -s, --site strings   Limit analysis to just the listed sites. Add multiple options to specify more than one site.
  -t, --timeout int    Time (in seconds) to wait for response to requests (default 10)

Global Flags:
  -v, --verbose   verbose output
```

单一用户名搜索:
```shell
./DetectDee detect -n piaolin
```

多用户名搜索:
```shell
./DetectDee detect -n piaolin,blue
```

指定网站搜索:
```shell
./DetectDee detect -n piaolin -s github,v2ex
```

## 贡献
希望您能帮助我们开发“神探狄仁杰”，您的每一份提交都会得到我的重视。

下面是一些我们希望得到您帮助的事情:

- 增加新的站点支持，可以通知我某个网站有可用的接口，或直接编写[JSON](t.json)
- 恢复对过去因误报而被删除的网站的支持

## 待办
- 对探测得到的社交网站进行**密码撞库**
- 更多的网站支持
- 保密

## 支持网站
### CyberSecurity
1. ![](https://www.google.com/s2/favicons?domain=https://www.freebuf.com/) [Freebuf](https://www.freebuf.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://hackerone.com/) [HackerOne](https://hackerone.com/)
3. ![](https://www.google.com/s2/favicons?domain=https://bugcrowd.com/) [BugCrowd](https://bugcrowd.com/)
4. ![](https://www.google.com/s2/favicons?domain=https://www.jarvisoj.com/) [Jarvisoj](https://www.jarvisoj.com/)
5. ![](https://www.google.com/s2/favicons?domain=https://vulfocus.cn/) [VulFocus](https://vulfocus.cn/)
6. ![](https://www.google.com/s2/favicons?domain=https://www.secrss.com/) [Secrss](https://www.secrss.com/)
7. ![](https://www.google.com/s2/favicons?domain=https://www.virustotal.com/) [VirusTotal](https://www.virustotal.com/)
8. ![](https://www.google.com/s2/favicons?domain=https://new.bugku.com/) [CTF论剑场](https://new.bugku.com/)
9. ![](https://www.google.com/s2/favicons?domain=https://www.yystv.cn/) [游研社](https://www.yystv.cn/)
10. ![](https://www.google.com/s2/favicons?domain=https://xz.aliyun.com) [先知社区](https://xz.aliyun.com)
11. ![](https://www.google.com/s2/favicons?domain=https://zone.huoxian.cn/) [火线社区](https://zone.huoxian.cn/)
12. ![](https://www.google.com/s2/favicons?domain=https://forum.ywhack.com/) [棱角社区](https://forum.ywhack.com/)
13. ![](https://www.google.com/s2/favicons?domain=https://www.thehackerworld.com/) [TheHackerWorld](https://www.thehackerworld.com/)
14. ![](https://www.google.com/s2/favicons?domain=https://x.threatbook.com) [微步在线](https://x.threatbook.com)
15. ![](https://www.google.com/s2/favicons?domain=https://www.secpulse.com) [安全脉搏](https://www.secpulse.com)
16. ![](https://www.google.com/s2/favicons?domain=https://hackerrank.com/) [HackerRank](https://hackerrank.com/)
17. ![](https://www.google.com/s2/favicons?domain=https://security.eastmoney.com) [东方财富SRC](https://security.eastmoney.com)
18. ![](https://www.google.com/s2/favicons?domain=https://www.sec-wiki.com/) [Sec-Wiki](https://www.sec-wiki.com/)
19. ![](https://www.google.com/s2/favicons?domain=https://forum.90sec.com/) [90sec](https://forum.90sec.com/)
20. ![](https://www.google.com/s2/favicons?domain=https://play.google.com) [Googleplay](https://play.google.com)
21. ![](https://www.google.com/s2/favicons?domain=https://www.bugbank.cn/) [漏洞银行](https://www.bugbank.cn/)
22. ![](https://www.google.com/s2/favicons?domain=https://www.ichunqiu.com) [i春秋](https://www.ichunqiu.com)
23. ![](https://www.google.com/s2/favicons?domain=https://www.seebug.org/) [Seebug](https://www.seebug.org/)
24. ![](https://www.google.com/s2/favicons?domain=https://0x00sec.org/) [0x00sec](https://0x00sec.org/)
25. ![](https://www.google.com/s2/favicons?domain=https://www.anquanke.com) [安全客](https://www.anquanke.com)
26. ![](https://www.google.com/s2/favicons?domain=https://www.infosecurity-magazine.com/) [Infosecurity-magazine](https://www.infosecurity-magazine.com/)
27. ![](https://www.google.com/s2/favicons?domain=http://www.xsssql.com/) [漏洞空间](http://www.xsssql.com/)
28. ![](https://www.google.com/s2/favicons?domain=https://www.t00ls.com/) [T00ls](https://www.t00ls.com/)
29. ![](https://www.google.com/s2/favicons?domain=https://www.52pojie.cn/) [吾爱破解](https://www.52pojie.cn/)
30. ![](https://www.google.com/s2/favicons?domain=https://www.mozhe.cn/) [墨者学院](https://www.mozhe.cn/)
31. ![](https://www.google.com/s2/favicons?domain=https://threatpost.com/) [ThreatPost](https://threatpost.com/)
32. ![](https://www.google.com/s2/favicons?domain=https://www.vulbox.com/) [漏洞盒子](https://www.vulbox.com/)
33. ![](https://www.google.com/s2/favicons?domain=https://www.bugku.com/) [BugKu](https://www.bugku.com/)
34. ![](https://www.google.com/s2/favicons?domain=https://forum.butian.net) [奇安信攻防社区](https://forum.butian.net)
35. ![](https://www.google.com/s2/favicons?domain=https://bbs.zkaq.cn/) [Track安全社区](https://bbs.zkaq.cn/)
36. ![](https://www.google.com/s2/favicons?domain=https://paper.seebug.org/) [Seebug-paper](https://paper.seebug.org/)
37. ![](https://www.google.com/s2/favicons?domain=https://tttang.com/) [跳跳糖](https://tttang.com/)
38. ![](https://www.google.com/s2/favicons?domain=https://tryhackme.com/) [TryHackMe](https://tryhackme.com/)
39. ![](https://www.google.com/s2/favicons?domain=https://www.aqniu.com/) [安全牛](https://www.aqniu.com/)

### Programmer
1. ![](https://www.google.com/s2/favicons?domain=https://opensource.com/) [OpenSource](https://opensource.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://www.infoq.cn/) [infoQ](https://www.infoq.cn/)
3. ![](https://www.google.com/s2/favicons?domain=https://www.twle.cn/) [简单教程](https://www.twle.cn/)
4. ![](https://www.google.com/s2/favicons?domain=https://quizlet.com/) [Quizlet](https://quizlet.com/)
5. ![](https://www.google.com/s2/favicons?domain=https://gitee) [码云](https://gitee)
6. ![](https://www.google.com/s2/favicons?domain=https://leetcode.com/) [力扣国际版](https://leetcode.com/)
7. ![](https://www.google.com/s2/favicons?domain=https://leetcode.cn/) [力扣中国版](https://leetcode.cn/)
8. ![](https://www.google.com/s2/favicons?domain=https://forum.sublimetext.com/) [SublimeForum](https://forum.sublimetext.com/)
9. ![](https://www.google.com/s2/favicons?domain=https://www.producthunt.com/) [Producthunt](https://www.producthunt.com/)
10. ![](https://www.google.com/s2/favicons?domain=https://studygolang.com/) [Go语言中文网](https://studygolang.com/)
11. ![](https://www.google.com/s2/favicons?domain=https://ld246.com/) [链滴](https://ld246.com/)
12. ![](https://www.google.com/s2/favicons?domain=https://my.oschina.net/) [开源中国](https://my.oschina.net/)
13. ![](https://www.google.com/s2/favicons?domain=https://github.com) [Github](https://github.com)
14. ![](https://www.google.com/s2/favicons?domain=https://github.io) [GithubBlog](https://github.io)
15. ![](https://www.google.com/s2/favicons?domain=https://segmentfault.com/) [思否](https://segmentfault.com/)
16. ![](https://www.google.com/s2/favicons?domain=https://news.ycombinator.com/) [HackerNews](https://news.ycombinator.com/)
17. ![](https://www.google.com/s2/favicons?domain=https://unsplash.com/) [unsplash](https://unsplash.com/)
18. ![](https://www.google.com/s2/favicons?domain=https://nextcloud.com/) [Nextcloud](https://nextcloud.com/)
19. ![](https://www.google.com/s2/favicons?domain=https://ruby-china.org/) [Ruby-China](https://ruby-china.org/)
20. ![](https://www.google.com/s2/favicons?domain=https://rubygems.org/) [RubyGems](https://rubygems.org/)
21. ![](https://www.google.com/s2/favicons?domain=https://packagist.org/) [packagist](https://packagist.org/)
22. ![](https://www.google.com/s2/favicons?domain=https://www.iteye.com/) [ITeye](https://www.iteye.com/)
23. ![](https://www.google.com/s2/favicons?domain=https://gitbook.com/) [GitBook](https://gitbook.com/)
24. ![](https://www.google.com/s2/favicons?domain=https://sspai.com/) [少数派](https://sspai.com/)
25. ![](https://www.google.com/s2/favicons?domain=https://gitlab.com/) [Gitlab](https://gitlab.com/)
26. ![](https://www.google.com/s2/favicons?domain=https://pypi.org/) [pypi](https://pypi.org/)
27. ![](https://www.google.com/s2/favicons?domain=https://trello.com/) [Trello](https://trello.com/)
28. ![](https://www.google.com/s2/favicons?domain=https://pastebin.com/) [Pastebin](https://pastebin.com/)
29. ![](https://www.google.com/s2/favicons?domain=https://community.oracle.com/) [Oracle](https://community.oracle.com/)

### Social
1. ![](https://www.google.com/s2/favicons?domain=https://dig.chouti.com/) [抽屉热榜](https://dig.chouti.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://tieba.baidu.com) [百度贴吧](https://tieba.baidu.com)
3. ![](https://www.google.com/s2/favicons?domain=https://t.me/) [Telegram](https://t.me/)
4. ![](https://www.google.com/s2/favicons?domain=https://www.douban.com/) [豆瓣](https://www.douban.com/)
5. ![](https://www.google.com/s2/favicons?domain=https://v2ex.com) [V2ex](https://v2ex.com)
6. ![](https://www.google.com/s2/favicons?domain=https://steamcommunity.com/) [SteamGroup](https://steamcommunity.com/)
7. ![](https://www.google.com/s2/favicons?domain=http://jandan.net/) [煎蛋](http://jandan.net/)
8. ![](https://www.google.com/s2/favicons?domain=https://my.zol.com.cn/) [中关村在线](https://my.zol.com.cn/)
9.  ![](https://www.google.com/s2/favicons?domain=https://www.zhihu.com/) [知乎](https://www.zhihu.com/)
10. ![](https://www.google.com/s2/favicons?domain=https://www.ifanr.com/) [爱范儿](https://www.ifanr.com/)
11. ![](https://www.google.com/s2/favicons?domain=https://www.snapchat.com/) [Snapchat](https://www.snapchat.com/)
12. ![](https://www.google.com/s2/favicons?domain=https://www.reddit.com/) [Reddit](https://www.reddit.com/)
13. ![](https://www.google.com/s2/favicons?domain=https://www.9gag.com/) [9GAG](https://www.9gag.com/)
14. ![](https://www.google.com/s2/favicons?domain=https://lolchess.gg/) [云顶之弈国际服](https://lolchess.gg/)
15. ![](https://www.google.com/s2/favicons?domain=https://tiktok.com/) [Tiktok](https://tiktok.com/)
16. ![](https://www.google.com/s2/favicons?domain=https://www.eastmoney.com) [东方财富](https://www.eastmoney.com)
17. ![](https://www.google.com/s2/favicons?domain=https://twitter.com/) [Twitter](https://twitter.com/)
18. ![](https://www.google.com/s2/favicons?domain=https://minecraft.net/) [我的世界](https://minecraft.net/)
19. ![](https://www.google.com/s2/favicons?domain=https://gaoloumi.cc/) [Gaoloumi](https://gaoloumi.cc/)
20. ![](https://www.google.com/s2/favicons?domain=https://zhidao.baidu.com) [百度知道](https://zhidao.baidu.com)
21. ![](https://www.google.com/s2/favicons?domain=https://www.quora.com/) [Quora](https://www.quora.com/)
22. ![](https://www.google.com/s2/favicons?domain=https://bbs.qyer.com) [穷游网](https://bbs.qyer.com)
23. ![](https://www.google.com/s2/favicons?domain=https://sourceforge.net/) [Sourceforge](https://sourceforge.net/)
24. ![](https://www.google.com/s2/favicons?domain=https://open.spotify.com/) [Spotify](https://open.spotify.com/)
25. ![](https://www.google.com/s2/favicons?domain=https://www.twitch.tv/) [Twitch](https://www.twitch.tv/)

### Video
1. ![](https://www.google.com/s2/favicons?domain=https://wordpress.com/) [WordPress](https://wordpress.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://www.bilibili.com) [哔哩哔哩](https://www.bilibili.com)
3. ![](https://www.google.com/s2/favicons?domain=https://www.youtube.com/) [Youtube](https://www.youtube.com/)
4. ![](https://www.google.com/s2/favicons?domain=https://www.wikipedia.org/) [维基百科](https://www.wikipedia.org/)

### Blog
1. ![](https://www.google.com/s2/favicons?domain=https://blog.csdn.net/) [CSDN](https://blog.csdn.net/)
2. ![](https://www.google.com/s2/favicons?domain=https://www.cnblogs.com/) [博客园](https://www.cnblogs.com/)

## 致谢
https://github.com/sherlock-project/sherlock