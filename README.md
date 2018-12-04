# GO - GRPC

Go ile basit bir grpc sunucu istemci ayağa kaldırıp bir web sayfasının parse edeceğiz


Go Kurulum

ubuntu ortamında aşağıdaki komutları çalıştırıp go yu kuruyoruz.

```sh
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt-get update
$ sudo apt-get install golang-go
```


Diğer ortamlar ve detaylar için bu adresi inceleyebilirsiniz
https://golang.org/doc/install

Deployment

Git Kurulu değil ise : 

```sh
$ sudo apt-get install git-core
$ git clone https://github.com/yilmazyusuf/go-grpc.git
```

#Dockerfile kullanarak

Docker kurulu değil ise : 

```sh
$ sudo apt-get install docker.io
```

Docker Image Oluşturuyoruz :

```sh
$ docker build -t parse_article .
```

Docker Image ı çalştırıyoruz:

```sh
$ docker run -p 8080:8080 -p 50051:50051 --name parse_article --rm parse_article
```


grpc protokolümüz 50051 portundan haberleşecek, 8080 portuna http client için ihtiyacımız var.

Bu portların kullanıldığına dair hata alırsanız dockerfile daki port tanımlarını değiştirip yeninden build ve run yapabilirsiniz.

Veya kullanılan portları boşaltabilirsiniz.
lsof -ti:8080 | xargs kill -9
lsof -ti:50051 | xargs kill -9

grpc server a istek yapmak için aşağıdaki http adresini kullanabilir veya console_client ile istek yapabilrisiniz.
Uzak sunucuda işlem yapıyorsanız localhost yerine sunucu IP nizi yazmanız gerekir.

```sh
http://localhost:8080/read_article/?url=https://medium.com/@yusufylmaz_3450/stateful-stateless-68cfae55653d
$ go run grpc/console_client.go https://medium.com/@yusufylmaz_3450/stateful-stateless-68cfae55653d
```

proto dosyasında geliştirme yapmak isterseniz bu şekilde tekrar compile edebilirsiniz

```sh
protoc --go_out=plugins=grpc:. protos/*.proto
```



Planlanan geliştirmeler
```sh
Authentication
```


