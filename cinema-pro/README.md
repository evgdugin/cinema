# Этапы разработки

## 1. Создаем модуль

``` 
go mod init github.com/evgdugin/cinema
```

## 2. Устанавливаем gRPC 

``` 
go get -u google.golang.org/grpc
```

## 3. Скачиваем protoc компилятор для платформы win64. https://github.com/google/protobuf/releases

``` 
копируем файл protoc в директорию go/bin
```

## 4. Вставляем папку include для работы компиляторы при работе с прото файлами

## 5. Потом нужно поставить плагин для protoc, который позволит компилировать proto файл в код GoLang:

``` 
go get -u github.com/golang/protobuf/protoc-gen-go
```

## 6. Проверяем настройку согласно readme https://github.com/evgdugin/cinema/tree/master/example/grpc

## 7. 

## 4. 

## 4. 

## 4. 

## 4. 

## 4. 

## 4. 
