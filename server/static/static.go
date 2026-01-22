package static

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/OpenListTeam/OpenList/v4/drivers/base"
	"github.com/OpenListTeam/OpenList/v4/internal/conf"
	"github.com/OpenListTeam/OpenList/v4/internal/setting"
	"github.com/OpenListTeam/OpenList/v4/pkg/utils"
	"github.com/OpenListTeam/OpenList/v4/public"
	"github.com/gin-gonic/gin"
)

type ManifestIcon struct {
	Src   string `json:"src"`
	Sizes string `json:"sizes"`
	Type  string `json:"type"`
}

type Manifest struct {
	Display  string         `json:"display"`
	Scope    string         `json:"scope"`
	StartURL string         `json:"start_url"`
	Name     string         `json:"name"`
	Icons    []ManifestIcon `json:"icons"`
}

var static fs.FS

func initStatic() {
	utils.Log.Debug("Initializing static file system...")
	if conf.Conf.DistDir == "" {
		dist, err := fs.Sub(public.Public, "dist")
		if err != nil {
			utils.Log.Fatalf("failed to read dist dir: %v", err)
		}
		static = dist
		utils.Log.Debug("Using embedded dist directory")
		return
	}
	static = os.DirFS(conf.Conf.DistDir)
	utils.Log.Infof("Using custom dist directory: %s", conf.Conf.DistDir)
}

func replaceStrings(content string, replacements map[string]string) string {
	for old, new := range replacements {
		content = strings.Replace(content, old, new, -1)
	}
	return content
}

func initIndex(siteConfig SiteConfig) {
	utils.Log.Debug("Initializing index.html...")
	// dist_dir is empty and cdn is not empty, and web_version is empty or beta or dev or rolling
	if conf.Conf.DistDir == "" && conf.Conf.Cdn != "" && (conf.WebVersion == "" || conf.WebVersion == "beta" || conf.WebVersion == "dev" || conf.WebVersion == "rolling") {
		utils.Log.Infof("Fetching index.html from CDN: %s/index.html...", siteConfig.Cdn)
		resp, err := base.RestyClient.R().
			SetHeader("Accept", "text/html").
			Get(fmt.Sprintf("%s/index.html", siteConfig.Cdn))
		if err != nil {
			utils.Log.Fatalf("failed to fetch index.html from CDN: %v", err)
		}
		if resp.StatusCode() != http.StatusOK {
			utils.Log.Fatalf("failed to fetch index.html from CDN, status code: %d", resp.StatusCode())
		}
		conf.RawIndexHtml = string(resp.Body())
		utils.Log.Info("Successfully fetched index.html from CDN")
	} else {
		utils.Log.Debug("Reading index.html from static files system...")
		indexFile, err := static.Open("index.html")
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				utils.Log.Fatalf("index.html not exist, you may forget to put dist of frontend to public/dist")
			}
			utils.Log.Fatalf("failed to read index.html: %v", err)
		}
		defer func() {
			_ = indexFile.Close()
		}()
		index, err := io.ReadAll(indexFile)
		if err != nil {
			utils.Log.Fatalf("failed to read dist/index.html")
		}
		conf.RawIndexHtml = string(index)
		utils.Log.Debug("Successfully read index.html from static files system")
	}
	utils.Log.Debug("Replacing placeholders in index.html...")
	// Construct the correct manifest path based on basePath
	manifestPath := "/manifest.json"
	if siteConfig.BasePath != "/" {
		manifestPath = siteConfig.BasePath + "/manifest.json"
	}
	replaceMap := map[string]string{
		"cdn: undefined":        fmt.Sprintf("cdn: '%s'", siteConfig.Cdn),
		"base_path: undefined":  fmt.Sprintf("base_path: '%s'", siteConfig.BasePath),
		`href="/manifest.json"`: fmt.Sprintf(`href="%s"`, manifestPath),
		"lang: undefined":       "lang: 'zh-CN'",
	}
	conf.RawIndexHtml = replaceStrings(conf.RawIndexHtml, replaceMap)
	UpdateIndex()
}

