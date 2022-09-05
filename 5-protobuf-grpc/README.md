# Protocol buffers - GRPC

## Главные недостатки JSON
- большой размер данных
- медленный парсинг

## Что такое Protocol Buffers
- (эффективный) бинарный формат
- для (де)кодирования нужно знать схему
- парсеры/сериалазторы генерируются **для многих языков**
- фокус на обратную совместимость

## Типы данных - скалярые
```protobuf
message TestMessage {
    bool some_value = 1;
    uint8 some_uint8 = 2;
    double some_double = 3;
    string some_string = 4;
    bytes binary_data = 5;
}
```

Все типы - [документация по scalar-ным типам](https://developers.google.com/protocol-buffers/docs/proto3#scalar) 

## Сообщение может быть полем
```protobuf
message InnerMessage {
    string name = 1;
}

message Message {
    InnerMessage inner = 1;
}
```

## Типы данных - repeated
```protobuf
message Comment {
    string text = 1;
}

message TestMessage {
    repeated Comment comments = 1;
    repeated int64 ids = 2;
}
```

## Типы данных - map
- ключ - любой скалярный тип, кроме bytes и чисел с плавающей точкой

```protobuf
message Comment {
    string text = 1;
}

message MessageWithMap {
    map<string, SomeMessage> my_map = 1;
}
```

## Типы данных - enum
```protobuf
enum Status {
    UNKNOWN = 1;
    ENABLED = 2;
    DISABLED = 3;
}

message Message {
    Status status = 1;
}
```

## Типы данных - oneof
```protobuf
message NumericValue {
    string name = 1;
    oneof test_oneof {
        int64 int = 2;
        double float = 3;
        FormattedFloat = 4;
    }
}

message FormattedFloat {
    double float = 1;
}
```

## Вложенность
```protobuf
message NestedDemo {
    message NestedMessage {
        enum NestedEnum {
            ONE = 1;
            TWO = 2;
        }

        NestedEnum value = 1;
    }

    NestedMessage.NestedEnum result = 1;
}
```