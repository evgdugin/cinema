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

## название базы данных MySQL

DB_SERVICE_MOVIE

Содержит:
---

## id 

MOVIE_ID

## название фильма

MOVIE_TITLE

## жанр 

MOVIE_CATEGORY

## год выпуска

MOVIE_RELEASE_DATE

## возрастной рейтинг NOTSET, P, PG, R

MOVIE_MPA

## статус фильма - видимость фильма

MOVIE_VISIBILITY

## дополнительно

MOVIE_CREATED_AT
MOVIE_UPDATED_AT

## SQL

``` sql
create table TABLE_SERVICE_MOVIE (
	MOVIE_ID INT,
	MOVIE_TITLE VARCHAR(50),
	MOVIE_CATEGORY VARCHAR(50),
	MOVIE_RELEASE_DATE DATE,
	MOVIE_MPA VARCHAR(50),
	MOVIE_VISIBILITY VARCHAR(50),
	MOVIE_CREATED_AT DATE,
	MOVIE_UPDATED_AT DATE
); 
```

``` s
insert into TABLE_SERVICE_MOVIE (

		MOVIE_ID,
		MOVIE_TITLE,
		MOVIE_CATEGORY,
		MOVIE_RELEASE_DATE,
		MOVIE_MPA,
		MOVIE_VISIBILITY,
		MOVIE_CREATED_AT,
		MOVIE_UPDATED_AT
	)
```

``` s
values (

		1,
		'Бойцовский клуб',
		'thriller',
		'20/08/2020',
		'PG',
		'NOTSET',
		'20/08/2020',
		'20/08/2020'
	);
```

![db_movie](png/db_movie.png)

Также на странице должен быть плеер с фильмом. Стриминг видео — это отдельная большая и сложная тема, которую мы не будем рассматривать в данном курсе. Поэтому встроим на страницу iframe YouTube с трейлером фильма.
