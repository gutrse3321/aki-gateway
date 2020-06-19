package route

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 16:13
 * @Title:
 * --- --- ---
 * @Desc:
 */

type RoutesOption struct {
	Routes map[string]string
}

func NewRoutesOption(v *viper.Viper) (*RoutesOption, error) {
	var err error

	opt := &RoutesOption{}
	if err = v.UnmarshalKey("app", opt); err != nil {
		return nil, errors.Wrap(err, "unmarshal app config error")
	}

	return opt, err
}

type RoutesMap struct {
	hashMap map[string]string
}

func NewRoutesMap(opt *RoutesOption) *RoutesMap {
	return &RoutesMap{hashMap: opt.Routes}
}

func (r *RoutesMap) Get(key string) string {
	return r.hashMap[key]
}
