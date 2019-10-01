# 表白墙服务 RPC接口
<!--TODO: 对应文档完成后，添加文档链接-->
以下类型在tars文件中定义, 参照[tars数据类型定义文档](#)
```
Message       留言信息
```

## 接口列表
* [PostMessage](#interface-postmessage)
* [GetMessageList](#interface-getmessagelist)
* [AddLike](#interface-addlike)
* [GetLike](#interface-getlike)

## 接口详情

### <a id="interface-postmessage"> PostMessage
发布表白

**定义**
```cpp
int PostMessage(Message Msg)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|Msg|Message|留言信息|

### <a id="interface-getmessagelist"> GetMessageList
获取表白留言列表

**定义**
```cpp
int GetMessageList(int Index, string Date, string WxId, out int NextIndex, out vector<Message> MsgList)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|Index|int|索引|
|Date|string|日期|
|WxId|string|用户id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|NextIndex|int|下一次的请求索引|
|MsgList|vector<Message>|留言列表|

### <a id="interface-addlike"> AddLike
点赞，对应留言点赞数+1

**定义**
```cpp
int AddLike(string MessageId)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|MessageId|string|留言id|

### <a id="interface-getlike"> GetLike
获取留言点赞数

**定义**
```cpp
int GetLike(string MessageId, out int LikeCount)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|MessageId|string|留言id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|LikeCount|int|点赞数|