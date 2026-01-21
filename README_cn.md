<div align="center">
  <img style="width: 128px; height: 128px;" src="https://raw.githubusercontent.com/KaolaList/Branding/main/logo.svg" alt="logo" />

  <h1>KaolaList</h1>
  <p><em>Premium Personal Cloud Portal | 高端私人云网盘网关</em></p>
  <p>KaolaList 是基于 OpenList 的二次开发版本，旨在提供更加安全、智能且极具美感的私人云访问体验。</p>

  <img src="https://img.shields.io/badge/Version-KaolaList--v1.0.0--alpha-68D391" alt="version" />
  <a href="./LICENSE"><img src="https://img.shields.io/github/license/OpenListTeam/OpenList" alt="License" /></a>
  <img src="https://img.shields.io/badge/OpenSource-AGPL--3.0-orange" alt="AGPL-3.0" />
</div>

---

[English](./README.md) | 中文 | [日本語](./README_ja.md) | [Dutch](./README_nl.md)

---

## 🐨 考拉系列核心特性 (Advanced Features)

KaolaList 在继承 OpenList 强大功能的基础上，深度集成了以下高级特性：

- **🛡️ 考拉保险箱 (Kaola Vault)**: 支持文件夹级别的 **MFA (多因素身份验证)**。为您的核心业务或私密数据文件夹加一层锁，支持 TOTP/WebAuthn。
- **📺 考拉剧场 (Kaola Theater)**: 极致的观影体验。支持 `rmvb`, `m2ts` 等多种专业视频格式预览，并实现 **全自动字幕扫描与智能匹配** (`.srt`, `.ass`, `.vtt`)。
- **🧠 本地智能 OCR (Local OCR)**: 插件化 OCR 架构。支持集成 **PaddleOCR** 本地推理，实现验证码识别零偏移、零外部 API 依赖，极致保护隐私。
- **🤖 考拉自动化 (Automation)**: 
  - **考拉分身 (Webhook)**: 文件变更实时推送到指定接口，轻松对接个人自动化流。
  - **自动分拣 (Auto-routing)**: 基于文件后缀自动归档（如自动将 `.mp4` 移动至 `/Videos`）。
- **🔒 安全加固**: 全面升级存储算法至 **bcrypt**，并提供平滑迁移补丁，确保账户密码万无一失。
- **✨ 极简美学**: 深度定制的 **玻璃拟态 (Glassmorphism)** UI 与 **考拉绿** 品牌设计，带来丝滑平衡的感官享受。

## 通用功能

- [x] **多种存储支持**: 本地、阿里云盘、OneDrive、天翼云、百度网盘、115、S3、FTP/SFTP 等 50+ 种。
- [x] **高性能遍历**: 优化了根目录与多存储切换时的并发加载速度。
- [x] **一键部署**: 提供 `kaolalist_deploy.sh` 脚本及 Docker 镜像，快速上线。
- [x] **深度编辑器定制**: 深度优化的 Monaco Editor / Markdown 实时预览。
- [x] **离线下载**: 支持磁力、BT 转存。
- [x] **全自动备份**: 核心项目 GitHub -> Gitee 实时同步。

## 🚀 快速开始

### Docker 部署 (推荐)
我们建议使用以下镜像，支持多架构 (amd64/arm64)：
```bash
# Docker Hub 官方镜像
docker pull canghaix/kaolalist:latest

# GitHub Container Registry (GHCR) 备份镜像
docker pull ghcr.io/bitkoala/kaolalist-git:latest
```

### 国内加速与镜像
如果您在中国大陆访问 GitHub 或 Docker Hub 较慢，可以使用我们的 Gitee 镜像：
🔗 **Gitee 仓库**: [https://gitee.com/bitekaola/KaolaList](https://gitee.com/bitekaola/KaolaList)

## 许可证与合规

本仓库代码基于 **AGPL-3.0** 许可证开源。

**核心声明**：
1. **传承与修改**：KaolaList 是 [OpenList](https://github.com/OpenListTeam/OpenList) 项目的衍生版本。我们严格遵守 AGPL-3.0 协议，所有对核心逻辑的修改均保持开源。
2. **免责声明**：本项目仅供技术交流与学习使用。用户在使用过程中应遵守当地法律法规，开发者不对用户存储内容的合法性及因使用本软件导致的任何损失负责。
3. **品牌保护**：“KaolaList”及其相关 Logo 资产受品牌保护，衍生开发请注明来源。

## 鸣谢

- 感谢原项目作者 [OpenList Team](https://github.com/OpenListTeam) 提供的坚实底座。
- 感谢 [AList](https://github.com/AlistGo/alist) 为个人云生态做出的开拓性贡献。

---
🐨 **KaolaList - 为安全而生，为体验而行。**
