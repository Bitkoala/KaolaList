<div align="center">
  <img style="width: 128px; height: 128px;" src="https://raw.githubusercontent.com/KaolaList/Branding/main/logo.svg" alt="logo" />

  <h1>KaolaList</h1>
  <p><em>プレミアム・パーソナル・クラウド・ポータル | 高機能クラウドストレージ・ゲートウェイ</em></p>
  <p>KaolaList は OpenList の高度なフォークであり、より安全で、インテリジェントで、審美的に優れたパーソナルクラウド体験を提供するために設計されています。</p>

  <img src="https://img.shields.io/badge/Version-KaolaList--v1.0.0--alpha-68D391" alt="version" />
  <a href="./LICENSE"><img src="https://img.shields.io/github/license/OpenListTeam/OpenList" alt="License" /></a>
  <img src="https://img.shields.io/badge/OpenSource-AGPL--3.0-orange" alt="AGPL-3.0" />
</div>

---

[English](./README.md) | [中文](./README_cn.md) | 日本語 | [Dutch](./README_nl.md)

---

- [コントリビュート](./CONTRIBUTING.md)
- [行動規範](./CODE_OF_CONDUCT.md)
- [ライセンス](./LICENSE)

## 免責事項

OpenListは、OpenListチームが独立して維持するオープンソースプロジェクトであり、AGPL-3.0ライセンスに従い、完全なコードの開放性と変更の透明性を維持することに専念しています。

コミュニティ内で、OpenListApp/OpenListAppなど、本プロジェクトと類似した名称を持つサードパーティプロジェクトや、同一または類似した命名を採用する有料専有ソフトウェアが出現していることを確認しています。ユーザーの誤解を避けるため、以下のように宣言いたします：

- OpenListは、いかなるサードパーティ派生プロジェクトとも公式な関連性はありません。

- 本プロジェクトのすべてのソフトウェア、コード、サービスはOpenListチームによって維持され、GitHubで無料で取得できます。

- プロジェクトドキュメントとAPIサービスは主にCloudflareが提供する公益リソースに依存しており、現在有料プランや商業展開はなく、既存機能の使用に費用は発生しません。

私たちはコミュニティの自由な使用と派生開発の権利を尊重しますが、下流プロジェクトに強く呼びかけます：

- 「OpenList」の名前で偽装宣伝や商業利益を得るべきではありません；

- OpenListベースのコードをクローズドソースで配布したり、AGPLライセンス条項に違反してはいけません。

エコシステムの健全な発展をより良く維持するため、以下を推奨します：

- プロジェクトの出典を明確に示し、オープンソース精神に合致する適切なオープンソースライセンスを選択する；

- 商業用途が関わる場合は、「OpenList」や混乱を招く可能性のある名前をプロジェクト名として使用することを避ける；

- OpenListTeam/Logo下の素材を使用する必要がある場合は、協定を遵守した上で修正して使用できます。

OpenListプロジェクトへのご支援とご理解をありがとうございます。

## 🐨 考拉（コアラ）シリーズのコア機能 (Advanced Features)

KaolaList は、OpenList の強力な基盤の上に、以下のハイエンド機能を統合しています：

- **🛡️ 考拉（コアラ）金庫 (Kaola Vault)**: フォルダレベルの **MFA (多要素認証)** をサポート。TOTP や WebAuthn を使用して、重要なデータフォルダにセキュリティの層を追加します。
- **📺 考拉（コアラ）シアター (Kaola Theater)**: 究極の視聴体験。`rmvb`, `m2ts` 等のプロフェッショナルなビデオ形式のプレビューに対応し、**字幕の自動スキャンとインテリジェント・マッチング** (`.srt`, `.ass`, `.vtt`) を実現。
- **🧠 ローカル・インテリジェント OCR (Local OCR)**: プラグインベースの OCR アーキテクチャ。**PaddleOCR** のローカル推論をサポートし、外部 API への依存をゼロにすることで、プライバシーを極限まで保護します。
- **🤖 考拉（コアラ）オートメーション (Automation)**: 
  - **考拉（コアラ）クローン (Webhook)**: ファイルの変更をリアルタイムで指定したエンドポイントに通知。
  - **自動ルーティング (Auto-routing)**: 拡張子に基づいてファイルを自動的に整理（例：`.mp4` を `/Videos` へ自動移動）。
- **🔒 セキュリティ強化**: ストレージアルゴリズムを **bcrypt** にアップグレード。既存ユーザー向けのシームレスな移行パッチも提供しています。
- **✨ ミニマリストの審美性**: 独自にカスタマイズされた **グラスモーフィズム (Glassmorphism)** UI と、**コアラ・グリーン** のブランドデザインにより、滑らかで上質な操作感を提供します。

## 一般機能

