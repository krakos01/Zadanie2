# -- ETAP 1 --
FROM golang:1.23-alpine3.19 AS builder

# Ustawienie katalogu roboczego
WORKDIR /app

# Kopiowanie przygotowanej aplikacji do kontenera
COPY projekt.go .

# Budowanie aplikacji w Go
RUN go build projekt.go



# -- ETAP 2 --
FROM scratch

# Deklaracja metadanych - imię i nazwisko autora
LABEL org.opencontainers.image.authors="Dawid Krajewski"

# Wykorzystuje obraz bazowy alpine
ADD alpine-minirootfs-3.20.3-x86_64.tar /

# Instalacja CURL'a
RUN apk add --update --no-cache curl

# Ustawienie katalogu roboczego
WORKDIR /app

# Kopiowanie aplikacji z etapu pierwszego
COPY --from=builder /app /app

# Deklaracja portu aplikacji w kontenerze
EXPOSE 8080

# Monitorowanie dostępności serwera
HEALTHCHECK --interval=10s --timeout=1s \
 CMD curl -f http://localhost:8080 || exit 1

# Deklaracja sposobu uruchomienia serwera
CMD ["./projekt"]