func UpdateIndex() {
	utils.Log.Debug("Updating index.html with settings...")
	favicon := setting.GetStr(conf.Favicon)
	logo := strings.Split(setting.GetStr(conf.Logo), "\n")[0]
	title := setting.GetStr(conf.SiteTitle)
	customizeHead := setting.GetStr(conf.CustomizeHead)
	customizeBody := setting.GetStr(conf.CustomizeBody)
	mainColor := setting.GetStr(conf.MainColor)
	utils.Log.Debug("Applying replacements for default pages...")
	replaceMap1 := map[string]string{
		"https://res.oplist.org/logo/logo.svg":     favicon,
		"https://res.oplist.org/logo/logo.png":     logo,
		"Loading...":                               title,
		"main_color: undefined":                    fmt.Sprintf("main_color: '%s'", mainColor),
		"OpenList":                                 "KaolaList",
		"AList":                                    "KaolaList",
		"由 OpenList 驱动":                            "由 KaolaList 驱动",
		"Powered by OpenList":                      "Powered by KaolaList",
		"OpenList 管理":                              "KaolaList 管理",
		"管理 OpenList":                              "管理 KaolaList",
		"https://openlist.team":                    "https://bitekaola.com",
		"https://docs.openlist.team":               "https://bitekaola.com",
		"https://github.com/OpenListTeam/OpenList": "https://gitee.com/bitekaola/KaolaList",
		"OpenList 官方文档":                            "KaolaList 官方文档",
		"OpenList 讨论区":                             "KaolaList 社区",
		"Management":                               "管理后台",
		"Settings center":                          "设置中心",
		"Settings":                                 "功能设置",
		"Storage":                                  "存储管理",
		"User":                                     "用户管理",
		"Profile":                                  "个人中心",
		"Logout":                                   "安全退出",
		"Back to home":                             "返回首页",
		"Guest":                                    "游客",
		"Search":                                   "搜索",
		"Download":                                 "下载",
		"Details":                                  "详情",
		"Login to":                                 "登录到",
		"<!-- customize head -->": `<style>
/* KaolaList Nuclear Premium Theme & Brand Purge */
:root {
  --hope-colors-primary9: #68D391 !important;
  --hope-colors-primary10: #48BB78 !important;
}
html, body, #root, [class*="hope-ui-"], [class*="hope-stack"] {
  background-image: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%) !important;
  background-color: transparent !important;
  background-attachment: fixed !important;
  background-size: cover !important;
  background-repeat: no-repeat !important;
  background-attachment: fixed !important;
}
/* Hide Footer / Branding */
footer, [class*="footer"], .footer, [class*="hope-stack"] > div > a[href*="github.com/alist-org"], [class*="hope-stack"] > div > a[href*="github.com/OpenListTeam"] {
  display: none !important;
  opacity: 0 !important;
  visibility: hidden !important;
}
body::before, body::after, [class*="hope-stack"]::before, [class*="hope-stack"]::after {
  display: none !important;
}
.hope-center, .hope-modal__content, aside, main, .hope-stack.hope-center > div {
  backdrop-filter: blur(20px) saturate(180%) !important;
  -webkit-backdrop-filter: blur(20px) saturate(180%) !important;
  background-color: rgba(255, 255, 255, 0.5) !important;
  border: 1px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.1) !important;
  border-radius: 18px !important;
}
.hope-ui-dark .hope-center, .hope-ui-dark .hope-modal__content, .hope-ui-dark aside, .hope-ui-dark main {
  background-color: rgba(23, 25, 35, 0.5) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
}
</style>
<script>
// 🐨 Kaola Brain - Ultimate UI Brand Purge & Localization Script
(function() {
  const trans = {
    'Ai gemini endpoint': 'Gemini API 终点',
    'Ai gemini keys': 'Gemini API 密钥',
    'Ai gemini keys-tips': '一行一个密钥，支持多 Key 轮询',
    'Ai doubao endpoint': '豆包 API 终点',
    'Ai doubao keys': '豆包 API 密钥',
    'Ai doubao keys-tips': '一行一个密钥，支持图片分析',
    'Ai doubao model': '豆包模型 ID',
    'Ai feature translation': 'AI 翻译功能引擎',
    'Ai feature summary': 'AI 摘要功能引擎',
    'Ai feature ocr plus': 'AI 识图增强模式',
    'OpenList Management': 'KaolaList 管理后台',
    'OpenList Manage': 'KaolaList 管理',
    'OpenList 管理': 'KaolaList 管理',
    '管理 OpenList': '管理 KaolaList',
    'OpenList Official Documentation': 'KaolaList 官方文档',
    'OpenList docs': 'KaolaList 文档',
    'OpenList discussion': 'KaolaList 社区',
    'OpenList': 'KaolaList',
    'AList': 'KaolaList'
  };

  const purge = () => {
    // 1. Text Nodes
    const walker = document.createTreeWalker(document.body, NodeFilter.SHOW_TEXT, null, false);
    let node;
    while(node = walker.nextNode()) {
      let val = node.nodeValue;
      for (const [eng, chi] of Object.entries(trans)) {
        if (val.includes(eng)) { val = val.split(eng).join(chi); }
      }
      if (val !== node.nodeValue) node.nodeValue = val;
    }
    // 2. Document Title
    for (const [eng, chi] of Object.entries(trans)) {
      if (document.title.includes(eng)) { document.title = document.title.split(eng).join(chi); }
    }
    // 3. Links & Redirection
    document.querySelectorAll('a').forEach(a => {
      if (a.href.includes('openlist.team') || a.href.includes('github.com/OpenListTeam')) {
        a.href = 'https://bitekaola.com';
        for (const [eng, chi] of Object.entries(trans)) {
          if (a.innerText.includes(eng)) { a.innerText = a.innerText.split(eng).join(chi); }
        }
      }
    });
  };

  const observer = new MutationObserver(purge);
  observer.observe(document.body, { childList: true, subtree: true, characterData: true });
  window.addEventListener('load', purge);
  purge(); 
})();
</script>`,
		"Copy":                   "复制",
		"Move":                   "移动",
		"Rename":                 "重命名",
		"Delete":                 "删除",
		"New Folder":             "新建文件夹",
		"Upload":                 "上传",
		"Refresh":                "刷新",
		"AI Features":            "考拉大脑 AI",
		"AI Gemini Key":          "Gemini 密钥",
		"AI Doubao Key":          "豆包密钥",
		"AI Feature Settings":    "AI 功能分发",
		"Ai gemini endpoint":     "Gemini API 终点",
		"Ai gemini keys":         "Gemini API 密钥",
		"Ai gemini keys-tips":    "一行一个密钥，支持自动轮询",
		"Ai doubao endpoint":     "豆包 API 终点",
		"Ai doubao keys":         "豆包 API 密钥",
		"Ai doubao keys-tips":    "多 Key 轮询，支持识图",
		"Ai doubao model":        "豆包模型 ID",
		"Ai feature translation": "翻译引擎分配",
		"Ai feature summary":     "摘要引擎分配",
		"Ai feature ocr plus":    "识图增强模式",
	}
	conf.ManageHtml = replaceStrings(conf.RawIndexHtml, replaceMap1)
	utils.Log.Debug("Applying replacements for manage pages...")
	replaceMap2 := map[string]string{
		"<!-- customize head -->": customizeHead,
		"<!-- customize body -->": customizeBody,
	}
	conf.IndexHtml = replaceStrings(conf.ManageHtml, replaceMap2)
	utils.Log.Debug("Index.html update completed")
}

