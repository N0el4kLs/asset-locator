<h2 align="center">assetlocator</h2>

## 🎯介绍

该工具用于目标资产归属查找，包括 `爱站权重查询` `公司备案信息查询`. 


### 🧑‍💻使用的第三方接口

|   接口名称   |                                  地址                                  |   说明    |
|:--------:|:--------------------------------------------------------------------:|:-------:|
|    ✅爱站    |                  https://baidurank.aizhan.com/baidu                  |  权重查询   |
| ✅pearktrue | https://api.pearktrue.cn/api/website/weight.php?domain=www.baidu.com |  权重查询   |
| ✅API Store |            https://apis.jxcxin.cn/api/icp?name=baidu.com             |  权重查询   |
|   ✅站长工具   |https://seo.chinaz.com/| ICP备案查询 |
| ☑️夏柔 - 分享的API永不收费 | https://api.aa1.cn/#%E6%9C%80%E6%96%B0ApiList |  ICP备案查询  |
| ☑️Webscan | https://www.webscan.cc/ | IP2domain |


感谢以上网站作者和开发者所作出的贡献, 没有你们就没有这个项目！🙏

如果您有更好的接口，欢迎提交 PR 或者 Issue.



## 命令行使用

安装
```bash
go get -u github.com/N0el4kLs/asset-locator/lib@latest

asset-locator -h

cat 1.txt | asset-locator -w
```


## 第三方包使用

```bash
go get -u github.com/N0el4kLs/asset-locator/lib@latest
```
