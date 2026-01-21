<div align="center">
  <img style="width: 128px; height: 128px;" src="https://raw.githubusercontent.com/KaolaList/Branding/main/logo.svg" alt="logo" />

  <h1>KaolaList</h1>
  <p><em>Premium Persoonlijk Cloudportaal | Hoogwaardige Netdisk Gateway</em></p>
  <p>KaolaList is een geavanceerde fork van OpenList, ontworpen om een veiligere, intelligentere en esthetisch aantrekkelijkere persoonlijke cloudervaring te bieden.</p>

  <img src="https://img.shields.io/badge/Version-KaolaList--v1.0.0--alpha-68D391" alt="version" />
  <a href="./LICENSE"><img src="https://img.shields.io/github/license/OpenListTeam/OpenList" alt="License" /></a>
  <img src="https://img.shields.io/badge/OpenSource-AGPL--3.0-orange" alt="AGPL-3.0" />
</div>

---

[English](./README.md) | [中文](./README_cn.md) | [日本語](./README_ja.md) | Dutch

---

## 🐨 Kaola-serie Kernfuncties (Geavanceerd)

KaolaList integreert de volgende geavanceerde functies bovenop de robuuste basis van OpenList:

- **🛡️ Kaola Vault**: Ondersteunt **MFA (Multi-Factor Authenticatie)** op mapniveau. Voeg een extra beveiligingslaag toe aan uw kritieke of privégegevens mappen met TOTP/WebAuthn.
- **📺 Kaola Theater**: Een ultieme kijkervaring. Ondersteunt professionele videoformaten zoals `rmvb`, `m2ts` en beschikt over **Volledig Gautomatiseerde Ondertitelingsscanning & Intelligente Matching** (`.srt`, `.ass`, `.vtt`).
- **🧠 Lokale Intelligente OCR**: Een op plug-ins gebaseerde OCR-architectuur. Ondersteunt **PaddleOCR** voor lokale inferentie, waardoor er geen externe API-afhankelijkheid is en maximale privacy wordt gegarandeerd.
- **🤖 Kaola Automatisering**: 
  - **Kaola Clone (Webhook)**: Real-time meldingen voor bestandsveranderingen, eenvoudig te integreren met uw persoonlijke automatiseringsworkflows.
  - **Auto-routing**: Archiveer bestanden automatisch op basis van extensies (bijv. verplaats `.mp4` automatisch naar `/Videos`).
- **🔒 Beveiligingsharding**: Opslagalgoritmen geüpgraded naar **bcrypt** met een naadloze migratiepatch om de accountbeveiliging te waarborgen.
- **✨ Minimalistische Esthetiek**: Diepgaand aangepaste **Glassmorphism** UI met **Kaola Green** branding voor een soepel, premium gevoel.

## Algemene Functies

- [x] **50+ Opslagmogelijkheden**: Lokaal, Aliyundrive, OneDrive, 189cloud, BaiduNetdisk, 115, S3, FTP/SFTP, enz.
- [x] **Hoge Prestaties**: Geoptimaliseerd gelijktijdig laden voor grote mappen en meerdere opslaglocaties.
- [x] **Implementatie met één klik**: Geleverd met `kaolalist_deploy.sh` script en Docker-images voor snelle installatie.
- [x] **Geavanceerde Editor**: Diepgaand aangepaste Monaco Editor en real-time Markdown-voorbeeld.
- [x] **Offline Download**: Ondersteuning voor Magnets en BT-overdrachten.
- [x] **Automatische Back-up**: Real-time GitHub naar Gitee project mirroring.

## 🚀 Snel aan de slag

### Docker Implementatie (Aanbevolen)
We raden aan om onze officiële images te gebruiken, die multi-arch (amd64/arm64) ondersteunen:
```bash
# Officiële Image op Docker Hub
docker pull canghaix/kaolalist:latest

# GitHub Container Registry (GHCR) Back-up
docker pull ghcr.io/bitkoala/kaolalist-git:latest
```

### Wereldwijde Mirrors
Voor gebruikers in regio's met beperkte toegang tot GitHub, gebruik onze Gitee mirror:
🔗 **Gitee Repository**: [https://gitee.com/bitekaola/KaolaList](https://gitee.com/bitekaola/KaolaList)

## Licentie & Naleving

Deze repository is open-source onder de **AGPL-3.0** licentie.

**Kernverklaringen**:
1. **Overerving & Wijziging**: KaolaList is een afgeleide van het [OpenList](https://github.com/OpenListTeam/OpenList) project. We houden ons strikt aan AGPL-3.0, en alle wijzigingen in de kernlogica blijven open-source.
2. **Disclaimer**: Dit project is alleen voor technische uitwisseling en leren. Gebruikers moeten voldoen aan de lokale wetgeving. De ontwikkelaars zijn niet verantwoordelijk voor door gebruikers opgeslagen inhoud of eventuele verliezen.
3. **Branding**: "KaolaList" en de bijbehorende Logo-activa zijn beschermd. Vermeld a.u.b. de bron voor afgeleide werken.

## Dankbetuigingen

- Dank aan het [OpenList Team](https://github.com/OpenListTeam) voor het bieden van de solide basis.
- Dank aan [AList](https://github.com/AlistGo/alist) voor het pionieren van het persoonlijke cloud-ecosysteem.

---
🐨 **KaolaList - Geboren voor veiligheid, gedreven door ervaring.**
