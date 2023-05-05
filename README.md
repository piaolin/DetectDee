# DetectDee
<font color="red">*Disclaimer: This article and this tool are for technical discussion and sharing only. Illegal use is strictly prohibited.</font>

DetectDee: Hunt down social media accounts by **username, email or phone** across [social networks](site.md)
![example.gif](https://s2.loli.net/2023/04/30/FZ1QtKoGud4xVPW.gif)
## Feat
- Includes sites frequently used by **CyberSecurity practitioners**
- Hunt down social media accounts by **username, email or phone**
- Precise thread control and custom request headers are used to prevent WAF recognition
- Extensible, simple, and easy-to-use template
- Integration of mobile versions of social networking sites

## Install
### Download(recommend)
https://github.com/piaolin/DetectDee/releases
### Compile
```shell
git clone https://github.com/piaolin/DetectDee.git
cd DetectDee
go mod tidy
go run .
```
## Usage
[English](README.md)

[中文文档](README_ZH.md)
### Detect
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

Please update the data for the first time
```shell
./DetectDee update
```

To search for only one user:
```shell
./DetectDee detect -n piaolin
```

To search for more than one user:
```shell
./DetectDee detect -n piaolin,blue
```

To search in specified site:
```shell
./DetectDee detect -n piaolin -s github,v2ex
```

## Contributing
We would love to have you help us with the development of DetectDee. Each and every contribution is greatly valued!

Here are some things we would appreciate your help on:

- Addition of new site support, You can notify me that a site has an interface available, or you can write [JSON](t.json) directly
- Bringing back site support of sites that have been removed in the past due to false positives

## Todo
- **Credential Stuffing** for result
- More site
- Secret

## Supported site
### CyberSecurity
1. ![](https://www.google.com/s2/favicons?domain=https://www.freebuf.com/) [Freebuf](https://www.freebuf.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://hackerone.com/) [HackerOne](https://hackerone.com/)
3. ![](https://www.google.com/s2/favicons?domain=https://bugcrowd.com/) [BugCrowd](https://bugcrowd.com/)
4. ![](https://www.google.com/s2/favicons?domain=https://www.jarvisoj.com/) [Jarvisoj](https://www.jarvisoj.com/)
5. ![](https://www.google.com/s2/favicons?domain=https://vulfocus.cn/) [VulFocus](https://vulfocus.cn/)
6. ![](https://www.google.com/s2/favicons?domain=https://www.secrss.com/) [Secrss](https://www.secrss.com/)
7. ![](https://www.google.com/s2/favicons?domain=https://www.virustotal.com/) [VirusTotal](https://www.virustotal.com/)
8. ![](https://www.google.com/s2/favicons?domain=https://new.bugku.com/) [newBugKu](https://new.bugku.com/)
9. ![](https://www.google.com/s2/favicons?domain=https://www.yystv.cn/) [yystv](https://www.yystv.cn/)
10. ![](https://www.google.com/s2/favicons?domain=https://xz.aliyun.com) [XZ](https://xz.aliyun.com)
11. ![](https://www.google.com/s2/favicons?domain=https://zone.huoxian.cn/) [huoxian](https://zone.huoxian.cn/)
12. ![](https://www.google.com/s2/favicons?domain=https://forum.ywhack.com/) [ywhack](https://forum.ywhack.com/)
13. ![](https://www.google.com/s2/favicons?domain=https://www.thehackerworld.com/) [TheHackerWorld](https://www.thehackerworld.com/)
14. ![](https://www.google.com/s2/favicons?domain=https://x.threatbook.com) [ThreatBook](https://x.threatbook.com)
15. ![](https://www.google.com/s2/favicons?domain=https://www.secpulse.com) [SecPulse](https://www.secpulse.com)
16. ![](https://www.google.com/s2/favicons?domain=https://hackerrank.com/) [HackerRank](https://hackerrank.com/)
17. ![](https://www.google.com/s2/favicons?domain=https://security.eastmoney.com) [EastMoney-src](https://security.eastmoney.com)
18. ![](https://www.google.com/s2/favicons?domain=https://www.sec-wiki.com/) [Sec-Wiki](https://www.sec-wiki.com/)
19. ![](https://www.google.com/s2/favicons?domain=https://forum.90sec.com/) [90sec](https://forum.90sec.com/)
20. ![](https://www.google.com/s2/favicons?domain=https://play.google.com) [Googleplay](https://play.google.com)
21. ![](https://www.google.com/s2/favicons?domain=https://www.bugbank.cn/) [BugBank](https://www.bugbank.cn/)
22. ![](https://www.google.com/s2/favicons?domain=https://www.ichunqiu.com) [ichunqiu](https://www.ichunqiu.com)
23. ![](https://www.google.com/s2/favicons?domain=https://www.seebug.org/) [Seebug](https://www.seebug.org/)
24. ![](https://www.google.com/s2/favicons?domain=https://0x00sec.org/) [0x00sec](https://0x00sec.org/)
25. ![](https://www.google.com/s2/favicons?domain=https://www.anquanke.com) [anquanke](https://www.anquanke.com)
26. ![](https://www.google.com/s2/favicons?domain=https://www.infosecurity-magazine.com/) [Infosecurity-magazine](https://www.infosecurity-magazine.com/)
27. ![](https://www.google.com/s2/favicons?domain=http://www.xsssql.com/) [xsssql](http://www.xsssql.com/)
28. ![](https://www.google.com/s2/favicons?domain=https://www.t00ls.com/) [T00ls](https://www.t00ls.com/)
29. ![](https://www.google.com/s2/favicons?domain=https://www.52pojie.cn/) [52pojie](https://www.52pojie.cn/)
30. ![](https://www.google.com/s2/favicons?domain=https://www.mozhe.cn/) [Mozhe](https://www.mozhe.cn/)
31. ![](https://www.google.com/s2/favicons?domain=https://threatpost.com/) [ThreatPost](https://threatpost.com/)
32. ![](https://www.google.com/s2/favicons?domain=https://www.vulbox.com/) [Vulbox](https://www.vulbox.com/)
33. ![](https://www.google.com/s2/favicons?domain=https://www.bugku.com/) [BugKu](https://www.bugku.com/)
34. ![](https://www.google.com/s2/favicons?domain=https://forum.butian.net) [BuTian](https://forum.butian.net)
35. ![](https://www.google.com/s2/favicons?domain=https://bbs.zkaq.cn/) [Track](https://bbs.zkaq.cn/)
36. ![](https://www.google.com/s2/favicons?domain=https://paper.seebug.org/) [Seebug-paper](https://paper.seebug.org/)
37. ![](https://www.google.com/s2/favicons?domain=https://tttang.com/) [tttang](https://tttang.com/)
38. ![](https://www.google.com/s2/favicons?domain=https://tryhackme.com/) [TryHackMe](https://tryhackme.com/)
39. ![](https://www.google.com/s2/favicons?domain=https://www.aqniu.com/) [aqniu](https://www.aqniu.com/)

### Programmer
1. ![](https://www.google.com/s2/favicons?domain=https://opensource.com/) [OpenSource](https://opensource.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://www.infoq.cn/) [infoQ](https://www.infoq.cn/)
3. ![](https://www.google.com/s2/favicons?domain=https://www.twle.cn/) [twle](https://www.twle.cn/)
4. ![](https://www.google.com/s2/favicons?domain=https://quizlet.com/) [Quizlet](https://quizlet.com/)
5. ![](https://www.google.com/s2/favicons?domain=https://gitee.com) [Gitee](https://gitee.com)
6. ![](https://www.google.com/s2/favicons?domain=https://leetcode.com/) [Leetcode](https://leetcode.com/)
7. ![](https://www.google.com/s2/favicons?domain=https://leetcode.cn/) [Leetcode-CN](https://leetcode.cn/)
8. ![](https://www.google.com/s2/favicons?domain=https://forum.sublimetext.com/) [SublimeForum](https://forum.sublimetext.com/)
9. ![](https://www.google.com/s2/favicons?domain=https://www.producthunt.com/) [Producthunt](https://www.producthunt.com/)
10. ![](https://www.google.com/s2/favicons?domain=https://studygolang.com/) [StudyGolang](https://studygolang.com/)
11. ![](https://www.google.com/s2/favicons?domain=https://ld246.com/) [ld246](https://ld246.com/)
12. ![](https://www.google.com/s2/favicons?domain=https://my.oschina.net/) [OSChina](https://my.oschina.net/)
13. ![](https://www.google.com/s2/favicons?domain=https://github.com) [Github](https://github.com)
14. ![](https://www.google.com/s2/favicons?domain=https://github.io) [GithubBlog](https://github.io)
15. ![](https://www.google.com/s2/favicons?domain=https://segmentfault.com/) [SegmentFault](https://segmentfault.com/)
16. ![](https://www.google.com/s2/favicons?domain=https://news.ycombinator.com/) [HackerNews](https://news.ycombinator.com/)
17. ![](https://www.google.com/s2/favicons?domain=https://unsplash.com/) [unsplash](https://unsplash.com/)
18. ![](https://www.google.com/s2/favicons?domain=https://nextcloud.com/) [Nextcloud](https://nextcloud.com/)
19. ![](https://www.google.com/s2/favicons?domain=https://ruby-china.org/) [Ruby-China](https://ruby-china.org/)
20. ![](https://www.google.com/s2/favicons?domain=https://rubygems.org/) [RubyGems](https://rubygems.org/)
21. ![](https://www.google.com/s2/favicons?domain=https://packagist.org/) [packagist](https://packagist.org/)
22. ![](https://www.google.com/s2/favicons?domain=https://www.iteye.com/) [ITeye](https://www.iteye.com/)
23. ![](https://www.google.com/s2/favicons?domain=https://gitbook.com/) [GitBook](https://gitbook.com/)
24. ![](https://www.google.com/s2/favicons?domain=https://sspai.com/) [SSPai](https://sspai.com/)
25. ![](https://www.google.com/s2/favicons?domain=https://gitlab.com/) [Gitlab](https://gitlab.com/)
26. ![](https://www.google.com/s2/favicons?domain=https://pypi.org/) [pypi](https://pypi.org/)
27. ![](https://www.google.com/s2/favicons?domain=https://trello.com/) [Trello](https://trello.com/)
28. ![](https://www.google.com/s2/favicons?domain=https://pastebin.com/) [Pastebin](https://pastebin.com/)
29. ![](https://www.google.com/s2/favicons?domain=https://community.oracle.com/) [Oracle](https://community.oracle.com/)

### Social
1. ![](https://www.google.com/s2/favicons?domain=https://dig.chouti.com/) [Chouti](https://dig.chouti.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://tieba.baidu.com) [Baidu-Tieba](https://tieba.baidu.com)
3. ![](https://www.google.com/s2/favicons?domain=https://t.me/) [Telegram](https://t.me/)
4. ![](https://www.google.com/s2/favicons?domain=https://www.douban.com/) [Douban](https://www.douban.com/)
5. ![](https://www.google.com/s2/favicons?domain=https://v2ex.com) [V2ex](https://v2ex.com)
6. ![](https://www.google.com/s2/favicons?domain=https://steamcommunity.com/) [SteamGroup](https://steamcommunity.com/)
7. ![](https://www.google.com/s2/favicons?domain=http://jandan.net/) [Jandan](http://jandan.net/)
8. ![](https://www.google.com/s2/favicons?domain=https://my.zol.com.cn/) [Zol](https://my.zol.com.cn/)
9.  ![](https://www.google.com/s2/favicons?domain=https://www.zhihu.com/) [Zhihu](https://www.zhihu.com/)
10. ![](https://www.google.com/s2/favicons?domain=https://www.ifanr.com/) [iFanr](https://www.ifanr.com/)
11. ![](https://www.google.com/s2/favicons?domain=https://www.snapchat.com/) [Snapchat](https://www.snapchat.com/)
12. ![](https://www.google.com/s2/favicons?domain=https://www.reddit.com/) [Reddit](https://www.reddit.com/)
13. ![](https://www.google.com/s2/favicons?domain=https://www.9gag.com/) [9GAG](https://www.9gag.com/)
14. ![](https://www.google.com/s2/favicons?domain=https://lolchess.gg/) [TFT](https://lolchess.gg/)
15. ![](https://www.google.com/s2/favicons?domain=https://tiktok.com/) [Tiktok](https://tiktok.com/)
16. ![](https://www.google.com/s2/favicons?domain=https://www.eastmoney.com) [EastMoney](https://www.eastmoney.com)
17. ![](https://www.google.com/s2/favicons?domain=https://twitter.com/) [Twitter](https://twitter.com/)
18. ![](https://www.google.com/s2/favicons?domain=https://minecraft.net/) [Minecraft](https://minecraft.net/)
19. ![](https://www.google.com/s2/favicons?domain=https://gaoloumi.cc/) [Gaoloumi](https://gaoloumi.cc/)
20. ![](https://www.google.com/s2/favicons?domain=https://zhidao.baidu.com) [Baidu-Zhidao](https://zhidao.baidu.com)
21. ![](https://www.google.com/s2/favicons?domain=https://www.quora.com/) [Quora](https://www.quora.com/)
22. ![](https://www.google.com/s2/favicons?domain=https://bbs.qyer.com) [qyer](https://bbs.qyer.com)
23. ![](https://www.google.com/s2/favicons?domain=https://sourceforge.net/) [Sourceforge](https://sourceforge.net/)
24. ![](https://www.google.com/s2/favicons?domain=https://open.spotify.com/) [Spotify](https://open.spotify.com/)
25. ![](https://www.google.com/s2/favicons?domain=https://www.twitch.tv/) [Twitch](https://www.twitch.tv/)

### Video
1. ![](https://www.google.com/s2/favicons?domain=https://wordpress.com/) [WordPress](https://wordpress.com/)
2. ![](https://www.google.com/s2/favicons?domain=https://www.bilibili.com) [Bilibili](https://www.bilibili.com)
3. ![](https://www.google.com/s2/favicons?domain=https://www.youtube.com/) [Youtube](https://www.youtube.com/)
4. ![](https://www.google.com/s2/favicons?domain=https://www.wikipedia.org/) [Wikipedia](https://www.wikipedia.org/)

### Blog
1. ![](https://www.google.com/s2/favicons?domain=https://blog.csdn.net/) [CSDN](https://blog.csdn.net/)
2. ![](https://www.google.com/s2/favicons?domain=https://www.cnblogs.com/) [cnblogs](https://www.cnblogs.com/)

## Thanks
https://github.com/sherlock-project/sherlock