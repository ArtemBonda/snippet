### Snippet - веб-приложение, в котором можно оставлять и делиться своими заметками в виде текста
(Это учебный проект)

Приложение разворачивается на локальном хосте "127.0.0.1" порт "8080"

Скачивание репозитария

```
git clone github.com/ArtemBonda/snippet
```

Создание бинарного файла

```
go build cmd/app/main.go
```
Запуск приложения в windows
```
.\main.exe
```
Запуск приложения в Linux
```
./main
```
Шаблоны HTML-страниц используется именование сдедующегно формата
```text
<название>.<роль>.tmpl
```
- роль - это одно из трех: page, partial, layout