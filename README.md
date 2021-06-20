## chaos-knight
一个简单的推荐框架(demo), 设计了推荐分层框架, 并提供了一个配置demo(`conf.json`), 通过`interpreter.go`
的实现将这些推荐的层关联起来.

使用者需要自己实现 `recall`, `feature`, `filter`, `model`, `rerank`
这5层下面的实例, 例如`hot.go`, `u2r.go`, `u2r2r`就是`recall.go`中定义的interface的具体实现.

另外还需要用户自己选择完成`rpc`, `log`, `message queue`, `db`, `kv`等模块来变成一个真正可用的推荐服务

值得一说的是模型层的实现, 根据经验来看, 它可以在golang本地实现, 可以引入第三方库实现
实现(例如[tensorflow](https://www.tensorflow.org/install/lang_go), [leaves](https://github.com/dmitryikh/leaves)), 也可以通过rpc调用其他推荐服务实现(例如`tensorflow serving`)

## 流量ab验证
未来验证数据是正交的, 可以运行
```bash
go build
./chaos-knight > log.csv
```
然后打开`data_validation.ipynb`查看简单的数据正交性的必要性验证