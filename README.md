# ●	веб-приложение, которое занимается рендерингом странички на сервере и выдачей готового HTML в браузер

# index

![index](png/index.png)

# signup

![signup](png/signup.png)

# signin

![signin](png/signin.png)

# profile

![profile](png/profile.png)

# watch

![watch](png/watch.png)

# payment

![payment](png/payment.png)

# Принцип работы

![payment](png/payment.png)

---

# Пользователь имеет 3 статуса

* анонимный, 
* зарегистрированный, 
* оформивший подписку

## Возможно требуется добавить сервисы:

### Сервис обратной связи support

### Сервис комментариев

---

# Базы данных

Базы данных будут содержать следующие микросервисы:

* user
* movie
* payment

## База данных movie

название базы данных MySQL

* DB_SERVICE_MOVIE

Содержит:

* MOVIE_TITLE

название фильма

* MOVIE_ID

id 

* MOVIE_CATEGORY

жанр 

* MOVIE_RELEASE_DATE

год выпуска

* MOVIE_MPA

возрастной рейтинг NOTSET, P, PG, R

* MOVIE_VISIBILITY

статус фильма - видимость фильма

* дополнительно

MOVIE_CREATED_AT
MOVIE_UPDATED_AT

Также на странице должен быть плеер с фильмом. Стриминг видео — это отдельная большая и сложная тема, которую мы не будем рассматривать в данном курсе. Поэтому встроим на страницу iframe YouTube с трейлером фильма.

