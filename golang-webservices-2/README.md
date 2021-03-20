## [Разработка веб-сервисов на Go, часть 2](https://www.coursera.org/learn/golang-webservices-2)

Go (golang) - современный язык программирования, предназначенный для разработки высококонкурентных приложений, работающих на многопроцессорных системах.
В курсе рассмотрены основы программирования на языке Go, а так же опыт применения языка в основных задачах, которые встречаются сегодня в серверной веб-разработке. 
В данной части курса рассмотрены основы языка и разработки веб-сервисов с использованием стандартной библиотеки

## Материалы для чтения к 5-му уроку

### Компоненты

**[https://github.com/avelino/awesome-go](https://github.com/avelino/awesome-go)**

### Шаблоны

- [https://github.com/SlinSo/goTemplateBenchmark](https://github.com/SlinSo/goTemplateBenchmark#full-featured-template-engines-2)

### Роутеры

- [https://github.com/gorilla/mux](https://github.com/gorilla/mux) - один из компонентов gorillatoolkit, из которых можно собрать себе полноценный фреймворк
- [https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [https://github.com/valyala/fasthttp](https://github.com/valyala/fasthttp)
- [https://github.com/julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)

### Фреймворки

- [https://beego.me/](https://beego.me/)
- [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [https://github.com/Massad/gin-boilerplate](https://github.com/Massad/gin-boilerplate)
- [https://github.com/gramework/gramework](https://github.com/gramework/gramework) - быстрый веб-ферймворк на основе fasthttp

### Логирование

- [https://github.com/uber-go/zap](https://github.com/uber-go/zap)
- [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
- [https://www.youtube.com/watch?v=c_MPDg2C9tg](https://www.youtube.com/watch?v=c_MPDg2C9tg) - видео по структурному логирования
- [https://habrahabr.ru/company/badoo/blog/328722/](https://habrahabr.ru/company/badoo/blog/328722/)

### Веб-сокеты

- [https://github.com/gorilla/websocket](https://github.com/gorilla/websocket)
- [https://github.com/gobwas/ws](https://github.com/gobwas/ws) - библиотека для низкоуровневой работы в веб-сокетами от Mail.ru, которая позволяет существенно сэкономить на памяти сервера
- [https://github.com/olahol/melody](https://github.com/olahol/melody)

### Управление зависимостями

- [https://github.com/golang/dep](https://github.com/golang/dep)
- [https://hackernoon.com/using-go-dep-as-a-project-maintainer-641d1f3006d7](https://hackernoon.com/using-go-dep-as-a-project-maintainer-641d1f3006d7)
- [https://about.sourcegraph.com/go/the-new-era-of-go-package-management/](https://about.sourcegraph.com/go/the-new-era-of-go-package-management/)
- [https://medium.freecodecamp.org/an-intro-to-dep-how-to-manage-your-golang-project-dependencies-7b07d84e7ba5](https://medium.freecodecamp.org/an-intro-to-dep-how-to-manage-your-golang-project-dependencies-7b07d84e7ba5)
- [https://blog.gopheracademy.com/advent-2015/vendor-folder/](https://blog.gopheracademy.com/advent-2015/vendor-folder/)

### Безопасность

- [https://github.com/Checkmarx/Go-SCP](https://github.com/Checkmarx/Go-SCP)

### Дополнительные материалы

- [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Материалы для чтения к 6-му уроку

- [http://www.vividcortex.com/hubfs/eBooks/The_Ultimate_Guide_To_Building_Database-Driven_Apps_with_Go.pdf](http://www.vividcortex.com/hubfs/eBooks/The_Ultimate_Guide_To_Building_Database-Driven_Apps_with_Go.pdf) - в удобной форме информация по основным аспектам работы с database/sql
- [https://golang.org/pkg/database/sql/](https://golang.org/pkg/database/sql/) - собственно сам интерфейс к базе
- [https://github.com/golang/go/wiki/SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers) - список поддерживаемых баз
- [https://github.com/golang/go/wiki/SQLInterface](https://github.com/golang/go/wiki/SQLInterface)
- [https://github.com/DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)
- [http://www.alexedwards.net/blog/configuring-sqldb](http://www.alexedwards.net/blog/configuring-sqldb)
- [http://go-database-sql.org/](http://go-database-sql.org/)
- [https://astaxie.gitbooks.io/build-web-application-with-golang/](https://astaxie.gitbooks.io/build-web-application-with-golang/)
- [https://github.com/thewhitetulip/web-dev-golang-anti-textbook/](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
- [https://codegangsta.gitbooks.io/building-web-apps-with-go/content/](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)
- [https://godoc.org/github.com/go-sql-driver/mysql](https://godoc.org/github.com/go-sql-driver/mysql)
- [https://godoc.org/github.com/lib/pq](https://godoc.org/github.com/lib/pq)
- [https://godoc.org/github.com/bradfitz/gomemcache/memcache](https://godoc.org/github.com/bradfitz/gomemcache/memcache)
- [https://godoc.org/github.com/garyburd/redigo/redis](https://godoc.org/github.com/garyburd/redigo/redis)
- [https://godoc.org/gopkg.in/mgo.v2](https://godoc.org/gopkg.in/mgo.v2)
- [http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/](http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/)
- [http://gorm.io/](http://gorm.io/)
- [http://motion-express.com/blog/gorm:-a-simple-guide-on-crud](http://motion-express.com/blog/gorm:-a-simple-guide-on-crud)
- [https://godoc.org/github.com/jinzhu/gorm](https://godoc.org/github.com/jinzhu/gorm)
- [https://habrahabr.ru/company/mailru/blog/266811/](https://habrahabr.ru/company/mailru/blog/266811/) - архи-полезная статья про устройство базы внутри
- [https://hackernoon.com/communicating-go-applications-through-redis-pub-sub-messaging-paradigm-df7317897b13](https://hackernoon.com/communicating-go-applications-through-redis-pub-sub-messaging-paradigm-df7317897b13)
- [https://medium.com/@shijuvar/introducing-nats-to-go-developers-3cfcb98c21d0](https://medium.com/@shijuvar/introducing-nats-to-go-developers-3cfcb98c21d0)
- [https://medium.com/@shijuvar/building-distributed-systems-and-microservices-in-go-with-nats-streaming-d8b4baa633a2](https://medium.com/@shijuvar/building-distributed-systems-and-microservices-in-go-with-nats-streaming-d8b4baa633a2)

## Материалы для чтения к 7-му уроку

- [https://about.sourcegraph.com/go/grpc-in-production-alan-shreve/](https://about.sourcegraph.com/go/grpc-in-production-alan-shreve/) + [https://www.youtube.com/watch?v=7FZ6ZyzGex0](https://www.youtube.com/watch?v=7FZ6ZyzGex0)
- [https://grpc.io/](https://grpc.io/) - общий сайт gRPC
- [https://github.com/grpc/grpc-go](https://github.com/grpc/grpc-go) - go-версия gRPC
- [https://github.com/grpc-ecosystem](https://github.com/grpc-ecosystem) - набор middleware для gRPC
- [https://outcrawl.com/getting-started-microservices-go-grpc-kubernetes/](https://outcrawl.com/getting-started-microservices-go-grpc-kubernetes/)
- [https://improbable.io/games/blog/grpc-web-moving-past-restjson-towards-type-safe-web-apis](https://improbable.io/games/blog/grpc-web-moving-past-restjson-towards-type-safe-web-apis)
- [https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/](https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/)
- [https://ops.tips/blog/sending-files-via-grpc/](https://ops.tips/blog/sending-files-via-grpc/)
- [https://github.com/mattn/ft](https://github.com/mattn/ft) - file transfer via gRPC
- [http://mhausenblas.info/fosdem2018-godevroom-networkingdeepdive/](http://mhausenblas.info/fosdem2018-godevroom-networkingdeepdive/) - всецело полезный доклад про работу с сетью в go
- [https://github.com/twitchtv/twirp](https://github.com/twitchtv/twirp) - RPC-фреймворк от Twitch, достаточно молодой
- [https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f) - статья про Twirp
- [https://about.sourcegraph.com/go/fallacies-of-distributed-gomputing/](https://about.sourcegraph.com/go/fallacies-of-distributed-gomputing/) - мощный доклад по распределённым системам на го
- [https://github.com/google/go-microservice-helpers](https://github.com/google/go-microservice-helpers)
- [https://github.com/vaporz/turbo](https://github.com/vaporz/turbo)
- [https://github.com/go-kit/kit](https://github.com/go-kit/kit) - мощный фреймворк для написания микросервисов
- [https://habrahabr.ru/post/276539/](https://habrahabr.ru/post/276539/) - "Это будущее". Обязательная для ознакомления статья если вы решили увлечься микросервисами по полной
- [https://medium.com/apis-and-digital-transformation/openapi-and-grpc-side-by-side-b6afb08f75ed](https://medium.com/apis-and-digital-transformation/openapi-and-grpc-side-by-side-b6afb08f75ed)
- [https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)
- [https://habrahabr.ru/company/beget/blog/348008/](https://habrahabr.ru/company/beget/blog/348008/)
- [https://www.ribice.ba/swagger-golang/](https://www.ribice.ba/swagger-golang/) - Create Golang API documentation with SwaggerUI
- [https://ewanvalentine.io/microservices-in-golang-part-1/](https://ewanvalentine.io/microservices-in-golang-part-1/) - большой туториал по микросервисам в го, охватывает множество сфер ( докер, авторизацию и прочее ) - на момент добавления в список вышла 7-я часть
- [https://github.com/MarquisIO/go-grpcmw](https://github.com/MarquisIO/go-grpcmw)
- [https://github.com/enricofoltran/hello-auth-grpc](https://github.com/enricofoltran/hello-auth-grpc)
- [https://blog.synq.fm/golang-microservice-starter-kit](https://blog.synq.fm/golang-microservice-starter-kit)
- [https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/](https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)
- [https://www.youtube.com/watch?v=s5l9ZdgxzXA](https://www.youtube.com/watch?v=s5l9ZdgxzXA)
- [http://www.minaandrawos.com/2016/05/14/udp-vs-tcp-in-golang/](http://www.minaandrawos.com/2016/05/14/udp-vs-tcp-in-golang/)
- [http://rodaine.com/2017/05/x-files-time-rate-golang/](http://rodaine.com/2017/05/x-files-time-rate-golang/)
- [https://blog.envoyproxy.io/introduction-to-modern-network-load-balancing-and-proxying-a57f6ff80236](https://blog.envoyproxy.io/introduction-to-modern-network-load-balancing-and-proxying-a57f6ff80236)
- [http://tumregels.github.io/Network-Programming-with-Go/](http://tumregels.github.io/Network-Programming-with-Go/)
