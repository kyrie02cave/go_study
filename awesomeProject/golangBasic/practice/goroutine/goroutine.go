package goroutine


/*
1.使用goroutine和channel实现一个计算int64随机数各位数和的程序。
 - 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
 - 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
 - 主goroutine从resultChan取出结果并打印到终端输出
2.为了保证业务代码的执行性能将之前写的日志库改写为异步记录日志方式。
**/