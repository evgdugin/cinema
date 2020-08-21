# Этапы разработки

## ✅ Подготовка
### 1.1 Создаем модуль

``` 
go mod init github.com/evgdugin/cinema
```

### 1.2 Устанавливаем gRPC 

``` 
go get -u google.golang.org/grpc
```

### 1.3 Скачиваем protoc компилятор для платформы win64. https://github.com/google/protobuf/releases

``` 
копируем файл protoc в директорию go/bin
```

### 1.4 Вставляем папку include для работы компиляторы при работе с прото файлами

### 1.5 Потом нужно поставить плагин для protoc, который позволит компилировать proto файл в код GoLang:

``` 
go get -u github.com/golang/protobuf/protoc-gen-go
```

### 1.6 Проверяем настройку согласно readme https://github.com/evgdugin/cinema/tree/master/example/grpc

## 7. 

## 4. 

## 4. 

## 4. 

## 4. 

## 4. 

## 4. 
