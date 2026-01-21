<div align="center">
  <img style="width: 128px; height: 128px;" src="https://jsd.cdn.zzko.cn/gh/AlistGo/alist-web@master/public/logo.svg" alt="logo" />

  <h1>KaolaList</h1>
  <p><em>Premium Personal Cloud Portal | High-End Netdisk Gateway</em></p>
  <p>KaolaList is an advanced fork of OpenList, designed to provide a more secure, intelligent, and aesthetically pleasing personal cloud experience.</p>

  <img src="https://img.shields.io/badge/Version-KaolaList--v1.0.0--alpha-68D391" alt="version" />
  <a href="./LICENSE"><img src="https://img.shields.io/github/license/OpenListTeam/OpenList" alt="License" /></a>
  <img src="https://img.shields.io/badge/OpenSource-AGPL--3.0-orange" alt="AGPL-3.0" />
</div>

---

English | [中文](./README_cn.md) | [日本語](./README_ja.md) | [Dutch](./README_nl.md)

---

## 🐨 Kaola Series Core Features (Advanced)

KaolaList integrates the following high-end features on top of the robust OpenList foundation:

- **🛡️ Kaola Vault**: Supports folder-level **MFA (Multi-Factor Authentication)**. Add an extra layer of security to your critical or private data using TOTP/WebAuthn.
- **📺 Kaola Theater**: An ultimate viewing experience. Supports professional video formats like `rmvb`, `m2ts` and features **Fully Automated Subtitle Scanning & Intelligent Matching** (`.srt`, `.ass`, `.vtt`).
- **🧠 Local Intelligent OCR**: A plugin-based OCR architecture. Supports **PaddleOCR** for local inference, ensuring zero external API dependency and maximum privacy.
- **🤖 Kaola Automation**: 
  - **Kaola Clone (Webhook)**: Real-time notifications for file changes, easily integrating with your personal automation workflows.
  - **Auto-routing**: Automatically archive files based on extensions (e.g., auto-move `.mp4` to `/Videos`).
- **🔒 Security Hardening**: Upgraded storage algorithms to **bcrypt** with a seamless migration patch to ensure account security.
- **✨ Minimalist Aesthetics**: Deeply customized **Glassmorphism** UI with **Kaola Green** branding for a smooth, premium feel.

## General Features

- [x] **50+ Storage Backends**: Local, Aliyundrive, OneDrive, 189cloud, BaiduNetdisk, 115, S3, FTP/SFTP, etc.
- [x] **High Performance**: Optimized concurrent loading for large directories and multiple storages.
- [x] **One-Click Deployment**: Provided `kaolalist_deploy.sh` script and Docker images for quick setup.
- [x] **Advanced Editor**: Deeply customized Monaco Editor and real-time Markdown preview.
- [x] **Offline Download**: Support for Magnets and BT transfers.
- [x] **Auto-Backup**: Real-time GitHub to Gitee project mirroring.

## 🚀 Quick Start

### Docker Deployment (Recommended)
We recommend using our official images, which support multi-arch (amd64/arm64):
```bash
# Official Image on Docker Hub
docker pull canghaix/kaolalist:latest

# GitHub Container Registry (GHCR) Backup
docker pull ghcr.io/bitkoala/kaolalist-git:latest
```

### Global Mirrors
For users in mainland China or other regions with restricted GitHub access, use our Gitee mirror:
🔗 **Gitee Repository**: [https://gitee.com/bitekaola/KaolaList](https://gitee.com/bitekaola/KaolaList)

## License & Compliance

This repository is open-source under the **AGPL-3.0** license.

**Core Statements**:
1. **Inheritance & Modification**: KaolaList is a derivative of the [OpenList](https://github.com/OpenListTeam/OpenList) project. We strictly adhere to AGPL-3.0, and all modifications to core logic remain open-source.
2. **Disclaimer**: This project is for technical exchange and learning only. Users must comply with local laws. The developers are not responsible for user-stored content or any losses.
3. **Branding**: "KaolaList" and its associated Logo assets are protected. Please credit the source for derivative works.

## Acknowledgments

- Credits to the [OpenList Team](https://github.com/OpenListTeam) for providing the solid foundation.
- Credits to [AList](https://github.com/AlistGo/alist) for pioneering the personal cloud ecosystem.

---
🐨 **KaolaList - Born for Security, Driven by Experience.**
