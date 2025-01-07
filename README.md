[![PL](https://img.shields.io/badge/lang-PL-blue.svg?logo=data:image/svg%2bxml;base64,PD94bWwgdmVyc2lvbj0iMS4wIj8+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB3aWR0aD0iNjQwIiBoZWlnaHQ9IjQwMCIgdmlld0JveD0iMCAwIDggNSI+CjxyZWN0IHdpZHRoPSI4IiBoZWlnaHQ9IjUiIGZpbGw9IiNkYzE0M2MiLz4KPHJlY3Qgd2lkdGg9IjgiIGhlaWdodD0iMi41IiBmaWxsPSIjZmZmIi8+Cjwvc3ZnPg==)](README.md)
[![EN](https://img.shields.io/badge/lang-EN-green.svg?logo=data:image/svg%2bxml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4NCjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgd2lkdGg9IjEyMzUiIGhlaWdodD0iNjUwIiB2aWV3Qm94PSIwIDAgNzQxMCAzOTAwIj4NCjxyZWN0IHdpZHRoPSI3NDEwIiBoZWlnaHQ9IjM5MDAiIGZpbGw9IiNiMjIyMzQiLz4NCjxwYXRoIGQ9Ik0wLDQ1MEg3NDEwbTAsNjAwSDBtMCw2MDBINzQxMG0wLDYwMEgwbTAsNjAwSDc0MTBtMCw2MDBIMCIgc3Ryb2tlPSIjZmZmIiBzdHJva2Utd2lkdGg9IjMwMCIvPg0KPHJlY3Qgd2lkdGg9IjI5NjQiIGhlaWdodD0iMjEwMCIgZmlsbD0iIzNjM2I2ZSIvPg0KPGcgZmlsbD0iI2ZmZiI+DQo8ZyBpZD0iczE4Ij4NCjxnIGlkPSJzOSI+DQo8ZyBpZD0iczUiPg0KPGcgaWQ9InM0Ij4NCjxwYXRoIGlkPSJzIiBkPSJNMjQ3LDkwIDMxNy41MzQyMzAsMzA3LjA4MjAzOSAxMzIuODczMjE4LDE3Mi45MTc5NjFIMzYxLjEyNjc4MkwxNzYuNDY1NzcwLDMwNy4wODIwMzl6Ii8+DQo8dXNlIHhsaW5rOmhyZWY9IiNzIiB5PSI0MjAiLz4NCjx1c2UgeGxpbms6aHJlZj0iI3MiIHk9Ijg0MCIvPg0KPHVzZSB4bGluazpocmVmPSIjcyIgeT0iMTI2MCIvPg0KPC9nPg0KPHVzZSB4bGluazpocmVmPSIjcyIgeT0iMTY4MCIvPg0KPC9nPg0KPHVzZSB4bGluazpocmVmPSIjczQiIHg9IjI0NyIgeT0iMjEwIi8+DQo8L2c+DQo8dXNlIHhsaW5rOmhyZWY9IiNzOSIgeD0iNDk0Ii8+DQo8L2c+DQo8dXNlIHhsaW5rOmhyZWY9IiNzMTgiIHg9Ijk4OCIvPg0KPHVzZSB4bGluazpocmVmPSIjczkiIHg9IjE5NzYiLz4NCjx1c2UgeGxpbms6aHJlZj0iI3M1IiB4PSIyNDcwIi8+DQo8L2c+DQo8L3N2Zz4=)](/README-en.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/michalsabaj/pldlchecker.svg)](https://pkg.go.dev/github.com/michalsabaj/pldlchecker)
---
# Sprawdz polskie uprawnienia do kierowania pojazdami
Założeniem projektu PLDLChecker (*Polish Driving Licence Checker*) jest umożliwenie łatwego i szybkiego sprawdzenia, czy osoba ma uprawnienia do kierowania pojazdami.
Wszystkie dane zostały pozyskane z apliakcji gov - [https://moj.gov.pl/uslugi/engine/ng/index?xFormsAppName=UprawnieniaKierowcow](https://moj.gov.pl/uslugi/engine/ng/index?xFormsAppName=UprawnieniaKierowcow)
## Jak to działa?
Otóż trzeba podać imię, nazwisko oraz numer prawa jazdy.  
Imię i nazwisko nie może mieć więcej niż 80 znaków, a prawo jazdy 8.  
Dane te następnie są normalizowane - usuwane są spacje i zamieniane są polskie znaki.  
Kolejnym krokiem jest zamienienie wszystkich liter, na duże litery i zahashowanie używając **MD5**.  
Tak przygotowany hash, jest kluczem przy odpytywaniu api.  
Jeśli kod odpowiedzi z API to **200**, oznacza to, że dane zostały znalezione.  
Kod **400** oznacza, że osoba o podanych danych **nie istnieje**.  
Odpowiedź przychodzi w formacie JSON.  
## Jak to użyć?
> [!IMPORTANT]  
> Projekt w bieżącej formie **NIE NADAJE SIĘ** do zastosowania produkcyjnego.  
> Szczególnym powodem, jest brak udostępnienia bardziej szczegółowych informacji przez Ministerstwo, jak i samo ryzyko zmiany działania algorytmu.  
> Sugeruje użyć kod jako "podkładkę" do własnych, wewnętrznych zastosowań.  
> Natomiast będę się starał cały czas rozwijać to repozytorium. 

```
Stwórz nowy projekt Go, następnie: 
go get github.com/michalsabaj/pldlchecker/ 

func IsDriverLicenseValid(firstName, lastName, driverLicenseNumber string) (bool, error)
```

> Projekt wykonano w GO v1.23.1.
---