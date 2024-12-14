Dawid Krajewski 97646

# Sprawozdanie - część obowiązkowa

### 1. Opis workflow
[Cały plik workflow dla części obowiązkowej](.github/workflows/gha_Zadanie2.yaml)

- Dodano zmienną REGISTRY: <br>
``` env: REGISTRY: ghcr.io ```
- W metadanych zmieniono <i>images</i> tak, żeby odpowiadał rejestrowi ghcr: <br>
``` images: ${{ env.REGISTRY }}/${{ github.actor }}/zadanie2_test ```
- Dodano logowanie do GitHuba:
```
name: Login to GitHub container registry
uses: docker/login-action@v3
with:
  registry: ${{ env.REGISTRY }}
  username: ${{ github.actor }}
  password: ${{ github.token }}
```
- Dodano test Docker Scout, który używa <i>docker/scout-action@v1.16.0</i>, skanuje budowany obraz i przefiltrowuje zagrożenia HIGH i CRITICAL. Jeśli jakieś będą, to wyjdzie z workflowa:
```
name: Docker Scout
uses: docker/scout-action@v1.16.0
with:
  command: cves
  only-severities: high, critical
  exit-code: true
```
- Na końcu obraz jest wysyłany na ghcr tak samo, jak we wzorcu

Przy uruchomieniu znalazło 6 błędów i obraz nie został wysłany na ghcr. Potwierdza to brak opublikowanych paczek. Użyte obrazy go i alpine są aktualne, więc nie można ich poprawić: <br>
![Screen1](https://github.com/user-attachments/assets/5543ea87-e806-4162-b4b8-d5f18bf76d26)
![Screen2](https://github.com/user-attachments/assets/e034dd05-a047-4b45-aca0-cbbc4aaf1e67)
