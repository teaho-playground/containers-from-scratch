package main

func main() {

	//https://www.lixueduan.com/posts/docker/05-namespace/#ipc

	//clone 和 unshare 的区别
	//
	//clone 和 unshare 的功能都是创建并加入新的 namespace， 他们的区别是：
	//
	//unshare 是使当前进程加入新的 namespace
	//clone 是创建一个新的子进程，然后让子进程加入新的 namespace，而当前进程保持不变

}
