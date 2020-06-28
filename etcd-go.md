```
func main(){
	del(kv,"/cron/jobs/job1")
	delPrex(kv,"/names/")
}
//  删除键值
func del(kv clientv3.KV,key string) {
	var (
		delRes *clientv3.DeleteResponse
		err error
	)
	if delRes,err = kv.Delete(context.TODO(),key); err !=nil{
		log.Fatal(err)
	}
	fmt.Println(delRes.Deleted)
}
//  删除相同前缀的键值
func delPrex(kv clientv3.KV,key string) {
	var (
		delRes *clientv3.DeleteResponse
		err error
	)
	if delRes,err = kv.Delete(context.TODO(),key,clientv3.WithPrefix()); err !=nil{
		log.Fatal(err)
	}
	fmt.Println(delRes.Deleted)
}
————————————————
版权声明：本文为CSDN博主「luslin1711」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/luslin1711/java/article/details/93160135
```
