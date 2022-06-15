
## Overview

Package snamer renames fields of a struct. It is useful to create snake_case json from a struct. In Golang, an initial character of a field must be uppercase to export the field.

## Example
```go
type User struct {
	UserId int
}
```

```go
user := User{UserId: 1}
result, _ := snamer.PascalStructToCamel(user) // map[userId:1]
data, _ := json.Marshal(result) // {"userId":1}
```

```go
user := User{UserId: 1}
result, _ := snamer.PascalStructToSnake(user) // map[user_id:1]
data, _ := json.Marshal(result) // {"user_id":1}
```

You can write any function to convert struct field name.
```go
user := User{UserId: 1}
result, _ := snamer.AnyStructToAny(user, func(s string) string {
    // PascalCase To CONSTANT_CASE
    re := regexp.MustCompile(`([A-Z])`)
    return strings.ToUpper(s[0:1] + re.ReplaceAllString(s[1:], `_$1`))
}) // map[USER_ID:1]
data, _ := json.Marshal(result) // {"USER_ID":1}
```

## Description
Various primitive types, nested struct, and pointer are supported.
See `struct_test.go`.