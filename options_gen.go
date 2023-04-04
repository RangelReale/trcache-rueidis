// Code generated by troptgen. DO NOT EDIT.

package trcache_rueidis

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

func WithDefaultClientSideDuration[K comparable, V any](duration time.Duration) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.DefaultClientSideDuration"
	const optionHash = uint64(0xa0dfb3eca67e6cf9)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDefaultClientSideDuration(duration)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithDefaultDuration[K comparable, V any](duration time.Duration) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.DefaultDuration"
	const optionHash = uint64(0xd15a576d42aa4dad)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDefaultDuration(duration)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.KeyCodec"
	const optionHash = uint64(0x1f5587bc972a6a57)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.RedisDelFunc"
	const optionHash = uint64(0x1f564119f8ee03d4)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.RedisGetFunc"
	const optionHash = uint64(0x69ed94d779c45687)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.RedisSetFunc"
	const optionHash = uint64(0xa484750639da0993)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.Validator"
	const optionHash = uint64(0xfa156d885e4a4f5c)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.RootOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/options.ValueCodec"
	const optionHash = uint64(0x9cbbdca355db186d)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithGetClientSideDuration[K comparable, V any](duration time.Duration) trcache.GetOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/getOptions.ClientSideDuration"
	const optionHash = uint64(0xa52219d712468772)
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptClientSideDuration(duration)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithGetCustomParams[K comparable, V any](customParams interface{}) trcache.GetOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/getOptions.CustomParams"
	const optionHash = uint64(0xec892b258f288a85)
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithGetRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.GetOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/getOptions.RedisGetFunc"
	const optionHash = uint64(0x6d62262658e666f3)
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithSetCustomParams[K comparable, V any](customParams interface{}) trcache.SetOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/setOptions.CustomParams"
	const optionHash = uint64(0x3fceea5eebedb4f1)
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithSetRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.SetOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/setOptions.RedisSetFunc"
	const optionHash = uint64(0x9f236e7c161100d3)
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithDeleteCustomParams[K comparable, V any](customParams interface{}) trcache.DeleteOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/deleteOptions.CustomParams"
	const optionHash = uint64(0xd087ad7569e445f8)
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case deleteOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithDeleteRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.DeleteOption {
	const optionName = "github.com/RangelReale/trcache-rueidis/deleteOptions.RedisDelFunc"
	const optionHash = uint64(0xd5104f27b66ca2a5)
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case deleteOptions[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
			return true
		}
		return false
	}, optionName, optionHash)
}

type rootOptionsImpl[K comparable, V any] struct {
	callDefaultDeleteOptions  []trcache.DeleteOption
	callDefaultGetOptions     []trcache.GetOption
	callDefaultSetOptions     []trcache.SetOption
	defaultClientSideDuration time.Duration
	defaultDuration           time.Duration
	keyCodec                  trcache.KeyCodec[K]
	name                      string
	redisDelFunc              RedisDelFunc[K, V]
	redisGetFunc              RedisGetFunc[K, V]
	redisSetFunc              RedisSetFunc[K, V]
	validator                 trcache.Validator[V]
	valueCodec                trcache.Codec[V]
}

var _ options[string, string] = &rootOptionsImpl[string, string]{}

func (o *rootOptionsImpl[K, V]) OptCallDefaultDeleteOptions(options ...trcache.DeleteOption) {
	o.callDefaultDeleteOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultGetOptions(options ...trcache.GetOption) {
	o.callDefaultGetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultSetOptions(options ...trcache.SetOption) {
	o.callDefaultSetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptDefaultClientSideDuration(duration time.Duration) {
	o.defaultClientSideDuration = duration
}
func (o *rootOptionsImpl[K, V]) OptDefaultDuration(duration time.Duration) {
	o.defaultDuration = duration
}
func (o *rootOptionsImpl[K, V]) OptKeyCodec(keyCodec trcache.KeyCodec[K]) {
	o.keyCodec = keyCodec
}
func (o *rootOptionsImpl[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptionsImpl[K, V]) OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) {
	o.redisDelFunc = redisDelFunc
}
func (o *rootOptionsImpl[K, V]) OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) {
	o.redisGetFunc = redisGetFunc
}
func (o *rootOptionsImpl[K, V]) OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) {
	o.redisSetFunc = redisSetFunc
}
func (o *rootOptionsImpl[K, V]) OptValidator(validator trcache.Validator[V]) {
	o.validator = validator
}
func (o *rootOptionsImpl[K, V]) OptValueCodec(valueCodec trcache.Codec[V]) {
	o.valueCodec = valueCodec
}

type getOptionsImpl[K comparable, V any] struct {
	clientSideDuration time.Duration
	customParams       interface{}
	redisGetFunc       RedisGetFunc[K, V]
}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

func (o *getOptionsImpl[K, V]) OptClientSideDuration(duration time.Duration) {
	o.clientSideDuration = duration
}
func (o *getOptionsImpl[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *getOptionsImpl[K, V]) OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) {
	o.redisGetFunc = redisGetFunc
}

type setOptionsImpl[K comparable, V any] struct {
	customParams interface{}
	duration     time.Duration
	redisSetFunc RedisSetFunc[K, V]
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}
func (o *setOptionsImpl[K, V]) OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) {
	o.redisSetFunc = redisSetFunc
}

type deleteOptionsImpl[K comparable, V any] struct {
	customParams interface{}
	redisDelFunc RedisDelFunc[K, V]
}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}

func (o *deleteOptionsImpl[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *deleteOptionsImpl[K, V]) OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) {
	o.redisDelFunc = redisDelFunc
}
