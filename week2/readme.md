#### 错误类型

#### 作业
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

sql.ErrNoRows:sql: no rows in result set
此错误属于正常错误，表示查询行的时候没有查询到数据，所以应该使用wrap这个error返回给上一层调用者