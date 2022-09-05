

# java知识体系
- ## 基础知识
	- 文件结构
		- 访问性
		- 数据类型
			- 集合类型
			- 值类型
			- 引用类型
		- io
			- 网络io
			- 文件io
		- 多线程
		- 异常
			- Throwable
				- Error
					- 不是程序可以控制的,直接宕机
				- Exception
					- runtime exception可以catch,又名非检查异常(Unchecked Exception)
						- 可以build success,但是run的时候可以try.
						- 经典实例
							- NullPointerException null
							- IndexOutOfBoundsException 数组越界
							- ClassCastException 类型转换错误
							- ArithmeticException 算术异常
							- OutofMemoryException 堆栈溢出
								- --Xmx
								- --Xms
					- no runtime exception,又名检查异常(Checked Exception)
						- 如果不处理，程序就不能编译通过.
						- 经典实例
							- IOException
							- ClassNotFoundException
							- 用户自定义Exception类型
			- 返回值的次序问题
				- try -> finally
				- try -> finally -> catch
				- try -> catch
### 语言概念
### 语言进阶
	- JVM
	- 线程
	- 网络与io
	- 语法糖
### 语言应用
	- 常用的工具类
		- [Apache Common](https://commons.apache.org/)
			- Commons BeanUtils
			- Commons Codec
			- Commons Collections
			- Commons Compress
			- Commons FileUpload
			- Commons IO
			- Commons Lang3
				- java.lang.*的加强包
-
- 项目实践
	- Tomcat与arthas
		- 利用Archas的aop能力排查tomcat容器中的服务慢问题;
		- https://blog.csdn.net/Armour__r/article/details/109733038
	- @Deprecated
		```java
			@Deprecated
			public void method1() {

			}
		```
	- hprof是什么文件？
		- 是java进程的内存镜像文件，里面包含了内存堆详细的使用信息。
		- 在 java 中，hprof 文件有2部分组成，一部分是 hprof head，一部分是 hprof body。其中 head 比较简单，由版本号，IDSize和时间组成。
	- [word转pdf异常之小括号](https://blog.csdn.net/u012998306/article/details/124803061)
	- [java call cmd](https://www.cnblogs.com/youmiyou/p/15779230.html)

## 动态代理

```java
public interface Person {
	void rentHouse();
}

public class Renter implements Person {
	@Override
	public void rentHouse() {
		System.out.println("rent success");
	}
}

public class RenterInvocationHandle<T> implements InvocationHandler {
	
	private T target;

	public RenterInvocationHandle(T target) {
		this.target = target;
	}

	/**
	 *
	 * 1. 代表动态代理对象
	 * 2. 要执行的方法
	 * 3. 入参
	 **/ 
	@Overrider
	public Object invoke(Object proxy, Method method, Object[] args) 
		throws Throwable {

			Object result = method.invoke(target, args);

			return result;
	}
}

public class Main {

	public void static main(String[] args) {
		Person person = new Renter();
		InvocationHandler renterHandler = new InvocationHandler<Person>(person);

		InvocationHandler invoke = (Person)InvocationHandler.newProxyInstance(Person.class.getClassLoader(),new Class<?>[]{Person.class}, renterHandler); 

		Class<?>

	}

}

```