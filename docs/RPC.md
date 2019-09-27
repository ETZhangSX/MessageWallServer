# 表白墙服务 RPC接口
<!--TODO: 对应文档完成后，添加文档链接-->
以下类型在tars文件中定义, 参照[tars数据类型定义文档](#)
* Message
## PostMessage(Message Msg)
发布表白

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|Msg|Message|留言信息|

## GetMessageList
获取表白留言列表

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

## AddLike
点赞，对应留言点赞数+1

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|MessageId|string|留言id|

## GetLike
获取留言点赞数

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|MessageId|string|留言id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|LikeCount|int|点赞数|