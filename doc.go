// JSON Processing project doc.go

/*
JSON Processing document
*/
package main

/*
如何将json数据与struct字段相匹配呢？例如JSON的key是Foo，那么怎么找对应的字段呢？

首先查找tag含有Foo的可导出的struct字段(首字母大写)
其次查找字段名是Foo的导出字段
最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段
聪明的你一定注意到了这一点：能够被赋值的字段必须是可导出字段(即首字母大写）。同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，这样的一个好处是：当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题。
*/

/*
我们看到上面的输出字段名的首字母都是大写的，如果你想用小写的首字母怎么办呢？把结构体的字段名改成首字母小写的？JSON输出的时候必须注意，只有导出的字段才会被输出，如果修改字段名，那么就会发现什么都不会输出，所以必须通过struct tag定义来实现：

type Server struct {
    ServerName string `json:"serverName"`
    ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
    Servers []Server `json:"servers"`
}
通过修改上面的结构体定义，输出的JSON串就和我们最开始定义的JSON串保持一致了。

针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:

字段的tag是"-"，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
举例来说：

type Server struct {
    // ID 不会导出到JSON中
    ID int `json:"-"`

    // ServerName2 的值会进行二次JSON编码
    ServerName  string `json:"serverName"`
    ServerName2 string `json:"serverName2,string"`

    // 如果 ServerIP 为空，则不输出到JSON串中
    ServerIP   string `json:"serverIP,omitempty"`
}

s := Server {
    ID:         3,
    ServerName:  `Go "1.0" `,
    ServerName2: `Go "1.0" `,
    ServerIP:   ``,
}
b, _ := json.Marshal(s)
os.Stdout.Write(b)
会输出以下内容：

{"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}
Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：

JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
Channel, complex和function是不能被编码成JSON的
嵌套的数据是不能编码的，不然会让JSON编码进入死循环
指针在编码的时候会输出指针指向的内容，而空指针会输出null
本小节，我们介绍了如何使用Go语言的json标准包来编解码JSON数据，同时也简要介绍了如何使用第三方包go-simplejson来在一些情况下简化操作，学会并熟练运用它们将对我们接下来的Web开发相当重要。
*/
