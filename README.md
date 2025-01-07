# Sprawdz polskie uprawnienia do kierowania pojazdami
Założeniem projektu PLDLChecker jest umożliwenie łatwego i szybkiego sprawdzenia, czy osoba ma uprawnienia do kierowania pojazdami.
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
## Jak to zbudować?
> [!IMPORTANT]  
> Projekt w bieżącej formie **NIE NADAJE SIĘ** do zastosowania profesjonalnego.
> Sugeruje użyć kod jako "podkładkę" do własnych, wewnętrznych zastosowań.

```
git clone https://github.com/michalsabaj/pldlchecker.git
cd pldlchecker
go run .
```
> Projekt zrobiono w GO v1.23.1.
---