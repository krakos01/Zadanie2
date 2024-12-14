Dawid Krajewski 97646

# Sprawozdanie
## Część obowiązkowa

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


## Część dodatkowa
[Cały plik workflow dla części dodatkowej](.github/workflows/gha_Zadanie2_dodatkowe.yaml)

- Najpierw jest "ręcznie" budowany lokalny obraz, który bedzie mógł zostać sprawdzony przez Trivy:
```
name: Build local image for trivy
id: build_trivy
run: |
  docker build -t local/krakos01/zadanie2dod:trivy .
```

- Następnie jest generowany SBOM poprzez użycie <i>aquasecurity/trivy-action@0.29.0</i>. Github-pat jest dodany, żeby można było przekazać wyniki do dependency graph. Jako obraz zostaje przekazany ten utwrzony w poprzednim kroku. W porównaniu do Docker Scouta, Trivy nie znajduje zagrożeń.
```
name: Generate SBOM using Trivy and submit results
uses: aquasecurity/trivy-action@0.29.0
with:
  image-ref: 'local/krakos01/zadanie2dod:trivy'
  format: 'github'
  output: 'dependency-results.sbom.json'
  github-pat: ${{ github.token }}
```

Wynik SBOM
```
Running Trivy with options: trivy image local/krakos01/zadanie2dod:trivy
2024-12-14T16:23:49Z	INFO	[vuln] Vulnerability scanning is enabled
2024-12-14T16:23:49Z	INFO	[secret] Secret scanning is enabled
2024-12-14T16:23:49Z	INFO	[secret] If your scanning is slow, please try '--scanners vuln' to disable secret scanning
2024-12-14T16:23:49Z	INFO	[secret] Please see also https://aquasecurity.github.io/trivy/v0.57/docs/scanner/secret#recommendation for faster secret detection
2024-12-14T16:23:51Z	INFO	Detected OS	family="alpine" version="3.21.0"
2024-12-14T16:23:51Z	WARN	This OS version is not on the EOL list	family="alpine" version="3.21"
2024-12-14T16:23:51Z	INFO	[alpine] Detecting vulnerabilities...	os_version="3.21" repository="3.21" pkg_num=24
2024-12-14T16:23:51Z	INFO	Number of language-specific files	num=1
2024-12-14T16:23:51Z	INFO	[gobinary] Detecting vulnerabilities...

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100  7755  100   168  100  7587    657  29707 --:--:-- --:--:-- --:--:-- 30292
 Uploading GitHub Dependency Snapshot{
  "id": 19330531,
  "created_at": "2024-12-14T16:23:51.566Z",
  "result": "SUCCESS",
  "message": "Dependency results for the repo have been successfully updated."
}
```

- Skanowanie zagrożeń przez Trivy. Jeśli znajdzie błędy o poziomie HIGH lub CRITICAL typu os lub library workflow zostanie zakończony. Wynik zostaje pokazany w formacie tabeli. <br> 
```
name: Scan image vulnerabilities using Trivy
uses: aquasecurity/trivy-action@0.29.0
with:
  image-ref: 'local/krakos01/zadanie2dod:trivy'
  format: 'table'
  severity: 'HIGH,CRITICAL'
  vuln-type: 'os,library'
  exit-code: '2'
  github-pat: ${{ github.token }}
```

Wynik skanowania:
```
2024-12-14T16:23:55Z	INFO	[gobinary] Detecting vulnerabilities...

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
```

- Jeśli skanowanie zagrożeń przejdzie pomyślnie, to następuje prawidłowe budowanie wieloarchitekturowego obrazu i wysłanie na rejestr ghcr.io, z użyciem <i>docker/build-push-action@v5</i>. W porównaniu do części obowiązkowej, został dodany input: ``` provenance: mode=max ```

Uruchomienie obrazu przechodzi pomyślnie, a w zakładce 'packages' pojawia się mój obraz.
![Screen3](https://github.com/user-attachments/assets/a2ca61b5-d608-46fe-b0ab-b13a6bceb360)
![Screen4](https://github.com/user-attachments/assets/a44e02f4-34cb-467b-87ad-12eac882b8f8)
![Screen5](https://github.com/user-attachments/assets/28aeae4e-aaba-419f-a919-48e846dc1730)


