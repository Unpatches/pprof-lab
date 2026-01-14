## Имя: Дорджиев Виктор
## Группа: ЭФМО-02-25
# Проект pprof-lab

Цели
1.	Научиться подключать и использовать профилировщик pprof для анализа CPU, памяти, блокировок и горутин.
2.	Освоить базовые техники измерения времени выполнения функций (ручные таймеры, бенчмарки).
3.	Научиться читать отчёты go tool pprof, строить графы вызовов и находить “узкие места”.
4.	Сформировать практические навыки оптимизации кода на основании метрик.

---

## Установка и запуск

(Необходимы предустановленные Go версии 1.22 и выше и Git)

Клонировать репозиторий:

```
git clone <URL_РЕПОЗИТОРИЯ>
cd pprof-lab
```

Команда запуска сервера:

```
go run ./cmd/api
```
------

## Структура проекта

```plaintext
pprof-lab/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── work/
│       ├── slow.go
│       ├── timer.go
│       └── slow_test.go
├── go.mod
└── README.md
    
```
------

## Отчёт о проделанной работе
### Нагрузочное тестирование
<img width="615" height="753" alt="image" src="https://github.com/user-attachments/assets/99663d4d-07d4-419b-a446-ef50831ccabc" />

### Профиль
<img width="610" height="107" alt="image" src="https://github.com/user-attachments/assets/cad93fbf-df57-4df9-a60f-751906a602eb" />
<img width="922" height="690" alt="image" src="https://github.com/user-attachments/assets/a824ef22-2d83-4280-8da1-065497ba0745" />
<img width="616" height="99" alt="image" src="https://github.com/user-attachments/assets/9c4ac1be-f768-46a0-be51-bbbf197a7052" />

### Отображение нагрузки в Web
<img width="853" height="991" alt="image" src="https://github.com/user-attachments/assets/1ab9a18d-31cb-4b5f-b4d7-0e61b869b145" />
<img width="619" height="884" alt="image" src="https://github.com/user-attachments/assets/eb6f0a9d-a42f-49e3-a363-cfa8d9cc9263" />
<img width="645" height="800" alt="image" src="https://github.com/user-attachments/assets/26920be0-c17f-4086-9c1a-a6382accce03" />
<img width="890" height="609" alt="image" src="https://github.com/user-attachments/assets/7e3d8cff-1d72-4a86-bc45-035282b89886" />
<img width="685" height="139" alt="image" src="https://github.com/user-attachments/assets/c1f7f933-9c42-4860-b01c-c2a9b7b59ab8" />
<img width="759" height="800" alt="image" src="https://github.com/user-attachments/assets/604b5393-ed5b-45c7-a27f-f3153d7e3d24" />
<img width="263" height="800" alt="image" src="https://github.com/user-attachments/assets/1fb775e0-5b77-4e1a-b442-0061a5270264" />
<img width="426" height="703" alt="image" src="https://github.com/user-attachments/assets/e660a013-1b95-417f-9c73-198bd67be3fb" />




