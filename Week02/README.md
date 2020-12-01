# 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

### 答
应该wrap这个error给service层，交由它判断如何处理这个error。
我的理解是当app没有这个数据的时候会报错则在service层返回错误并记录日志，当app可以接受空数据的时候则在service层判断转为空数据正常200返回。
具体实现由app与下游协议决定。