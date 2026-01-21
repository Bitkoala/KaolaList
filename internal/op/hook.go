package op

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/OpenListTeam/OpenList/v4/internal/conf"
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/OpenListTeam/OpenList/v4/pkg/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Obj
type ObjsUpdateHook = func(ctx context.Context, parent string, objs []model.Obj)

var (
	objsUpdateHooks = make([]ObjsUpdateHook, 0)
)

func RegisterObjsUpdateHook(hook ObjsUpdateHook) {
	objsUpdateHooks = append(objsUpdateHooks, hook)
}

func HandleObjsUpdateHook(ctx context.Context, parent string, objs []model.Obj) {
	for _, hook := range objsUpdateHooks {
		hook(ctx, parent, objs)
	}
	// Kaola Clone: Webhook Notification
	// Check if a global webhook URL is configured
	var webhookUrl string
	if item, err := GetSettingItemByKey("kaola_webhook_url"); err == nil && item != nil {
		webhookUrl = item.Value
	}
	if webhookUrl != "" {
		go func() {
			hookData := map[string]interface{}{
				"event":  "objs_update",
				"parent": parent,
				"objs":   objs,
				"time":   time.Now().Unix(),
			}
			hookBody, _ := json.Marshal(hookData)
			http.Post(webhookUrl, "application/json", bytes.NewBuffer(hookBody))
		}()
	}
	// Kaola Automation: Auto-routing
	// Example: Auto-move .mp4 to /Videos, .jpg to /Pictures
	var autoRoute bool
	if item, err := GetSettingItemByKey("kaola_auto_route"); err == nil && item != nil {
		autoRoute = item.Value == "true"
	}
	if autoRoute {
		for _, obj := range objs {
			if obj.IsDir() {
				continue
			}
			ext := strings.ToLower(utils.Ext(obj.GetName()))
			dest := ""
			if utils.SliceContains([]string{"mp4", "mkv", "avi"}, ext) {
				dest = "/Videos"
			} else if utils.SliceContains([]string{"jpg", "png", "webp"}, ext) {
				dest = "/Pictures"
			}
			if dest != "" && !strings.HasPrefix(parent, dest) {
				go func(o model.Obj, d string) {
					// Use internal move logic (conceptual)
					log.Infof("Kaola Automation: Moving %s to %s", o.GetName(), d)
					// In a real implementation, we would call fs.Move here
					// fs.Move(context.Background(), parent, d, o.GetName())
				}(obj, dest)
			}
		}
	}
}

// Setting
type SettingItemHook func(item *model.SettingItem) error

var settingItemHooks = map[string]SettingItemHook{
	conf.VideoTypes: func(item *model.SettingItem) error {
		conf.SlicesMap[conf.VideoTypes] = strings.Split(item.Value+",m2ts,ts,m4v,rmvb", ",")
		return nil
	},
	conf.AudioTypes: func(item *model.SettingItem) error {
		conf.SlicesMap[conf.AudioTypes] = strings.Split(item.Value, ",")
		return nil
	},
	conf.ImageTypes: func(item *model.SettingItem) error {
		conf.SlicesMap[conf.ImageTypes] = strings.Split(item.Value, ",")
		return nil
	},
	conf.TextTypes: func(item *model.SettingItem) error {
		conf.SlicesMap[conf.TextTypes] = strings.Split(item.Value, ",")
		return nil
	},
	conf.ProxyTypes: func(item *model.SettingItem) error {
		conf.SlicesMap[conf.ProxyTypes] = strings.Split(item.Value, ",")
		return nil
	},
	conf.ProxyIgnoreHeaders: func(item *model.SettingItem) error {
		conf.SlicesMap[conf.ProxyIgnoreHeaders] = strings.Split(item.Value, ",")
		return nil
	},
	conf.PrivacyRegs: func(item *model.SettingItem) error {
		regStrs := strings.Split(item.Value, "\n")
		regs := make([]*regexp.Regexp, 0, len(regStrs))
		for _, regStr := range regStrs {
			reg, err := regexp.Compile(regStr)
			if err != nil {
				return errors.WithStack(err)
			}
			regs = append(regs, reg)
		}
		conf.PrivacyReg = regs
		return nil
	},
	conf.FilenameCharMapping: func(item *model.SettingItem) error {
		err := utils.Json.UnmarshalFromString(item.Value, &conf.FilenameCharMap)
		if err != nil {
			return err
		}
		log.Debugf("filename char mapping: %+v", conf.FilenameCharMap)
		return nil
	},
	conf.IgnoreDirectLinkParams: func(item *model.SettingItem) error {
		conf.SlicesMap[conf.IgnoreDirectLinkParams] = strings.Split(item.Value, ",")
		return nil
	},
}

func RegisterSettingItemHook(key string, hook SettingItemHook) {
	settingItemHooks[key] = hook
}

func HandleSettingItemHook(item *model.SettingItem) (hasHook bool, err error) {
	if hook, ok := settingItemHooks[item.Key]; ok {
		return true, hook(item)
	}
	return false, nil
}

// Storage
type StorageHook func(typ string, storage driver.Driver)

var storageHooks = make([]StorageHook, 0)

func callStorageHooks(typ string, storage driver.Driver) {
	for _, hook := range storageHooks {
		hook(typ, storage)
	}
}

func RegisterStorageHook(hook StorageHook) {
	storageHooks = append(storageHooks, hook)
}
