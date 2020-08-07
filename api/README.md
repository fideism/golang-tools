# API

## http response 返回数据格式
- `ToJSON` ToJSON 转换为JSON数据下发

- `OKJSON` 正常200状态下的数据返回

- `EmptyStruct`  空结构体 200

- `EmptyCollection` 空集合 200

- `BadRequest` 客户端错误 400

- `Unauthorized` 未登陆 401 

- `Forbidden` 访问权限限制 403

- `NotFound` 请求信息不存在 404

- `PreconditionFailed` 请求前置条件不满足 412

- `Unprocessable` 上行数据格式不正确 422

- `ServerError` 服务器出错 500

- `Error` 自定义