func ManifestJSON(c *gin.Context) {
	// Get site configuration to ensure consistent base path handling
	siteConfig := getSiteConfig()

	// Get site title from settings
	siteTitle := setting.GetStr(conf.SiteTitle)

	// Get logo from settings, use the first line (light theme logo)
	logoSetting := setting.GetStr(conf.Logo)
	logoUrl := strings.Split(logoSetting, "\n")[0]

	// Use base path from site config for consistency
	basePath := siteConfig.BasePath

	// Determine scope and start_url
	// PWA scope and start_url should always point to our application's base path
	// regardless of whether static resources come from CDN or local server
	scope := basePath
	startURL := basePath

	manifest := Manifest{
		Display:  "standalone",
		Scope:    scope,
		StartURL: startURL,
		Name:     siteTitle,
		Icons: []ManifestIcon{
			{
				Src:   logoUrl,
				Sizes: "512x512",
				Type:  "image/png",
			},
		},
	}

	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "public, max-age=3600") // cache for 1 hour

	if err := json.NewEncoder(c.Writer).Encode(manifest); err != nil {
		utils.Log.Errorf("Failed to encode manifest.json: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate manifest"})
		return
	}
}

func Static(r *gin.RouterGroup, noRoute func(handlers ...gin.HandlerFunc)) {
	utils.Log.Debug("Setting up static routes...")
	siteConfig := getSiteConfig()
	initStatic()
	initIndex(siteConfig)
	folders := []string{"assets", "images", "streamer", "static"}

	if conf.Conf.Cdn == "" {
		utils.Log.Debug("Setting up static file serving...")
		r.Use(func(c *gin.Context) {
			for _, folder := range folders {
				if strings.HasPrefix(c.Request.RequestURI, fmt.Sprintf("/%s/", folder)) {
					c.Header("Cache-Control", "public, max-age=15552000")
				}
			}
		})
		for _, folder := range folders {
			sub, err := fs.Sub(static, folder)
			if err != nil {
				utils.Log.Fatalf("can't find folder: %s", folder)
			}
			utils.Log.Debugf("Setting up route for folder: %s", folder)
			r.StaticFS(fmt.Sprintf("/%s/", folder), http.FS(sub))
		}
	} else {
		// Ensure static file redirected to CDN
		for _, folder := range folders {
			r.GET(fmt.Sprintf("/%s/*filepath", folder), func(c *gin.Context) {
				filepath := c.Param("filepath")
				c.Redirect(http.StatusFound, fmt.Sprintf("%s/%s%s", siteConfig.Cdn, folder, filepath))
			})
		}
	}

	utils.Log.Debug("Setting up catch-all route...")
	noRoute(func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "POST" {
			c.Status(405)
			return
		}
		c.Header("Content-Type", "text/html")
		c.Status(200)
		if strings.HasPrefix(c.Request.URL.Path, "/@manage") {
			_, _ = c.Writer.WriteString(conf.ManageHtml)
		} else {
			_, _ = c.Writer.WriteString(conf.IndexHtml)
		}
		c.Writer.Flush()
		c.Writer.WriteHeaderNow()
	})
}
