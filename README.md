hashcash
工作量认证算法
========
## proof of work by hashcash



给出如下字符串，让机器工作3秒钟
~~~ go
"taopopoo@126.com"
~~~

开始工作
~~~ go
key := hashcash.Work("taopopoo@126.com", 20)
~~~
工作大约3秒钟后得到key


验证工作是否完成
~~~ go
ok := hashcash.Check("taopopoo@126.com", 20, key)
if ok{
	fmt.Println("work done!")
}
~~~