- [x] 複数ストレージ
  - [x] ローカルストレージ
  - [x] [Aliyundrive](https://www.alipan.com)
  - [x] OneDrive / Sharepoint ([グローバル](https://www.microsoft.com/en-us/microsoft-365/onedrive/online-cloud-storage), [中国](https://portal.partner.microsoftonline.cn), DE, US)
  - [x] [189cloud](https://cloud.189.cn)（個人、家族）
  - [x] [GoogleDrive](https://drive.google.com)
  - [x] [123pan](https://www.123pan.com)
  - [x] [FTP / SFTP](https://en.wikipedia.org/wiki/File_Transfer_Protocol)
  - [x] [PikPak](https://www.mypikpak.com)
  - [x] [S3](https://aws.amazon.com/s3)
  - [x] [Seafile](https://seafile.com)
  - [x] [UPYUN Storage Service](https://www.upyun.com/products/file-storage)
  - [x] [WebDAV](https://en.wikipedia.org/wiki/WebDAV)
  - [x] Teambition([中国](https://www.teambition.com), [国際](https://us.teambition.com))
  - [x] [Mediatrack](https://www.mediatrack.cn)
  - [x] [ProtonDrive](https://proton.me/drive)
  - [x] [139yun](https://yun.139.com)（個人、家族、グループ）
  - [x] [YandexDisk](https://disk.yandex.com)
  - [x] [BaiduNetdisk](http://pan.baidu.com)
  - [x] [Terabox](https://www.terabox.com/main)
  - [x] [UC](https://drive.uc.cn)
  - [x] [Quark](https://pan.quark.cn)
  - [x] [Thunder](https://pan.xunlei.com)
  - [x] [Lanzou](https://www.lanzou.com)
  - [x] [ILanzou](https://www.ilanzou.com)
  - [x] [Google photo](https://photos.google.com)
  - [x] [Mega.nz](https://mega.nz)
  - [x] [Baidu photo](https://photo.baidu.com)
  - [x] [SMB](https://en.wikipedia.org/wiki/Server_Message_Block)
  - [x] [115](https://115.com)
  - [x] [Cloudreve](https://cloudreve.org)
  - [x] [Dropbox](https://www.dropbox.com)
  - [x] [FeijiPan](https://www.feijipan.com)
  - [x] [dogecloud](https://www.dogecloud.com/product/oss)
  - [x] [Azure Blob Storage](https://azure.microsoft.com/products/storage/blobs)
  - [x] [Chaoxing](https://www.chaoxing.com)
  - [x] [CNB](https://cnb.cool/)
  - [x] [Degoo](https://degoo.com)
  - [x] [Doubao](https://www.doubao.com)
  - [x] [Febbox](https://www.febbox.com)
  - [x] [GitHub](https://github.com)
  - [x] [OpenList](https://github.com/OpenListTeam/OpenList)
  - [x] [Teldrive](https://github.com/tgdrive/teldrive)
  - [x] [Weiyun](https://www.weiyun.com)
  - [x] [MediaFire](https://www.mediafire.com)
- [x] 簡単にデプロイでき、すぐに使える
- [x] ファイルプレビュー（PDF、markdown、コード、テキストなど）
- [x] ギャラリーモードでの画像プレビュー
- [x] ビデオ・オーディオプレビュー、歌詞・字幕対応
- [x] Officeドキュメントプレビュー（docx、pptx、xlsxなど）
- [x] `README.md` プレビュー表示
- [x] ファイルのパーマリンクコピーと直接ダウンロード
- [x] ダークモード
- [x] 国際化対応
- [x] 保護されたルート（パスワード保護と認証）
- [x] WebDAV
- [x] Dockerデプロイ
- [x] Cloudflare Workersプロキシ
- [x] ファイル/フォルダのパッケージダウンロード
- [x] Webアップロード（訪問者のアップロード許可可）、削除、フォルダ作成、リネーム、移動、コピー
- [x] オフラインダウンロード
- [x] ストレージ間のファイルコピー
- [x] 単一ファイルのマルチスレッドダウンロード/ストリーム加速

## ドキュメント

- 📘 [グローバルサイト](https://doc.oplist.org)
- 📚 [バックアップサイト](https://doc.openlist.team)
- 🌏 [CNサイト](https://doc.oplist.org.cn)

## デモ

- 🌎 [グローバルデモ](https://demo.oplist.org)
- 🇨🇳 [CNデモ](https://demo.oplist.org.cn)

## ディスカッション

一般的な質問は [*Discussions*](https://github.com/OpenListTeam/OpenList/discussions) をご利用ください。***Issues* はバグ報告と機能リクエスト専用です。**

## スポンサー

[![VPS.Town](https://vps.town/static/images/sponsor.png)](https://vps.town "VPS.Town - Trust, Effortlessly. Your Cloud, Reimagined.")

## ライセンスとコンプライアンス

本リポジトリのコードは **AGPL-3.0** ライセンスの下で公開されています。

**重要な声明**：
1. **継承と変更**：KaolaList は [OpenList](https://github.com/OpenListTeam/OpenList) プロジェクトの派生版です。AGPL-3.0 プロトコルを厳守しており、コアロジックに対するすべての変更はオープンソースとして維持されます。
2. **免責事項**：本プロジェクトは技術交流および学習用です。利用にあたっては現地の法令を遵守してください。開発者は、ユーザーが保存したコンテンツの合法性、および本ソフトウェアの使用に起因するいかなる損失についても責任を負いません。
3. **ブランド保護**：「KaolaList」および関連するロゴ資産はブランド保護されています。派生開発を行う際は出典を明記してください。

## お問い合わせ

- [@GitHub](https://github.com/OpenListTeam)
- [Telegram グループ](https://t.me/OpenListTeam)
- [Telegram チャンネル](https://t.me/OpenListOfficial)

## コントリビューター

オリジナルプロジェクト [AlistGo/alist](https://github.com/AlistGo/alist) の作者 [Xhofe](https://github.com/Xhofe) およびその他すべての貢献者に心より感謝いたします。

素晴らしい皆様に感謝します：

[![Contributors](https://contrib.rocks/image?repo=OpenListTeam/OpenList)](https://github.com/OpenListTeam/OpenList/graphs/contributors)

## 謝辞

- 強固な基盤を提供してくれた [OpenList Team](https://github.com/OpenListTeam) に感謝します。
- パーソナルクラウドのエコシステムを切り拓いた [AList](https://github.com/AlistGo/alist) に感謝します。

---
🐨 **KaolaList - 安全のために生まれ、日常のために進化する。**
